package pkg

type Calculator struct {
	Sum       int
	Histories []History
}

type History struct {
	operation string
	param     int
	result    int
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c Calculator) Add(a int) int {
	c.Sum += a
	c.Histories = append(c.Histories, History{operation: "Add", param: a, result: c.Sum})
	return c.Sum
}

func (c Calculator) Subtract(a int) int {
	c.Sum -= a
	c.Histories = append(c.Histories, History{operation: "Subtract", param: a, result: c.Sum})
	return c.Sum
}

func (c Calculator) Multiply(a int) int {
	c.Sum *= a
	c.Histories = append(c.Histories, History{operation: "Multiply", param: a, result: c.Sum})
	return c.Sum
}

func (c Calculator) Divide(a int) int {
	c.Sum /= a
	c.Histories = append(c.Histories, History{operation: "Divide", param: a, result: c.Sum})
	return c.Sum
}

func (c Calculator) GetSum() int {
	return c.Sum
}

func (c Calculator) GetHistories() []History {
	return c.Histories
}
