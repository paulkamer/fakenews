# Fakenews

_Fakenews_ a fake/dummy RSS & ATOM news feed generator, for integration testing purposes.

## Prerequisites

_Fakenews_ serves HTTP on port 8080 and HTTPS on port `8443`. For HTTPS you need to generate a self-signed certificate with:

```sh
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```

Save the `.pem` files in the same folder as `main.go`.

## Usage

Start with: `go run main.go`, or with [Air](https://github.com/air-verse/air), using simply: `air`.

To run as a Docker container, run:

```sh
docker build . -t fakenews:latest

docker run -p 8080:8080 -p 8443:8443 fakenews:latest
```

## TODO

- [x] Add endpoints for semi-invalid responses
  - [x] RSS response but with Atom header and vice versa
  - [x] RSS response but with HTML header
- [x] Add endpoints for invalid responses
  - [x] HTML (error page or whatever) returned as RSS
  - [x] Malformed XML
- [x] Support Atom 1.0
- [x] Weird redirects
- [x] http -> https and vice versa
- [x] Dockerize solution
- [ ] Support Docker compose
- [ ] Upgrade dependencies
- [ ] Publish as Go module
- [ ] Responses with wrong/unexpected status codes
- [ ] Endpoint that randomly returns a vaid/semi-valid/invalid response

## Useful links

- https://en.wikipedia.org/wiki/RSS
- https://en.wikipedia.org/wiki/Atom_(web_standard)
- https://gin-gonic.com/
- https://github.com/brianvoe/gofakeit
- https://github.com/air-verse/air