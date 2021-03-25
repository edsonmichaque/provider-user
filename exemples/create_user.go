import (
  "gitlab.com/ulombe/sdk"
  "gitlab.com/ulombe/provider-user"
  "gitlab.com/ulombe/provider-user/linux"
  "gitlab.com/ulombe/provider-user/linux/create"
)

userProvider := user.DefaultResource.Provider()

newOperation, err := userProvider(
  user.Linux,
  user.Create,
  create.WithName("beto"),
  create.WithPassword("beto"),
  create.WithGroup("michaque"),
  create.WithGroups([]string{"admin", "sudo"})
)

engine.Push(newOperation)
