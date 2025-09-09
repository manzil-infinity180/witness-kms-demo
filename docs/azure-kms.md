# Azure KMS

# Install Azure CLI
`az login`
### Or set environment variables for service principal authentication
`export AZURE_TENANT_ID="your-tenant-id"`
`export AZURE_CLIENT_ID="your-client-id"` 
`export AZURE_CLIENT_SECRET="your-client-secret"`

```rs
# Create a resource group (if needed)
az group create --name witness-rg --location eastus

# Create a Key Vault
az keyvault create \
  --name witness-vault \
  --resource-group witness-rg \
  --location eastus

# Create a key for signing (EC P-256 curve, suitable for signing)
az keyvault key create \
  --vault-name witness-vault \
  --name witness-signing-key \
  --kty EC \
  --curve P-256 \
  --ops sign verify

# Note the key identifier from the response
# Format: https://witness-vault.vault.azure.net/keys/witness-signing-key/{version}
```

```rs
# Get the public key from Azure Key Vault
az keyvault key show \
  --vault-name witness-vault \
  --name witness-signing-key \
  --query "key" > azure-key.json

# Convert to PEM format (you may need to write a small script for this conversion)
# Azure returns the key in JWK format, which needs conversion to PEM
```

## 3. Create attestation
```sh
witness run -s test -o test.json \
  --signer-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key \
  -- echo "hello world" > hello.txt
```

## 4. Sign policy
```rs
witness sign -f policy.json -o policy-signed.json \
  --signer-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key
```

## 5. Verify
```rs
witness verify -p policy-signed.json -a test.json \
  --verifier-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key \
  -f hello.txt
```



---
```js
witness (main) $ ./bin/witness run -s test -o test.json --signer-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key -- echo "hello world" > hello.txt

witness (main) $ az keyvault key download \                              
  --vault-name witness-vault \
  --name witness-signing-key \
  --file azure-public-key.pem \
  --encoding PEM
witness (main) $ cat azure-public-key.pem | base64 | tr -d '\n'

## Create policy and replace the public key

witness (main) $ ./bin/witness sign \                                    
-f policy.json \
-o policy-signed.json \
--signer-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key
witness (main) $ cat policy-signed.json | jq -r .payload | base64 -d | jq

witness (main) $ ./bin/witness verify \
  -p policy-signed.json \
  -a test.json \
  --verifier-kms-ref=azurekms://witness-vault.vault.azure.net/witness-signing-key \
  -f hello.txt
INFO    Starting verify attestors stage...           
INFO    Starting policyverify attestor...            
INFO    policy signature verified                    
INFO    Finished policyverify attestor... (12.308939875s) 
INFO    Completed verify attestors stage...          
INFO    Verification succeeded                       
INFO    Evidence:                                    
INFO    Step: test                                   
INFO    0: test.json                                 
INFO    1: test.json                                 
INFO    2: test.json                                 




```
