package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/guneyin/todo-app/todo-api/handler"
	"github.com/stretchr/testify/require"

	"os"
	"testing"

	pact "github.com/pact-foundation/pact-go/v2/provider"
)

func mockServer() *httptest.Server {
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return httptest.NewServer(mux)
}

func Test_E2E(t *testing.T) {
	ms := mockServer()
	defer ms.Close()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../todo-ui/pacts", dir)

	verifier := pact.NewVerifier()
	err := verifier.VerifyProvider(t, pact.VerifyRequest{
		ProviderBaseURL: ms.URL,
		PactFiles:       []string{fmt.Sprintf("%s/todo-ui-todo-api.json", pactDir)},
	})
	require.NoError(t, err)
}
