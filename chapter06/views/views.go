package views

import "html/template"

// NewView Returns a new view
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/footer.html",
		"views/layouts/bootstrap.html",
		"views/layouts/navbar.html",
	)

	template, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)
	}

	return &View{
		Template: template,
		Layout:   layout,
	}
}

// View This will be the view we display
type View struct {
	Template *template.Template
	Layout   string
}
