package app

import (
	"sync/atomic"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
)

var Ready atomic.Bool
var Detector *fraud.Detector
var Vectorize *fraud.Vectorizer
