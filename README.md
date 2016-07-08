# EGOL

## Dependencies
---
- [Go](https://golang.org/) programming language binaries with the `GOPATH` environment variable specified.
- [Go](https://golang.org/) version 1.6, or version 1.5 with the `GO15VENDOREXPERIMENT` environment variable set to `1`.
- [NodeJS](http://nodejs.org/) JavaScript runtime.
- [gulp](http://http://gulpjs.com/) build toolkit (npm install gulp -g).

## Development
---
Clone the repository:

```bash
mkdir $GOPATH/src/github.com/unchartedsoftware
cd $GOPATH/src/github.com/unchartedsoftware
git clone git@github.com:unchartedsoftware/egol.git
```

Install dependencies

```bash
cd egol
make deps
```

Run the server:

```bash
gulp
```
