package resolver

import (
	"context"

	userdto "github.com/bharath0292/quantdrey/internal/domains/user/dto"
	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
	qerrors "github.com/bharath0292/quantdrey/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (r *queryResolver) LoginUser(ctx context.Context, email, password string) (*userentity.User, error) {
	return r.userService.LoginUser(ctx, email, password)
}

func (r *mutationResolver) SignupUser(ctx context.Context, input userdto.UserInput) (*userentity.User, error) {

	if input.Email == nil {
		return nil, qerrors.New(qerrors.CodeValidation, "missing email", nil)
	}

	if input.Password == nil {
		return nil, qerrors.New(qerrors.CodeValidation, "missing password", nil)
	}

	return r.userService.SignupUser(ctx, &input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input userdto.UserInput) (*userentity.User, error) {
	return r.userService.UpdateUser(ctx, id, &input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	return r.userService.DeleteUser(ctx, id)
}

func (r *mutationResolver) CreateBroker(ctx context.Context, input userdto.UserBrokerConfigInput) (bson.ObjectID, error) {
	switch {
	case input.Flattrade != nil:
		if input.Flattrade.ClientId == "" {
			return bson.NilObjectID, qerrors.New(qerrors.CodeValidation, "missing client id", nil)
		}
		if input.Flattrade.ApiKey == "" {
			return bson.NilObjectID, qerrors.New(qerrors.CodeValidation, "missing api key", nil)
		}
		if input.Flattrade.ApiSecret == "" {
			return bson.NilObjectID, qerrors.New(qerrors.CodeValidation, "missing api secret", nil)
		}
	}
	return r.userService.CreateBroker(ctx, &input)
}

func (r *queryResolver) GetUsersBrokers(ctx context.Context) ([]*userentity.UserBrokerConfig, error) {
	return r.userService.GetBrokers(ctx, 1)
}

func (r *mutationResolver) UpdateBroker(
	ctx context.Context,
	brokerId bson.ObjectID,
	input userdto.UserBrokerConfigInput,
) (*userentity.UserBrokerConfig, error) {
	return r.userService.UpdateBroker(ctx, brokerId, &input)
}

func (r *mutationResolver) DeleteBroker(ctx context.Context, brokerId bson.ObjectID) (bool, error) {
	if err := r.userService.DeleteBroker(ctx, brokerId); err != nil {
		return false, err
	}
	return true, nil
}
