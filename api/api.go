package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	e "rooftop-challenge/error"
	"rooftop-challenge/models"
)

const (
	baseUrl = "https://rooftop-career-switch.herokuapp.com"
	mail    = "usuario@gmail.com"
)

type RooftopApi interface {
	GetToken() (string, error)
	GetData(token string) ([]string, error)
	CheckBlock(request models.CheckRequest, token string) (bool, error)
	MockCheckBlock(request models.CheckRequest) (bool, error)
}

type RooftopApiImpl struct{}

func (r RooftopApiImpl) GetToken() (string, error) {
	resp, err := http.Get(baseUrl + "/token?email=" + mail)
	if err != nil {
		return "", errors.New(e.ErrorGettingToken)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var token models.Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", errors.New(e.ErrorDecodingResponse)
	}

	return token.Token, nil
}

func (r RooftopApiImpl) GetData(token string) ([]string, error) {
	resp, err := http.Get(baseUrl + "/blocks?token=" + token)
	if err != nil {
		return nil, errors.New(e.ErrorGettingData)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var blocks models.Blocks
	err = json.NewDecoder(resp.Body).Decode(&blocks)
	if err != nil {
		return nil, errors.New(e.ErrorDecodingResponse)
	}

	return blocks.Data, nil
}

func (r RooftopApiImpl) CheckBlock(request models.CheckRequest, token string) (bool, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return false, errors.New(e.ErrorMarshalingData)
	}

	resp, err := http.Post(baseUrl+"/check?token="+token, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, errors.New(e.ErrorCheckingBlock)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, err
	}

	var message models.CheckResponse
	err = json.NewDecoder(resp.Body).Decode(&message)
	if err != nil {
		return false, errors.New(e.ErrorDecodingResponse)
	}

	return message.Message, nil
}

func (r RooftopApiImpl) MockCheckBlock(request models.CheckRequest) (bool, error) {
	blockA := request.Blocks[0]
	blockB := request.Blocks[1]

	if blockA == "qwer" && blockB == "asdf" {
		return true, nil
	} else if blockA == "asdf" && blockB == "zcvf" {
		return true, nil
	} else {
		return false, nil
	}
}
