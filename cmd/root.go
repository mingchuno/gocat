package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var noBuffer bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&noBuffer, "nobuffer", "u", false, "Disable output buffering.")
}

func check(e error) {
	if e != nil {
		fmt.Println("gocat:", e)
		os.Exit(1)
	}
}

func catFile(filename string) {
	file, err := os.Open(filename)
	check(err)
	catFileInternal(file)
	err = file.Close()
	check(err)
}

func catFileInternal(reader io.Reader) {
	buffer := make([]byte, 1)
	if !noBuffer {
		// let internal choose buffer size
		buffer = nil
	}
	_, err := io.CopyBuffer(os.Stdout, reader, buffer)
	check(err)
}

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "gocat is cat",
	Long:  `gocat is cat`,
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			catFileInternal(os.Stdin)
		} else {
			for i:= 0; i < len(args); i++ {
				catFile(args[i])
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
