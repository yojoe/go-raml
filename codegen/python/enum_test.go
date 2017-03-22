package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEnum(t *testing.T) {
	Convey("Test Enum", t, func() {
		apiDef := new(raml.APIDefinition)
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("Enum from RAML", func() {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			So(err, ShouldBeNil)

			err = generateWtfClasses(apiDef.Types, targetDir)
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/wtf_class/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"EnumEnumCityEnum_homeNum.py", "EnumEnumCityEnum_homeNum.py"}, // enum of integer
				{"EnumEnumCityEnum_parks.py", "EnumEnumCityEnum_parks.py"},     // Enum of string
				{"EnumString.py", "EnumString.py"},
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
