package linux

import (
	"errors"
	"gitlab.com/ulombe/sdk"
	"gitlab.com/ulombe/sdk/provider"
	"gitlab.com/ulombe/provider-user/linux"
)

type UserProvider struct {
}

func New() *UserProvider {
	return &UserProvider{}
}

func(u UserProvider) Name() string {
	return "user"
}

func(u UserProvider) Provides() map[]string {
	return []string{
		"linux",
	}
}

func(u UserProvider) Execute(levels ...string, args ...provider.Argument) (provider.Operation, error) {
	if len(level) != 2 {
		return nil, errors.New("Must have 2 levels")
	}

	switch levels[0] {
	case "linux":
		switch levels[1] {
		case "create":
			return linux.CreateUser(args...)
		case "update":
			return linux.UpdateUser(args...)
		case "delete":
			return linux.DeleteUser(args...)
		}
	}

	return nil, errors.New("Cannot find provider")
}
