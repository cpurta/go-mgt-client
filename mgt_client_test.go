package gomgtclient

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestPlayerDrawingCheckIn(t *testing.T) {
	var (
		client                       = NewClient(http.DefaultClient, "http://localhost", "fakelicense")
		playerDrawingCheckInResponse = `{"CheckedIn":true,"Message":"Your tickets have been deposited in the drum for today’s drawing."}`
		playerID                     = "1234"
		drawingID                    = "1234"
		checkIn                      *CheckIn
		expectedCheckIn              = &CheckIn{
			CheckedIn: true,
			Message:   "Your tickets have been deposited in the drum for today’s drawing.",
		}
		err error
	)

	httpmock.Activate()
	defer httpmock.Deactivate()

	stringResponder := httpmock.NewStringResponder(http.StatusOK, playerDrawingCheckInResponse)

	httpmock.RegisterResponder("PUT", fmt.Sprintf("http://%s/mgt/restservices/player/%s/drawing/%s/checkin", "localhost", playerID, drawingID), stringResponder)

	checkIn, err = client.PlayerDrawingCheckIn(playerID, drawingID, "fakeUUID")

	assert.NoError(t, err, "should not have got an error")

	assert.Equal(t, expectedCheckIn, checkIn)
}

func TestGetPromotions(t *testing.T) {
	var (
		client             = NewClient(http.DefaultClient, "http://localhost", "fakelicense")
		promotionsResponse = `{
    "availablePromotions": {
        "promotion": [
            {
                "mtgId": 181,
                "tgId": 408,
                "siteId": 1,
                "mPromoDescription": "JAN20 Birthday Cash Bash",
                "mPromoButtonDescription": "1500 Birthday Points!",
                "mPromoMediaURL": "",
                "mPromoActive": true,
                "tierGroupDescription": "1500 Birthday Points",
                "tierGroupStartDate": "2020-01-01T00:00:00",
                "tierGroupEndDate": "2020-01-31T00:00:00",
                "tierGroupActive": true,
                "participationRangeStartDate": "2020-01-01T00:00:00",
                "participationRangeEndDate": "2020-01-31T00:00:00",
                "participationRangeStartTime1": "00:00",
                "participationRangeEndTime1": "23:59",
                "participationRangeStartTime2": null,
                "participationRangeEndTime2": null,
                "participationRangeStartTime3": null,
                "participationRangeEndTime3": null,
                "entryType": "One Entry per Range",
                "promoSummary": null,
                "promoImagePath": null,
                "promoQuickTag": null,
                "promoQuickTimes": null,
                "promoQuickCriteria": null,
                "promoBonusTag": null,
                "promoRules": null,
                "IconURL": null
            },
            {
                "mtgId": 181,
                "tgId": 410,
                "siteId": 1,
                "mPromoDescription": "JAN20 Birthday Cash Bash",
                "mPromoButtonDescription": "$30 Reward Play",
                "mPromoMediaURL": "",
                "mPromoActive": true,
                "tierGroupDescription": "$30 Reward Play",
                "tierGroupStartDate": "2020-01-01T00:00:00",
                "tierGroupEndDate": "2020-01-31T00:00:00",
                "tierGroupActive": true,
                "participationRangeStartDate": "2020-01-01T00:00:00",
                "participationRangeEndDate": "2020-01-31T00:00:00",
                "participationRangeStartTime1": "00:00",
                "participationRangeEndTime1": "23:59",
                "participationRangeStartTime2": null,
                "participationRangeEndTime2": null,
                "participationRangeStartTime3": null,
                "participationRangeEndTime3": null,
                "entryType": "One Entry per Range",
                "promoSummary": null,
                "promoImagePath": null,
                "promoQuickTag": null,
                "promoQuickTimes": null,
                "promoQuickCriteria": null,
                "promoBonusTag": null,
                "promoRules": null,
                "IconURL": null
            },
            {
                "mtgId": 182,
                "tgId": 411,
                "siteId": 1,
                "mPromoDescription": "JAN20 - $47K Hot Seats & Slot Tourney",
                "mPromoButtonDescription": "$47K Hot Seats & Slot Tourney",
                "mPromoMediaURL": "",
                "mPromoActive": true,
                "tierGroupDescription": "$47K Hot Seats & Slot Tourney",
                "tierGroupStartDate": "2020-01-02T00:00:00",
                "tierGroupEndDate": "2020-01-23T00:00:00",
                "tierGroupActive": true,
                "participationRangeStartDate": "2020-01-02T00:00:00",
                "participationRangeEndDate": "2020-01-23T00:00:00",
                "participationRangeStartTime1": "15:00",
                "participationRangeEndTime1": "19:10",
                "participationRangeStartTime2": null,
                "participationRangeEndTime2": null,
                "participationRangeStartTime3": null,
                "participationRangeEndTime3": null,
                "entryType": "One Entry per Day",
                "promoSummary": null,
                "promoImagePath": null,
                "promoQuickTag": null,
                "promoQuickTimes": null,
                "promoQuickCriteria": null,
                "promoBonusTag": null,
                "promoRules": null,
                "IconURL": null
            }
        ]
    }
}`
		promotions         []*Promotion
		expectedPromotions = []*Promotion{
			{
				MtgID:                        181,
				TgID:                         408,
				SiteID:                       1,
				MPromoDescription:            "JAN20 Birthday Cash Bash",
				MPromoButtonDescription:      "1500 Birthday Points!",
				MPromoMediaURL:               "",
				MPromoActive:                 true,
				TierGroupDescription:         "1500 Birthday Points",
				TierGroupStartDate:           "2020-01-01T00:00:00",
				TierGroupEndDate:             "2020-01-31T00:00:00",
				TierGroupActive:              true,
				ParticipationRangeStartDate:  "2020-01-01T00:00:00",
				ParticipationRangeEndDate:    "2020-01-31T00:00:00",
				ParticipationRangeStartTime1: "00:00",
				ParticipationRangeEndTime1:   "23:59",
				ParticipationRangeStartTime2: nil,
				ParticipationRangeEndTime2:   nil,
				ParticipationRangeStartTime3: nil,
				ParticipationRangeEndTime3:   nil,
				EntryType:                    "One Entry per Range",
				PromoSummary:                 nil,
				PromoImagePath:               nil,
				PromoQuickTag:                nil,
				PromoQuickTimes:              nil,
				PromoQuickCriteria:           nil,
				PromoBonusTag:                nil,
				PromoRules:                   nil,
				IconURL:                      nil,
			},
			{
				MtgID:                        181,
				TgID:                         410,
				SiteID:                       1,
				MPromoDescription:            "JAN20 Birthday Cash Bash",
				MPromoButtonDescription:      "$30 Reward Play",
				MPromoMediaURL:               "",
				MPromoActive:                 true,
				TierGroupDescription:         "$30 Reward Play",
				TierGroupStartDate:           "2020-01-01T00:00:00",
				TierGroupEndDate:             "2020-01-31T00:00:00",
				TierGroupActive:              true,
				ParticipationRangeStartDate:  "2020-01-01T00:00:00",
				ParticipationRangeEndDate:    "2020-01-31T00:00:00",
				ParticipationRangeStartTime1: "00:00",
				ParticipationRangeEndTime1:   "23:59",
				ParticipationRangeStartTime2: nil,
				ParticipationRangeEndTime2:   nil,
				ParticipationRangeStartTime3: nil,
				ParticipationRangeEndTime3:   nil,
				EntryType:                    "One Entry per Range",
				PromoSummary:                 nil,
				PromoImagePath:               nil,
				PromoQuickTag:                nil,
				PromoQuickTimes:              nil,
				PromoQuickCriteria:           nil,
				PromoBonusTag:                nil,
				PromoRules:                   nil,
				IconURL:                      nil,
			},
			{
				MtgID:                        182,
				TgID:                         411,
				SiteID:                       1,
				MPromoDescription:            "JAN20 - $47K Hot Seats & Slot Tourney",
				MPromoButtonDescription:      "$47K Hot Seats & Slot Tourney",
				MPromoMediaURL:               "",
				MPromoActive:                 true,
				TierGroupDescription:         "$47K Hot Seats & Slot Tourney",
				TierGroupStartDate:           "2020-01-02T00:00:00",
				TierGroupEndDate:             "2020-01-23T00:00:00",
				TierGroupActive:              true,
				ParticipationRangeStartDate:  "2020-01-02T00:00:00",
				ParticipationRangeEndDate:    "2020-01-23T00:00:00",
				ParticipationRangeStartTime1: "15:00",
				ParticipationRangeEndTime1:   "19:10",
				ParticipationRangeStartTime2: nil,
				ParticipationRangeEndTime2:   nil,
				ParticipationRangeStartTime3: nil,
				ParticipationRangeEndTime3:   nil,
				EntryType:                    "One Entry per Day",
				PromoSummary:                 nil,
				PromoImagePath:               nil,
				PromoQuickTag:                nil,
				PromoQuickTimes:              nil,
				PromoQuickCriteria:           nil,
				PromoBonusTag:                nil,
				PromoRules:                   nil,
				IconURL:                      nil,
			},
		}
		err error
	)

	httpmock.Activate()
	defer httpmock.Deactivate()

	stringResponder := httpmock.NewStringResponder(http.StatusOK, promotionsResponse)

	httpmock.RegisterResponder("GET", fmt.Sprintf("http://%s/mgt/restservices/promotions", "localhost"), stringResponder)

	promotions, err = client.GetPromotions("fakeUUID")

	assert.NoError(t, err, "should not have got an error")

	assert.Equal(t, expectedPromotions, promotions)
}

