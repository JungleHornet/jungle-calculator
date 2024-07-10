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

// pythagCmd represents the pythag command
var pythagCmd = &cobra.Command{
	Use:   "pythag [leg 1 length] [leg 2 length]",
	Short: "Pythagorean theorem",
	Long:  `Uses the pythagorean theorem to find the hypotenuse of a right triangle using the length of the legs of the triangle.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pythag called")
	},
}

func init() {
	rootCmd.AddCommand(pythagCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pythagCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pythagCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
