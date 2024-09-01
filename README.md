# Fake news generator

- https://gin-gonic.com/
- https://github.com/brianvoe/gofakeit
- https://en.wikipedia.org/wiki/RSS
- https://en.wikipedia.org/wiki/Atom_(web_standard)

Serves HTTPS on port `8443`. Generate a self-signed certificate with:

```sh
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

## TODO

- [x] Add endpoints for semi-invalid responses
  - [x] RSS response but with Atom header and vice versa
  - [x] RSS response but with HTML header
- [x] Add endpoints for invalid responses
  - [x] HTML (error page or whatever) returned as RSS
  - [x] Malformed XML
- [x] Support Atom 1.0
- [ ] Responses with wrong/unexpected status codes
- [ ] Weird redirects
- [ ] http -> https and vice versa
