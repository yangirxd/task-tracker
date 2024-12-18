/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yangirxd/task-cli/todo"
	"log"
	"strconv"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update will update todo",
	Long:  `Update will update description of todo or change status of it`,
	Run:   updateRun,
}

func updateRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	id, err := args[1]

	if err != nil {
		log.Printf("%v", err)
	}

	for _, v := range items {
		if id == v.Id {

		}
	}

	if err := todo.SaveItems(viper.GetString("dataFile"), items); err != nil {
		fmt.Errorf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	//updateCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: in progress 1, todo 2, done 3")
}
