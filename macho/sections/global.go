package sections

import (
	"fmt"
	"io"
	"strings"
)

type GlobalSection struct {
	Names []string
}

func (section GlobalSection) WriteTo(writer io.Writer) (int64, error) {
	if len(section.Names) > 0 {
		n, err := fmt.Fprintf(writer, ".global %s\n", strings.Join(section.Names, ", "))
		return int64(n), err
	} else {
		return 0, nil
	}
}
