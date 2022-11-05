package cmd

import (
	"fmt"
	"hewenda/go-rei/storage"

	"github.com/spf13/cobra"
)

var addToken string
var deleteToken string

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User token manager",
	Run: func(cmd *cobra.Command, args []string) {
		if len(addToken) > 0 {
			storage.InsertUser(addToken)
		} else if len(deleteToken) > 0 {
			storage.DeleteUser(deleteToken)
		}

		for _, user := range storage.QueryUser() {
			fmt.Println(user.Token)
		}
	},
}

func init() {
	userCmd.Flags().StringVarP(&addToken, "add", "a", "", "user token to add")
	userCmd.Flags().StringVarP(&deleteToken, "del", "d", "", "user token to delete")
	rootCmd.AddCommand(userCmd)
}
