package linux

import (
	"github.com/edsonmichaque/ulombe/pkg/types"
)

type CreateUserOperationOption func(*CreateUserOperation)

type CreateUserOperation struct {
	UID string
	GID string
	Name string
	Password string
	Expire string
	Comment string
	Group string
	Groups []string
	System bool
	Shell string
	Home bool
	HomeDir string
}

func NewCreateUserOperation(name string, options ...CreateUserOperationOption) *CreateUserOperation {
	h := &CreateUserOperation{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	}
}

func CreateUser(args ...provider.Argument) (*provider.Operation, error) {
	user := CreateUserOperation{}

	for _, arg := range args {
		if arg.Name == "name" {
			user.Name = arg.Value.(string)
		}

		if arg.Name == "shell" {
			user.Shell = arg.Value.(string)
		}

		if arg.Name == "password" {
			user.Password = arg.Value.(string)
		}

		if arg.Name == "comment" {
			user.Comment = arg.Value.(string)
		}
	}

	return &user
}

func WithBash() CreateUserOperationOption {
	return func(u *CreateUserOperation) {
		u.Shell = "/usr/bin/bash"
	}
}

func WithZsh() CreateUserOperationOption {
	return func(u *CreateUserOperation) {
		u.Shell = "/usr/bin/zsh"
	}
}

func System() CreateUserOperationOption {
	return func(u *CreateUserOperation) {
		u.System = true
	}
}

func WithGroups(list []string) CreateUserOperationOption {
	return func(u *CreateUserOperation) {
		u.Groups = list
	}
}

func WithGroup(group string) CreateUserOperationOption {
	return func(u *CreateUserOperation) {
		u.Group = group
	}
}

func (p CreateUserOperation) Validate() error {
}

func (p CreateUserOperation) command() string {
	command := make([]string, 0)

	command = append(command, "useradd")

	if u.System {
		command = append(command, "-r")
	}

	if u.Shell != "" {
		command = append(command, "-s", u.Shell)
	}

	if u.Group != "" {
		command = append(command, "-g", u.Group)
	}

	if len(u.Groups) > 0 {
		command = append(command, "-G", strings.Join(u.Groups, ","))
	}

	if u.Comment != "" {
		command = append(command, "-c", fmt.Sprintf("\"%s\"", u.Comment))
	}

	if u.Name != "" {
		command = append(command, u.Name)
	}

	if u.Home && (u.HomeDir == "") {
		command = append(command, "-m")
		u.HomeDir = fmt.Sprintf("/home/%s", u.Name)
	}

	if u.HomeDir != "" {
		command = append(command, "-m")
		command = append(command, "-d", u.HomeDir)
	}

	return strings.Join(command, " ")
}

func (u *CreateUserOperation) Render() {
	tmpl := template.Must(template.ParseFiles("templates/create.tmpl"))
	tmpl.Execute(os.Stdout, p)
}
