package fraud

import (
	"time"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
)

type Vector [14]float32

type Vectorizer struct {
	Normalization dataset.Normalization
	MccRisk       map[string]float32
}

func NewVectorizer(normalization dataset.Normalization, mccRisk map[string]float32) *Vectorizer {
	return &Vectorizer{
		Normalization: normalization,
		MccRisk:       mccRisk,
	}
}

func clamp(v float32) float32 {
	if v < 0 {
		return 0
	}

	if v > 1 {
		return 1
	}

	return v
}

func (v *Vectorizer) Vectorize(payload Payload) Vector {
	var vec Vector

	/* 0 - amount */
	vec[0] = clamp(
		payload.Transaction.Amount /
			v.Normalization.MaxAmount,
	)

	/* 1 - installments */
	vec[1] = clamp(
		float32(payload.Transaction.Installments) /
			v.Normalization.MaxInstallments,
	)

	/* 2 - amount_vs_avg */
	vec[2] = clamp(
		(payload.Transaction.Amount /
			payload.Customer.AvgAmount) /
			v.Normalization.AmountVsAvgRatio,
	)

	/* 3 - hour_of_day */
	hour := payload.Transaction.
		RequestedAt.
		UTC().
		Hour()

	vec[3] = float32(hour) / 23.0

	/*
		4 - day_of_week

		Rinha:
		seg=0
		ter=1
		qua=2
		qui=3
		sex=4
		sab=5
		dom=6
	*/
	vec[4] = weekday(
		payload.Transaction.RequestedAt,
	)

	/*
		5 - minutes_since_last_tx
		6 - km_from_last_tx
	*/
	if payload.LastTransaction == nil {

		vec[5] = -1
		vec[6] = -1

	} else {

		minutes := float32(
			payload.Transaction.
				RequestedAt.
				Sub(
					payload.LastTransaction.Timestamp,
				).
				Minutes(),
		)

		vec[5] = clamp(
			minutes /
				v.Normalization.MaxMinutes,
		)

		vec[6] = clamp(
			payload.LastTransaction.KmFromCurrent /
				v.Normalization.MaxKm,
		)
	}

	/*
		7 - km_from_home
	*/
	vec[7] = clamp(
		payload.Terminal.KmFromHome /
			v.Normalization.MaxKm,
	)

	/*
		8 - tx_count_24h
	*/
	vec[8] = clamp(
		float32(payload.Customer.TxCount24h) /
			v.Normalization.MaxTxCount24h,
	)

	/*
		9 - is_online
	*/
	if payload.Terminal.IsOnline {
		vec[9] = 1
	}

	/*
		10 - card_present
	*/
	if payload.Terminal.CardPresent {
		vec[10] = 1
	}

	/*
		11 - unknown_merchant
	*/
	vec[11] = unknownMerchant(payload)

	/*
		12 - mcc_risk
	*/
	vec[12] = v.mccRisk(
		payload.Merchant.Mcc,
	)

	/*
		13 - merchant_avg_amount
	*/
	vec[13] = clamp(
		payload.Merchant.AvgAmount /
			v.Normalization.MaxMerchantAvgAmount,
	)

	return vec
}

func weekday(t time.Time) float32 {
	w := t.UTC().Weekday()

	if w == time.Sunday {
		return 1.0
	}

	return float32(w-1) / 6.0
}

func (v *Vectorizer) mccRisk(
	mcc string,
) float32 {

	risk, ok := v.MccRisk[mcc]

	if !ok {
		return 0.5
	}

	return risk
}

func unknownMerchant(
	payload Payload,
) float32 {

	for _, merchant := range payload.Customer.KnownMerchants {

		if merchant == payload.Merchant.ID {
			return 0
		}
	}

	return 1
}
