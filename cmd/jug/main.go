package jug

import "fmt"

var (
	name    = "reverb"
	version = "03.Jan.2020"
	author  = "Harry Scells"
)

type args struct {
	Pipeline     string   `help:"Path to boogie experimental pipeline file"`
	Port         string   `help:"Port to run server on" arg:"-p"`
	Hosts        []string `help:"When in client mode, list of reverb servers to distribute the pipeline across" arg:"-s,separate"`
	Mode         string   `help:"Mode to run reverb in [client/server]" arg:"required,positional"`
	TemplateArgs []string `help:"Additional arguments to pass to experimental pipeline file" arg:"positional"`
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

		

}
