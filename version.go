package main

import (
	"fmt"
)

// Version the version of supervisor
var Version = "0.7.8"

// VersionCommand implement the flags.Commander interface
type VersionCommand struct {
}

var versionCommand VersionCommand

// Execute implement Execute() method defined in flags.Commander interface, executes the given command
func (v VersionCommand) Execute(args []string) error {
	fmt.Println(Version)
	return nil
}

func init() {
	parser.AddCommand("version",
		"show the version of supervisor",
		"display the supervisor version",
		&versionCommand)
}
