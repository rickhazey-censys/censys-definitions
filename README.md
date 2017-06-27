# censys-definitions
Censys internal Protobuf objects

## Regenerating
Only do this on OS X in a clean Python 2.7 virtualenv.

The build script requires that [sh](https://amoffat.github.io/sh/) is installed
in your virtualenv.

```console
$ # In a virtualenv
$ pip install sh
```

If you already have grpc or protobuf installed, remove it. Note that there's
several formula's for grpc and protobuf that may be installed in Homebrew.

```console
$ brew uninstall grpc/grpc/grpc
$ brew uninstall grpc
$ brew uninstall google-protobuf
$ brew uninstall protobuf
```

Install the versions from the [ZMap Homebrew
Tap](https://github.com/zmap/homebrew-formula).

```console
$ brew tap zmap/homebrew-formula
$ brew install zmap/formula/grpc@1.2
$ ./build.py
```

Install Protobuf for Golang
```console
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

Regenerate everything:
```console
$ ./build.py
```
