/*
Copyright © 2023 Helge Falch <helge.falch@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Prints CPU usage",
	Long:  `Prints CPU usage visually in a bar graph.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cpu called")
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
