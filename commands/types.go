package commands

import (
	"time"
)

//Date represent RFC3399 date
type Date time.Time

//MarshalJSON override marshalJSON
func (t *Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*t).Format(`"` + time.RFC3339 + `"`)), nil
}

//MarshalText override marshalText
func (t *Date) MarshalText() ([]byte, error) {
	return []byte(time.Time(*t).Format(`"` + time.RFC3339 + `"`)), nil
}

//UnmarshalJSON override unmarshalJSON
func (t *Date) UnmarshalJSON(b []byte) error {
	ts, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = Date(ts)
	return nil
}

//UnmarshalText override unmarshalText
func (t *Date) UnmarshalText(b []byte) error {
	ts, err := time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return err
	}

	*t = Date(ts)
	return nil
}

func (t *Date) String() string {
	return time.Time(*t).String()
}

//ConvertToGoType handle convert from raml to go type
func convertToGoType(source string) string {
	var result string
	switch source {
	case "string":
		result = "string"
	case "number":
		result = "float"
	case "integer":
		result = "int"
	case "boolean":
		result = "bool"
	case "date":
		result = "Date"
	case "enum":
		result = "[]string"
	case "file":
		result = "string"
	default:
		result = source
	}
	return result
}
