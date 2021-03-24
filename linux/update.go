package linux

import (
	"gitlab.com/ulombe/sdk"
	"gitlab.com/ulombe/sdk/provider"
)

const (
		updateMainCommand = "usermod"
		updateTemplateFile = "templates/update.tmpl"
)

const (
	updateSystemFlag = "-r"
	updateShellFlag = "-s"
	updateHomeDirFlag = "-d"
	updateHomeFlag = "-m"
	updateGroupFlag = "-g"
	updateGroupsFlag = "-G"
	updateAppendFlag = "-a"
)

const (
	UpdateUID = "uid"
	UpdateGID = "gid"
	UpdateName = "name"
	UpdateNewName = "new_name"
	UpdateMoveHome = "move_home"
	UpdateAppend = "append"
	UpdateUnlock = "unlock"
	UpdateLock = "lock"
	UpdatePassword = "password"
	UpdateExpire = "expire"
	UpdateComment = "comment"
	UpdateGroup = "group"
	UpdateGroups = "groups"
	UpdateSystem = "system"
	UpdateShell = "shell"
	UpdateHome = "home"
	UpdateHomeDir = "home_dir"
)

type UpdateUserOperationOption func(*UpdateUserOperation)

type UpdateUserOperation struct {
	UID string
	GID string
	Name string
	NewName string
	MoveHome bool
	Append bool
	Unlock bool
	Lock bool
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

func NewUpdateUserOperation(name string, options ...UpdateUserOperationOption) *UpdateUserOperation {
	h := &UpdateUserOperation{
		Name: name,
	}

	for _, option := range options {
		option(handler)
	}
}

func UpdateUser(args ...provider.Argument) (*provider.Operation, error) {
	user := UpdateUserOperation{}

	for _, arg := range args {
		if arg.Name == UpdateName {
			user.Name = arg.Value.(string)
		}

		if arg.Name == UpdateNewName {
			user.NewName = arg.Value.(string)
		}

		if arg.Name == UpdateMoveHome" {
			user.MoveHome = arg.Value.(bool)
		}

		if arg.Name == UpdateUnlock {
			user.Unlock = arg.Value.(bool)
		}

		if arg.Name == UpdateLock {
			user.Lock = arg.Value.(bool)
		}

		if arg.Name == UpdateShell {
			user.Shell = arg.Value.(string)
		}

		if arg.Name == UpdatePassword {
			user.Password = arg.Value.(string)
		}

		if arg.Name == UpdateComment {
			user.Comment = arg.Value.(string)
		}
	}

	return &user
}

func WithBash() UpdateUserOperationOption {
	return func(u *UpdateUserOperation) {
		u.Shell = "/usr/bin/bash"
	}
}

func WithZsh() UpdateUserOperationOption {
	return func(u *UpdateUserOperation) {
		u.Shell = "/usr/bin/zsh"
	}
}

func System() UpdateUserOperationOption {
	return func(u *UpdateUserOperation) {
		u.System = true
	}
}

func WithGroups(list []string) UpdateUserOperationOption {
	return func(u *UpdateUserOperation) {
		u.Groups = list
	}
}

func WithGroup(group string) UpdateUserOperationOption {
	return func(u *UpdateUserOperation) {
		u.Group = group
	}
}

func (p UpdateUserOperation) Validate() error {
}

func (p UpdateUserOperation) Command() string {
	command := make([]string, 0)

	command = append(command, updateMainCommand)

	if u.System {
		command = append(command, updateSystemFlag)
	}

	if u.Shell != "" {
		command = append(command, updateShellFlag, u.Shell)
	}

	if u.Group != "" {
		command = append(command, updateGroupFlag, u.Group)
	}

	if u.Append {
		command = append(command, updateAppendFlag)
	}

	if len(u.Groups) > 0 {
		command = append(command, updateGroupsFlag, strings.Join(u.Groups, ","))
	}

	if u.Comment != "" {
		command = append(command, updateCommentFlag, fmt.Sprintf("\"%s\"", u.Comment))
	}

	if u.Name != "" {
		command = append(command, u.Name)
	}

	if u.Home && (u.HomeDir == "") {
		command = append(command, updateHomeDirFlag)
		u.HomeDir = fmt.Sprintf("/home/%s", u.Name)
	}

	if u.HomeDir != "" {
		command = append(command, updateHomeDirFlag)
		command = append(command, updateHomeFlag, u.HomeDir)
	}

	return strings.Join(command, " ")
}

func (u *UpdateUserOperation) Render() {
	tmpl := template.Must(template.ParseFiles(updateTemplateFile))
	tmpl.Execute(os.Stdout, p)
}
