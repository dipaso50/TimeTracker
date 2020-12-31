/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"tracker/application"

	"github.com/spf13/cobra"
)

var (
	projectName        string
	taskName           string
	projectDescription string
	taskDescription    string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Añade Proyectos o tareas a proyectos",
	Long: `Comando para añadir proyectos o tareas a proyectos. 

Por ejemplo:

	//añade un nuevo proyecto
	$ tracker add -p "proyectName" -d "Descripción del proyecto"
	
	//añade una nueva tarea y descripción de la misma, se añade al proyecto donde se haya hecho la última modificación.
	$ tracker add -t "taskName" -d "Descripción de la tarea"      

	//añade una nueva tarea a un proyecto ya existente, si el proyecto no existe lo crea.
	$ tracker add -t "taskName" -p "proyectName"   

	 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if projectName == "" && taskName == "" {
			cmd.Help()
			return
		}

		process()
	},
}

func process() {
	application.CreateOrUpdateProject(projectName, taskName, projectDescription, taskDescription)
}

func init() {

	addCmd.Flags().StringVarP(&projectName, "projectName", "p", "", "Nombre del proyecto.")
	addCmd.Flags().StringVarP(&taskName, "taskName", "t", "", "Nombre de la tarea.")
	addCmd.Flags().StringVarP(&taskDescription, "taskDescription", "d", "", "Breve descripción de la tarea")
	addCmd.Flags().StringVarP(&projectDescription, "projectDescription", "D", "", "Breve descripción del proyecto")

	rootCmd.AddCommand(addCmd)
}
