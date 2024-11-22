# grpc-dll
A simple gRPC server written in Go, packaged as a DLL for use with C# (or any other language) applications.

## Compile DLL

- Compiling for windows using linux:
```sh
sudo apt-get install gcc-mingw-w64

env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o cmd_app/lib/fixture-bridge.dll -buildmode=c-shared dll.go
```

- Compiling for windows on windows
```sh
go build -o cmd_app/lib/fixture-bridge.dll -buildmode=c-shared dll.go
```

