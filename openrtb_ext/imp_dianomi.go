package openrtb_ext

import (
	"encoding/json"
)

type ExtImpDianomi struct {
	smartadId json.Number `json:"smartadId,omitempty"`
	PriceType string      `json:"priceType,omitempty"`
}
