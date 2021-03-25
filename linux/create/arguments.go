package create

import (
  "github.com/ulombe/sdk"
  "github.com/ulombe/sdk/provide"
  "github.com/ulombe/provider-user/linux"
)

func WithBash() *provider.Argument {
  return &provider.Argument{
    Name: linux.CreateShell,
    Value: "/usr/bin/bash"
  }
}

func WithZsh() *provider.Argument {
  argument := &provider.Argument{
    Name: linux.CreateShell,
    Value: "/usr/bin/bash"
  }

  argument.AddArgumentValidator(
    provider.StringValidator,
    func(a *provider.Argument) error {

    }
  )
  return argument
}

func WithName(name string) *provider.Argument {
  argument := &provider.Argument{
    Name: linux.CreateName,
    Value: name,
  }

  argument.AddArgumentValidator(provider.StringValidator)
  return argument
}

func WithPassword(password string) *provider.Argument {
  return &provider.Argument{
    Name: linux.CreatePassword,
    Value: password,
  }
}

func WithHome() *provider.Argument {
  return &provider.Argument{
    Name: linux.CreateHome,
    Value: true,
  }
}

func WithHomeDir(dir string) *provider.Argument {
  return &provider.Argument{
    Name: linux.CreateHomeDir,
    Value: dir,
  }
}
