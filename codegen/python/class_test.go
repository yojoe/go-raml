package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratePythonClass(t *testing.T) {
	Convey("generate python class from raml", t, func() {
		apiDef := new(raml.APIDefinition)
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("python class from raml Types", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			globAPIDef = apiDef

			_, err = generateAllClasses(apiDef, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/"
			files := []string{
				"Cage.py",
				"EnumCity.py",
				"EnumString.py",
				"SingleInheritance.py",
				"MultipleInheritance.py",
				"animal.py",
				"petshop.py",
				"Catanimal.py",
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
