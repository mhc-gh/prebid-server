package openrtb_ext

import (
	"encoding/json"
)

type ExtImpDianomi struct {
	SmartadId       json.Number `json:"mid,omitempty"`
	PriceType         string      `json:"priceType,omitempty"`
}
