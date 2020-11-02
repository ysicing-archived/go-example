// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package command

import (
	"app/constants"
	"fmt"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version of example",
		Run:   versionCommandFunc,
	}
}

func versionCommandFunc(cmd *cobra.Command, args []string) {
	fmt.Printf("release: %s, build date: %s, commit: %s", constants.Release, constants.Date, constants.Commit)
}
