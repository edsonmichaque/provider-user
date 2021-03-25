import (
  "gitlab.com/ulombe/sdk"
  "gitlab.com/ulombe/provider-user"
  "gitlab.com/ulombe/provider-user/linux"
  "gitlab.com/ulombe/provider-user/linux/create"
)

newUserProvider := userProvider.New()

newOperation, err := newUserProvider.Execute(
  userProvider.Linux,
  userProvider.Create,
  create.WithName("beto"),
  create.WithPassword("beto"),
  create.WithGroup("michaque"),
  create.WithGroups([]string{"admin", "sudo"})
)

newOperation.Render(writer)
