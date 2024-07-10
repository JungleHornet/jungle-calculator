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
	"os"
	"sort"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jcalc [variable name] [function]",
	Short: "Get info and do operations on a variable",
	Long: `Get info and do operations on a variable. All functions:
[no function] - View a variable and it's values
delete - Delete a variable 
` + "\033[1;32mLine specific:\033[0m" + `
                len/length - Measure the length of the line
` + "\033[1;32mAngle Specific:\033[0m" + `
                measure - Get the measure of the angle
` + "\033[1;32mTriangle Specific:\033[0m" + `
                orthocenter - Get the orthocenter of the triangle 
                circumcenter - Get the circumcenter of the triangle 
                centroid - Get the centroid of the triangle 
                incenter - Get the incenter of the triangle 
                orthocenter - Get the orthocenter of the triangle 
                parts - Get the info on each angle and side of the triangle.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				fmt.Println("\033[1;31mError displaying help for command \"jcalc\"\033[0m")
			}
			return
		}
		varfile, err := util.GetVarfile()
		if err != nil {
			fmt.Println(err)
			return
		}
		argLen := len(args)
		if util.GetVarRaw(args[0], varfile) != nil {
			varMap := util.GetVarRaw(args[0], varfile)
			varType := util.GetVarType(varMap, args[0])
			if argLen > 1 {
				switch varType {
				case "junglemath.Line":
					lineVar := junglemath.ToLine(util.GetVarOfType(args[0], varType, varfile), args[0])
					if args[1] == "len" || args[1] == "length" {
						fmt.Println("\033[1;32m", lineVar.Length(), "\033[0m")
						return
					}
				case "junglemath.Angle":
					angleVar := junglemath.ToAngle(util.GetVarOfType(args[0], varType, varfile), args[0])
					if args[0] == "measure" {
						fmt.Println("\033[1;32m", angleVar.Measure(), "\033[0m")
						return
					}
				case "junglemath.Triangle":
					triangleVar := junglemath.ToTriangle(util.GetVarOfType(args[0], varType, varfile), args[0])
					switch args[1] {
					case "orthocenter":
						fmt.Println("\033[1;32m", triangleVar.Orthocenter(), "\033[0m")
						return
					case "circumcenter":
						fmt.Println("\033[1;32m", triangleVar.Circumcenter(), "\033[0m")
						return
					case "centroid":
						fmt.Println("\033[1;32m", triangleVar.Centroid(), "\033[0m")
						return
					case "incenter":
						fmt.Println("\033[1;32m", triangleVar.Incenter(), "\033[0m")
						return
					case "parts":
						util.Parts(triangleVar, args[0])
						return
					}
				}
				fmt.Println("Error: Invalid variable type")
				return
			}
			fmt.Println("Variable " + args[0] + ":")
			fmt.Println("Type:", varType)
			delete(varMap, "type")
			var names []string
			for name := range varMap {
				names = append(names, name)
			}
			sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })
			for _, name := range names {
				fmt.Println(name+":", varMap[name])
			}
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jcalc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
