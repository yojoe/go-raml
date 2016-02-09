package commands

import (
	"os"
	"os/exec"

	"strings"
	"text/template"
)

func doNormalizeURI(URI string) string {
	normalizeSlash := strings.Replace(URI, "/", " ", -1)
	return normalizeBracket(normalizeSlash)
}

func normalizeURI(URI string) string {
	return strings.Replace(doNormalizeURI(URI), " ", "", -1)
}

func normalizeURITitle(URI string) string {
	titleString := strings.Title(doNormalizeURI(URI))
	return strings.Replace(titleString, " ", "", -1)

}

func normalizeBracket(URI string) string {
	normalizeLeftBracket := strings.Replace(URI, "{", "", -1)
	return strings.Replace(normalizeLeftBracket, "}", "", -1)
}

// generate Go file from a template.
// if file already exist and override==false, file won't be regenerated
// funcMap = this parameter is used for passing go function to the template
func generateFile(data interface{}, tmplFile, tmplName, filename string, override bool) error {
	if !override && isFileExist(filename) {
		return nil
	}

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	t, err := template.New(tmplName).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := t.ExecuteTemplate(f, tmplName, data); err != nil {
		return err
	}
	return runGoFmt(filename)
}

// create directory if not exist
func checkCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}

// cek if a file exist
func isFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsExist(err) {
		return true
	}

	return false
}

// run `go fmt` command to a file
func runGoFmt(filePath string) error {
	args := []string{"fmt", filePath}

	return exec.Command("go", args...).Run()
}
