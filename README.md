# TLSencrypt
A simple implementation of a TCP/TLS connection respecting modern versions of Golang clients

```
[req]
default_bits = 2048
prompt = no
default_md = sha256
distinguished_name = dn
req_extensions = req_ext

[dn]
C = GR
ST = Attica
L = Athens
O = MyCompany
CN = localhost

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
IP.1 = 127.0.0.1
```

```bash
touch server.cnf
# Add the required configurations with SANS implementation

# Generate a private key
openssl genrsa -out server.key 2048

# Generate a Certificate Signing Request (CSR)
openssl req -new -key server.key -out server.csr -config server.cnf

# Generate a Self-Signed Certificate (valid for 1 year)
openssl x509 -req -in server.csr -signkey server.key -out server.crt -days 365 -extfile server.cnf -extensions req_ext


go run client/client.go
go run server/server.go
```
