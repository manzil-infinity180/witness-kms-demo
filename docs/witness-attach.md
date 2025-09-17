# Witness Attach Attestation Command Demo 

## Using DockerHub
### Build your image 
```
docker build -t witness-kms-attach-cmd-test .
```

### Tag your image and push to OCI
- I am using dockerhub here 
```
docker tag witness-kms-attach-cmd-test manzilrahul/witness-kms-attach-cmd-test:latest
docker push manzilrahul/witness-kms-attach-cmd-test:latest
```

### Create a Keypair
DOCS: https://witness.dev/docs/docs/tutorials/getting-started#1-create-a-keypair
```
openssl genpkey -algorithm ed25519 -outform PEM -out attachTest.pem
openssl pkey -in attachTest.pem -pubout > attachPublic.pem
```

### Generate attestation about the build process
```
witness run --step build -k attachTest.pem -o build-attach-attestation.json -- docker build .
```

### Attach the attestation to the docker image with verification
```ts
/Users/rahulxf/PleaseHelpMeGod/witness/bin/witness attach attestation \
--attestation build-attach-attestation.json \
manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072
```

```ts
witness attach attestation \
  --registry-username "manzilrahul" \
  --registry-password "<docker_pat_token>" \
  --attestation build-attach-attestation.json \
  manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072
```

<img width="1048" height="902" alt="Screenshot 2025-09-17 at 6 13 12 PM" src="https://github.com/user-attachments/assets/5845097d-1700-47d9-80be-dacfee234640" />

<img width="1048" height="902" alt="Screenshot 2025-09-17 at 6 13 36 PM" src="https://github.com/user-attachments/assets/f02c19de-ae7a-4dcb-8d0c-b49a8599c8ec" />

<details>
<summary>Full logs</summary>

