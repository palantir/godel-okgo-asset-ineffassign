package amalgomated

import (
	"amalgomated_flag"

	"./pkg/ineffassign"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func AmalgomatedMain() {
	singlechecker.Main(ineffassign.Analyzer)
}

func init() {
	flag.Bool("n", false, "don't recursively check paths (deprecated)")
}
