package tarantool

import (
	"github.com/Jumpscale/go-raml/codegen"
	"github.com/Jumpscale/go-raml/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServer(t *testing.T) {
	Convey("server generator", t, func() {
		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)
		Convey("Tarantool server", func() {
			s := codegen.Server{
				RAMLFile:   "../fixtures/server/user_api/api.raml",
				Kind:       "",
				Dir:        targetdir,
				Lang:       "tarantool",
				APIDocsDir: "apidocs",
			}
			err := s.Generate()
			So(err, ShouldBeNil)

			rootFixture := "./fixtures/server/"
			checks := []string{
				"main.lua",
				"handlers/helloworld_handler.lua",
				"handlers/users_handler.lua",
				"handlers/usersuserId_handler.lua",
				"handlers/usersuserIdaddressaddressId_handler.lua",
				"schemas/Address.capnp",
				"schemas/EnumUserNames.capnp",
				"schemas/User.capnp",
				//"schemas/schema.lua",
			}
			for _, check := range checks {
				s, err := utils.TestLoadFileRemoveID(filepath.Join(targetdir, check))
				So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFileRemoveID(filepath.Join(rootFixture, check))
				So(err, ShouldBeNil)

				So(s, ShouldEqual, tmpl)
			}

		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
