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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority,
		"priority",
		"p",
		2,
		"Priority: in progress 1, todo 2, done 3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Id: uuid.New().String(), Description: x, CreatedAt: time.Now(), UpdatedAt: time.Now()}
		item.SetPriority(priority)
		items = append(items, item)
		fmt.Printf("Task added successfully (ID: %v)\n", item.Id)
	}

	if err := todo.SaveItems(viper.GetString("dataFile"), items); err != nil {
		fmt.Errorf("%v", err)
	}
}
