/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"doollm/server"

	"github.com/spf13/cobra"
)

var (
	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Start service",
		Run:   upFunc,
	}
)

func init() {

}

func upFunc(cmd *cobra.Command, args []string) {
	server.Run()
}
