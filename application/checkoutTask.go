package application

import (
	"fmt"
	"tracker/infrastructure/sqlite"
)

//Checkout Cambia la tarea del contexto actual
func Checkout(taskName string) {

	if taskName == "" {
		return
	}

	hrepo := sqlite.NewHeadsRepo()
	trepo := sqlite.NewTaskRepo()

	t, err := trepo.GetTaskByName(taskName)

	if err != nil {
		fmt.Printf("Error %v", err)
		return
	}

	hrepo.SaveCurrentTask(*t)
	fmt.Println("Tarea cambiada correctamente")
}
