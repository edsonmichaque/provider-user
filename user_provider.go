package userProvider

import (
	"errors"
	"gitlab.com/ulombe/sdk"
	"gitlab.com/ulombe/sdk/provider"
	"gitlab.com/ulombe/provider-user/linux"
)

const (
	providerLevels = 2
	providerLevelOperation = 1
	providerLevelOS = 0
	linuxProvider = "linux"
	linuxCreateOperation = "create"
	linuxUpdateOperation = "update"
	linuxDeleteOperation = "delete"
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
	if len(levels) != providerLevels {
		return nil, errors.New("Must have 2 levels")
	}

	switch levels[providerLevelOS] {
	case linuxProvider:
		switch levels[providerLevelOperation] {
		case linuxCreateOperation:
			return linux.CreateUser(args...)
		case linuxUpdateOperation:
			return linux.UpdateUser(args...)
		case linuxDeleteOperation:
			return linux.DeleteUser(args...)
		}
	}

	return nil, errors.New("Cannot find provider")
}
