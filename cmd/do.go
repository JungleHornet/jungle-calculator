/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/junglehornet/jungle-calculator/util"
	"github.com/junglehornet/junglemath"
	"sort"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
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
			fmt.Println("\033[1;31mError: No arguments given. Use \"jcalc help do\" for usage\033[0m")
			return
		}
		varfile, err := util.GetVarfile()
		if err != nil {
			fmt.Println(err)
			return
		}
		argLen := len(args)
		if argLen > 1 {
			if args[1] == "delete" {
				util.RmVar(args[0], varfile)
				fmt.Println("\033[1;32mSuccessfully deleted " + args[0] + ".\033[0m")
				return
			}
		}
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
		} else {
			fmt.Println("\033[1;31mWarning: Variable " + args[0] + " does not exist.\033[0m")
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
