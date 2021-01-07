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
Hoy          : {{.TimeToday}}
Esta semana  : {{.TimeThisWeek}}
Este mes     : {{.TimeThisMonth}}
Total        : {{.TimeTotal}}

`

type status struct {
	CurrentProjectName string
	CurrentTaskName    string
	PID                int
	TID                int
	TimeToday          string
	TimeThisWeek       string
	TimeThisMonth      string
	TimeTotal          string
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
		TimeToday:          minHours(today),
		TimeThisWeek:       minHours(week),
		TimeThisMonth:      minHours(month),
		TimeTotal:          minHours(total),
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

func minHours(allMin int) string {
	hh := int(allMin / 60)
	min := allMin % 60

	return fmt.Sprintf("%02dh:%02dm (%03d min)", hh, min, allMin)
}
