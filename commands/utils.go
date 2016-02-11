package commands

import (
	"os"
	"os/exec"
	"strings"
	"text/template"

	log "github.com/Sirupsen/logrus"

	"github.com/Jumpscale/go-raml/raml"
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

func _completeResourceURI(r *raml.Resource, completeURI string) string {
	if r == nil {
		return completeURI
	}
	return _completeResourceURI(r.Parent, r.URI+completeURI)
}

// get complete URI of a resource
func completeResourceURI(r *raml.Resource) string {
	return _completeResourceURI(r, "")
}

func _getResourceParams(r *raml.Resource, params []string) []string {
	if r == nil {
		return params
	}
	if strings.HasPrefix(r.URI, "/{") && strings.HasSuffix(r.URI, "}") {
		params = append(params, r.URI[2:len(r.URI)-1])
	}
	return _getResourceParams(r.Parent, params)
}

// get all params of a resource
// examples:
// /users  							  : no params
// /users/{userId}					  : params 1 = userId
// /users/{userId}/address/{addressId : params 1= userId, param 2= addressId
func getResourceParams(r *raml.Resource) []string {
	params := []string{}
	return _getResourceParams(r, params)
}

func paramizingURI(URI string) string {
	// replace { with "+
	uri := strings.Replace(URI, "{", `"+`, -1)

	// if ended with }/" or }", remove trailing "
	if strings.HasSuffix(uri, `}/"`) || strings.HasSuffix(uri, `}"`) {
		uri = uri[:len(uri)-1]
	}

	// replace } with +"
	uri = strings.Replace(uri, "}", `+"`, -1)

	// clean trailing +"
	if strings.HasSuffix(uri, `+"`) {
		uri = uri[:len(uri)-2]
	}
	return uri
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

	var t *template.Template
	var err error

	if testMode {
		t, err = template.New(tmplName).Funcs(funcMap).ParseFiles(tmplFile)
	} else {
		tmplFile = strings.Replace(tmplFile, "./", "", -1)
		data, err := Asset(tmplFile)
		if err != nil {
			return err
		}

		//get string from byte
		t, err = template.New(tmplName).Funcs(funcMap).Parse(string(data))
	}

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

//templatingResourcePath build resourcePath for templating purpose
//Input : raw string, ex : /users/{userId}/address/{addressId}
//Output : "/users/"+userId+"/address/"+addressId
func templatingResourcePath(rawURI string) string {
	replaceLeftBracket := strings.Replace(rawURI, "{", "\"+", -1)
	finalURI := strings.Replace(replaceLeftBracket, "}", "+\"", -1)

	if strings.HasSuffix(finalURI, "+\"") {
		finalURI = "\"" + finalURI[:len(finalURI)-2]
	} else {
		finalURI = "\"" + finalURI + "\""
	}

	return finalURI
}

//convert interface type to string
//example :
//1. string type, result would be string
//2. []interface{} type, result would be array of string. ex: a,b,c
//Please add other type as needed
func interfaceToString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case []interface{}:
		interfaceArr := data.([]interface{})
		resultStr := ""
		for _, v := range interfaceArr {
			resultStr += interfaceToString(v) + ","
		}
		return resultStr[:len(resultStr)-1]
	default:
		return ""
	}
}
