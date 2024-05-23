package starforce

type (
	MatrixRow struct {
		Success float64
		Drop    float64
		Boom    float64
	}
)

func (m *MatrixRow) IsValid() bool {
	return m.Success+m.Drop+m.Boom <= 1
}