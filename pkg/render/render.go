package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mariogarzac/go-learngo/pkg/config"
	"github.com/mariogarzac/go-learngo/pkg/models"
)

var functions = template.FuncMap {
}

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
    app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
    return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

    var tc map[string]*template.Template
    if app.UseCache {
        tc = app.TemplateCache
    }else{
        tc,_ = CreateTemplateCache()
    }


    // get requested template from cache
    t, ok := tc[tmpl]

    if !ok {
        log.Fatal("Could not get template from template cache")
    }

    td = AddDefaultData(td)

    buf := new(bytes.Buffer)

    _ = t.Execute(buf,nil)

    _, err := buf.WriteTo(w)
    if err != nil {
        log.Println("Error writing template to browser" ,err)
    }
}

func CreateTemplateCache() (map[string]*template.Template, error){
    cache := map[string]*template.Template{}

    //get all html files
    pages, err := filepath.Glob("./templates/*.page.html")
    if err != nil {
        return cache, err
    }

    // range through all files
    for _, page := range pages {
        filename := filepath.Base(page)
        ts, err := template.New(filename).ParseFiles(page)
        if err != nil {
            log.Println(err)
        }
        matches, err := filepath.Glob("./templates/*.layout.html")

        if err != nil {
            return cache, err
        }

        if len(matches) > 0 {
            ts, err = ts.ParseGlob("./templates/*layout.html")
        }
        cache[filename] = ts
    }

    return cache, nil
}
