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
	"github.com/junglehornet/junglemath"

	"github.com/spf13/cobra"
)

// calcCmd represents the calc command
var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Opens the calculator",
	Long: `Opens the general calculator where you can solve equations with order of operations.
To Add: +
	Ex. 2+4 = 6
To Subtract: -
	Ex. 2-4 = -2
To Multiply: *
	Ex. 2*4 = 8
To Divide: /
	Ex. 2/4 = 0.5
To # Root: #r<number>
	Ex. 2r4 = 2
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opened calculator. Enter equation and enter to solve, or q to quit. For more info on how to use the calculator, run \"jcalc help calc\"")
		junglemath.OpenCalculator()
	},
}

func init() {
	rootCmd.AddCommand(calcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// calcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// calcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