func TestGetPlayerPromotions(t *testing.T) {
	var (
		client             = NewClient(http.DefaultClient, "http://localhost", "fakelicense")
		promotionsResponse = `{"playerId":"6","playerPromotions":{"playerPromotion":[{"mtgId":10,"tgId":49,"promoDescription":"WeekDay Bonus Free Play","buttonDescription":"WeekDay Bonus Free Play","IconURL":"http://server:8087/images/promoimage.jpg"},{"mtgId":9,"tgId":48,"promoDescription":"Press 2 Win","buttonDescription":"Press 2 Win","IconURL":"http://server:8087/images/promoimage.jpg"}]}}`
		playerID           = "6"
		promotions         []*PlayerPromotion
		expectedPromotions = []*PlayerPromotion{
			{
				MTGID:             10,
				TGID:              49,
				Description:       "WeekDay Bonus Free Play",
				ButtonDescription: "WeekDay Bonus Free Play",
				IconURL:           "http://server:8087/images/promoimage.jpg",
			},
			{
				MTGID:             9,
				TGID:              48,
				Description:       "Press 2 Win",
				ButtonDescription: "Press 2 Win",
				IconURL:           "http://server:8087/images/promoimage.jpg",
			},
		}
		err error
	)

	httpmock.Activate()
	defer httpmock.Deactivate()

	stringResponder := httpmock.NewStringResponder(http.StatusOK, promotionsResponse)

	httpmock.RegisterResponder("GET", fmt.Sprintf("http://%s/mgt/restservices/player/%s/promotions", "localhost", playerID), stringResponder)

	promotions, err = client.GetPlayerPromotions(playerID, "fakeUUID")

	assert.NoError(t, err, "should not have got an error")

	assert.Equal(t, expectedPromotions, promotions)
}

