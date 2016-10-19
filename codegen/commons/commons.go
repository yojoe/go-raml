package commons

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/templates"
)

// ParseDescription create string slice from an RAML description.
// each element is a  description line
func ParseDescription(desc string) []string {
	// we need to trim it because our parser usually give
	// space after last newline
	desc = strings.TrimSpace(desc)

	if desc == "" {
		return []string{}
	}

	return strings.Split(desc, "\n")
}

// GenerateFile generates file from a template.
// if file already exist and override==false, file won't be regenerated
// funcMap = this parameter is used for passing go function to the template
func GenerateFile(data interface{}, tmplFile, tmplName, filename string, override bool) error {
	if !override && isFileExist(filename) {
		log.Infof("file %v already exist and override=false, no need to regenerate", filename)
		return nil
	}

	// pass Go function to template
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	// all template files path is relative to current directory (./)
	// while go-bindata files exist in templates directory
	tmplFile = strings.Replace(tmplFile, "./", "", -1)

	byteData, err := templates.Asset(tmplFile)
	if err != nil {
		return err
	}

	t, err := template.New(tmplName).Funcs(funcMap).Parse(string(byteData))
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Infof("generating file %v", filename)
	if err := t.ExecuteTemplate(f, tmplName, data); err != nil {
		return err
	}

	if strings.HasSuffix(filename, ".go") {
		return runGoFmt(filename)
	}
	return nil
}

// cek if a file exist
func isFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// run `go fmt` command to a file
func runGoFmt(filePath string) error {
	args := []string{"fmt", filePath}

	if out, err := exec.Command("go", args...).CombinedOutput(); err != nil {
		log.Errorf("Error running go fmt on '%s' failed:\n%s", filePath, string(out))
		return errors.New("go fmt failed")
	}
	return nil
}
