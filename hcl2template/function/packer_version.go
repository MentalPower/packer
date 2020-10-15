package function

import (
	"github.com/hashicorp/packer/version"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

// PackerVersionFunc constructs a function that returns a string representation
// of the current Packer version
var PackerVersionFunc = function.New(&function.Spec{
	Params: []function.Parameter{},
	Type:   function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return cty.StringVal(version.FormattedVersion()), nil
	},
})

// PackerVersion returns a string representation of the current Packer version
func PackerVersion() (cty.Value, error) {
	return PackerVersionFunc.Call([]cty.Value{})
}
