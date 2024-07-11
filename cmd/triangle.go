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
	"github.com/junglehornet/jungle-calculator/util"
	"github.com/junglehornet/junglemath"

	"github.com/spf13/cobra"
)

// triangleCmd represents the triangle command
var triangleCmd = &cobra.Command{
	Use:   "triangle [name] [point1] [point2] [point3]",
	Short: "Sets a triangle variable",
	Long: `Sets a triangle variable.
Ex. jcalc set triangle t1 p1 p2 p3`,
	Run: func(cmd *cobra.Command, args []string) {
		varfile, err := util.GetVarfile()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(args) > 3 {
			varName := args[0]
			if util.GetVarRaw(args[1], varfile) != nil && util.GetVarRaw(args[1], varfile)["type"] == "junglemath.Point" {
				if util.GetVarRaw(args[2], varfile) != nil && util.GetVarRaw(args[1], varfile)["type"] == "junglemath.Point" {
					if util.GetVarRaw(args[3], varfile) != nil && util.GetVarRaw(args[1], varfile)["type"] == "junglemath.Point" {
						p1 := junglemath.ToPoint(util.GetVarOfType(args[1], "junglemath.Point", varfile), args[1])
						p2 := junglemath.ToPoint(util.GetVarOfType(args[2], "junglemath.Point", varfile), args[2])
						p3 := junglemath.ToPoint(util.GetVarOfType(args[3], "junglemath.Point", varfile), args[3])
						util.WriteVar(varName, junglemath.Triangle{A: p1, B: p2, C: p3}, varfile)
						fmt.Println("\033[1;32mSuccessfully set triangle " + varName + ".\033[0m")
					} else {
						fmt.Println("\033[1;31mError: Variable " + args[3] + " does not exist or is of wrong type.\033[0m")
					}
				} else {
					fmt.Println("\033[1;31mError: Variable " + args[2] + " does not exist or is of wrong type.\033[0m")
				}
			} else {
				fmt.Println("\033[1;31mError: Variable " + args[1] + " does not exist or is of wrong type.\033[0m")
			}
			return
		} else {
			fmt.Println("\033[1;31mError: Not enough arguments provided. Run \"jcalc help set triangle\" for usage.\033[0m")
		}
	},
}

func init() {
	setCmd.AddCommand(triangleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// triangleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// triangleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
