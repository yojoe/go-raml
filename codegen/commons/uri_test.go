package commons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamizingURI(t *testing.T) {
	type fixture struct {
		input  string
		output string
		sep    string
	}

	fixtures := []fixture{
		fixture{
			input:  `/users/{id}`, // simple uri
			output: `"/users/"+id`,
			sep:    "+",
		},
		fixture{
			input:  `/users/{id}/address`,
			output: `"/users/"+id+"/address"`,
			sep:    "+",
		},
		fixture{
			input:  `/users/{id}/address/`, // with trailing slash
			output: `"/users/"+id+"/address/"`,
			sep:    "+",
		},
		fixture{
			input:  `/users/{userId}/address/folder{addressId}test{addressId2}`,
			output: `"/users/"+userId+"/address/folder"+addressId+"test"+addressId2`,
			sep:    "+",
		},
		fixture{
			input:  `/users/{user-id}/address`,     // with non alphanumeric (dash)
			output: `"/users/"+user_id+"/address"`, // alphanumeric changed into underscore
			sep:    "+",
		},
	}

	for _, f := range fixtures {
		out := ParamizingURI(f.input, f.sep)
		assert.Equal(t, f.output, out)
	}
}
