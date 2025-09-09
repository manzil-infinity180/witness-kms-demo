# AWS KMS 

```bash
# Install AWS CLI
aws configure
# Or set environment variables
export AWS_ACCESS_KEY_ID="your-access-key"
export AWS_SECRET_ACCESS_KEY="your-secret-key"
export AWS_DEFAULT_REGION="us-east-1"
```

```bash
# Create a KMS key for signing
aws kms create-key \
  --description "Witness signing key" \
  --key-usage SIGN_VERIFY \
  --key-spec ECC_NIST_P256

# Note the KeyId from the response, e.g., 1234abcd-12ab-34cd-56ef-1234567890ab

# Create an alias (optional but recommended)
aws kms create-alias \
  --alias-name alias/witness-signing-key \
  --target-key-id 1234abcd-12ab-34cd-56ef-1234567890ab
```

### Signing
* Creating the attestation `test.json`
```bash
witness run -s test -o test.json --signer-kms-ref=awskms:///arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd -- echo "hello world" > hello.txt
```
### Create witness policy

```bash
aws kms get-public-key \
  --key-id "arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd" \
  --query PublicKey --output text | base64 -d > public-key.der

```

```bash
openssl rsa -pubin -inform DER -in public-key.der -outform PEM -out public-key.pem
```

```bash
openssl ec -pubin -inform DER -in public-key.der -outform PEM -out public-key.pem
```
* Put the publickey as base64 in `publickeys.key`

```json
{
  "expires": "2035-12-17T23:57:40-05:00",
  "steps": {
    "test": {
      "name": "test",
      "attestations": [
        {
          "type": "https://witness.dev/attestations/command-run/v0.1"
        },
        {
          "type": "https://witness.dev/attestations/product/v0.1"
        },
        {
          "type": "https://witness.dev/attestations/environment/v0.1"
        }
      ],
      "functionaries": [
        {
          "type": "publickey",
          "publickeyid": "awskms:///arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd"
        }
      ]
    }
  },
  "publickeys": {
    "awskms:///arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd": {
      "keyid": "awskms:///arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd",
      "key":"LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFQTNyY2FnU2dyQ1VUYW5BbVZYYTN0V1NhRWZCZgpKMkVRZUF4Q29LRnJGd1JLWG5DTFJQTi9lcWw0U25xSG5RZ2RFNHVod0hvK3pUNzFQSHFNUnlpbXZRPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg=="
    }
  }
}
```

```bash
witness sign -f policy.json -o policy-signed.json --signer-kms-ref=awskms:///arn:aws:kms:us-east-1:756546933635:key/fcc0689a-5b7f-4779-983f-39be43cba0bd
```

### Verifying 
```bash
witness verify -p policy-signed.json -a test.json --verifier-kms-ref=gcpkms://projects/test-project/locations/europe-west2/keyRings/test-keyring/cryptoKeys/test-key -f test.txt
```