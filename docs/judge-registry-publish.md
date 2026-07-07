# Judge platform: sync attestations to OCI registries (registry-publish)

Testing guide for the Judge platform's attestation registry-publish feature
(and the `witness attach` CLI, kept at the end). Judge automatically publishes
attestations ingested into its embedded Archivista to the OCI registry named by
the attestation's subject, filtered by predicate type, laid out
cosign-compatibly (DSSE envelope layers under the `sha256-<image digest>.att`
tag in the same repository as the image).

## Verified results

| Registry | Status | Evidence |
|---|---|---|
| Local (ggcr in-memory) | ✅ PASS | `.att` tag + payload round-trip + filter negative |
| Docker Hub | ✅ PASS | `manzilrahul/witness-kms-attach-cmd-test:sha256-8477b500….att` |
| GitHub (ghcr) | ✅ PASS | `ghcr.io/manzil-infinity180/witness-attach-test:sha256-8477b500….att` |
| AWS ECR | ⏳ pending | see ECR section |

---

## How it works (30 seconds)

```
attestation upload → embedded Archivista → Redis stream → subscriber
  → predicate-type filter + subject → RegistryPublishWorkflow (durable)
  → push DSSE to <subject repo>:sha256-<digest>.att
```

- **Which registry gets pushed** is decided by the attestation's **subject**: a
  subject named `ghcr.io/owner/app` with a `sha256` digest publishes to that
  repo. Subjects without an explicit registry host (`myapp`, attestor URIs) are
  skipped by design.
- **Filter**: only attestations whose predicateType matches a configured prefix
  are published (e.g. "only SLSA provenance and SBOMs").
- **Idempotent**: re-publishing the same envelope is a no-op (no duplicate
  layers), so retries are safe.

## Configuration

In `testifysec.env` (gitignored; judge-api loads it at boot):

```bash
REGISTRY_PUBLISH_ENABLED=true
REGISTRY_PUBLISH_PREDICATE_TYPES=https://slsa.dev/provenance,https://spdx.dev/Document
# credentials — see below; often not needed at all
```

**Credentials — simplest first (verified working):** if you're already
`docker login`'d to the target registry and standalone was started with
`sudo -E` (HOME preserved), Judge's keychain fallback picks up your existing
credentials — no config needed. This is how both the Docker Hub and ghcr runs
above authenticated.

