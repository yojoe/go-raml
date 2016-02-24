package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructBodyFromRaml(t *testing.T) {
	Convey("generate struct body from raml", t, func() {
		apiDef, err := raml.ParseFile("./fixtures/struct/struct.raml")
		So(err, ShouldBeNil)

		targetdir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("simple body", func() {
			err := generateBodyStructs(apiDef, targetdir, "main")
			So(err, ShouldBeNil)

			//load and compare UsersIdGetRespBody
			s2, err := testLoadFile(filepath.Join(targetdir, "usersidgetrespbody.go"))
			So(err, ShouldBeNil)

			tmpl2, err := testLoadFile("./fixtures/usersidgetrespbody.txt")
			So(err, ShouldBeNil)

			So(tmpl2, ShouldEqual, s2)

			//load and compare usersgetreqbody
			tUserGetReqBody, err := testLoadFile(filepath.Join(targetdir, "usersgetreqbody.go"))
			So(err, ShouldBeNil)

			fUserGetReqBody, err := testLoadFile("./fixtures/usersgetreqbody.txt")
			So(err, ShouldBeNil)

			So(tUserGetReqBody, ShouldEqual, fUserGetReqBody)
		})

		Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