```rs
witness-kms-demo (main) $ docker build -t witness-kms-attach-cmd-test .
[+] Building 1.7s (1/3)                                                                                           docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                              0.0s
[+] Building 45.2s (14/14) FINISHED                                                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => => transferring dockerfile: 598B                                                                                              0.0s
 => WARN: FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)                                                    0.0s
 => [internal] load metadata for cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5  16.7se => [internal] load metadata for cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4   16.2s
 => [auth] chainguard/go:pull token for cgr.dev                                                                                   0.0s
 => [auth] chainguard/static:pull token for cgr.dev                                                                               0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => => transferring context: 2B                                                                                                   0.0s
 => [builder 1/4] FROM cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4              0.0s
 => => resolve cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4                      0.0s
 => CACHED [stage-1 1/2] FROM cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511   0.0s
 => => resolve cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511                  0.0s
 => [internal] load build context                                                                                                 0.1s
 => => transferring context: 17.18MB                                                                                              0.1s
 => CACHED [builder 2/4] WORKDIR /build                                                                                           0.0s
 => [builder 3/4] COPY . .                                                                                                        1.1s
 => [builder 4/4] RUN go build -o bin/software                                                                                   27.1s
 => [stage-1 2/2] COPY --from=builder /build/bin/software /software                                                               0.0s 
 => exporting to image                                                                                                            0.1s 
 => => exporting layers                                                                                                           0.1s
 => => exporting manifest sha256:57e852d38bffcf80796bde40e1bef866b365a75284abd2452e67d1600fda4a50                                 0.0s
 => => exporting config sha256:d51dbd3636c24bd193405ed956792c5d7039e8716c2c4b3cce476fb9ad3044c9                                   0.0s
 => => exporting attestation manifest sha256:bbdca7c4bd7eb1b8fac48b23571108f6f0fcc2d908e6e42d7fe3164f1802b0ae                     0.0s
 => => exporting manifest list sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072                            0.0s
 => => naming to docker.io/library/witness-kms-attach-cmd-test:latest                                                             0.0s
 => => unpacking to docker.io/library/witness-kms-attach-cmd-test:latest                                                          0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/wg9hip3n9b34ptsso7lfm5ibh

 1 warning found (use docker --debug to expand):
 - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)


witness-kms-demo (main) $ docker tag witness-kms-attach-cmd-test manzilrahul/witness-kms-attach-cmd-test:latest
witness-kms-demo (main) $ docker push manzilrahul/witness-kms-attach-cmd-test:latest

The push refers to repository [docker.io/manzilrahul/witness-kms-attach-cmd-test]
cea5e50a208c: Pushed 
f04769f37838: Pushed 
a93cf116b643: Mounted from manzilrahul/software 
latest: digest: sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072 size: 855


witness-kms-demo (main) $ openssl genpkey -algorithm ed25519 -outform PEM -out attachTest.pem
witness-kms-demo (main) $ openssl pkey -in attachTest.pem -pubout > attachPublic.pem
witness-kms-demo (main) $ witness run --step build -k attachTest.pem -o build-attach-attestation.json -- docker build .

INFO    Starting prematerial attestors stage...      
INFO    Starting git attestor...                     
INFO    Starting environment attestor...             
INFO    Finished environment attestor... (0.000123625s) 
INFO    Finished git attestor... (0.062413208s)      
INFO    Completed prematerial attestors stage...     
INFO    Starting material attestors stage...         
INFO    Starting material attestor...                
INFO    Finished material attestor... (0.027784459s) 
INFO    Completed material attestors stage...        
INFO    Starting execute attestors stage...          
INFO    Starting command-run attestor...             
#0 building with "desktop-linux" instance using docker driver

#1 [internal] load build definition from Dockerfile
#1 transferring dockerfile: 598B done
#1 WARN: FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)
#1 DONE 0.0s

#2 [internal] load metadata for cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511
#2 DONE 0.0s

#3 [internal] load metadata for cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4
#3 DONE 0.0s

#4 [internal] load .dockerignore
#4 transferring context: 2B done
#4 DONE 0.0s

#5 [stage-1 1/2] FROM cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511
#5 resolve cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511 done
#5 DONE 0.0s

#6 [builder 1/4] FROM cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4
#6 resolve cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4 done
#6 DONE 0.0s

#7 [internal] load build context
#7 transferring context: 13.73kB done
#7 DONE 0.0s

#8 [builder 2/4] WORKDIR /build
#8 CACHED

#9 [builder 3/4] COPY . .
#9 DONE 0.1s

#10 [builder 4/4] RUN go build -o bin/software
#10 0.117 go: downloading go1.24.2 (linux/arm64)
#10 18.94 go: downloading github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
#10 DONE 22.3s

#11 [stage-1 2/2] COPY --from=builder /build/bin/software /software
#11 CACHED

#12 exporting to image
#12 exporting layers done
#12 exporting manifest sha256:57e852d38bffcf80796bde40e1bef866b365a75284abd2452e67d1600fda4a50 done
#12 exporting config sha256:d51dbd3636c24bd193405ed956792c5d7039e8716c2c4b3cce476fb9ad3044c9 done
#12 exporting attestation manifest sha256:8b55fd77457200b5166e03edbf4fc4fb901881339d10ec94e579aca5eced444d done
#12 exporting manifest list sha256:fba5d62954a18461e32d36cbe7532ce87429c43084c7152ca9ebe873be53a26d done
#12 naming to moby-dangling@sha256:fba5d62954a18461e32d36cbe7532ce87429c43084c7152ca9ebe873be53a26d done
#12 unpacking to moby-dangling@sha256:fba5d62954a18461e32d36cbe7532ce87429c43084c7152ca9ebe873be53a26d done
#12 DONE 0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/lemt1glpr62lafdamq57f0qot

 1 warning found (use docker --debug to expand):
 - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)
INFO    Finished command-run attestor... (22.867333166999998s) 
INFO    Completed execute attestors stage...         
INFO    Starting product attestors stage...          
INFO    Starting product attestor...                 
INFO    Finished product attestor... (0.0407865s)    
INFO    Completed product attestors stage...         
INFO    Starting postproduct attestors stage...      
INFO    Completed postproduct attestors stage...  


witness-kms-demo (main) $ /Users/rahulxf/PleaseHelpMeGod/witness/bin/witness attach attestation --attestation build-attach-attestation.json manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072

INFO    manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072 
INFO    att                                          
INFO    ref: manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072 
INFO    desc mediaType: application/vnd.oci.image.index.v1+json 
INFO    index                                        
INFO    AttachAttestationToUnknown:                  
INFO    index Attestations                           
INFO    attestations                                 
INFO    sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072 
INFO    img %s<nil>                                  
INFO    &{0x14000166180}                             
INFO    remote.write                                 
INFO    index.docker.io/manzilrahul/witness-kms-attach-cmd-test:sha256-8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072.att 
```
</details>

---

## Using Github Registry 

### Build your image 
```
docker build -t witness-kms-attach-cmd-ghcr-test .
```

### Login
```
export GHCR_PAT="<github_pat_token>"
echo $GHCR_PAT | docker login ghcr.io -u manzilrahul --password-stdin
```

