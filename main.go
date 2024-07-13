package main

import (
	"flag"
	"os"

	"fortio.org/cli"
	"fortio.org/log"
	"github.com/ldemailly/gorepl/eval"
	"github.com/ldemailly/gorepl/repl"
)

func main() {
	showParse := flag.Bool("parse", false, "show parse tree")
	showEval := flag.Bool("eval", true, "show eval results")
	sharedState := flag.Bool("shared-state", false, "All files share same interpreter state (default is new state for each)")
	cli.ArgsHelp = "*.gr files to interpret or no arg for stdin repl..."
	cli.MaxArgs = -1
	cli.Main()
	options := repl.Options{
		ShowParse: *showParse,
		ShowEval:  *showEval,
	}
	nArgs := len(flag.Args())
	if nArgs == 0 {
		repl.Interactive(os.Stdin, os.Stdout, options)
		return
	}
	options.All = true
	s := eval.NewState()
	for _, file := range flag.Args() {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Infof("Running %s", file)
		repl.EvalAll(s, f, os.Stdout, options)
		f.Close()
		if !*sharedState {
			s = eval.NewState()
		}
	}
	log.Infof("All done")
}
