package main

import (
	"github.com/rhiadc1/script/event"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "app"}
	rootCmd.AddCommand(event.Cmd())
	rootCmd.Execute()
}
