package codegen

import (
	"io/ioutil"
	//"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)
		Convey("simple Go server", func() {
			err := GenerateServer("./fixtures/server/user_api/api.raml", targetdir, "main", "go", "apidocs", "examples.com/ramlcode", true)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server/user_api/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"main.go", "main.txt"},
				{"users_if.go", "users_if.txt"},
				{"users_api.go", "users_api.txt"},
				{"helloworld_if.go", "helloworld_if.txt"},
				{"helloworld_api.go", "helloworld_api.txt"},
				// goraml package
				{"goraml/datetime.go", "goraml/datetime.txt"},
				{"goraml/date_only.go", "goraml/date_only.txt"},
				{"goraml/datetime_only.go", "goraml/datetime_only.txt"},
				{"goraml/datetime_rfc2616.go", "goraml/datetime_rfc2616.txt"},
				{"goraml/time_only.go", "goraml/time_only.txt"},
				{"goraml/struct_input_validator.go", "goraml/struct_input_validator.txt"},
			}
			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetdir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("Congo python server", func() {
			err := GenerateServer("./fixtures/congo/api.raml", targetdir, "main", "python", "apidocs", "", true)
			So(err, ShouldBeNil)

			// check drones API implementation
			s, err := testLoadFile(filepath.Join(targetdir, "drones.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/congo/python_server/drones.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check deliveries API implementation
			s, err = testLoadFile(filepath.Join(targetdir, "deliveries.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/congo/python_server/deliveries.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile(filepath.Join(targetdir, "app.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/congo/python_server/app.py")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			//os.RemoveAll(targetdir)
		})
	})
}
