package dataset

import (
	"compress/gzip"
	"encoding/json"
	"os"
)

func Load(path string) ([]Reference, error) {
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

	var references []Reference

	decoder := json.NewDecoder(gzipReader)
	if err := decoder.Decode(&references); err != nil {
		return nil, err
	}

	return references, nil
}
