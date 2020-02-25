package validators

import (
	"testing"

	"github.com/oussama4/validate/v4"
	"github.com/stretchr/testify/require"
)

func Test_BytesArePresent(t *testing.T) {
	r := require.New(t)

	v := BytesArePresent{Name: "Name", Field: []byte("Mark")}
	errors := validate.NewErrors()
	v.IsValid(errors)
	r.Equal(errors.Count(), 0)

	v = BytesArePresent{Name: "Name", Field: []byte("")}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Name can not be blank."})

	errors = validate.NewErrors()
	v = BytesArePresent{Name: "Name", Field: []byte(""), Message: "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})

	errors = validate.NewErrors()
	v = BytesArePresent{"Name", []byte(""), "Field can't be blank."}
	v.IsValid(errors)
	r.Equal(errors.Count(), 1)
	r.Equal(errors.Get("name"), []string{"Field can't be blank."})
}
