package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	restapi "github.com/stackpath/backend-developer-tests/rest-service/pkg/app/api/rest"
	restservice "github.com/stackpath/backend-developer-tests/rest-service/pkg/services/rest"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	group := &errgroup.Group{}
	ctx := context.Background()
	// RESTful web service
	group.Go(func() error {
		return startRestAPIService(ctx)
	})

	// Wait here
	if err := group.Wait(); err != nil {
		log.Panic(fmt.Sprintf("main: Failed starting REST server. %s", err))
	}
}

// startRestAPIService starts REST API service
func startRestAPIService(ctx context.Context) error {
	service, err := restservice.New()
	if err != nil {
		return err
	}

	engine, err := restapi.New(service)
	if err != nil {
		return err
	}

	if err := engine.Run(); err != nil {
		return errors.New("Failed to start REST API")
	}

	return nil
}
