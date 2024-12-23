package arangodb

import (
	"github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/connection"
)

// Credentials holds the connection information
// for the ArangoDB client.
// Endpoint: The URL of the ArangoDB server.
// SkipVerify: Whether to skip verification of the server's certificate.
type Credentials struct {
	Endpoints  []string
	Username   string
	Password   string
	SkipVerify bool
}

type Client arangodb.Client

func NewClient(c Credentials) (Client, error) {

	// Create a connection
	endpoint := connection.NewRoundRobinEndpoints(c.Endpoints)
	conn := connection.NewHttp2Connection(
		connection.DefaultHTTP2ConfigurationWrapper(endpoint, c.SkipVerify))

	// Add authentication
	auth := connection.NewBasicAuth(c.Username, c.Password)
	err := conn.SetAuthentication(auth)
	if err != nil {
		return nil, err
	}

	// Create a client
	return arangodb.NewClient(conn), nil
}
