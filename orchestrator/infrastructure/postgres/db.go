package postgresFactory

import "context"

type PostgresConfig struct{}

type PostgresClient struct{}

func NewPostgresClient(ctx context.Context, config PostgresConfig) (*PostgresClient, error) {
	return &PostgresClient{}, nil
}

func (pc *PostgresClient) Close() error {
	return nil
}
