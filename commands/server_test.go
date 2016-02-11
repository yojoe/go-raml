package commands

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServerGeneration(t *testing.T) {
	Convey("test command server generattion", t, func() {
		Convey("Test run server command using go language", func() {
			cmd := ServerCommand{
				Dir:      "./test_server_command",
				RamlFile: "./fixtures/server/user_api/api.raml",
			}
			err := cmd.Execute()
			So(err, ShouldBeNil)

			// check users api implementation
			s, err := testLoadFile("./test_server_command/users_api.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/server/user_api/users_api.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = testLoadFile("./test_server_command/users_if.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/users_if.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)

			// check main file
			s, err = testLoadFile("./test_server_command/main.go")
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/server/user_api/main.txt")
			So(err, ShouldBeNil)
			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			//cleanup
			os.RemoveAll("./test_server_command")
		})
	})
}
