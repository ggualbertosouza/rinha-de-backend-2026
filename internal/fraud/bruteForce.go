package fraud

import (
	"sort"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
)

type BruteForce struct {
	Dataset []dataset.Reference
}

func NewBruteForce(rfs []dataset.Reference) *BruteForce {
	return &BruteForce{Dataset: rfs}
}

func (b *BruteForce) Search(query Vector, k int) []Neighbor {
	neighbors := make([]Neighbor, 0, len(b.Dataset))

	for _, ref := range b.Dataset {
		dist := Distance(
			query,
			ref.Vector,
		)

		neighbors = append(
			neighbors,
			Neighbor{
				Distance: dist,
				Fraud:    ref.Fraud,
			},
		)
	}

	sort.Slice(
		neighbors,
		func(i, j int) bool {
			return neighbors[i].Distance <
				neighbors[j].Distance
		},
	)

	if len(neighbors) > k {
		return neighbors[:k]
	}

	return neighbors
}
