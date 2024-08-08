package views

import (
    "embed"
    "html/template"
    "io/fs"
    "strings"
)

//go:embed *
var templatesFS embed.FS

var tmpl *template.Template

func init() {
    tmpl = template.New("")
    // Walk through the embedded file system.
    err := fs.WalkDir(templatesFS, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        if !d.IsDir() && strings.HasSuffix(d.Name(), ".html") {
            // Read and parse the file content into the template.
            // The name of the template will be the relative file path.
            content, err := fs.ReadFile(templatesFS, path)
            if err != nil {
                return err
            }
            _, err = tmpl.New(path).Parse(string(content))
            if err != nil {
                return err
            }
        }

        return nil
    })

    if err != nil {
        panic(err)
    }
}

func Get() *template.Template {
    return tmpl
}
