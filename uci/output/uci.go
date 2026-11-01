package output

import (
	"fmt"
)

const (
	name    = "GOndon"
	version = "1.0"
	author  = "Tanzor-Disco (Pomoyniy Guru)"
)

type UCIOutput struct {
}

func (o UCIOutput) GetOutput() string {
	return fmt.Sprintf("id name %s %s\nid author %s\nuciok\n", name, version, author)
}
