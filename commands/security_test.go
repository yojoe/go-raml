package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOauth2Middleware(t *testing.T) {
	Convey("oauth2 middleware", t, func() {

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("middleware generation test", func() {
			apiDef, err := raml.ParseFile("./fixtures/security/dropbox.raml")
			So(err, ShouldBeNil)

			err = generateSecurity(apiDef, targetdir, "main", langGo)
			So(err, ShouldBeNil)

			// oauth 2 facebook
			s, err := testLoadFile(filepath.Join(targetdir, "oauth2_Facebook_middleware.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/oauth2_Facebook_middleware.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// oauth 2 dropbox
			s, err = testLoadFile(filepath.Join(targetdir, "oauth2_Dropbox_middleware.go"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/security/oauth2_Dropbox_middleware.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Convey("Go routes generation", func() {
			apiDef, err := raml.ParseFile("./fixtures/security/dropbox.raml")
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef, targetdir, "main", langGo)
			So(err, ShouldBeNil)

			// check route
			s, err := testLoadFile(filepath.Join(targetdir, "deliveries_if.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/deliveries_if.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Convey("With included .raml file", func() {
			apiDef, err := raml.ParseFile("./fixtures/security/dropbox_with_include.raml")
			So(err, ShouldBeNil)

			err = generateSecurity(apiDef, targetdir, "main", langGo)
			So(err, ShouldBeNil)

			// oauth 2 middleware
			s, err := testLoadFile(filepath.Join(targetdir, "oauth2_DropboxIncluded_middleware.go"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/oauth2_DropboxIncluded_middleware.txt")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Convey("python middleware generation test", func() {
			apiDef, err := raml.ParseFile("./fixtures/security/dropbox.raml")
			So(err, ShouldBeNil)

			err = generateSecurity(apiDef, targetdir, "main", langPython)
			So(err, ShouldBeNil)

			// oauth 2 in dropbox
			s, err := testLoadFile(filepath.Join(targetdir, "oauth2_Dropbox.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/oauth2_Dropbox.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// oauth 2 facebook
			s, err = testLoadFile(filepath.Join(targetdir, "oauth2_Facebook.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/security/oauth2_Facebook.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

			// scope matching
			s, err = testLoadFile(filepath.Join(targetdir, "oauth2_Facebook_ADMINISTRATOR.py"))
			So(err, ShouldBeNil)

			tmpl, err = testLoadFile("./fixtures/security/oauth2_Facebook_ADMINISTRATOR.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)

		})

		Convey("Python routes generation", func() {
			apiDef, err := raml.ParseFile("./fixtures/security/dropbox.raml")
			So(err, ShouldBeNil)

			_, err = generateServerResources(apiDef, targetdir, "main", langPython)
			So(err, ShouldBeNil)

			// check route
			s, err := testLoadFile(filepath.Join(targetdir, "deliveries.py"))
			So(err, ShouldBeNil)

			tmpl, err := testLoadFile("./fixtures/security/deliveries.py")
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
