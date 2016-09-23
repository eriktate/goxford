package goxford

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Client is a collection of API structs that can be used to communicate with different Project Oxford APIs.
type Client struct {
	face *Face
}

//Face is a receiver for methods related to talking to the Project Oxford Face API.
type Face struct {
	key string
}

//InitFace prepares the Client to make calls to the Project Oxford Face API.
func (c *Client) InitFace(key string) {
	c.face = &Face{key: key}
}

//DetectURL calls the Project Oxford Face API to perform a detection using a URL to an image.
//TODO: Need to figure out if this is how I want to handle error responses.
func (c *Client) DetectURL(url, returnFaceAttributes string, returnFaceID, returnFaceLandmarks bool) (*FaceDetectResponse, error) {
	//TODO: Do this better...
	reqURL := fmt.Sprintf("https://api.projectoxford.ai/face/v1.0/detect?returnFaceId=%t&returnFaceLandmarks=%t&returnFaceAttributes%s",
		returnFaceID,
		returnFaceLandmarks,
		returnFaceAttributes)

	req, err := generateRequestWithPayload("POST", reqURL, c.face.key, []byte(fmt.Sprintf("{url:\"%s\"}", url)))

	if err != nil {
		return nil, err
	}

	//TODO: Might create a separate client from the default later.
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	detRes := &FaceDetectResponse{}

	err = decoder.Decode(detRes)

	if err != nil {
		return nil, err
	}

	return detRes, nil
}

func generateRequest(method, url, key string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)

	if err != nil {
		return r, err
	}

	r.Header.Add(AuthHeader, key)
	r.Header.Add("Content-Type", "application/json")

	return r, nil
}

func generateRequestWithPayload(method, url, key string, payload []byte) (*http.Request, error) {
	body := bytes.NewReader(payload)
	r, err := http.NewRequest(method, url, body)

	if err != nil {
		return r, err
	}

	r.Header.Add(AuthHeader, key)
	r.Header.Add("Content-Type", "application/json")

	return r, nil
}
