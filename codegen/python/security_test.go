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

func TestOauth2Middleware(t *testing.T) {
	Convey("oauth2 middleware", t, func() {

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("middleware generation test", func() {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox.raml", apiDef)
			So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "", true, nil)
			err = fs.generateOauth2(fs.APIDef.SecuritySchemes, targetdir)
			So(err, ShouldBeNil)

			// oauth 2 in dropbox
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Dropbox.py"))
			So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("./fixtures/security/oauth2_Dropbox.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// oauth 2 facebook
			s, err = utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Facebook.py"))
			So(err, ShouldBeNil)

			tmpl, err = utils.TestLoadFile("./fixtures/security/oauth2_Facebook.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Convey("routes generation", func() {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox.raml", apiDef)
			So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", true, nil)
			err = fs.generateResources(targetdir)
			So(err, ShouldBeNil)

			// check route
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "deliveries_api.py"))
			So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("./fixtures/security/deliveries_api.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
