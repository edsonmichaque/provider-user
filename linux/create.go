package linux

import (
	"errors"
	"github.com/ulombe/sdk"
	"github.com/ulombe/sdk/provider"
)

const (
		createMainCommand = "useradd"
		createTemplateFile = "templates/create.tmpl"
)

const (
		createGroupFlag = "-g"
		createGroupsFlag = "-G"
		createSystemFlag = "-r"
		createShellFlag = "-s"
		createCommentFlag = "-c"
		createHomeFlag = "-m"
		createHomeDirFlag = "-d"
)

const (
	CreateUID = "uid"
	CreateGID = "gid"
	CreateName = "name"
	CreateLock = "lock"
	CreatePassword = "password"
	CreateExpire = "expire"
	CreateComment = "comment"
	CreateGroup = "group"
	CreateGroups = "groups"
	CreateSystem = "system"
	CreateShell = "shell"
	CreateHome = "home"
	CreateHomeDir = "home_dir"
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

func (u *CreateUserOperation) addValidators() {
		u.validators = []Validator{
			func(a *Validable) error {
				createUser := a.(CreateUserOperation)

				if a.Name == "" {
					return errors.New("Username shouldn't be empty")
				}
			},
		}
}

func CreateUser(args ...provider.Argument) (*provider.Operation, error) {
	user := CreateUserOperation{}

	for _, arg := range args {
		if arg.Name == CreateName {
			user.Name = arg.Value.(string)
		}

		if arg.Name == CreateShell {
			user.Shell = arg.Value.(string)
		}

		if arg.Name == CreatePassword {
			user.Password = arg.Value.(string)
		}

		if arg.Name == CreateComment {
			user.Comment = arg.Value.(string)
		}
	}

	user.addValidators()
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

func (p *CreateUserOperation) Validate() error {
	for _, validator := range p.validators {
		if err := validator(p); err != nil {
			return err
		}
	}

	return nil
}

func (p CreateUserOperation) Command() string {
	command := make([]string, 0)

	command = append(command, createMainCommand)

	if u.System {
		command = append(command, createSystemFlag)
	}

	if u.Shell != "" {
		command = append(command, createShellFlag, u.Shell)
	}

	if u.Group != "" {
		command = append(command, createGroupFlag, u.Group)
	}

	if len(u.Groups) > 0 {
		command = append(command, createGroupsFlag, strings.Join(u.Groups, ","))
	}

	if u.Comment != "" {
		command = append(command, createCommentFlag, fmt.Sprintf("\"%s\"", u.Comment))
	}

	if u.Name != "" {
		command = append(command, u.Name)
	}

	if u.Home && (u.HomeDir == "") {
		command = append(command, createHomeFlag)
		u.HomeDir = fmt.Sprintf("/home/%s", u.Name)
	}

	if u.HomeDir != "" {
		command = append(command, createHomeFlag)
		command = append(command, createHomeDirFlag, u.HomeDir)
	}

	return strings.Join(command, " ")
}

func (u *CreateUserOperation) Render() {
	tmpl := template.Must(template.ParseFiles(createTemplateFile))
	tmpl.Execute(os.Stdout, p)
}
