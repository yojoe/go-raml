package commands

import (
	"os"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		Convey("simple server", func() {
			apiDef, err := raml.ParseFile("./fixtures/server/user_api/api.raml")
			So(err, ShouldBeNil)

			err = generateServer(apiDef, "./tmp")
			So(err, ShouldBeNil)

			// check users api implementation
			s, err := testLoadFile("./tmp/users_api.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server/user_api/users_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = testLoadFile("./tmp/users_if.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/users_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile("./tmp/main.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/main.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll("./tmp")
		})
	})
}
