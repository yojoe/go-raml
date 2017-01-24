package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClassFromBody(t *testing.T) {
	Convey("Class from method body", t, func() {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("from RAML", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateClassesFromBodies(getAllResources(apiDef, true), targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"UsersPostReqBody.py", "UsersPostReqBody.py"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Convey("from RAML with JSON", func() {
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateClassesFromBodies(getAllResources(apiDef, true), targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/json/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"PersonPostReqBody.py", "PersonPostReqBody.py"},
			}

			for _, check := range checks {
				s, err := testLoadFile(filepath.Join(targetDir, check.Result))
				So(err, ShouldBeNil)

				tmpl, err := testLoadFile(filepath.Join(rootFixture, check.Expected))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
