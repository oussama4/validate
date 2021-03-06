package validators

import (
	"fmt"
	"time"

	"github.com/oussama4/validate/v4"
)

type TimeIsBeforeTime struct {
	FirstName  string
	FirstTime  time.Time
	SecondName string
	SecondTime time.Time
	Message    string
}

// IsValid adds an error if the FirstTime is after the SecondTime.
func (v *TimeIsBeforeTime) IsValid(errors *validate.Errors) {
	if v.FirstTime.UnixNano() <= v.SecondTime.UnixNano() {
		return
	}

	if len(v.Message) > 0 {
		errors.Add(GenerateKey(v.FirstName), v.Message)
		return
	}

	errors.Add(GenerateKey(v.FirstName), fmt.Sprintf("%s must be before %s.", v.FirstName, v.SecondName))
}
