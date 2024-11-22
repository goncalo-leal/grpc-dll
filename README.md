# grpc-dll
A simple gRPC server written in Go, packaged as a DLL for use with C# (or any other language) applications.

## Compile DLL
```bash
go build -o cmd_app/lib/fixture-bridge.dll -buildmode=c-shared dll.go
```
