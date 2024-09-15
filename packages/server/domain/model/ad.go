package model

import (
	"server/infrastructure/ent"
)

type Ad struct {
	EntAd *ent.Ad
}

type NewAdParams struct {
	EntAd      *ent.Ad
	OptionList []func(*Ad)
}

func NewAd(params NewAdParams) *Ad {
	Ad := &Ad{EntAd: params.EntAd}

	for _, option_i := range params.OptionList {
		option_i(Ad)
	}

	return Ad
}
