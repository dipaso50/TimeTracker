package application

import (
	"fmt"
	"tracker/infrastructure/sqlite"
)

//Delete a project o a task
func Delete(projectID, taskID int) {

	if projectID > 0 {
		deleteProject(projectID)
	}

	if taskID > 0 {
		deleteTask(taskID)
	}
}

//DeleteProject delete given project
func deleteProject(projectID int) {
	prepo := sqlite.NewProjectRepo()

	if err := prepo.Delete(projectID); err != nil {
		fmt.Printf("Error %v \n", err)
		return
	}
	fmt.Println("Proyecto borrado correctamente")
}

//DeleteTask delete given task
func deleteTask(taskID int) {

	trepo := sqlite.NewTaskRepo()

	trepo.DeleteTask(taskID)
	fmt.Println("Tarea borrada correctamente")
}
