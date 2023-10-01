/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Ahmad940/dropify/platform/cli"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
	role     string
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long:  `Manage users right away with this command to create, update or even delete a  user`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	userCmd.AddCommand(newUserCmd())
	userCmd.AddCommand(resetUserCmd())
	userCmd.AddCommand(deleteUserCmd())

	rootCmd.AddCommand(userCmd)
}

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			// new user here
			cli.NewUser(&username, &password, &role)
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "The username of the new user")
	cmd.Flags().StringVarP(&password, "password", "p", "", "The password of the new user")
	cmd.Flags().StringVarP(&password, "role", "r", "", "The role of the new user")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")

	return cmd
}

func resetUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset",
		Short: "Reset the password of a user",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Resetting password for user...")
			// Reset the password of the user here
			cli.ResetUser(&username, &password)
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "The username of the user to reset the password for")
	cmd.Flags().StringVarP(&password, "password", "p", "", "The new password for the user")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")

	return cmd
}

func deleteUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a user",
		Run: func(cmd *cobra.Command, args []string) {
			// Delete the user here
			cli.DeleteUser(&username)
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "The username of the user to delete")
	cmd.MarkFlagRequired("username")

	return cmd
}
