package userProvider

import (
	"errors"
	"gitlab.com/ulombe/sdk"
	"gitlab.com/ulombe/sdk/provider"
	"gitlab.com/ulombe/provider-user/linux"
)

const (
	depth = 2
	operation = 1
	os = 0
	Linux = "linux"
	Create = "create"
	Update = "update"
	Delete = "delete"
)

func NewProvider() sdk.Provider {
	return func(levels ...string, args ...provider.Argument) (provider.Operation, error) {
		if len(levels) != depth {
			return nil, errors.New("Must have 2 levels")
		}

		switch levels[os] {
		case Linux:
			switch levels[operation] {
			case Create:
				return linux.CreateUser(args...)
			case Update:
				return linux.UpdateUser(args...)
			case Delete:
				return linux.DeleteUser(args...)
			}
		}

		return nil, errors.New("Cannot find provider")
	}
}
