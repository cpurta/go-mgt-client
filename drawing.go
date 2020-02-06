package gomgtclient

type PlayerDrawingResponse struct {
	PlayerID       int64            `json:"playerId"`
	PlayerDrawings []*PlayerDrawing `json:"playerDrawings"`
}

type PlayerDrawing struct {
	CheckInAvailable bool    `json:"checkInAvailable"`
	CheckInStartTime string  `json:"checkInStartTime"`
	Name             string  `json:"name"`
	CheckedIn        bool    `json:"checkedIn"`
	DWID             float64 `json:"DWID"`
	CheckInEndTime   string  `json:"checkInEndTime"`
	NumberOfEntries  int     `json:"numberOfEntries"`
	DrawingState     int     `json:"drawingState"`
	IconURL          string  `json:"IconURL"`
	CheckInDev       string  `json:"checkInDev"`
	CheckInTime      string  `json:"checkInTime"`
	Description      string  `json:"description"`
}

type DrawingResponse struct {
	AvailableDrawings *AvailableDrawings `json:"availableDrawings"`
}

type AvailableDrawings struct {
	Drawings []*Drawing `json:"drawing"`
}

type Drawing struct {
	ID                    int                    `json:"id,omitempty"`
	DWID                  int                    `json:"DWID"`
	Name                  string                 `json:"name"`
	SiteID                int                    `json:"siteId"`
	DrawingDescription    string                 `json:"drawingDescription"`
	DrawingMediaURL       string                 `json:"drawingMediaURL"`
	DrawingStartDate      string                 `json:"drawingStartDate"`
	DrawingEndDate        string                 `json:"drawingEndDate"`
	DrawingCheckInPeriods *DrawingCheckInPeriods `json:"drawingCheckInPeriods"`
	DrawingSummary        interface{}            `json:"drawingSummary"`
	DrawingImagePath      interface{}            `json:"drawingImagePath"`
	DrawingQuickTag       interface{}            `json:"drawingQuickTag"`
	DrawingQuickTimes     interface{}            `json:"drawingQuickTimes"`
	DrawingQuickCriteria  interface{}            `json:"drawingQuickCriteria"`
	DrawingBonusTag       interface{}            `json:"drawingBonusTag"`
	DrawingRules          interface{}            `json:"drawingRules"`
	IconURL               interface{}            `json:"IconURL"`
}

type DrawingCheckInPeriods struct {
	DrawingCheckInPeriod []*DrawingCheckInPeriod `json:"drawingCheckInPeriod"`
}

type DrawingCheckInPeriod struct {
	CheckInStartDate string `json:"checkInStartDate"`
	CheckInEndDate   string `json:"checkInEndDate"`
}
