package golang

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClientMethodWithSpecialChars(t *testing.T) {
	Convey("TestClientMethodWithSpecialChars", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		So(err, ShouldBeNil)

		client, err := NewClient(apiDef, "theclient", "examples.com/libro", targetDir, nil)
		So(err, ShouldBeNil)

		err = client.Generate()
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/client"
		files := []string{
			"escape_type_service",
			"uri_service",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f+".go"))
			So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		}

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
