package gomgtclient

type AvailablePromotionsResponse struct {
	AvailablePromotions *AvailablePromotions `json:"availablePromotions"`
}

type AvailablePromotions struct {
	Promotion []*Promotion `json:"promotion"`
}

type Promotion struct {
	MtgID                        int         `json:"mtgId"`
	TgID                         int         `json:"tgId"`
	SiteID                       int         `json:"siteId"`
	MPromoDescription            string      `json:"mPromoDescription"`
	MPromoButtonDescription      string      `json:"mPromoButtonDescription"`
	MPromoMediaURL               string      `json:"mPromoMediaURL"`
	MPromoActive                 bool        `json:"mPromoActive"`
	TierGroupDescription         string      `json:"tierGroupDescription"`
	TierGroupStartDate           string      `json:"tierGroupStartDate"`
	TierGroupEndDate             string      `json:"tierGroupEndDate"`
	TierGroupActive              bool        `json:"tierGroupActive"`
	ParticipationRangeStartDate  string      `json:"participationRangeStartDate"`
	ParticipationRangeEndDate    string      `json:"participationRangeEndDate"`
	ParticipationRangeStartTime1 string      `json:"participationRangeStartTime1"`
	ParticipationRangeEndTime1   string      `json:"participationRangeEndTime1"`
	ParticipationRangeStartTime2 interface{} `json:"participationRangeStartTime2"`
	ParticipationRangeEndTime2   interface{} `json:"participationRangeEndTime2"`
	ParticipationRangeStartTime3 interface{} `json:"participationRangeStartTime3"`
	ParticipationRangeEndTime3   interface{} `json:"participationRangeEndTime3"`
	EntryType                    string      `json:"entryType"`
	PromoSummary                 interface{} `json:"promoSummary"`
	PromoImagePath               interface{} `json:"promoImagePath"`
	PromoQuickTag                interface{} `json:"promoQuickTag"`
	PromoQuickTimes              interface{} `json:"promoQuickTimes"`
	PromoQuickCriteria           interface{} `json:"promoQuickCriteria"`
	PromoBonusTag                interface{} `json:"promoBonusTag"`
	PromoRules                   interface{} `json:"promoRules"`
	IconURL                      interface{} `json:"IconURL"`
}

type PlayerPromotionResponse struct {
	PlayerID         string             `json:"playerId"`
	PlayerPromotions []*PlayerPromotion `json:"playerPromotions"`
}

type PlayerPromotion struct {
	MTGID             int    `json:"mtgId"`
	TGID              int    `json:"tgId"`
	Description       string `json:"promoDescription"`
	ButtonDescription string `json:"buttonDescription"`
	IconURL           string `json:"IconURL"`
}
