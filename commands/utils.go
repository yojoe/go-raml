package commands

import (
	"os"
	"strings"
	"text/template"
)

func normalizeURI(URI string) string {
	normalizeSlash := strings.Replace(URI, "/", " ", -1)
	normalizeLeftBracket := strings.Replace(normalizeSlash, "{", "", -1)
	normalizeRightBracket := strings.Replace(normalizeLeftBracket, "}", "", -1)
	return strings.Replace(normalizeRightBracket, " ", "", -1)
}

func generateFile(data interface{}, tmplFile, tmplName, filename string) error {
	t, err := template.ParseFiles(tmplFile)
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
	return runGoTools(cmdFormatCode, filename)
}

func checkCreateDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}
