package commons

import (
	"errors"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/codegen/templates"
)

const (
	// ReqBodySuffix is suffix name for request body object
	ReqBodySuffix = "ReqBody"

	// RespBodySuffix is suffix name for response body object
	RespBodySuffix = "RespBody"
)

// doNormalizeURI removes `{`, `}`, and `/` from an URI
func doNormalizeURI(URI string) string {
	s := strings.Replace(URI, "/", " ", -1)
	s = strings.Replace(s, "{", "", -1)
	return strings.Replace(s, "}", "", -1)
}

// NormalizeURI removes `{`, `}`, `/`, and space from an URI
func NormalizeURI(URI string) string {
	return strings.Replace(doNormalizeURI(URI), " ", "", -1)
}

// NormalizeURITitle does NormalizeURI with first character in upper case
func NormalizeURITitle(URI string) string {
	s := strings.Title(doNormalizeURI(URI))
	return strings.Replace(s, " ", "", -1)

}

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

// InterfaceToString converts interface type to string.
//
// We can't simply do this using type casting.
// example :
// 1. string type, result would be string
// 2. []interface{} type, result would be array of string. ex: a,b,c
// Please add other type as needed.
func InterfaceToString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case []interface{}:
		interfaceArr := data.([]interface{})
		results := []string{}
		for _, v := range interfaceArr {
			results = append(results, InterfaceToString(v))
		}
		return strings.Join(results, ",")
	default:
		return ""
	}
}

// MapToSortedStrings returns sorted string arrays from a map
func MapToSortedStrings(m map[string]struct{}) []string {
	ss := []string{}
	for k := range m {
		ss = append(ss, k)
	}
	sort.Strings(ss)
	return ss
}

// cek if a file exist
func isFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

// ParamizingURI creates parameterized URI
// Input : raw string, ex : /users/{userId}/address/{addressId}
// Output : "/users/"+userId+"/address/"+addressId
func ParamizingURI(URI, sep string) string {
	uri := `"` + URI + `"`
	// replace { with "+
	uri = strings.Replace(uri, "{", `"`+sep, -1)

	// if ended with }/" or }", remove trailing "
	if strings.HasSuffix(uri, `}/"`) || strings.HasSuffix(uri, `}"`) {
		uri = uri[:len(uri)-1]
	}

	// replace } with +"
	uri = strings.Replace(uri, "}", sep+`"`, -1)

	// clean trailing +"
	if strings.HasSuffix(uri, sep+`"`) {
		uri = uri[:len(uri)-2]
	}
	return uri
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
