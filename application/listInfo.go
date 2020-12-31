package application

import (
	"fmt"
	"html/template"
	"os"
	"tracker/domain"
	"tracker/infrastructure/sqlite"
)

const listTemplate = `{{$root := .}}{{range .All}}{{if $root.MatchProjectName .Name}}-------------------------------------------------
Proyecto(ID {{.ID}}, {{len .Tasks}} tareas): {{.Name}}
{{range .Tasks}}Tarea(ID {{.ID}}): {{.Name}}
{{end}}{{end}}
{{end}}
`

type lstReport struct {
	All           []domain.Project
	FilterProject string
	FilterTask    string
}

func (lre lstReport) MatchProjectName(prName string) bool {
	if lre.FilterProject == "" {
		return true
	}

	return lre.FilterProject == prName
}

//ListAll registered information
func ListAll(projectName string) {
	prepo := sqlite.NewProjectRepo()

	var all []domain.Project
	var err error

	if all, err = prepo.GetAll(); err != nil {
		fmt.Printf("Error : %v", err)
	}

	lstRep := lstReport{
		All:           all,
		FilterProject: projectName,
	}

	report, err := template.New("reportList").Parse(listTemplate)

	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	err = report.Execute(os.Stdout, lstRep)

	if err != nil {
		fmt.Printf("Error %v", err)
	}
}
