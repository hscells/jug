package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/hscells/jug"
	"os"
)

var (
	name    = "jug"
	version = "06.Feb.2020"
	author  = "Harry Scells"
)

type args struct {
	File  string `help:"JSON file" arg:"required,positional"`
	Query string `help:"Query" arg:"required,positional"`
}

func (args) Version() string {
	return version
}

func (args) Description() string {
	return fmt.Sprintf(`
   /  _ 
(_/(/(/ 
    _/
%s
@ %s
# %s
`, name, author, version)
}

func main() {
	var args args
	arg.MustParse(&args)

	f, err := os.OpenFile(args.File, os.O_RDONLY, 0664)
	if err != nil {
		panic(err)
	}

	var frame jug.DataFrame
	err = json.NewDecoder(f).Decode(&frame)
	if err != nil {
		panic(err)
	}

	plan, err := jug.Parse(args.Query)
	if err != nil {
		panic(err)
	}

	result, err := plan.Execute(frame)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(result.Frame, "", "  ")
	if err != nil {
		panic(err)
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		panic(err)
	}
}
