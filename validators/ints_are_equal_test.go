package validators

import (
	"testing"

	"github.com/oussama4/validate/v4"
	"github.com/stretchr/testify/require"
)

func Test_IntsAreEqual(t *testing.T) {
	r := require.New(t)

	v := IntsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 1}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = IntsAreEqual{Name: "Number", ValueOne: 1, ValueTwo: 2}
	errors = validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("number"), []string{"1 is not equal to 2"})
}