func TestGetDrawings(t *testing.T) {
	var (
		client           = NewClient(http.DefaultClient, "http://localhost", "fakelicense")
		drawingsResponse = `{"availableDrawings":{"drawing":[{"DWID":111,"name":"Six Figure Snowflakes","siteId":1,"drawingDescription":"Six Figure Snowflakes","drawingMediaURL":"","drawingStartDate":"2020-01-01T00:00:00","drawingEndDate":"2020-01-30T00:00:00","drawingCheckInPeriods":{"drawingCheckInPeriod":[{"checkInStartDate":"2020-01-30T16:00:00","checkInEndDate":"2020-01-30T19:45:00"}]},"drawingSummary":null,"drawingImagePath":null,"drawingQuickTag":null,"drawingQuickTimes":null,"drawingQuickCriteria":null,"drawingBonusTag":null,"drawingRules":null,"IconURL":null},{"DWID":115,"name":"$47K Hot Seats & Slot Tourney","siteId":1,"drawingDescription":"$47K Hot Seats & Slot Tourney","drawingMediaURL":"","drawingStartDate":"2020-01-02T00:00:00","drawingEndDate":"2020-01-30T00:00:00","drawingCheckInPeriods":{"drawingCheckInPeriod":[{"checkInStartDate":"2020-01-02T15:00:00","checkInEndDate":"2020-01-02T19:09:00"},{"checkInStartDate":"2020-01-09T15:00:00","checkInEndDate":"2020-01-09T19:09:00"},{"checkInStartDate":"2020-01-16T15:00:00","checkInEndDate":"2020-01-16T19:09:00"},{"checkInStartDate":"2020-01-23T15:00:00","checkInEndDate":"2020-01-23T19:09:00"}]},"drawingSummary":null,"drawingImagePath":null,"drawingQuickTag":null,"drawingQuickTimes":null,"drawingQuickCriteria":null,"drawingBonusTag":null,"drawingRules":null,"IconURL":null},{"DWID":116,"name":"$113K Decide Your Ride Drawing","siteId":1,"drawingDescription":"$113K Decide Your Ride Drawing","drawingMediaURL":"","drawingStartDate":"2020-02-01T00:00:00","drawingEndDate":"2020-03-26T00:00:00","drawingCheckInPeriods":{"drawingCheckInPeriod":[{"checkInStartDate":"2020-03-26T16:00:00","checkInEndDate":"2020-03-26T19:59:00"}]},"drawingSummary":null,"drawingImagePath":null,"drawingQuickTag":null,"drawingQuickTimes":null,"drawingQuickCriteria":null,"drawingBonusTag":null,"drawingRules":null,"IconURL":null},{"DWID":117,"name":"FEB20 - Valentine's Day Feel The Love","siteId":1,"drawingDescription":"FEB20 - Valentine's Day Feel The Love","drawingMediaURL":"","drawingStartDate":"2020-02-14T00:00:00","drawingEndDate":"2020-02-14T00:00:00","drawingCheckInPeriods":{"drawingCheckInPeriod":[{"checkInStartDate":"2020-02-14T18:00:00","checkInEndDate":"2020-02-14T20:00:00"}]},"drawingSummary":null,"drawingImagePath":null,"drawingQuickTag":null,"drawingQuickTimes":null,"drawingQuickCriteria":null,"drawingBonusTag":null,"drawingRules":null,"IconURL":null}]}}`
		drawings         *AvailableDrawings
		expectedDrawings = &AvailableDrawings{
			Drawings: []*Drawing{
				{
					DWID:               111,
					Name:               "Six Figure Snowflakes",
					SiteID:             1,
					DrawingDescription: "Six Figure Snowflakes",
					DrawingMediaURL:    "",
					DrawingStartDate:   "2020-01-01T00:00:00",
					DrawingEndDate:     "2020-01-30T00:00:00",
					DrawingCheckInPeriods: &DrawingCheckInPeriods{
						DrawingCheckInPeriod: []*DrawingCheckInPeriod{
							{
								CheckInStartDate: "2020-01-30T16:00:00",
								CheckInEndDate:   "2020-01-30T19:45:00",
							},
						},
					},
					DrawingSummary:       nil,
					DrawingImagePath:     nil,
					DrawingQuickTag:      nil,
					DrawingQuickTimes:    nil,
					DrawingQuickCriteria: nil,
					DrawingBonusTag:      nil,
					DrawingRules:         nil,
					IconURL:              nil,
				},
				{
					DWID:               115,
					Name:               "$47K Hot Seats & Slot Tourney",
					SiteID:             1,
					DrawingDescription: "$47K Hot Seats & Slot Tourney",
					DrawingMediaURL:    "",
					DrawingStartDate:   "2020-01-02T00:00:00",
					DrawingEndDate:     "2020-01-30T00:00:00",
					DrawingCheckInPeriods: &DrawingCheckInPeriods{
						DrawingCheckInPeriod: []*DrawingCheckInPeriod{
							{
								CheckInStartDate: "2020-01-02T15:00:00",
								CheckInEndDate:   "2020-01-02T19:09:00",
							},
							{
								CheckInStartDate: "2020-01-09T15:00:00",
								CheckInEndDate:   "2020-01-09T19:09:00",
							},
							{
								CheckInStartDate: "2020-01-16T15:00:00",
								CheckInEndDate:   "2020-01-16T19:09:00",
							},
							{
								CheckInStartDate: "2020-01-23T15:00:00",
								CheckInEndDate:   "2020-01-23T19:09:00",
							},
						},
					},
					DrawingSummary:       nil,
					DrawingImagePath:     nil,
					DrawingQuickTag:      nil,
					DrawingQuickTimes:    nil,
					DrawingQuickCriteria: nil,
					DrawingBonusTag:      nil,
					DrawingRules:         nil,
					IconURL:              nil,
				},
				{
					DWID:               116,
					Name:               "$113K Decide Your Ride Drawing",
					SiteID:             1,
					DrawingDescription: "$113K Decide Your Ride Drawing",
					DrawingMediaURL:    "",
					DrawingStartDate:   "2020-02-01T00:00:00",
					DrawingEndDate:     "2020-03-26T00:00:00",
					DrawingCheckInPeriods: &DrawingCheckInPeriods{
						DrawingCheckInPeriod: []*DrawingCheckInPeriod{
							{
								CheckInStartDate: "2020-03-26T16:00:00",
								CheckInEndDate:   "2020-03-26T19:59:00",
							},
						},
					},
					DrawingSummary:       nil,
					DrawingImagePath:     nil,
					DrawingQuickTag:      nil,
					DrawingQuickTimes:    nil,
					DrawingQuickCriteria: nil,
					DrawingBonusTag:      nil,
					DrawingRules:         nil,
					IconURL:              nil,
				},
				{
					DWID:               117,
					Name:               "FEB20 - Valentine's Day Feel The Love",
					SiteID:             1,
					DrawingDescription: "FEB20 - Valentine's Day Feel The Love",
					DrawingMediaURL:    "",
					DrawingStartDate:   "2020-02-14T00:00:00",
					DrawingEndDate:     "2020-02-14T00:00:00",
					DrawingCheckInPeriods: &DrawingCheckInPeriods{
						DrawingCheckInPeriod: []*DrawingCheckInPeriod{
							{
								CheckInStartDate: "2020-02-14T18:00:00",
								CheckInEndDate:   "2020-02-14T20:00:00",
							},
						},
					},
					DrawingSummary:       nil,
					DrawingImagePath:     nil,
					DrawingQuickTag:      nil,
					DrawingQuickTimes:    nil,
					DrawingQuickCriteria: nil,
					DrawingBonusTag:      nil,
					DrawingRules:         nil,
					IconURL:              nil,
				},
			},
		}
		err error
	)

	httpmock.Activate()
	defer httpmock.Deactivate()

	stringResponder := httpmock.NewStringResponder(http.StatusOK, drawingsResponse)

	httpmock.RegisterResponder("GET", fmt.Sprintf("http://%s/mgt/restservices/drawings", "localhost"), stringResponder)

	drawings, err = client.GetDrawings("fakeUUID")

	assert.NoError(t, err, "should not have got an error")

	assert.Equal(t, expectedDrawings, drawings)
}

