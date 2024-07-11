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
	"strconv"

	"github.com/spf13/cobra"
)

// pointCmd represents the point command
var pointCmd = &cobra.Command{
	Use:   "point [name] [x value] [y value] [z value]",
	Short: "Sets a point variable",
	Long: `Sets a point variable.
Note: x, y, and z values are optional and will default to 0 if not specified.
Ex. jcalc set point p1 1 54`,
	Run: func(cmd *cobra.Command, args []string) {
		varfile, err := util.GetVarfile()
		if err != nil {
			fmt.Println(err)
			return
		}
		if len(args) > 3 {
			x, _ := strconv.ParseFloat(args[1], 64)
			y, _ := strconv.ParseFloat(args[2], 64)
			z, _ := strconv.ParseFloat(args[3], 64)
			util.WriteVar(args[0], junglemath.Point{X: x, Y: y, Z: z}, varfile)
			fmt.Println("\033[1;32mSuccessfully set point " + args[0] + ".\033[0m")
			return
		} else if len(args) > 2 {
			x, _ := strconv.ParseFloat(args[1], 64)
			y, _ := strconv.ParseFloat(args[2], 64)
			util.WriteVar(args[0], junglemath.Point{X: x, Y: y, Z: 0}, varfile)
			fmt.Println("\033[1;32mSuccessfully set point " + args[0] + ".\033[0m")
			return
		} else {
			fmt.Println("\033[1;31mError: Not enough arguments provided. Run \"jcalc help set point\" for usage.\033[0m")
		}
		return
	},
}

func init() {
	setCmd.AddCommand(pointCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pointCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pointCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
