package cmd

import (
	"fmt"

	_ "hewenda/go-rei/storage"

	"github.com/spf13/cobra"
)

var CommandAddLink = &cobra.Command{
	Use:   "add",
	Short: "Add a link",
	Run: func(cmd *cobra.Command, args []string) {
		for _, val := range args {
			fmt.Println(val)
		}
	},
}
