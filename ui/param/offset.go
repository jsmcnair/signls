package param

import (
	"fmt"

	"signls/core/common"
	"signls/core/node"

	"golang.org/x/text/unicode/norm"
)

type Offset struct {
	nodes []common.Node
}

func (o Offset) Name() string {
	return "off"
}

func (o Offset) Display() string {
	if o.nodes[0].(*node.EuclidEmitter).Offset.RandomAmount() != 0 {
		return norm.NFC.String(
			fmt.Sprintf(
				"%d%+d\u033c",
				o.nodes[0].(*node.EuclidEmitter).Offset.Value(),
				o.nodes[0].(*node.EuclidEmitter).Offset.RandomAmount(),
			),
		)
	}
	return fmt.Sprintf("%d", o.Value())
}

func (o Offset) Value() int {
	return o.nodes[0].(*node.EuclidEmitter).Offset.Value()
}

func (o Offset) AltValue() int {
	return o.nodes[0].(*node.EuclidEmitter).Offset.RandomAmount()
}

func (o Offset) Up() {
	o.Set(o.Value() + 1)
}

func (o Offset) Down() {
	o.Set(o.Value() - 1)
}

func (o Offset) Left() {
	o.SetAlt(o.nodes[0].(*node.EuclidEmitter).Offset.RandomAmount() - 1)
}

func (o Offset) Right() {
	o.SetAlt(o.nodes[0].(*node.EuclidEmitter).Offset.RandomAmount() + 1)
}

func (o Offset) AltUp() {}

func (o Offset) AltDown() {}

func (o Offset) AltLeft() {}

func (o Offset) AltRight() {}

func (o Offset) Set(value int) {
	if value < 0 {
		return
	}
	for _, n := range o.nodes {
		if value > n.(*node.EuclidEmitter).Steps.Value() {
			continue
		}
		n.(*node.EuclidEmitter).Offset.Set(value)
	}
}

func (o Offset) SetAlt(value int) {
	for _, n := range o.nodes {
		n.(*node.EuclidEmitter).Offset.SetRandomAmount(value)
	}
}
