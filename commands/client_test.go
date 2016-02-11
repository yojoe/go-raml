package commands

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClientGeneration(t *testing.T) {
	Convey("test command client generattion", t, func() {
		Convey("Test run client command using go language", func() {
			cmd := ClientCommand{
				Language: "go",
				Dir:      "./test_client_command",
				RamlFile: "./fixtures/client_resources/client.raml",
			}
			err := cmd.Execute()
			So(err, ShouldBeNil)

			s, err := testLoadFile("./test_client_command/client_structapitest.go")
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/client_resources/client_structapitest.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)
		})

		Reset(func() {
			//cleanup
			os.RemoveAll("./test_client_command")
		})
	})
}
