package linux

import (
	"github.com/edsonmichaque/ulombe/pkg/types"
)

type DeleteUserOperationOption func(*DeleteUserOperation)

type DeleteUserOperation struct {
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

func NewDeleteUserOperation(name string, options ...DeleteUserOperationOption) *DeleteUserOperation {
	h := &DeleteUserOperation{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	}
}

func DeleteUser(args ...provider.Argument) (*provider.Operation, error) {
	user := DeleteUserOperation{}

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

func WithBash() DeleteUserOperationOption {
	return func(u *DeleteUserOperation) {
		u.Shell = "/usr/bin/bash"
	}
}

func WithZsh() DeleteUserOperationOption {
	return func(u *DeleteUserOperation) {
		u.Shell = "/usr/bin/zsh"
	}
}

func System() DeleteUserOperationOption {
	return func(u *DeleteUserOperation) {
		u.System = true
	}
}

func WithGroups(list []string) DeleteUserOperationOption {
	return func(u *DeleteUserOperation) {
		u.Groups = list
	}
}

func WithGroup(group string) DeleteUserOperationOption {
	return func(u *DeleteUserOperation) {
		u.Group = group
	}
}

func (p DeleteUserOperation) Validate() error {
}

func (p DeleteUserOperation) command() string {
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

func (u *DeleteUserOperation) Render() {
	tmpl := template.Must(template.ParseFiles("templates/create.tmpl"))
	tmpl.Execute(os.Stdout, p)
}
