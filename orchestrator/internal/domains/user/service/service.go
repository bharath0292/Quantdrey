package userservice

import (
	"context"
	"errors"
	"time"

	userdto "github.com/bharath0292/quantdrey/internal/domains/user/dto"
	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
	userrepository "github.com/bharath0292/quantdrey/internal/domains/user/repository"
	qerrors "github.com/bharath0292/quantdrey/pkg/errors"
	util "github.com/bharath0292/quantdrey/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type userService struct {
	repo userrepository.IUserRepository
}

type IUserService interface {
	LoginUser(ctx context.Context, email, password string) (*userentity.User, error)
	SignupUser(ctx context.Context, input *userdto.UserInput) (*userentity.User, error)
	UpdateUser(ctx context.Context, id int, input *userdto.UserInput) (*userentity.User, error)
	DeleteUser(ctx context.Context, id int) (bool, error)

	CreateBroker(ctx context.Context, config *userdto.UserBrokerConfigInput) (bson.ObjectID, error)
	GetBrokers(ctx context.Context, userId int) ([]*userentity.UserBrokerConfig, error)
	UpdateBroker(ctx context.Context, brokerId bson.ObjectID, config *userdto.UserBrokerConfigInput) (*userentity.UserBrokerConfig, error)
	DeleteBroker(ctx context.Context, brokerId bson.ObjectID) error
}

func NewUserService(
	repo userrepository.IUserRepository,
) IUserService {
	return &userService{repo: repo}
}

func (u *userService) LoginUser(ctx context.Context, email, password string) (*userentity.User, error) {
	user := userentity.User{}
	err := u.repo.GetUserByEmail(ctx, email, &user)
	if err != nil {
		if errors.Is(err, qerrors.ErrNotFound) {
			return nil, qerrors.New(qerrors.CodeNotFound, "user not found", nil)
		}

		return nil, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}

	if !util.CheckPasswordHash(password, *user.Password) {
		return nil, qerrors.New(qerrors.CodeAuthentication, "invalid email or password", nil)
	}

	return &user, nil
}

func (u *userService) SignupUser(ctx context.Context, input *userdto.UserInput) (*userentity.User, error) {
	hashedPassword, err := util.HashPassword(*input.Password)
	if err != nil {
		return nil, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}

	user := userentity.User{
		Email:    *input.Email,
		Password: &hashedPassword,
		DOB:      input.DOB,
		Country:  input.Country,
	}

	_, err = u.repo.CreateUser(ctx, &user)
	if err != nil {
		if errors.Is(err, qerrors.ErrAlreadyExists) {
			return nil, qerrors.New(qerrors.CodeConflict, "email already exists", nil)
		}

		return nil, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}
	return &user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, input *userdto.UserInput) (*userentity.User, error) {
	userEntity := input.ToEntity()

	if userEntity.Password != nil {
		// Seperate func for updating password
		userEntity.Password = nil
	}

	err := s.repo.UpdateUser(ctx, id, userEntity)
	if err != nil {
		if errors.Is(err, qerrors.ErrAlreadyExists) {
			return nil, qerrors.New(qerrors.CodeConflict, "email already exists", nil)
		}
		if errors.Is(err, qerrors.ErrNotFound) {
			return nil, qerrors.New(qerrors.CodeNotFound, "user not found", nil)
		}

		return nil, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}

	return userEntity, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) (bool, error) {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, qerrors.ErrNotFound) {
			return false, qerrors.New(qerrors.CodeNotFound, "user not found", nil)
		}
		return false, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}
	return true, nil
}

func (u *userService) CreateBroker(ctx context.Context, config *userdto.UserBrokerConfigInput) (bson.ObjectID, error) {
	if config == nil {
		return bson.NilObjectID, qerrors.New(qerrors.CodeValidation, "broker config missing", nil)
	}

	marshalToMap := func(in any) (map[string]any, error) {
		bsonBytes, err := bson.Marshal(in)
		if err != nil {
			return nil, err
		}
		var m map[string]any
		if err := bson.Unmarshal(bsonBytes, &m); err != nil {
			return nil, err
		}
		return m, nil
	}

	var brokerName string
	var credMap map[string]any
	var err error

	switch {
	case config.Flattrade != nil:
		credMap, err = marshalToMap(config.Flattrade)
		if err != nil {
			return bson.NilObjectID, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
		}
		brokerName = "flattrade"
	default:
		return bson.NilObjectID, qerrors.New(qerrors.CodeValidation, "unsupported broker config", nil)
	}

	userbroker := userentity.UserBrokerConfig{
		UserId:      1,
		BrokerName:  brokerName,
		Credentials: credMap,
		CreatedAt:   time.Now(),
	}

	insertedId, err := u.repo.CreateBroker(ctx, &userbroker)
	if err != nil {
		return bson.NilObjectID, qerrors.New(qerrors.CodeInternal, qerrors.ErrInternal.Error(), err)
	}

	return insertedId, nil
}

func (u *userService) GetBrokers(ctx context.Context, userId int) ([]*userentity.UserBrokerConfig, error) {
	return u.repo.GetBrokers(ctx, userId)
}

func (u *userService) UpdateBroker(
	ctx context.Context,
	brokerId bson.ObjectID,
	input *userdto.UserBrokerConfigInput,
) (*userentity.UserBrokerConfig, error) {

	setFields := bson.M{}

	switch {
	case input.Flattrade != nil:
		if input.Flattrade.ClientId != "" {
			setFields["credentials.clientId"] = input.Flattrade.ClientId
		}
		if input.Flattrade.ApiKey != "" {
			setFields["credentials.apiKey"] = input.Flattrade.ApiKey
		}
		if input.Flattrade.ApiSecret != "" {
			setFields["credentials.apiSecret"] = input.Flattrade.ApiSecret
		}
	}

	setFields["updatedAt"] = time.Now()

	updateInput := bson.M{"$set": setFields}

	broker, err := u.repo.UpdateBroker(ctx, brokerId, updateInput)
	if err != nil {
		return nil, err
	}

	return broker, nil
}

func (u *userService) DeleteBroker(ctx context.Context, brokerId bson.ObjectID) error {
	return u.repo.DeleteBroker(ctx, brokerId)
}
