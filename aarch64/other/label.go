package other

import "fmt"

// TODO: perhaps not all strings are valid labels?

type Label struct {
	Name string
}

func (label Label) String() string {
	return fmt.Sprintf("%s:", label.Name)
}
