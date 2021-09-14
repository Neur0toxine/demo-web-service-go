package cmd

import "github.com/jessevdk/go-flags"

// Parser will parse command line
var Parser *flags.Parser

func init() {
	Parser = flags.NewParser(nil, flags.Default)
}
