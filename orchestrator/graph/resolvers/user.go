package resolver

import (
	"context"

	userdto "github.com/bharath0292/quantdrey/internal/domains/user/dto"
	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
	qerrors "github.com/bharath0292/quantdrey/pkg/errors"
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
