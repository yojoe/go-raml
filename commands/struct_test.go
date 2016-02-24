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
		apiDef, err := raml.ParseFile("./fixtures/struct/struct.raml")
		So(err, ShouldBeNil)
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Simple struct from raml", func() {
			err = generateStructs(apiDef, targetdir, "main")
			So(err, ShouldBeNil)

			//first test
			s, err := testLoadFile(filepath.Join(targetdir, "enumcity.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/struct/enumcity.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			//second test
			s, err = testLoadFile(filepath.Join(targetdir, "animal.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/animal.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			//third test, single inheritance
			s, err = testLoadFile(filepath.Join(targetdir, "singleinheritance.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/singleinheritance.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			//fourth test, multiple inheritance
			s, err = testLoadFile(filepath.Join(targetdir, "multipleinheritance.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/multipleinheritance.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			//fifth test, array of object
			s, err = testLoadFile(filepath.Join(targetdir, "arrayofcats.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/arrayofcats.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// bidimensional array
			s, err = testLoadFile(filepath.Join(targetdir, "bidimensionalarrayofcats.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/bidimensionalarrayofcats.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// map type
			s, err = testLoadFile(filepath.Join(targetdir, "mapofcats.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/struct/mapofcats.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
