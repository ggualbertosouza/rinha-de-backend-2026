package dataset

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Normalization struct {
	MaxAmount            float32 `json:"max_amount"`
	MaxInstallments      float32 `json:"max_installments"`
	AmountVsAvgRatio     float32 `json:"amount_vs_avg_ratio"`
	MaxMinutes           float32 `json:"max_minutes"`
	MaxKm                float32 `json:"max_km"`
	MaxTxCount24h        float32 `json:"max_tx_count_24h"`
	MaxMerchantAvgAmount float32 `json:"max_merchant_avg_amount"`
}

func LoadNormalization(path string) (Normalization, error) {
	start := time.Now()

	var normalization Normalization

	file, err := os.Open(path)
	if err != nil {
		return normalization, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&normalization)

	log.Printf("load normalization in %s", time.Since(start))

	return normalization, err
}
