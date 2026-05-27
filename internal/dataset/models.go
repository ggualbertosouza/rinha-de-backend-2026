package dataset

type Reference struct {
	Vector []float32 `json:"vector"`
	Label  string    `json:"label"`
}
