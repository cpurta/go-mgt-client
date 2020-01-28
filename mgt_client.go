package gomgtclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Client = &mgtClient{}

const (
	basePath = "/mgt/restservices/"
)

type mgtClient struct {
	httpClient *http.Client
	mgtHost    string
	baseHeader http.Header
}

func NewClient(httpClient *http.Client, host, license string) *mgtClient {
	return &mgtClient{
		httpClient: httpClient,
		mgtHost:    host,
		baseHeader: http.Header{
			"Accept":       []string{"application/json"},
			"Content-Type": []string{"application/json"},
			"licenseKey":   []string{license},
		},
	}
}

func (client *mgtClient) makeRequest(path, uuid, method, requestBody string) (*http.Request, error) {
	var (
		url  = fmt.Sprintf("%s%s%s", client.mgtHost, basePath, path)
		body io.Reader
		r    *http.Request
		err  error
	)

	if requestBody != "" {
		body = strings.NewReader(requestBody)
	}

	if r, err = http.NewRequest(method, url, body); err != nil {
		return nil, err
	}

	r.Header = client.baseHeader

	r.Header.Add("X-UUID", uuid)

	if body == nil && method == http.MethodPut {
		r.Header.Add("Content-Length", "0")
	}

	return r, nil
}

func (client *mgtClient) GetOffers(playerID, uuid string) ([]*Offer, error) {
	var (
		path     = fmt.Sprintf("player/%s/offers", playerID)
		request  *http.Request
		response *http.Response
		body     []byte
		offers   = make([]*Offer, 0)
		err      error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodPut, ""); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &offers); err != nil {
		return nil, err
	}

	return offers, nil
}

func (client *mgtClient) PlayerDrawingCheckIn(playerID, drawingID, uuid string) (*CheckIn, error) {
	var (
		path        = fmt.Sprintf("player/%s/drawing/%s/checkin", playerID, drawingID)
		request     *http.Request
		response    *http.Response
		requestBody = fmt.Sprintf(`{"playerId":"%s","drawingId":"%s"}`, playerID, drawingID)
		body        []byte
		checkIn     *CheckIn
		err         error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodPut, requestBody); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &checkIn); err != nil {
		return nil, err
	}

	return checkIn, nil
}

func (client *mgtClient) GetPromotions(uuid string) ([]*Promotion, error) {
	var (
		path          = "promotions"
		request       *http.Request
		response      *http.Response
		body          []byte
		promoResponse *AvailablePromotionsResponse
		err           error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodGet, ""); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &promoResponse); err != nil {
		return nil, err
	}

	return promoResponse.AvailablePromotions.Promotion, nil
}

func (client *mgtClient) GetPlayerPromotions(playerID, uuid string) ([]*PlayerPromotion, error) {
	var (
		path          = fmt.Sprintf("player/%s/promotions", playerID)
		request       *http.Request
		response      *http.Response
		body          []byte
		promoResponse *PlayerPromotionResponse
		err           error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodGet, ""); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &promoResponse); err != nil {
		return nil, err
	}

	return promoResponse.PlayerPromotions, nil
}

func (client *mgtClient) GetDrawings(uuid string) (*AvailableDrawings, error) {
	var (
		path     = "drawings"
		request  *http.Request
		response *http.Response
		body     []byte
		drawing  = &DrawingResponse{
			AvailableDrawings: &AvailableDrawings{
				Drawings: make([]*Drawing, 0),
			},
		}
		err error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodGet, ""); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &drawing); err != nil {
		return nil, err
	}

	return drawing.AvailableDrawings, nil
}

func (client *mgtClient) GetPlayerDrawings(playerID, uuid string) ([]*PlayerDrawing, error) {
	var (
		path     = fmt.Sprintf("player/%s/drawings", playerID)
		request  *http.Request
		response *http.Response
		body     []byte
		drawing  *PlayerDrawingResponse
		err      error
	)

	if request, err = client.makeRequest(path, uuid, http.MethodGet, ""); err != nil {
		return nil, err
	}

	if response, err = client.httpClient.Do(request); err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &drawing); err != nil {
		return nil, err
	}

	return drawing.PlayerDrawings, nil
}
