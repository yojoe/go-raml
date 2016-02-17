package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructFromRaml(t *testing.T) {
	Convey("generate struct from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct.raml")
		So(err, ShouldBeNil)
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = generateStructs(apiDef, targetdir, "main")
			So(err, ShouldBeNil)

			//first test
			s, err := testLoadFile(filepath.Join(targetdir, "city.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/struct/city.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			//second test
			s, err = testLoadFile(filepath.Join(targetdir, "animal.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/animal.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			//third test, single inheritance
			s, err = testLoadFile(filepath.Join(targetdir, "mammal.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/mammal.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			//fourth test, multiple inheritance
			s, err = testLoadFile(filepath.Join(targetdir, "anggora.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/anggora.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)

			//fifth test, array of object
			s, err = testLoadFile(filepath.Join(targetdir, "catcat.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/catcat.txt")
			So(err, ShouldBeNil)

			So(tmpl, ShouldEqual, s)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
