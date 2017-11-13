package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratePythonClientFromRaml(t *testing.T) {
	Convey("Python client", t, func() {
		apiDef := new(raml.APIDefinition)
		err := raml.ParseFile("./fixtures/client/client.raml", apiDef)
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("requests client", func() {
			client := NewClient(apiDef, "", false)
			err = client.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/client/requests_client"
			// cek with generated with fixtures
			checks := []struct {
				Result   string
				Expected string
			}{
				{"client.py", "client.py"},
				{"__init__.py", "__init__.py"},
				{"client_utils.py", "client_utils.py"},
				{"users_service.py", "users_service.py"},
				{"Address.py", "Address.py"},
				{"City.py", "City.py"},
				{"GetUsersReqBody.py", "GetUsersReqBody.py"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("requests client with unmarshall response", func() {
			client := NewClient(apiDef, "", true)
			err = client.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/client/requests_client/unmarshall_response"
			// cek with generated with fixtures
			files := []string{
				"__init__.py",
				"users_service.py",
				"unmarshall_error.py",
				"api_response.py",
			}

			for _, file := range files {
				s, err := testLoadFile(filepath.Join(targetDir, file))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, file))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("aiohttp client", func() {
			client := NewClient(apiDef, clientNameAiohttp, false)
			err = client.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/client/aiohttp_client"
			// cek with generated with fixtures
			files := []string{
				"client.py",
				"__init__.py",
				"client_utils.py",
				"users_service.py",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Convey("aiohttp client with unmarshall response", func() {
			client := NewClient(apiDef, clientNameAiohttp, true)
			err = client.Generate(targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/client/aiohttp_client/unmarshall_response"
			// cek with generated with fixtures
			files := []string{
				"__init__.py",
				"users_service.py",
				"unmarshall_error.py",
				"api_response.py",
			}

			for _, f := range files {
				s, err := testLoadFile(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}

func testLoadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	return string(b), err
}
