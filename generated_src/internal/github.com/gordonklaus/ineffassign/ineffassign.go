package amalgomated

import (
	"github.com/palantir/godel-okgo-asset-ineffassign/generated_src/internal/github.com/gordonklaus/ineffassign/amalgomated_flag"

	"github.com/palantir/godel-okgo-asset-ineffassign/generated_src/internal/github.com/gordonklaus/ineffassign/pkg/ineffassign"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func AmalgomatedMain() {
	singlechecker.Main(ineffassign.Analyzer)
}

func init() {
	flag.Bool("n", false, "don't recursively check paths (deprecated)")
}
