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
	defaultMinutes = 25
	minutes        int
	taskID         int
	tName          string
)

// trackCmd represents the track command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Inicia un pomodoro y registra el tiempo en una tarea determinada",
	Long:  `Inicia un pomodoro y registra el tiempo en una tarea determinada`,
	Run: func(cmd *cobra.Command, args []string) {

		application.Track(minutes, tName, taskID)
	},
}

func init() {
	trackCmd.Flags().IntVarP(&minutes, "minutes", "m", defaultMinutes, "Número de minutos del pomodoro")
	trackCmd.Flags().IntVarP(&taskID, "taskID", "i", 0, "Id de la tarea (default current task)")
	trackCmd.Flags().StringVarP(&tName, "taskName", "n", "", "Nombre de la tarea (default current task)")
	rootCmd.AddCommand(trackCmd)
}
