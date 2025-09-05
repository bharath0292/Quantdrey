package postgresFactory

type PostgresConfig struct{}

type PostgresClient struct{}

func NewPostgresClient(config PostgresConfig) (*PostgresClient, error) {
	return &PostgresClient{}, nil
}

func (pc *PostgresClient) Close() error {
	return nil
}
