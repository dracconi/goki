package cmd

import (
	"fmt"
	"flag"
	"os"
)

type Command struct {
	Name string

	UsageLine string

	Short string

	Long string

	AddFlags func(fs *flag.FlagSet)

	FlagParse func(fs *flag.FlagSet, args []string) error
}

