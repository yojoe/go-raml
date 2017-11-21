package mypy

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"fmt"
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
			err := raml.ParseFile("./fixtures/types.raml", apiDef)
			So(err, ShouldBeNil)

			err = GenerateMyPy(apiDef, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/"
			files := []string{
				"EnumCity.%s",
				"Animal.%s",
				"Cage.%s",
				"SingleInheritance.%s",
				"PlainObject.%s",
				"NumberFormat.%s",
				"Cat.%s",
				"MultipleInheritance.%s",
				"EnumString.%s",
				"Petshop.%s",
				"WithDateTime.%s",
			}

			for _, f := range files {
				// check the python classes
				class := fmt.Sprintf(f, "py")
				s, err := utils.TestLoadFile(filepath.Join(targetDir, class))
				So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, "class", class))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)

				// check the capnp schemas
				schema := fmt.Sprintf(f, "capnp")
				s, err = utils.TestLoadFileRemoveID(filepath.Join(targetDir, schema))
				So(err, ShouldBeNil)

				tmpl, err = utils.TestLoadFileRemoveID(filepath.Join(rootFixture, "capnp", schema))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			return
			os.RemoveAll(targetDir)
		})

	})
}
