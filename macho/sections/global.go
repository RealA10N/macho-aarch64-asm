package sections

import (
	"fmt"
	"io"
	"strings"
)

type GlobalSection struct {
	Names []string
}

func (section GlobalSection) WriteAssembly(writer io.Writer) (n int, err error) {
	if len(section.Names) > 0 {
		return fmt.Fprintf(writer, ".global %s\n", strings.Join(section.Names, ", "))
	} else {
		return 0, nil
	}
}
