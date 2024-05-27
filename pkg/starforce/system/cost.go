package system

type Cost struct {
	Meso       float64
	Boom       float64
	MaplePoint float64
	Star       float64
}

// Non-mutational summation of costs
func CostSum(costs ...*Cost) (result *Cost) {
	result.Add(costs...)
	return
}

// Non-mutational scalar multiplication of a Cost
func (c *Cost) Weight(m float64) *Cost {
	return &Cost{c.Meso * m, c.Boom * m, c.MaplePoint * m, c.Star * m}
}

// Mutational addition of costs
func (c *Cost) Add(costs ...*Cost) {
	for _, c2 := range costs {
		c.Meso = c.Meso + c2.Meso
		c.Boom = c.Boom + c2.Boom
		c.MaplePoint = c.MaplePoint + c2.MaplePoint
		c.Star = c.Star + c2.Star
	}
}