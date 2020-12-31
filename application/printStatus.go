package application

import (
	"fmt"
	"html/template"
	"os"
	"tracker/domain"
	"tracker/infrastructure/sqlite"
)

const statusTemplate = `
-------------------------------------------------
Proyecto Actual :(ID {{.PID}}) {{.CurrentProjectName}}
Tarea Actual    :(ID {{.TID}}) {{.CurrentTaskName}}     

Tiempo Registrado -------------------------------
Hoy          : {{.TimeToday | printf "%02d"}}
Esta semana  : {{.TimeThisWeek | printf "%02d"}}
Este mes     : {{.TimeThisMonth | printf "%02d"}}
Total        : {{.TimeTotal | printf "%02d"}}

`

type status struct {
	CurrentProjectName string
	CurrentTaskName    string
	PID                int
	TID                int
	TimeToday          int
	TimeThisWeek       int
	TimeThisMonth      int
	TimeTotal          int
}

//PrintStatus print different status information
func PrintStatus() {
	hrepo := sqlite.NewHeadsRepo()
	trepo := sqlite.NewTaskRepo()

	var pr *domain.Project
	var t *domain.Task
	var err error

	if t, pr, err = hrepo.GetCurrentTask(); err != nil {
		fmt.Printf("Error :%v\n", err)
		return
	}

	taskName := "No definida"
	tid := 0

	if t != nil {
		taskName = t.Name
		tid = int(t.ID)
	}

	projectName := "No definido"
	pid := 0

	if pr != nil {
		projectName = pr.Name
		pid = int(pr.ID)
	}

	today, week, month, total, _ := trepo.GetAggregates(tid)

	st := status{
		CurrentProjectName: projectName,
		CurrentTaskName:    taskName,
		PID:                pid,
		TID:                tid,
		TimeToday:          today,
		TimeThisWeek:       week,
		TimeThisMonth:      month,
		TimeTotal:          total,
	}

	report, err := template.New("report").Parse(statusTemplate)

	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	err = report.Execute(os.Stdout, st)

	if err != nil {
		fmt.Printf("Error %v", err)
	}
}
