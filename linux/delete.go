package linux

import (
	"errors"
	"github.com/ulombe/ulombe/sdk"
	"github.com/ulombe/ulombe/sdk/provider"
)

const (
		deleteMainCommand = "userdel"
		deleteTemplateFile = "templates/delete.tmpl"
)

const (
	forceFlag = "-f"
)

const (
	DeleteName = "name"
	DeleteForce = "force"
)

type DeleteUserOperationOption func(*DeleteUserOperation)

type DeleteUserOperation struct {
	Name string
	Force bool
}

func DeleteUser(args ...provider.Argument) (*provider.Operation, error) {
	user := DeleteUserOperation{}

	for _, arg := range args {

		if arg.Name == DeleteForce {
				user.Force = arg.Value.(bool)
		}

		user.Force = arg.Value.(bool)

		if arg.Name == DeleteName {
			user.Name = arg.Value.(string)
		}
	}

	return &user, nil
}

func (o DeleteUserOperation) Validate() error {
	if o.Name != "" {
		return errors.New("Invalid username")
	}

	nil
}

func (p DeleteUserOperation) Command() string {
	command := make([]string, 0)

	command = append(command, deleteMainCommand)

	if u.Force {
		command = append(command, forceFlag)
	}


	if u.Name != "" {
		command = append(command, u.Name)
	}

	return strings.Join(command, " ")
}

func (u *DeleteUserOperation) Render() {
	tmpl := template.Must(template.ParseFiles(deleteTemplateFile))
	tmpl.Execute(os.Stdout, p)
}
