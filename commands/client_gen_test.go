package commands

import (
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClientFromRaml(t *testing.T) {
	Convey("generate struct from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/client_resources/client.raml")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = GenerateClient(apiDef, "./test")
			So(err, ShouldBeNil)

			s, err := testLoadFile("./test/client_structapitest.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/client_resources/client_structapitest.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)
		})

		Reset(func() {
			cleanTestingDir()
		})
	})
}
