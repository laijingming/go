package tree

import (
	"fmt"
)

type TN struct {
	Value       int
	Left, Right *TN
}

func (n TN) Print() {
	fmt.Println(n.Value)
}
