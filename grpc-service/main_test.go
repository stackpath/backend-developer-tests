package main

import (
	"context"
	"net"
	"sort"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	uuid "github.com/satori/go.uuid"
	"github.com/stackpath/backend-developer-tests/grpc-service/pkg/models"
	"github.com/stackpath/backend-developer-tests/grpc-service/pkg/proto/schema"
	"github.com/stackpath/backend-developer-tests/grpc-service/pkg/proto/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func startServerAndGetClient(t *testing.T) (service.PersonClient, *grpc.ClientConn) {
	// Create a listener on a random available port
	listener, err := net.Listen("tcp", ":")
	require.NoError(t, err)

	// Start the GRPC Server
	grpcServer := grpc.NewServer()
	service.RegisterPersonServer(grpcServer, &personServer{})
	go grpcServer.Serve(listener)

	conn, err := grpc.Dial(listener.Addr().String(), grpc.WithInsecure())
	require.NoError(t, err)

	return service.NewPersonClient(conn), conn
}

// TestPeopleServer_GetPeople will create the grpc server and test requests to GetPeople.
func TestPeopleServer_GetPeople(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, conn := startServerAndGetClient(t)
	defer conn.Close()

	// Test Getting all people with no filters
	resp, err := client.GetPeople(ctx, &service.GetPeopleRequest{})
	require.NoError(t, err)
	compareGetPeopleResults(t, models.AllPeople(), resp)

	// Test Getting people with last name smith
	t.Run("all people", func(t *testing.T) {
		resp, err = client.GetPeople(ctx, &service.GetPeopleRequest{LastName: &wrappers.StringValue{Value: "Smith"}})
		require.NoError(t, err)
		compareGetPeopleResults(t, models.FindPeopleByName("", "Smith"), resp)
	})

	// Test Getting people with last name smith and first name Jane
	t.Run("last name Smith and frist name Jane", func(t *testing.T) {
		resp, err = client.GetPeople(ctx, &service.GetPeopleRequest{
			FirstName: &wrappers.StringValue{Value: "Jane"},
			LastName:  &wrappers.StringValue{Value: "Smith"},
		})
		require.NoError(t, err)
		compareGetPeopleResults(t, models.FindPeopleByName("Jane", "Smith"), resp)
	})

	// Test Getting people with the phone number +1 (800) 555-1212
	t.Run("phone number +1 (800) 555-1212", func(t *testing.T) {
		resp, err = client.GetPeople(ctx, &service.GetPeopleRequest{
			PhoneNumber: &wrappers.StringValue{Value: "+1 (800) 555-1212"},
		})
		require.NoError(t, err)
		compareGetPeopleResults(t, models.FindPeopleByPhoneNumber("+1 (800) 555-1212"), resp)
	})

	// Test Getting people with the last name Doe and phone number +1 (800) 555-1414
	t.Run("last name Doe and phone number +1 (800) 555-1414", func(t *testing.T) {
		resp, err = client.GetPeople(ctx, &service.GetPeopleRequest{
			LastName:    &wrappers.StringValue{Value: "Doe"},
			PhoneNumber: &wrappers.StringValue{Value: "+1 (800) 555-1414"},
		})
		require.NoError(t, err)

		expectedPerson, err := models.FindPersonByID(uuid.FromStringOrNil("135af595-aa86-4bb5-a8f7-df17e6148e63"))
		require.NoError(t, err)
		compareGetPeopleResults(t, []*models.Person{expectedPerson}, resp)
	})

	// Test filter with no people matching
	t.Run("no results found for filter", func(t *testing.T) {
		resp, err = client.GetPeople(ctx, &service.GetPeopleRequest{
			FirstName:   &wrappers.StringValue{Value: "Does not exist"},
			PhoneNumber: &wrappers.StringValue{Value: "+1 (800) 555-1212"},
		})
		require.NoError(t, err)
		require.Equal(t, 0, len(resp.Results))
	})

}

// TestPeopleServer_GetPerson will create the grpc server and test requests to GetPerson.
func TestPeopleServer_GetPerson(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, conn := startServerAndGetClient(t)
	defer conn.Close()

	t.Run("Find person with id 5b81b629-9026-450d-8e46-da4f8c7bd513", func(t *testing.T) {
		resp, err := client.GetPerson(ctx, &service.GetPersonRequest{
			Id: "5b81b629-9026-450d-8e46-da4f8c7bd513",
		})

		expected, err := models.FindPersonByID(uuid.FromStringOrNil("5b81b629-9026-450d-8e46-da4f8c7bd513"))
		require.NoError(t, err)
		require.Equal(t, expected, schemaToModel(resp.Person))
	})

	t.Run("Find person with id 000ebe58-b659-422b-ab48-a0d0d40bd8f9", func(t *testing.T) {
		resp, err := client.GetPerson(ctx, &service.GetPersonRequest{
			Id: "000ebe58-b659-422b-ab48-a0d0d40bd8f9",
		})

		expected, err := models.FindPersonByID(uuid.FromStringOrNil("000ebe58-b659-422b-ab48-a0d0d40bd8f9"))
		require.NoError(t, err)
		require.Equal(t, expected, schemaToModel(resp.Person))
	})

	t.Run("person not found with id 45f83c10-7cce-4a77-b32f-e58297ed6b95", func(t *testing.T) {
		resp, err := client.GetPerson(ctx, &service.GetPersonRequest{
			Id: "45f83c10-7cce-4a77-b32f-e58297ed6b95",
		})
		require.EqualError(
			t,
			err,
			"rpc error: code = Unknown desc = user ID 45f83c10-7cce-4a77-b32f-e58297ed6b95 not found",
		)
		require.Nil(t, resp)
	})
}

func compareGetPeopleResults(t *testing.T, expected []*models.Person, response *service.GetPeopleResponse) {
	responsePeople := make([]*models.Person, len(response.Results))
	for i, person := range response.Results {
		responsePeople[i] = schemaToModel(person)
	}

	sort.Slice(responsePeople, func(a, b int) bool {
		return sort.StringsAreSorted([]string{responsePeople[a].ID.String(), responsePeople[a].ID.String()})
	})
	sort.Slice(expected, func(a, b int) bool {
		return sort.StringsAreSorted([]string{expected[a].ID.String(), expected[a].ID.String()})
	})

	require.Equal(t, expected, responsePeople)
}

func schemaToModel(person *schema.Person) *models.Person {
	return &models.Person{
		ID:          uuid.FromStringOrNil(person.Id),
		FirstName:   person.FirstName,
		LastName:    person.LastName,
		PhoneNumber: person.PhoneNumber,
	}
}
