package gomgtclient

type Client interface {
	GetOffers(playerID, uuid string) ([]*Offer, error)
	PlayerDrawingCheckIn(playerID, drawingID, uuid string) (*CheckIn, error)
	GetPromotions(uuid string) ([]*Promotion, error)
	GetPlayerPromotions(playerID, uuid string) ([]*PlayerPromotion, error)
	GetDrawings(uuid string) (*AvailableDrawings, error)
	GetPlayerDrawings(playerID, uuid string) ([]*PlayerDrawing, error)
}
