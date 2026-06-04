package dataset

import (
	"compress/gzip"
	"io"
	"log"
	"os"
	"time"

	"github.com/bytedance/sonic"
)

type Reference struct {
	Vector [14]float32
	Fraud  bool
}

type rawReference struct {
	Vector [14]float32 `json:"vector"`
	Label  string      `json:"label"`
}

func LoadReferences(path string) ([]Reference, error) {
	start := time.Now()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	var rawRefs []rawReference

	bytes, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	err = sonic.Unmarshal(bytes, &rawRefs)
	if err != nil {
		return nil, err
	}

	refs := make([]Reference, len(rawRefs))

	for i, raw := range rawRefs {
		refs[i] = Reference{
			Vector: raw.Vector,
			Fraud:  raw.Label == "fraud",
		}
	}

	log.Printf(
		"loaded %d references in %s",
		len(refs),
		time.Since(start),
	)

	return refs, nil
}
