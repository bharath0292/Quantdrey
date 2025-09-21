package postgresFactory

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

func (pc *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		pc.Host, pc.User, pc.Password, pc.DBName, pc.Port, pc.SSLMode, pc.TimeZone,
	)
}

type PostgresClient struct {
	client *gorm.DB
}

func NewPostgresClient(config *PostgresConfig) (*PostgresClient, error) {
	if config == nil {
		return nil, errors.New("missing config")
	}

	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Warn),
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database object: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresClient{client: db}, nil
}

// func (pc *PostgresClient) SelectQuery(ctx context.Context, query string, args ...any) error {
// 	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()

// 	result := pc.client.WithContext(ctxWithTimeout).Raw(query, args...)
// 	if result.Error != nil {
// 		return fmt.Errorf("failed to select record: %w", result.Error)
// 	}
// 	return nil
// }

func (pc *PostgresClient) Select(ctx context.Context, conditions map[string]any, value any) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result := pc.client.WithContext(ctxWithTimeout).Find(value, conditions)
	if result.Error != nil {
		return fmt.Errorf("failed to select record: %w", result.Error)
	}
	return nil
}

func (pc *PostgresClient) UpdateOne(ctx context.Context, conditions map[string]any, model any) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := pc.client.WithContext(ctxWithTimeout).Model(model)
	if len(conditions) > 0 {
		tx = tx.Where(conditions)
	}

	result := tx.Updates(model)
	if result.Error != nil {
		return fmt.Errorf("failed to update record: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pc *PostgresClient) InsertOne(ctx context.Context, value any) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result := pc.client.WithContext(ctxWithTimeout).Create(value)
	if result.Error != nil {
		return fmt.Errorf("failed to create record: %w", result.Error)
	}

	return nil
}

func (pc *PostgresClient) Delete(ctx context.Context, conditions map[string]interface{}, model any) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx := pc.client.WithContext(ctxWithTimeout).Model(model)
	if len(conditions) > 0 {
		tx = tx.Where(conditions)
	}

	result := tx.Delete(model)
	if result.Error != nil {
		return fmt.Errorf("failed to delete record: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pc *PostgresClient) Close() error {
	if pc.client != nil {
		sqlDB, err := pc.client.DB()
		if err != nil {
			return fmt.Errorf("failed to get database object: %w", err)
		}
		sqlDB.Close()
	}
	return nil
}
