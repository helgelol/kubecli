/*
Copyright © 2023 Helge Falch <helge.falch@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

// memCmd represents the mem command
var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "Prints memory usage",
	Long:  `Prints memory usage visually in a bar graph.`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintMemUsage()
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
}

func PrintMemUsage() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	usedPercent := int(memInfo.UsedPercent)

	// Create colored text
	barColor := color.New(color.FgGreen)

	// Draw the memory usage bar
	bar := makeBar(usedPercent, barColor)

	// Print the memory usage with color
	fmt.Printf("\nMemory Usage: %s %d%%\n", bar, usedPercent)

	time.Sleep(1 * time.Second)
}

func makeBar(usedPercent int, barColor *color.Color) string {
	const barWidth = 20
	blockCount := int((usedPercent * barWidth) / 100)
	bar := "["

	for i := 0; i < barWidth; i++ {
		if i < blockCount {
			bar += barColor.Sprint("#")
		} else {
			bar += " "
		}
	}

	bar += "]"

	return bar
}
