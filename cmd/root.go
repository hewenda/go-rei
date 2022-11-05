package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var addUrl string

var rootCmd = &cobra.Command{
	Use:   "rei",
	Short: "rei spider",
	Run: func(cmd *cobra.Command, args []string) {
		if len(addUrl) > 0 {
			SkuAdd(addUrl)
		} else {
			SkuMonit()
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&addUrl, "add", "a", "", "Add a url to monit")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
