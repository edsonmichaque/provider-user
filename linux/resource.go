package linux

import (
  "gitlab.com/ulombe/sdk"
  "gitlab.com/ulombe/sdk/attribute"
  "gitlab.com/ulombe/provider-user/linux/operation"
)

var DefaultResource = &Resource{
  Attributes: map[*Operation]AttributeTypeMap{
    operation.Create: AttributeTypeMap{
      "uid": attribute.OptionalString(),
      "gid": attribute.OptionalString(),
      "name": attribute.String(),
      "group": attribute.OptionalString(),
      "groups": attribute.OptionalArray(),
      "password": attribute.OptionalString(),
    },
    operation.Update: AttributeTypeMap{
      "uid": attribute.OptionalString(),
      "gid": attribute.OptionalString(),
      "name": attribute.String(),
      "newName": attribute.OptionalString(),
      "group": attribute.OptionalString(),
      "groups": attribute.OptionalArray(),
      "password": attribute.OptionalString(),
    },
    operation.Delete: AttributeTypeMap{
      "name": attribute.String(),
    },
  },
  Validators: NewValidators(),
}

DefaultResource.Validator(operation.Create, func(res *sdk.Resource, o *sdk.Operation, c *sdk.Change) error {
  if rules, ok := res.Attributes[o]; ok {
    for key, rule := range rules {
      if rule.Required() {
        if value, err := c.Get(key); err == nil {
          if rule.Kind() != reflect.ValueOf(value).Kind() {
            return errors.New(fmt.Sprintf("Kinds %s don't match %s", rule.Kind(), reflect.ValueOf(value).Kind()))
          }
        } else {
          return errors.New(fmt.Sprintf("%s is required", key))
        }
      } else {
        if value, err := c.Get(key); err == nil {
          if rule.Kind() != reflect.ValueOf(value).Kind() {
            return errors.New(fmt.Sprintf("Kinds %s don't match %s", rule.Kind(), reflect.ValueOf(value).Kind()))
          }
        }
      }
    }
  }

  return nil
})
