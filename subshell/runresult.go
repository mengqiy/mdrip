package subshell

import (
	"fmt"
	"os"
	"strings"

	"github.com/monopole/mdrip/base"
	"github.com/monopole/mdrip/program"
)

type status int

const (
	yep status = iota
	nope
)

// BlockOutput pairs success status (yes or no) with the output
// collected from a stream (i.e. stderr or stdout) as a result of
// executing a command block (or as much as could be executed before
// it failed).
//
// Output can appear on stderr without neccessarily being associated
// with shell failure, so it's collected even in successful runs.
type BlockOutput struct {
	success status
	output  string
}

func (x BlockOutput) Succeeded() bool {
	return x.success == yep
}

func (x BlockOutput) Output() string {
	return x.output
}

func NewFailureOutput(output string) *BlockOutput {
	return &BlockOutput{nope, output}
}

func NewSuccessOutput(output string) *BlockOutput {
	return &BlockOutput{yep, output}
}

// RunResult pairs BlockOutput with meta data about shell execution.
type RunResult struct {
	BlockOutput
	fileName base.FilePath     // File in which the error occurred.
	index    int               // Command block index.
	block    *program.BlockPgm // Content of actual command block.
	problem  error             // Error, if any.
	message  string            // Detailed error message, if any.
}

func NewRunResult() *RunResult {
	blockOutput := NewFailureOutput("")
	return &RunResult{
		*blockOutput, "", -1,
		program.NewEmptyBlockPgm(),
		nil, ""}
}

// For tests.
func NoCommandsRunResult(
	blockOutput *BlockOutput, path base.FilePath, index int, message string) *RunResult {
	return &RunResult{
		*blockOutput, path, index,
		program.NewEmptyBlockPgm(),
		nil, message}
}

func (x *RunResult) FileName() base.FilePath {
	return x.fileName
}

func (x *RunResult) Problem() error {
	return x.problem
}

func (x *RunResult) SetProblem(e error) *RunResult {
	x.problem = e
	return x
}

func (x *RunResult) Message() string {
	return x.message
}

func (x *RunResult) SetMessage(m string) *RunResult {
	x.message = m
	return x
}

func (x *RunResult) SetOutput(m string) *RunResult {
	x.output = m
	return x
}

func (x *RunResult) Index() int {
	return x.index
}

func (x *RunResult) SetIndex(i int) *RunResult {
	x.index = i
	return x
}

func (x *RunResult) SetBlock(b *program.BlockPgm) *RunResult {
	x.block = b
	return x
}

func (x *RunResult) SetFileName(n base.FilePath) *RunResult {
	x.fileName = n
	return x
}

func (x *RunResult) Print(selectedLabel base.Label) {
	delim := strings.Repeat("-", 70) + "\n"
	fmt.Fprintf(os.Stderr, delim)
	x.block.Print(os.Stderr, "Error", x.index+1, selectedLabel, x.fileName)
	fmt.Fprintf(os.Stderr, delim)
	printCapturedOutput("Stdout", delim, x.output)
	if len(x.message) > 0 {
		printCapturedOutput("Stderr", delim, x.message)
	}
}

func printCapturedOutput(name, delim, output string) {
	fmt.Fprintf(os.Stderr, "\n%s capture:\n", name)
	fmt.Fprintf(os.Stderr, delim)
	fmt.Fprintf(os.Stderr, output)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, delim)
}
