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
	delTaskID int
	delProjID int
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Borra una tarea o un proyecto",
	Long:  `Permite borrar tareas o proyectos enteros`,
	Run: func(cmd *cobra.Command, args []string) {

		application.Delete(delProjID, delTaskID)

	},
}

func init() {

	deleteCmd.Flags().IntVarP(&delProjID, "projectID", "p", 0, "Borra el proyecto con este ID")
	deleteCmd.Flags().IntVarP(&delTaskID, "taskID", "t", 0, "Borra la tarea con este ID")

	rootCmd.AddCommand(deleteCmd)
}
