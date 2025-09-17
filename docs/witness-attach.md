# Witness Attach Attestation Command Demo 

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
```
/Users/rahulxf/PleaseHelpMeGod/witness/bin/witness attach attestation --attestation build-attach-attestation.json manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072
```

```
witness attach attestation \
  --registry-username "manzilrahul" \
  --registry-password "<docker_pat_token>" \
  --attestation build-attach-attestation.json \
  manzilrahul/witness-kms-attach-cmd-test@sha256:8abe274a0afab7843bc09b22ff31dd35011e299bd1904c8e26554df6cae19072
```

---

### Full log 

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