package commands

import (
	"os"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	bodyPropertyMap = map[string]raml.Property{
		"ID": raml.Property{
			Type:     "string",
			Required: true,
		},
		"age": raml.Property{
			Type:     "integer",
			Required: false,
		},
	}
)

func TestGenerateStructBodyFromRaml(t *testing.T) {
	Convey("generate struct body from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct.raml")
		So(err, ShouldBeNil)

		dir, err := os.Getwd()
		So(err, ShouldBeNil)

		Convey("simple body", func() {
			err := GenerateBodyStruct(apiDef, dir+"/test")
			So(err, ShouldBeNil)

			//load and compare TestusersIdGetReqBody
			s, err := testLoadFile("./test/UsersIdGetReqBody.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/UsersIdGetReqBody.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			//load and compare TestusersIdGetRespBody
			s2, err := testLoadFile("./test/UsersIdGetRespBody.go")
			So(err, ShouldBeNil)

			tmpl2, err := testLoadFile("./fixtures/UsersIdGetRespBody.txt")
			So(err, ShouldBeNil)

			So(tmpl2, ShouldEqual, s2)

			//load and compare TestusersGetReqBody
			tUserGetReqBody, err := testLoadFile("./test/UsersGetReqBody.go")
			So(err, ShouldBeNil)

			fUserGetReqBody, err := testLoadFile("./fixtures/UsersGetReqBody.txt")
			So(err, ShouldBeNil)

			So(tUserGetReqBody, ShouldEqual, fUserGetReqBody)

			//load and compare TestusersGetRespBody
			tUserGetRespBody, err := testLoadFile("./test/UsersGetRespBody.go")
			So(err, ShouldBeNil)

			fUserGetRespBody, err := testLoadFile("./fixtures/UsersGetRespBody.txt")
			So(err, ShouldBeNil)

			So(tUserGetRespBody, ShouldEqual, fUserGetRespBody)
		})

		Reset(func() {
			cleanTestingDir()
		})
	})
}
