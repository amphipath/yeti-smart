package calculation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amphipath/yeti-smart/pkg/starforce/calculation"

)

func TestCost(t *testing.T) {
	t.Run("CostSum", func(t *testing.T) {
		c1 := &calculation.Cost{1, 2, 3, 4}
		c2 := &calculation.Cost{2,4,6,8}
		c3 := &calculation.Cost{3,6,9,12}
		c4 := &calculation.Cost{4,8,12,16}
		sum := calculation.CostSum(c1, c2, c3, c4)
		assert.Equal(t, sum, &calculation.Cost{10,20,30,40})
	})
	
	t.Run("Weight", func(t *testing.T) {
		c := &calculation.Cost{10, 1, 2, 0.5}
		assert.Equal(t, c.Weight(0.2), &calculation.Cost{2, 0.2, 0.4, 0.1})
	})

	t.Run("Add", func(t *testing.T) {
		c1 := &calculation.Cost{1, 2, 3, 4}
		c2 := &calculation.Cost{2,4,6,8}
		c1.Add(c2)
		assert.Equal(t, c1.Meso, 3)
		assert.Equal(t, c1.Boom, 6)
		assert.Equal(t, c1.Boom, 9)
		assert.Equal(t, c1.Star, 12)
	})
}