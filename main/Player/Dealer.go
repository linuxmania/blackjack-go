package Player

import (
	"fmt"
)

type Dealer struct {
	P Player
}

func (d Dealer) ReportShowing() {
	fmt.Println(d.P.Name, "showing:")
	d.P.Hand[1].PrintCard()
}
