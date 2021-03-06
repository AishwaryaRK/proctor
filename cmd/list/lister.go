package list

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gojektech/proctor/daemon"
	"github.com/gojektech/proctor/io"
	"github.com/spf13/cobra"
	"github.com/gojektech/proctor/utility/sort"
)

func NewCmd(printer io.Printer, proctorDClient daemon.Client) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "List procs available for execution",
		Long:    "List procs available for execution",
		Example: "proctor list",
		Run: func(cmd *cobra.Command, args []string) {
			procList, err := proctorDClient.ListProcs()
			if err != nil {
				printer.Println(err.Error(), color.FgRed)
				return
			}
			printer.Println("List of Procs:\n", color.FgGreen)
			sort.Procs(procList)

			for _, proc := range procList {
				printer.Println(fmt.Sprintf("%-40s %-100s", proc.Name, proc.Description), color.Reset)
			}

			printer.Println("\nFor detailed information of any proc, run:\nproctor describe <proc_name>", color.FgGreen)
		},
	}
}
