package dataset

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"time"
)

func Load(path string) ([]Reference, error) {
	startTotal := time.Now()

	// =========
	startFile := time.Now()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	log.Printf("open file took: %s", time.Since(startFile))

	// ========
	startDecompress := time.Now()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	log.Printf("decompress took: %s", time.Since(startDecompress))

	// ==========
	startDecode := time.Now()
	var references []Reference

	decoder := json.NewDecoder(gzipReader)
	if err := decoder.Decode(&references); err != nil {
		return nil, err
	}

	log.Printf("decode took: %s", time.Since(startDecode))

	log.Printf("total time took: %s", time.Since(startTotal))

	return references, nil
}
