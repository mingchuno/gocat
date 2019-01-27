package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("gocat:", e)
		os.Exit(1)
	}
}

const (
	BufferSize = 4
)

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "gocat is cat",
	Long:  `gocat is cat`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		file, err := os.Open(args[0])
		check(err)
		buffer := make([]byte, BufferSize)
		n := 1
		for n > 0 {
			n, err = file.Read(buffer)
			if  err != nil && err == io.EOF {
				break
			}
			_, err := os.Stdout.Write(buffer[:n])
			check(err)
		}
		err = file.Close()
		check(err)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
