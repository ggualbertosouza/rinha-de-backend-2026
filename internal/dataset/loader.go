package dataset

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"time"
)

func Load(path string) ([]Reference, error) {
	startFile := time.Now()

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	log.Printf("open file took: %s", time.Since(startFile))

	// =========

	startGz := time.Now()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	log.Printf("decompress gz took: %s", time.Since(startGz))

	startDecode := time.Now()

	var references []Reference

	decoder := json.NewDecoder(gzipReader)
	if err := decoder.Decode(&references); err != nil {
		return nil, err
	}

	log.Printf("decode took: %s", time.Since(startDecode))

	return references, nil
}
