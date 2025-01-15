package cmd

import (
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
var updatedDescription string

func updateRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	id, err := uuid.Parse(args[0])

	if err != nil {
		log.Printf("%v", err)
	}

	for i := range items {
		if items[i].Id == id.String() {
			items[i].UpdatedAt = time.Now()
			items[i].Description = updatedDescription
			items[i].SetPriority(updatedPriority)
			break
		}
	}

	if err := todo.SaveItems(viper.GetString("dataFile"), items); err != nil {
		log.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&updatedPriority, "priority", "p", 2, "Priority: in progress 1, todo 2, done 3")
	updateCmd.Flags().StringVarP(&updatedDescription, "description", "d", "none", "Update description")
}