### Tag your image and push to OCI
- I am using GitHub registry here 
```
docker tag witness-kms-attach-cmd-ghcr-test ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:latest
docker push ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:latest
```

### Generate attestation about the build process
```
witness run --step build -k attachTest.pem -o build-attach-attestation2.json -- docker build .
```

### Attach the attestation to the docker image with verification
Here I am attaching multiple attestations

```ts
witness attach attestation \
--attestation <first_attestation.json> \
--attestation <second_attestation.json> \
<image_uri>
```

```ts
/Users/rahulxf/PleaseHelpMeGod/witness/bin/witness attach attestation \
--attestation build-attach-attestation.json \
--attestation build-attach-attestation2.json \
ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3
```

<img width="1048" height="902" alt="Screenshot 2025-09-17 at 6 39 01 PM" src="https://github.com/user-attachments/assets/244ce576-8e63-4efa-ab94-d6739ee77325" />

<img width="1048" height="902" alt="Screenshot 2025-09-17 at 6 38 54 PM" src="https://github.com/user-attachments/assets/0f818b88-e1f2-4aca-ae58-c3d8976ac41b" />

<details>
<summary>Full logs</summary>
  
```js
witness-kms-demo (main) $ docker build -t witness-kms-attach-cmd-ghcr-test .
[+] Building 25.5s (14/14) FINISHED                                                                               docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                                              0.0s
 => => transferring dockerfile: 598B                                                                                              0.0s
 => WARN: FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)                                                    0.0s
 => [internal] load metadata for cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a55  3.7s
 => [internal] load metadata for cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4    3.7s
 => [auth] chainguard/go:pull token for cgr.dev                                                                                   0.0s
 => [auth] chainguard/static:pull token for cgr.dev                                                                               0.0s
 => [internal] load .dockerignore                                                                                                 0.0s
 => => transferring context: 2B                                                                                                   0.0s
 => [builder 1/4] FROM cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4              0.0s
 => => resolve cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4                      0.0s
 => CACHED [stage-1 1/2] FROM cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511   0.0s
 => => resolve cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511                  0.0s
 => [internal] load build context                                                                                                 0.1s
 => => transferring context: 5.86MB                                                                                               0.1s
 => CACHED [builder 2/4] WORKDIR /build                                                                                           0.0s
 => [builder 3/4] COPY . .                                                                                                        0.1s
 => [builder 4/4] RUN go build -o bin/software                                                                                   21.3s
 => [stage-1 2/2] COPY --from=builder /build/bin/software /software                                                               0.0s 
 => exporting to image                                                                                                            0.1s 
 => => exporting layers                                                                                                           0.1s
 => => exporting manifest sha256:c4b7c75482c0cc4cad1728c0f78a0fea12e0a692d6b0d60aa84ea311476e11c7                                 0.0s
 => => exporting config sha256:8c785e889b5848a591d6890ec8e053c51d9895ffbb387a242919d5e0b0da2453                                   0.0s
 => => exporting attestation manifest sha256:5a3930f15dc0f8e5ba8636115a92c7118567337900661f879b1e029151178f8e                     0.0s
 => => exporting manifest list sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3                            0.0s
 => => naming to docker.io/library/witness-kms-attach-cmd-ghcr-test:latest                                                        0.0s
 => => unpacking to docker.io/library/witness-kms-attach-cmd-ghcr-test:latest                                                     0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/96k9qj84w6p58mqirqzju01so

 1 warning found (use docker --debug to expand):
 - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)


witness-kms-demo (main) $ export GHCR_PAT="<gh_pat_token>"
witness-kms-demo (main) $ echo $GHCR_PAT | docker login ghcr.io -u manzilrahul --password-stdin
Login Succeeded


witness-kms-demo (main) $ docker tag witness-kms-attach-cmd-ghcr-test ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:latest
witness-kms-demo (main) $ docker push ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:latest


The push refers to repository [ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test]
0d72d00cd212: Pushed 
2777e2d58afc: Pushed 
a93cf116b643: Pushed 
latest: digest: sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 size: 855


witness-kms-demo (main) $ witness run --step build -k attachTest.pem -o build-attach-attestation2.json -- docker build .

INFO    Starting prematerial attestors stage...      
INFO    Starting git attestor...                     
INFO    Starting environment attestor...             
INFO    Finished environment attestor... (0.000167042s) 
INFO    Finished git attestor... (0.11030075s)       
INFO    Completed prematerial attestors stage...     
INFO    Starting material attestors stage...         
INFO    Starting material attestor...                
INFO    Finished material attestor... (0.028796959s) 
INFO    Completed material attestors stage...        
INFO    Starting execute attestors stage...          
INFO    Starting command-run attestor...             
#0 building with "desktop-linux" instance using docker driver

#1 [internal] load build definition from Dockerfile
#1 transferring dockerfile: 598B done
#1 WARN: FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)
#1 DONE 0.0s

#2 [internal] load metadata for cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4
#2 DONE 0.0s

#3 [internal] load metadata for cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511
#3 DONE 0.0s

#4 [internal] load .dockerignore
#4 transferring context: 2B done
#4 DONE 0.0s

#5 [stage-1 1/2] FROM cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511
#5 resolve cgr.dev/chainguard/static@sha256:a432665213f109d5e48111316030eecc5191654cf02a5b66ac6c5d6b310a5511 done
#5 DONE 0.0s

#6 [builder 1/4] FROM cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4
#6 resolve cgr.dev/chainguard/go@sha256:c87a8cf30c9a4e58df04712e1bb0b98d8d0421cc924e88f6fca4a6fabf45c6b4 done
#6 DONE 0.0s

#7 [internal] load build context
#7 transferring context: 18.38kB done
#7 DONE 0.0s

#8 [builder 2/4] WORKDIR /build
#8 CACHED

#9 [builder 3/4] COPY . .
#9 CACHED

#10 [builder 4/4] RUN go build -o bin/software
#10 CACHED

#11 [stage-1 2/2] COPY --from=builder /build/bin/software /software
#11 CACHED

#12 exporting to image
#12 exporting layers done
#12 exporting manifest sha256:c4b7c75482c0cc4cad1728c0f78a0fea12e0a692d6b0d60aa84ea311476e11c7 done
#12 exporting config sha256:8c785e889b5848a591d6890ec8e053c51d9895ffbb387a242919d5e0b0da2453 done
#12 exporting attestation manifest sha256:2d54cadd943f383a158bf3130ed87a2a9f4cdda2b46dd2fdf28f43471aa32ade done
#12 exporting manifest list sha256:de290a04664ddabf6f2ac88fa5e67fed3e576ed3a2e379daef8c1e6da2ef6222 done
#12 naming to moby-dangling@sha256:de290a04664ddabf6f2ac88fa5e67fed3e576ed3a2e379daef8c1e6da2ef6222 done
#12 unpacking to moby-dangling@sha256:de290a04664ddabf6f2ac88fa5e67fed3e576ed3a2e379daef8c1e6da2ef6222 done
#12 DONE 0.0s

View build details: docker-desktop://dashboard/build/desktop-linux/desktop-linux/36yvw4mm8j4dr37mt2y7nwkig

 1 warning found (use docker --debug to expand):
 - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 1)
INFO    Finished command-run attestor... (0.537953s) 
INFO    Completed execute attestors stage...         
INFO    Starting product attestors stage...          
INFO    Starting product attestor...                 
INFO    Finished product attestor... (0.011959875s)  
INFO    Completed product attestors stage...         
INFO    Starting postproduct attestors stage...      
INFO    Completed postproduct attestors stage...

  
witness-kms-demo (main) $ /Users/rahulxf/PleaseHelpMeGod/witness/bin/witness attach attestation --attestation build-attach-attestation.json --attestation build-attach-attestation2.json ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3

INFO    ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    att                                          
INFO    ref: ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    desc mediaType: application/vnd.oci.image.index.v1+json 
INFO    index                                        
INFO    AttachAttestationToUnknown:                  
INFO    index Attestations                           
INFO    attestations                                 
INFO    sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    img %s<nil>                                  
INFO    &{0x14000150480}                             
INFO    remote.write                                 
INFO    ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:sha256-317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3.att 
INFO    ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    att                                          
INFO    ref: ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test@sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    desc mediaType: application/vnd.oci.image.index.v1+json 
INFO    index                                        
INFO    AttachAttestationToUnknown:                  
INFO    index Attestations                           
INFO    attestations                                 
INFO    sha256:317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3 
INFO    img %s&{0x1400005f900 ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:sha256-317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3.att} 
INFO    &{0x140007a0600}                             
INFO    remote.write                                 
INFO    ghcr.io/manzil-infinity180/witness-kms-attach-cmd-ghcr-test:sha256-317fc563c5b95d640c443a673c01d33fd58e7fc3d1585370976b8d27b99a23c3.att 
witness-kms-demo (main) $ 
```
<details>
