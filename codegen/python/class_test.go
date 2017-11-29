package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"
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
				"EnumCity.py",
				"animal.py",
				"Cage.py",
				"SingleInheritance.py",
				"PlainObject.py",
				"NumberFormat.py",
				"Cat.py",
				"MultipleInheritance.py",
				"EnumString.py",
				"petshop.py",
				"Catanimal.py",
				"UsersByIdGetRespBody.py",
				"UsersPostReqBody.py",
				"WithDateTime.py",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
				So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
