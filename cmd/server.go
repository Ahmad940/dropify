/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ahmad940/dropify/pkg/config"
	"github.com/spf13/cobra"
)

var (
	defaultPort int
	port        int
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if defaultPort != port {
			config.UpdateEnv(config.ENV_KEY_PORT, fmt.Sprintf("%v", port))
		}
	},
}

func init() {
	// retrieving default port
	var err error
	defaultPort, err = strconv.Atoi(strings.Split(config.GetEnv().PORT, ":")[1])
	if err != nil {
		fmt.Println("Error while converting port to int", err)
	}

	serverCmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Port to listen on")
	rootCmd.AddCommand(serverCmd)
}
