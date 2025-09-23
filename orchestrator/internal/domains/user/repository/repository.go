package userrepository

import (
	"context"
	"errors"
	"fmt"

	mongoFactory "github.com/bharath0292/quantdrey/infrastructure/mongo"
	postgresFactory "github.com/bharath0292/quantdrey/infrastructure/postgres"
	redisFactory "github.com/bharath0292/quantdrey/infrastructure/redis"
	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
	qerrors "github.com/bharath0292/quantdrey/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"gorm.io/gorm"
)

type userRepository struct {
	postgre *postgresFactory.PostgresClient
	mongo   *mongoFactory.MongoClient
	cache   *redisFactory.RedisClient
}

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string, user *userentity.User) error
	CreateUser(ctx context.Context, input *userentity.User) (int, error)
	UpdateUser(ctx context.Context, id int, input *userentity.User) error
	DeleteUser(ctx context.Context, id int) error

	CreateBroker(ctx context.Context, config *userentity.UserBrokerConfig) (bson.ObjectID, error)
	GetBrokers(ctx context.Context, userId int) ([]*userentity.UserBrokerConfig, error)
	UpdateBroker(ctx context.Context, brokerId bson.ObjectID, input bson.M) (*userentity.UserBrokerConfig, error)
	DeleteBroker(ctx context.Context, brokerId bson.ObjectID) error
}

func NewUserRepository(
	postgre *postgresFactory.PostgresClient,
	mongo *mongoFactory.MongoClient,
	cache *redisFactory.RedisClient,
) IUserRepository {
	return &userRepository{
		postgre: postgre,
		mongo:   mongo,
		cache:   cache,
	}
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string, user *userentity.User) error {
	conditions := map[string]any{
		"email": email,
	}

	if err := u.postgre.Select(ctx, conditions, user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return qerrors.ErrNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	return nil
}

func (u *userRepository) CreateUser(ctx context.Context, user *userentity.User) (int, error) {
	if err := u.postgre.InsertOne(ctx, &user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return -1, qerrors.ErrAlreadyExists
		}
		return -1, fmt.Errorf("failed to create user: %w", err)
	}

	return user.ID, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, id int, input *userentity.User) error {
	conditions := map[string]any{
		"id": id,
	}

	if err := u.postgre.UpdateOne(ctx, conditions, input); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return qerrors.ErrAlreadyExists
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			return qerrors.ErrNotFound
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	conditions := map[string]any{
		"id": id,
	}
	if err := u.postgre.Delete(ctx, conditions, &userentity.User{}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return qerrors.ErrNotFound
		}
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (u *userRepository) CreateBroker(ctx context.Context, config *userentity.UserBrokerConfig) (bson.ObjectID, error) {
	insertedId, err := u.mongo.CreateDocument(ctx, "userBroker", config)
	if err != nil {
		return bson.NilObjectID, err
	}

	return insertedId, nil
}

func (u *userRepository) GetBrokers(ctx context.Context, userId int) ([]*userentity.UserBrokerConfig, error) {
	var brokers []*userentity.UserBrokerConfig

	conditions := bson.M{"userId": userId}
	err := u.mongo.ReadDocuments(ctx, "userBroker", conditions, &brokers)
	if err != nil {
		return nil, fmt.Errorf("failed to get brokers: %w", err)
	}

	return brokers, nil
}

func (u *userRepository) UpdateBroker(ctx context.Context, brokerId bson.ObjectID, input bson.M) (*userentity.UserBrokerConfig, error) {
	var broker userentity.UserBrokerConfig

	conditions := bson.M{"_id": brokerId}
	_, err := u.mongo.UpdateDocument(ctx, "userBroker", conditions, input, &broker)
	if err != nil {
		return nil, fmt.Errorf("failed to update brokers: %w", err)
	}

	return &broker, nil
}

func (u *userRepository) DeleteBroker(ctx context.Context, brokerId bson.ObjectID) error {

	conditions := bson.M{"_id": brokerId}
	err := u.mongo.DeleteOne(ctx, "userBroker", conditions)
	if err != nil {
		return fmt.Errorf("failed to delete brokers: %w", err)
	}

	return nil
}
