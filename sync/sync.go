package sync

type Counter struct {
	amount int
}

func (c Counter) Increment() {
	c.amount += 1
}

func (c Counter) Value() int {
	return c.amount
}
