package runtime

import (
	"github.com/spf13/cobra"
)

var (
	Level    string
	Location string
)

var rootCmd = &cobra.Command{
	Use:   "merkbot",
	Short: "Runtime for translated msl script",
}

func AddEvent(name string, f func()) {
	cmd := &cobra.Command{
		Use:   name,
		Short: "event for " + name,
		Run:   func(cmd *cobra.Command, args []string) { f() },
	}
	cmd.Flags().StringVarP(&Level, "level", "", "*", "User level")
	cmd.Flags().StringVarP(&Location, "location", "", "*", "Event location")
	rootCmd.AddCommand(cmd)
}

func Start() {
	rootCmd.Execute()
}
