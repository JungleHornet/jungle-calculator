/*
JungleCalculator - An open-source Go calculator for advanced math functions.
Copyright (c) 2023-present  JungleHornet

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// varsCmd represents the vars command
var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "Used for various operations involving variables",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vars called")
	},
}

func init() {
	rootCmd.AddCommand(varsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// varsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// varsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
