package tarantool

import (
	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)
		Convey("Tarantool server", func() {
			apiDef := new(raml.APIDefinition)
			err = raml.ParseFile("../fixtures/server/user_api/api.raml", apiDef)
			So(err, ShouldBeNil)

			s := Server{
				apiDef:     apiDef,
				TargetDir:  targetDir,
				APIDocsDir: "apidocs",
			}
			s.generateResources()
			err := s.generateMain()
			So(err, ShouldBeNil)
			err = s.generateHandlers()
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server/"
			checks := []string{
				"main.lua",
				"handlers/helloworld_handler.lua",
				"handlers/users_handler.lua",
				"handlers/usersuserId_handler.lua",
				"handlers/usersuserIdaddressaddressId_handler.lua",
			}
			for _, check := range checks {
				s, err := utils.TestLoadFileRemoveID(filepath.Join(targetDir, check))
				So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFileRemoveID(filepath.Join(rootFixture, check))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
