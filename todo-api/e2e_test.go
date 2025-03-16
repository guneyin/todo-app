package main_test

import (
	"fmt"
	"github.com/guneyin/todo-app/todo-api/server"
	"github.com/stretchr/testify/require"

	pact "github.com/pact-foundation/pact-go/v2/provider"
	"log"
	"os"
	"testing"
)

func Test_Todo(t *testing.T) {
	go startMockServer()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../todo-ui/pacts", dir)

	verifier := pact.NewVerifier()
	err := verifier.VerifyProvider(t, pact.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:3000",
		PactFiles:       []string{fmt.Sprintf("%s/todo-ui-todo-api.json", pactDir)},
	})
	require.NoError(t, err)
}

func startMockServer() {
	if err := server.StartServer("3000"); err != nil {
		log.Fatal(err)
	}
}
