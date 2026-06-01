package dataset

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func LoadMccRisk(path string) (map[string]float32, error) {
	start := time.Now()

	mccRisk := make(map[string]float32)

	file, err := os.Open(path)
	if err != nil {
		return mccRisk, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&mccRisk)

	log.Printf("load mccRisk in %s", time.Since(start))

	return mccRisk, err
}
