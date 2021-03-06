// IMPORTANT DEV NOTES:
// - Still need to decide on a better way of setting up the PO endpoints. A lot of the URL is defined within the functions
//	 which means that a direct change to goxford.go is required if those endpoints ever change. Would probably be better/easier
//   to just build some config ahead of time that determines the URL, params, and method outside of this code.

package goxford

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Client is a collection of API structs that can be used to communicate with different Project Oxford APIs.
type Client struct {
	face *Face
}

// Face is a receiver for methods related to talking to the Project Oxford Face API.
type Face struct {
	key string
}

// InitFace prepares the Client to make calls to the Project Oxford Face API.
func (c *Client) InitFace(key string) {
	c.face = &Face{key: key}
}

// GetURL returns the base URL for the Project Oxford Face API.
func (f *Face) GetURL() string {
	return fmt.Sprintf("%s/face/v1.0", BaseURL)
}

// DetectURL calls the Project Oxford Face API to perform a detection using a URL to an image.
// TODO: Need to figure out if this is how I want to handle error responses.
func (c *Client) DetectURL(url, returnFaceAttributes string, returnFaceID, returnFaceLandmarks bool) ([]*FaceDetectResponse, error) {
	// TODO: Do this better...
	reqURL := fmt.Sprintf("%s/detect?returnFaceId=%t&returnFaceLandmarks=%t&returnFaceAttributes=%s",
		c.face.GetURL(),
		returnFaceID,
		returnFaceLandmarks,
		returnFaceAttributes)

	log.Printf("Request URL: %s", reqURL)

	req, err := generateRequestWithJSONPayload("POST", reqURL, c.face.key, []byte(fmt.Sprintf("{url:\"%s\"}", url)))

	if err != nil {
		return nil, err
	}

	// TODO: Might create a separate client from the default later.
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	detRes := []*FaceDetectResponse{}

	err = decoder.Decode(&detRes)

	if err != nil {
		return nil, err
	}

	return detRes, nil
}

// DetectPath calls the Project Oxford Face API to perform a detection using a local path to an image.
func (c *Client) DetectPath(path, returnFaceAttributes string, returnFaceID, returnFaceLandmarks bool) ([]*FaceDetectResponse, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return c.DetectBytes(data, returnFaceAttributes, returnFaceID, returnFaceLandmarks)
}

// DetectBytes calls the Project Oxford Face API to perform a detection using a slice of bytes representing an image.
func (c *Client) DetectBytes(image []byte, returnFaceAttributes string, returnFaceID, returnFaceLandmarks bool) ([]*FaceDetectResponse, error) {
	// TODO: Need to do this better.
	reqURL := fmt.Sprintf("%s/detect?returnFaceId=%t&returnFaceLandmarks=%t&returnFaceAttributes=%s",
		c.face.GetURL(),
		returnFaceID,
		returnFaceLandmarks,
		returnFaceAttributes)

	r, err := generateRequestWithOctetPayload("POST", reqURL, c.face.key, image)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(r)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)
	detRes := []*FaceDetectResponse{}
	err = decoder.Decode(&detRes)

	if err != nil {
		return nil, err
	}

	return detRes, nil
}

// CreatePersonGroup calls the Project Oxford Face API to create a new person group that persons can be added to.
func (c *Client) CreatePersonGroup(personGroupID, displayName, userData string) (bool, error) {
	reqURL := fmt.Sprintf("%s/personGroups/%s", c.face.GetURL(), strings.ToLower(personGroupID))
	payload := fmt.Sprintf("{\"name\":\"%s\",\"userData\":\"%s\"}", displayName, userData)

	r, err := generateRequestWithJSONPayload("PUT", reqURL, c.face.key, []byte(payload))

	if err != nil {
		return false, err
	}

	resp, err := http.DefaultClient.Do(r)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Failed to create PersonGroup. Status code: %d", resp.StatusCode)
	}

	return true, nil
}

// GetPersonGroup calls the Project Oxford Face API to get a person group with the given personGroupID.
func (c *Client) GetPersonGroup(personGroupID string) (*PersonGroup, error) {
	reqURL := fmt.Sprintf("%s/personGroups/%s", c.face.GetURL(), strings.ToLower(personGroupID))

	r, err := generateRequest("GET", reqURL, c.face.key)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	pgRes := &PersonGroup{}
	err = decoder.Decode(pgRes)

	if err != nil {
		return nil, err
	}

	return pgRes, nil
}

// CreatePerson calls the Project Oxford Face API to create a person within an existing personGroup.
func (c *Client) CreatePerson(personGroupID, name, userData string) (*Person, error) {
	reqURL := fmt.Sprintf("%s/personGroups/%s/persons", c.face.GetURL(), strings.ToLower(personGroupID))
	payload := fmt.Sprintf("{\"name\":\"%s\",\"userData\":\"%s\"}", name, userData)
	r, err := generateRequestWithJSONPayload("POST", reqURL, c.face.key, []byte(payload))

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	pRes := &Person{}
	err = decoder.Decode(pRes)

	return pRes, err
}

// GetPerson calls the Project Oxford Face API to get a person within an existing personGroup with the given personID.
func (c *Client) GetPerson(personGroupID, personID string) (*Person, error) {
	reqURL := fmt.Sprintf("%s/personGroups/%s/persons/%s", c.face.GetURL(), strings.ToLower(personGroupID), personID)
	r, err := generateRequest("GET", reqURL, c.face.key)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(res.Body)
	pRes := &Person{}
	err = decoder.Decode(pRes)

	return pRes, err
}

func generateRequest(method, url, key string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)

	if err != nil {
		return r, err
	}

	r.Header.Add(AuthHeader, key)

	return r, nil
}

func generateRequestWithJSONPayload(method, url, key string, payload []byte) (*http.Request, error) {
	body := bytes.NewReader(payload)
	r, err := http.NewRequest(method, url, body)

	if err != nil {
		return r, err
	}

	r.Header.Add(AuthHeader, key)
	r.Header.Add("Content-Type", "application/json")

	return r, nil
}

func generateRequestWithOctetPayload(method, url, key string, payload []byte) (*http.Request, error) {
	body := bytes.NewReader(payload)
	r, err := http.NewRequest(method, url, body)

	if err != nil {
		return r, err
	}

	r.Header.Add(AuthHeader, key)
	r.Header.Add("Content-Type", "application/octet-stream")

	return r, nil
}
