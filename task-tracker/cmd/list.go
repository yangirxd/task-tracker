package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yangirxd/task-cli/todo"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the todos",
	Long:  `list will show all todos`,
	Run:   listRun,
}

var (
	doneOpt       bool
	todoOpt       bool
	inProgressOpt bool
	allOpt        bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'done' Todos")
	listCmd.Flags().BoolVar(&todoOpt, "todo", false, "Show 'todo' Todos")
	listCmd.Flags().BoolVar(&inProgressOpt, "in-progress", false, "Show 'in progress' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todos")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for i, v := range items {
		if allOpt || (doneOpt && v.Status == 3) || (inProgressOpt && v.Status == 1) || (todoOpt && v.Status == 2) {
			fmt.Fprintln(w,
				strconv.Itoa(i+1)+".\t"+v.Description+"\t"+"Status:"+"\t"+v.PrettyP())
		}
	}
	w.Flush()
}
