package starforce_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amphipath/yeti-smart/pkg/starforce"
)

func TestMatrixRow(t *testing.T) {
	t.Run("IsValid", func(t *testing.T) {
		testcases := []struct {
			name string
			row  starforce.MatrixRow
			want bool
		}{
			{"sums to more than 1", starforce.MatrixRow{0.4, 0.4, 0.4}, false},
			{"sums to exactly 1", starforce.MatrixRow{0.5,0.5,0}, true},
			{"sums to less than 1", starforce.MatrixRow{0.1, 0, 0}, true},
		}
		for _, c := range testcases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.row.IsValid(), c.want)
			})
		}
	})
}