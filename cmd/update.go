package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yangirxd/task-cli/todo"
	"log"
	"time"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update will update todo",
	Long:  `Update will update description of todo or change status of it`,
	Run:   updateRun,
}

var updatedPriority int

func updateRun(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Error, not enough arguments, please type id and description")
		return
	}

	id, parseErr := uuid.Parse(args[0])
	if parseErr != nil {
		log.Printf("%v", parseErr)
		return
	}

	items, readErr := todo.ReadItems(viper.GetString("datafile"))
	if readErr != nil {
		log.Printf("%v", readErr)
		return
	}

	updatedDescription := args[1]

	for i := range items {
		if items[i].Id == id.String() {
			items[i].UpdatedAt = time.Now()
			items[i].Description = updatedDescription
			if cmd.Flags().Changed("priority") {
				items[i].SetPriority(updatedPriority)
			}
			break
		}
	}

	if err := todo.SaveItems(viper.GetString("dataFile"), items); err != nil {
		log.Printf("%v", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&updatedPriority, "priority", "p", 2, "Priority: in progress 1, todo 2, done 3")
}
