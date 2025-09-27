package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the gomod check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw lint compiles

Output:
generated_src/internal/github.com/gordonklaus/ineffassign/ineffassign.go:1: : # github.com/palantir/godel-okgo-asset-ineffassign/generated_src/internal/github.com/gordonklaus/ineffassign
generated_src/internal/github.com/gordonklaus/ineffassign/ineffassign.go:39:38: cannot use func(f *flag.Flag) {…} (value of type func(f *"github.com/palantir/godel-okgo-asset-ineffassign/generated_src/internal/amalgomated_flag".Flag)) as func(*"flag".Flag) value in argument to ineffassign.Analyzer.Flags.VisitAll (typecheck)
package amalgomated
1 issues:
* typecheck: 1

*/
