package debug

import (
	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/dgraph/cmd"
	"github.com/spf13/cobra"
)

type options struct {
	dir      string
	truncate bool
}

var opt options
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Run Dgraph debugger",
	RunE:  doRun,
}

func init() {
	cmd.RootCmd.AddCommand(debugCmd)
	debugCmd.Flags().StringVarP(&opt.dir, "dir", "d", "", "Badger data directory.")
	debugCmd.Flags().BoolVarP(&opt.truncate, "truncate", "t", false,
		"Allow truncation of data to get rid of invalid entries.")
}

func doRun(cmd *cobra.Command, args []string) error {
	bopts := badger.DefaultOptions
	bopts.Dir = opt.dir
	bopts.ValueDir = opt.dir
	bopts.Truncate = opt.truncate

	db, err := badger.Open(bopts)
	if err != nil {
		return err
	}

	return nil
}
