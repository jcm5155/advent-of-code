package common

import (
	"os"
	"text/tabwriter"
	"text/template"
)

func DisplaySolutions(year string, solutions []*Solution) {
	tmpl, _ := template.ParseFiles("common/templates/results.tpl")
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 4, ' ', 0)
	_ = tmpl.Execute(w, map[string]interface{}{
		"Year":      year,
		"Solutions": solutions,
	})
}