func TestGetPlayerDrawings(t *testing.T) {
	var (
		client           = NewClient(http.DefaultClient, "http://localhost", "fakelicense")
		drawingsResponse = `{"playerId":"108427","playerDrawings":[{"checkInAvailable":false,"checkInStartTime":"1900-01-01T00:00:00","name":"Six Figure Snowflakes","checkedIn":false,"DWID":111,"checkInEndTime":"1900-01-01T00:00:00","numberOfEntries":0,"drawingState":0,"IconURL":"","checkInDev":"","checkInTime":"0001-01-01T00:00:00","description":"Six Figure Snowflakes"},{"checkInAvailable":true,"checkInStartTime":"2020-01-09T19:27:41","name":"$47K Hot Seats & Slot Tourney","checkedIn":false,"DWID":115,"checkInEndTime":"2020-01-09T19:09:00.187","numberOfEntries":0,"drawingState":1,"IconURL":"","checkInDev":"","checkInTime":"0001-01-01T00:00:00","description":"$47K Hot Seats & Slot Tourney"}]}`
		playerID         = "6"
		drawings         []*PlayerDrawing
		expectedDrawings = []*PlayerDrawing{
			{
				CheckInAvailable: false,
				CheckInStartTime: "1900-01-01T00:00:00",
				Name:             "Six Figure Snowflakes",
				CheckedIn:        false,
				DWID:             111,
				CheckInEndTime:   "1900-01-01T00:00:00",
				NumberOfEntries:  0,
				DrawingState:     0,
				IconURL:          "",
				CheckInDev:       "",
				CheckInTime:      "0001-01-01T00:00:00",
				Description:      "Six Figure Snowflakes",
			},
			{
				CheckInAvailable: true,
				CheckInStartTime: "2020-01-09T19:27:41",
				Name:             "$47K Hot Seats & Slot Tourney",
				CheckedIn:        false,
				DWID:             115,
				CheckInEndTime:   "2020-01-09T19:09:00.187",
				NumberOfEntries:  0,
				DrawingState:     1,
				IconURL:          "",
				CheckInDev:       "",
				CheckInTime:      "0001-01-01T00:00:00",
				Description:      "$47K Hot Seats & Slot Tourney",
			},
		}
		err error
	)

	httpmock.Activate()
	defer httpmock.Deactivate()

	stringResponder := httpmock.NewStringResponder(http.StatusOK, drawingsResponse)

	httpmock.RegisterResponder("GET", fmt.Sprintf("http://%s/mgt/restservices/player/%s/drawings", "localhost", playerID), stringResponder)

	drawings, err = client.GetPlayerDrawings(playerID, "fakeUUID")

	assert.NoError(t, err, "should not have got an error")

	assert.Equal(t, expectedDrawings, drawings)
}
