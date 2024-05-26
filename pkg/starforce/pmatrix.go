package starforce

type (
	// MatrixRow describes the base data type that describes the base outcome chances of a star attempt per star per sf system
	MatrixRow struct {
		Success float64
		Drop    float64
		Boom    float64
	}
)

func (m *MatrixRow) IsValid() bool {
	return m.Success+m.Drop+m.Boom <= 1
}