Explicit credentials (CI, or when the keychain isn't available):

| Registry | `REGISTRY_PUBLISH_USERNAME` | `REGISTRY_PUBLISH_PASSWORD` |
|---|---|---|
| Docker Hub | `manzilrahul` | Docker Hub PAT |
| ghcr | `manzil-infinity180` | GitHub PAT (`write:packages`) |
| AWS ECR | `AWS` | `aws ecr get-login-password --region us-east-1` (12h expiry) |

POC limitation: one global credential set → test one registry at a time (or
use the keychain, which resolves per-host). Per-tenant registry credentials
are the productization follow-up.

After editing `testifysec.env`, reload:
```bash
sudo -E env "PATH=$HOME/go/bin:$PATH" jade dev reload
```

## Test procedure

### 1. Push an image, note its digest

```bash
# Docker Hub
docker push manzilrahul/witness-kms-attach-cmd-test:latest

# ghcr (package auto-creates on first push)
docker login ghcr.io -u manzil-infinity180        # PAT with write:packages
docker push ghcr.io/manzil-infinity180/witness-attach-test:latest

# ECR (repo must pre-exist!)
aws ecr create-repository --repository-name witness-ecr-testing --region us-east-1
aws ecr get-login-password --region us-east-1 | \
  docker login --username AWS --password-stdin 756546933635.dkr.ecr.us-east-1.amazonaws.com
docker push 756546933635.dkr.ecr.us-east-1.amazonaws.com/witness-ecr-testing:latest
```

> No local image? `crane append -t <repo>:latest -f any.tar` pushes a minimal
> one without docker:
> `go run github.com/google/go-containerregistry/cmd/crane@v0.20.7 append ...`

### 2. Create an attestation whose subject names the image

```bash
python3 - "<registry>/<repo>" "<sha256-hex-no-prefix>" <<'EOF'
import base64, json, sys
repo, digest = sys.argv[1], sys.argv[2]
statement = {
    "_type": "https://in-toto.io/Statement/v0.1",
    "subject": [{"name": repo, "digest": {"sha256": digest}}],
    "predicateType": "https://spdx.dev/Document",   # must match the filter
    "predicate": {"spdxVersion": "SPDX-2.3", "name": "demo-sbom", "packages": []},
}
env = {
    "payload": base64.b64encode(json.dumps(statement).encode()).decode(),
    "payloadType": "application/vnd.in-toto+json",
    "signatures": [{"keyid": "test-key", "sig": base64.b64encode(b"test-sig").decode()}],
}
json.dump(env, open("att.json", "w"))
print("wrote att.json for", repo, "@sha256:" + digest)
EOF
```

Docker Hub subjects need the full host: `index.docker.io/manzilrahul/<repo>`.

### 3. Upload to judge

Two auth quirks: the internal token goes in a **`Token:` header** (not
`Authorization: Bearer`), and an explicit **tenant header** is required.

```bash
TOKEN=$(cd ~/work-dir/judge && jade dev token)
curl -sk -X POST "https://platform.testifysec.localhost/archivista/upload" \
  -H "Token: $TOKEN" \
  -H "X-Archivista-TenantID: <tenant-oid>" \
  -H "Content-Type: application/json" \
  --data-binary @att.json
# → {"gitoid":"<gitoid>"}  HTTP 200
```

Find your tenant oid:
`sqlite3 ~/work-dir/judge/.standalone/judge.db "select oid, name from tenants;"`

### 4. Verify — registry side

The workflow is async; the tag appears within a few seconds.

```bash
CRANE="go run github.com/google/go-containerregistry/cmd/crane@v0.20.7"

# the .att tag exists next to the image
$CRANE ls <registry>/<repo>
# → latest, sha256-<digest>.att

# manifest: one DSSE layer with the witness annotation
$CRANE manifest <registry>/<repo>:sha256-<digest>.att
# → layers[0].mediaType == application/vnd.dsse.envelope.v1+json
# → layers[0].annotations has dev.witnessproject.witness/signature

# round-trip the payload: pull the blob and decode the statement
LAYER=$($CRANE manifest <registry>/<repo>:sha256-<digest>.att | python3 -c "import json,sys; print(json.load(sys.stdin)['layers'][0]['digest'])")
$CRANE blob <registry>/<repo>@$LAYER | python3 -c "import json,sys,base64; e=json.load(sys.stdin); s=json.loads(base64.b64decode(e['payload'])); print(s['predicateType'], s['subject'][0])"

# bonus: cosign discovers it too (same layout)
cosign tree <registry>/<repo>@sha256:<digest>

# ECR-native check
aws ecr describe-images --repository-name witness-ecr-testing --region us-east-1 \
  --query 'imageDetails[].imageTags'
```

Docker Hub / ghcr UIs also show the `.att` tag next to `latest`.

### 5. Verify — judge side (the receipts)

**The attestation was stored** (download it back by gitoid):

```bash
curl -sk "https://platform.testifysec.localhost/archivista/download/<gitoid>" \
  -H "Token: $TOKEN" -H "X-Archivista-TenantID: <tenant-oid>" \
  | python3 -c "import json,sys,base64; e=json.load(sys.stdin); print(base64.b64decode(e['payload']).decode())"
```

**The workflow ran and what it pushed** (durable-engine state, straight from
the DB):

```bash
sqlite3 ~/work-dir/judge/.standalone/judge.db \
  "select name, runtime_status, created_time, output from workflow_instances
   where name='RegistryPublishWorkflow' order by created_time desc limit 5;"
# → COMPLETED | {"succeeded":true,"published":["<repo>:sha256-<digest>.att"]}
```

Boot-log check (feature registered — activities fail silently if not):

```
registering workflow  workflow=RegistryPublishWorkflow
registering activity  activity=RegistryPublishActivity
```

### 6. Negative probes (worth demoing)

- **Filter**: upload an attestation with `predicateType:
  https://witness.dev/attestations/git/v0.1` (outside the filter) → ingested by
  Archivista, but **no workflow row is created** and no tag is pushed. The
  filter blocks at dispatch, before any work happens.
- **Idempotency**: upload the same `att.json` twice → second workflow completes
  with the reference in `skipped` instead of `published`; registry layer count
  unchanged.
- **Bad/expired credentials** (e.g. ECR after 12h) → workflow fails with the
  registry's 401/403 in `failure_details`; durabletask retries per policy.

## Gotchas

- `jade dev token` / `jade dev reload` are **cwd- and user-scoped** — run from
  the judge worktree; a sudo-started standalone needs a sudo reload.
- Started standalone with `--skip-build`? The running binary may predate the
  feature — check the boot log for `RegistryPublishWorkflow` registration.
- ECR: the image repo must pre-exist (`aws ecr create-repository`); the `.att`
  tag lands in the *same* repo, so no extra repo is needed.
- `REGISTRY_PUBLISH_ALLOW_HTTP=true` is only for local registries.

---

## Appendix: witness attach CLI (same layout, manual flow)

The CLI equivalent (branch `attach-cmd-v2`,
[manzil-infinity180/witness#4](https://github.com/manzil-infinity180/witness/pull/4)).
Two changes vs the original PR #661 flow documented in
[witness-attach.md](./witness-attach.md):

1. attestation files are **positional**, the image is `--image-uri`;
2. **subject verification is on by default** — the attestation subject digest
   must match the image digest, or pass `--skip-verification`.

```bash
# Docker Hub
witness attach attestation att.json \
  --registry-username manzilrahul --registry-password <docker_pat> \
  --image-uri manzilrahul/witness-kms-attach-cmd-test@sha256:<digest>

# ghcr
witness attach attestation att.json \
  --registry-username manzil-infinity180 --registry-password <github_pat> \
  --image-uri ghcr.io/manzil-infinity180/witness-attach-test@sha256:<digest>

# AWS ECR
witness attach attestation att.json \
  --registry-username AWS \
  --registry-password "$(aws ecr get-login-password --region us-east-1)" \
  --image-uri 756546933635.dkr.ecr.us-east-1.amazonaws.com/witness-ecr-testing@sha256:<digest>

# local/HTTP registry
witness attach attestation att.json --allow-http-registry \
  --image-uri 127.0.0.1:1338/demo/app@sha256:<digest>
```

Omit the credential flags to use your `docker login` keychain. Verify with the
same registry-side commands as step 4 above.
