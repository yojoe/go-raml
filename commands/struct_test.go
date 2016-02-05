package commands

import (
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	structTestData = map[string]raml.Type{
		"User": raml.Type{
			Properties: map[string]raml.Property{
				"ID": raml.Property{
					Type:     "string",
					Required: true,
				},
				"age": raml.Property{
					Type:     "integer",
					Required: false,
				},
				"active": raml.Property{
					Type:     "boolean",
					Required: true,
				},
				"city": raml.Property{
					Type:     "City",
					Required: false,
				},
			},
		},
		"City": raml.Type{
			Properties: map[string]raml.Property{
				"names": raml.Property{
					Type:     "enum",
					Required: false,
				},
			},
		},
	}
)

func setupGenerateStruct() *raml.APIDefinition {
	apiDef := new(raml.APIDefinition)
	apiDef.Types = structTestData
	return apiDef
}

func TestGenerateStruct(t *testing.T) {
	Convey("generate struct from body", t, func() {
		apiDef := setupGenerateStruct()

		Convey("Simple struct", func() {
			err := GenerateStruct("./test", apiDef)
			So(err, ShouldBeNil)

			s, err := testLoadFile("./test/City.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/City.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			Reset(func() {
				cleanTestingDir()
			})
		})
	})
}

func TestGenerateStructFromRaml(t *testing.T) {
	Convey("generate struct from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct.raml")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = GenerateStruct("./test", apiDef)
			So(err, ShouldBeNil)

			s, err := testLoadFile("./test/City.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/City.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)
		})

		Reset(func() {
			cleanTestingDir()
		})
	})
}
