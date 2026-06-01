package fraud

type Neighbor struct {
	Distance float32
	Fraud    bool
}

type SearchIndex interface {
	// Recebe um vetor de consulta
	// e retorna os K vizinhos mais próximos.
	Search(
		query Vector,
		k int,
	) []Neighbor
}
