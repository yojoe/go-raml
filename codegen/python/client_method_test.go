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

func TestClientMethodWithComplexBody(t *testing.T) {
	Convey("TestClientMethodWithComplexBody", t, func() {
		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/body.raml", apiDef)
		So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameRequests, true)

		err = client.Generate(targetDir)
		So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/client/complex_body/requests_unmarshall"
		files := []string{
			"arrays_service.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			So(err, ShouldBeNil)

			So(s, ShouldEqual, tmpl)
		}

		Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
