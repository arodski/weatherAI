package weather

const (
	coldThreshold = 40.0
	hotThreshold  = 80.0
)

type Temperature struct {
	Value float64
}

func (t *Temperature) Classify() string {
	switch {
	case t.Value < coldThreshold:
		return "cold"
	case t.Value > hotThreshold:
		return "hot"
	default:
		return "moderate"
	}
}
