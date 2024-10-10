/*
Copyright © 2023 Your Name <your.email@example.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:     "split [input]",
	Aliases: []string{"s"},
	Short:   "Split a video file at the specified duration",
	Long: `The split command splits an MP4 video file into two parts at a specified duration.
This command is useful for handling large video files efficiently. 

Usage:
  video-splitter split [input] --duration 12:00:00

Flags:
  -d, --duration  Duration for the first part (default: 12:00:00)
`,
	Run: func(cmd *cobra.Command, args []string) {
		var input string
		if len(args) > 0 {
			input = args[0] // Take the first positional argument as input
		} else {
			fmt.Println("Error: Input file is required.")
			cmd.Help() // Show usage information
			os.Exit(1)
		}

		duration, _ := cmd.Flags().GetString("duration")

		// If duration is not provided, set it to default
		if duration == "" {
			duration = "12:00:00"
		}

		startTime := duration // Set start time to be the same as duration

		if err := splitVideo(input, duration, startTime); err != nil {
			fmt.Printf("Error splitting video: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	// Define flags for the split command with a default value
	splitCmd.Flags().StringP("duration", "d", "12:00:00", "Duration for the first part")
}

// splitVideo splits the input video file into two parts
func splitVideo(input, duration, startTime string) error {
	// Get the directory and base name of the input file
	dir := filepath.Dir(input)
	baseName := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))

	// Create output file paths using the original file name as a prefix
	part1 := filepath.Join(dir, fmt.Sprintf("%s-part1.mp4", baseName))
	part2 := filepath.Join(dir, fmt.Sprintf("%s-part2.mp4", baseName))

	// Split the video into part 1
	cmd1 := exec.Command("ffmpeg", "-i", input, "-t", duration, "-c", "copy", part1)
	if err := cmd1.Run(); err != nil {
		return fmt.Errorf("failed to create %s: %w", part1, err)
	}

	// Split the video into part 2
	cmd2 := exec.Command("ffmpeg", "-i", input, "-ss", startTime, "-c", "copy", part2)
	if err := cmd2.Run(); err != nil {
		return fmt.Errorf("failed to create %s: %w", part2, err)
	}

	fmt.Printf("Successfully split the video into %s and %s\n", part1, part2)
	return nil
}
