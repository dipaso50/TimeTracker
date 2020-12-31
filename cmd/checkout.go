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
	tcheckName string
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Cambia el contexto a una tarea determinada",
	Long:  `Cambia el contexto (la tarea actual) a otra tarea ya existente para poder registrar o consultar tiempos`,
	Run: func(cmd *cobra.Command, args []string) {

		application.Checkout(tcheckName)
	},
}

func init() {
	checkoutCmd.Flags().StringVarP(&tcheckName, "taskName", "n", "", "Nombre de la tarea")
	checkoutCmd.MarkFlagRequired("taskName")

	rootCmd.AddCommand(checkoutCmd)
}
