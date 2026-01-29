package template

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	shopTemplates  *template.Template
	adminTemplates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if tmpl := t.shopTemplates.Lookup(name); tmpl != nil {
		return tmpl.Execute(w, data)
	}
	if tmpl := t.adminTemplates.Lookup(name); tmpl != nil {
		return tmpl.Execute(w, data)
	}
	return fmt.Errorf("template %s not found", name)
}

func RenderTemplates() *Template {
	t := &Template{
		shopTemplates:  template.Must(template.ParseGlob("../template/shop/*.html")),
		adminTemplates: template.Must(template.ParseGlob("../template/admin/*.html")),
	}

	fmt.Println("Shop templates:")
	for _, tmpl := range t.shopTemplates.Templates() {
		fmt.Printf("  - %s\n", tmpl.Name())
	}

	fmt.Println("Admin templates:")
	for _, tmpl := range t.adminTemplates.Templates() {
		fmt.Printf("  - %s\n", tmpl.Name())
	}

	return t
}

func (t *Template) HasTemplate(name string, isAdmin bool) bool {
	if isAdmin {
		return t.adminTemplates.Lookup(name) != nil
	}
	return t.shopTemplates.Lookup(name) != nil
}
