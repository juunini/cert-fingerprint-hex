# Cert fingerprint hex

## Introduction

This code is local test for WebTransport API.  
Download pre-built binary from [release page](https://github.com/juunini/cert-fingerprint-hex/releases)

```js
const fingerprintHexString = "<fingerprint.hex content>";

let fingerprint = [];
for (let c = 0; c < fingerprintHexString.length-1; c += 2) {
    fingerprint.push(parseInt(fingerprintHexString.substring(c, c+2), 16));
}

const wt = new WebTransport("https://localhost/webtransport", {
  allowPooling: false,
  serverCertificateHashes: [{
    algorithm: "sha-256",
    value: new Uint8Array(fingerprint),
  }],
});
```

## Usage

You can get custom mkcert from https://github.com/juunini/mkcert/releases  
(added `-days` option)  

```bash
./mkcert -ecdsa -install
./mkcert -ecdsa -cert-file cert.pem -key-file key.pem -days 13 localhost 127.0.0.1 ::1
./fingerprint -cert cert.pem -output fingerprint.hex
```

Result of `fingerprint.hex` is equal of this script

```bash
openssl x509 -in cert.pem -outform der | openssl dgst -sha256 > ./fingerprint.hex
```

## Caution

mkcert option `days` must be set under **13**  
(see: https://developer.mozilla.org/en-US/docs/Web/API/WebTransport/WebTransport#servercertificatehashes)

> The certificate must be an X.509v3 certificate that has a validity period of less that 2 weeks, and the current time must be within that validity period. The format of the public key in the certificate depends on the implementation, but must minimally include ECDSA with the secp256r1 (NIST P-256) named group, and must not include RSA keys. An ECSDA key is therefore an interoperable default public key format. A user agent may add further requirements; these will be listed in the browser compatibility section if known.
