package gomgtclient

type Offer struct {
	PlayerID     int64          `json:"playerId"`
	PlayerOffers []*PlayerOffer `json:"playerOffers"`
}

type PlayerOffer struct {
	ID                     int    `json:"id,omitempty"`
	OfferID                int    `json:"offerId"`
	OfferStartDate         string `json:"offerStartDate"`
	OfferEndDate           string `json:"offerEndDate"`
	OfferDescription       string `json:"offerDescription"`
	OfferButtonDescription string `json:"offerButtonDescription"`
	PrizeID                int    `json:"prizeId"`
	PrizeDescription       string `json:"prizeDescription"`
}
