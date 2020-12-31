package application

import (
	"fmt"
	"tracker/domain"
	"tracker/infrastructure/sqlite"
)

//CreateOrUpdateProject creates or updates a project and his tasks
func CreateOrUpdateProject(projectName, taskName, projectDescription, taskDescription string) {

	prepo := sqlite.NewProjectRepo()
	hrepo := sqlite.NewHeadsRepo()

	var pr *domain.Project
	var err error

	if pr, err = getProject(projectName, projectDescription, hrepo); err != nil {
		fmt.Printf("Error %v\n", err)
	}

	if taskName != "" {
		fmt.Println("With task " + taskName)

		pr.Tasks = append(pr.Tasks, domain.Task{
			Name: taskName,
		})
	}

	pro, _ := prepo.GetProjectByName(projectName)

	if pro == nil {
		var p domain.Project
		if p, err = prepo.AddProject(*pr); err != nil {
			fmt.Printf("Error %v", err)
			return
		}
		fmt.Printf("Proyecto %s añadido correctamente", projectName)
		pro = &p
	} else {
		pr.ID = pro.ID
		if _, err = prepo.UpdateProject(*pr); err != nil {
			fmt.Printf("Error %v", err)
			return
		}
	}

	hrepo.SaveCurrentProject(*pro)

	if len(pro.Tasks) > 0 {
		hrepo.SaveCurrentTask(pro.Tasks[0])
	}
}

func getProject(projectName, projectDescription string, hrepo domain.HeadsRepo) (*domain.Project, error) {

	if projectName != "" {
		pr := domain.Project{
			Name:        projectName,
			Description: projectDescription,
		}

		return &pr, nil
	}

	var pr *domain.Project
	var err error

	if _, pr, err = hrepo.GetCurrentTask(); err != nil {
		return nil, err
	}

	if projectDescription != "" {
		pr.Description = projectDescription
	}

	return pr, nil
}
