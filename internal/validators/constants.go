package mistvalidator

import (
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type NineTypesValidator interface {
	validator.Bool
	validator.Float64
	validator.Int64
	validator.List
	validator.Map
	validator.Number
	validator.Object
	validator.Set
	validator.String
}
