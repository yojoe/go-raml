package raml

import (
	"strings"

	"bitbucket.org/pkg/inflect"
	chuck_inflect "github.com/chuckpreslar/inflect"
)

// singularize returns singular version of a word
func singularize(s string) string {
	return inflect.Singularize(s)
}

// pluralize returns plural version of a word
func pluralize(s string) string {
	return chuck_inflect.Pluralize(s)
}

// upperCase returns upper case version of a word
func upperCase(s string) string {
	return strings.ToUpper(s)
}

// lowerCase returns lower case version of a word
func lowerCase(s string) string {
	return strings.ToLower(s)
}

// lowerCamelCase returns camel case version of a word
// with lower case first character
func lowerCamelCase(s string) string {
	return inflect.CamelizeDownFirst(s)
}

// upperCamelCase returns camel case version of a word
// with upper case first character
func upperCamelCase(s string) string {
	return inflect.Camelize(s)
}

// lowerUnderScoreCase returns lower & underscore case version of a word
func lowerUnderScoreCase(s string) string {
	return strings.ToLower(inflect.Underscore(s))
}

// lowerUnderScoreCase returns upper & underscore case version of a word
// with all characters changed to upper case
func upperUnderScoreCase(s string) string {
	return strings.ToUpper(inflect.Underscore(s))
}

// lowerHyphenCase returns lower case hyphenated version of a word
func lowerHyphenCase(s string) string {
	return strings.ToLower(chuck_inflect.Hyphenate(s))
}

// upperHyphenCase returns upper case hyphenated version of a word
func upperHyphenCase(s string) string {
	return strings.ToUpper(chuck_inflect.Hyphenate(s))
}
