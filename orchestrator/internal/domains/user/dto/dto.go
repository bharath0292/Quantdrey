package userdto

import (
	"time"

	userentity "github.com/bharath0292/quantdrey/internal/domains/user/entity"
)

type UserInput struct {
	Email    *string
	Password *string
	DOB      *time.Time
	Country  *string
}

func (u *UserInput) ToEntity() *userentity.User {
	user := &userentity.User{}

	if u.Email != nil {
		user.Email = *u.Email
	}

	if u.Password != nil {
		user.Password = u.Password
	}

	if u.DOB != nil {
		user.DOB = u.DOB
	}

	if u.Country != nil {
		user.Country = u.Country
	}

	return user
}
