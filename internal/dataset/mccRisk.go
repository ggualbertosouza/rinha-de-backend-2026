package dataset

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/bytedance/sonic"
)

func LoadMccRisk(path string) (map[string]float32, error) {
	start := time.Now()

	mccRisk := make(map[string]float32)

	file, err := os.Open(path)
	if err != nil {
		return mccRisk, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return mccRisk, err
	}

	err = sonic.Unmarshal(bytes, &mccRisk)
	if err != nil {
		return mccRisk, err
	}

	log.Printf("load mccRisk in %s", time.Since(start))

	return mccRisk, err
}
