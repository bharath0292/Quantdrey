package userservice

import (
	"context"
	"errors"

	userdto "github.com/bharath0292/quantdrey/internal/domains/user/dto"
	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
	userrepository "github.com/bharath0292/quantdrey/internal/domains/user/repository"
	qerrors "github.com/bharath0292/quantdrey/pkg/errors"
	util "github.com/bharath0292/quantdrey/pkg/utils"
)

type userService struct {
	repo userrepository.IUserRepository
}

type IUserService interface {
	LoginUser(ctx context.Context, email, password string) (*userentity.User, error)
	SignupUser(ctx context.Context, input *userdto.UserInput) (*userentity.User, error)
	UpdateUser(ctx context.Context, id int, input *userdto.UserInput) (*userentity.User, error)
	DeleteUser(ctx context.Context, id int) (bool, error)
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

		return nil, qerrors.New(qerrors.CodeInternal, "something went wrong", err)
	}

	if !util.CheckPasswordHash(password, *user.Password) {
		return nil, qerrors.New(qerrors.CodeAuthentication, "invalid email or password", nil)
	}

	return &user, nil
}

func (u *userService) SignupUser(ctx context.Context, input *userdto.UserInput) (*userentity.User, error) {
	hashedPassword, err := util.HashPassword(*input.Password)
	if err != nil {
		return nil, qerrors.New(qerrors.CodeInternal, "something went wrong", err)
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

		return nil, qerrors.New(qerrors.CodeInternal, "something went wrong", err)
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

		return nil, qerrors.New(qerrors.CodeInternal, "something went wrong", err)
	}

	return userEntity, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) (bool, error) {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, qerrors.ErrNotFound) {
			return false, qerrors.New(qerrors.CodeNotFound, "user not found", nil)
		}
		return false, qerrors.New(qerrors.CodeInternal, "something went wrong", err)
	}
	return true, nil
}
