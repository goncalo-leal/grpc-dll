// This file describes a DLL that exposes a gRPC server that listens on port.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"unsafe"

	pb "github.com/goncalo-leal/go-fixture/proto/data"

	"google.golang.org/grpc"
)

/*
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

typedef void (*Callback)(const char* data);

// Wrapper function to safely call the callback
static inline void invokeCallback(Callback cb, const char* data) {
    if (cb != NULL) {
        cb(data);
    }
}
*/
import "C"

// Global variable to store the C callback
var callbackFunction C.Callback

// server is used to implement each service.
type server struct {
	pb.UnimplementedDataServiceServer
}

// DataCallback implements the service.
func (s *server) DataCallback(ctx context.Context, req *pb.DataReceived) (*pb.DataResponse, error) {

	// transform the data into a string
	data := fmt.Sprintf("%v", req.Data)

	// Call the C# callback using the stored function pointer
	if callbackFunction != nil {
		cStr := C.CString(data) // Convert Go string to C string
		defer C.free(unsafe.Pointer(cStr))

		// Explicitly invoke the C function pointer
		// Use the C wrapper to invoke the callback
		C.invokeCallback(callbackFunction, cStr)
	} else {
		log.Println("No callback registered")
	}

	// Send a response back to the gRPC client
	return &pb.DataResponse{Status: "Success"}, nil
}

//export StartServer
func StartServer(port *C.char) {
	lis, err := net.Listen("tcp", C.GoString(port))
	if err != nil {
		log.Fatalf("failed to listen on port 8080: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//export SetCallback
func SetCallback(f C.Callback) {
	callbackFunction = f
}

func main() {}
