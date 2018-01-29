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

func TestServerMethodWithSpecialChars(t *testing.T) {
	Convey("TestServerMethodWithSpecialChars", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		So(err, ShouldBeNil)

		gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, targetDir, nil)
		_, err = gs.generateServerResources(targetDir)
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/server"
		files := []string{
			filepath.Join(serverAPIDir, "escape_type", "escape_type_api_Post"),
			filepath.Join(serverAPIDir, "uri", "uri_api_Users_idGet"),
			"uri_if",
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
