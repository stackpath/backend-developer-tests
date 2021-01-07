package main

import (
	"fmt"

	"github.com/stackpath/backend-developer-tests/grpc-service/pkg/proto/service"
)

// personServer should implement service.PersonServer and interact with the backend model data for serving requests.
//
// TODO: Implement service.PersonServer methods.
//   - GetPeople(context.Context, *service.GetPeopleRequest) (*service.GetPeopleResponse, error)
//   - GetPerson(context.Context, *service.GetPersonRequest) (*service.GetPersonResponse, error)
type personServer struct {
	service.PersonServer
}

// main will create the GRPC Server and run a series of tests on the server to verify the proper output is given.
func main() {
	fmt.Println("SP// Backend Developer Test - GRPC Service")
	fmt.Println()

	// TODO: Register and start the grpc server.
}
