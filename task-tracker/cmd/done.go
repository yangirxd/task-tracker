package cmd

import (
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark item as Done",
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

/*\func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "Is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")
		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("dataFile"), items)
	} else {
		log.Println(i, "doesn't match any item")
	}
}*/
