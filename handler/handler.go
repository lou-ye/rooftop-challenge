package handler

import (
	"rooftop-challenge/api"
	"rooftop-challenge/models"
	"strings"
)

type Handler struct {
	RooftopApi api.RooftopApi
	Mocked     func() string
}

func (handler Handler) HandleRequest() (bool, error) {
	token, err := handler.RooftopApi.GetToken()
	if err != nil {
		return false, err
	}

	data, err := handler.RooftopApi.GetData(token)
	if err != nil {
		return false, err
	}

	block, err := handler.check(data, token)
	if err != nil {
		return false, err
	}

	result, err := handler.RooftopApi.CheckBlock(models.CheckRequest{Encoded: handler.concatenateBlocks(block)}, token)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (handler Handler) check(data []string, token string) ([]string, error) {
	blocks := []string{data[0]}

	for range data {
		for _, block := range data {
			result, err := handler.areSequential(models.CheckRequest{Blocks: []string{blocks[len(blocks)-1], block}}, token)
			if err != nil {
				return nil, err
			}

			if result {
				blocks = append(blocks, block)
				break
			}
		}
	}

	return blocks, nil
}

func (handler Handler) areSequential(request models.CheckRequest, token string) (bool, error) {
	if handler.Mocked != nil {
		return handler.RooftopApi.MockCheckBlock(request)
	}

	return handler.RooftopApi.CheckBlock(request, token)
}

func (handler Handler) concatenateBlocks(block []string) string {
	return strings.Join(block[:], "")
}
