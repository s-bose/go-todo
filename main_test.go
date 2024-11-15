package main

import (
	"log"
	"testing"

	"github.com/ory/dockertest/v3"
)

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")

	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker, %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=postgres",
			"listen_addresses = '*'",
		},
	})
}
