/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create would create new repository on github",
	Long: `create would create new repository on github:
We can pass repository name, description and visibility as arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
			return
		}
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println(err)
			return
		}
		private, err := cmd.Flags().GetBool("private")
		if err != nil {
			fmt.Println(err)
			return
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Creating new repository on github")
		fmt.Println(name, description, private, token)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	//createCmd.Flags().StringP // setting up flags for the app, StringP has additional parameter to accept short form of the flag as well
	createCmd.Flags().StringP("name", "n", "", "name of the repository")
	createCmd.Flags().StringP("description", "d", "random Desc", "description of the repository")
	createCmd.Flags().BoolP("private", "p", true, "visibility of the repository, private is by default , set it to false to make it public")
	createCmd.Flags().StringP("token", "t", "", "GitHub authentication token")

	// marking the flags as required, so user must pass otherwise cli would quite
	err := createCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = createCmd.MarkFlagRequired("token")
	if err != nil {
		fmt.Println(err)
		return
	}

}
