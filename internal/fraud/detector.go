package fraud

type Result struct {
	Approved   bool    `json:"approved"`
	FraudScore float32 `json:"fraud_score"`
}

type Detector struct {
	Index SearchIndex
}

func NewDetector(idx SearchIndex) *Detector {
	return &Detector{Index: idx}
}

func (d *Detector) Detect(query Vector) Result {
	neighbors := d.Index.Search(query, 5)

	frauds := 0

	for _, n := range neighbors {
		if n.Fraud {
			frauds++
		}
	}

	score := float32(frauds) / 5.0

	return Result{
		Approved:   score < 0.6,
		FraudScore: score,
	}
}
