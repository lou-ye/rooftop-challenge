package handler

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"rooftop-challenge/api"
	e "rooftop-challenge/error"
	"rooftop-challenge/models"
	"testing"
)

const (
	blockA               = "qwer"
	blockB               = "zcvf"
	blockC               = "asdf"
	getTokenMethodName   = "GetToken"
	getDataMethodName    = "GetData"
	checkBlockMethodName = "CheckBlock"
)

func Test_WhenThereAreNoErrorsQueryingApi_TrueIsReturned(t *testing.T) {
	mockRooftopApi := new(MockRooftopApi)

	mockRooftopApi.On(getTokenMethodName).Return("", nil)
	mockRooftopApi.On(getDataMethodName, mock.Anything).Return(getData(), nil)
	mockRooftopApi.On(checkBlockMethodName, mock.Anything, mock.Anything).Return(true, nil)

	handler := Handler{RooftopApi: mockRooftopApi}

	result, err := handler.HandleRequest()

	assert.True(t, result)
	assert.Nil(t, err)
}

func Test_WhenGetTokenFails_ErrorIsReturned(t *testing.T) {
	mockRooftopApi := new(MockRooftopApi)

	mockRooftopApi.On(getTokenMethodName).Return(nil, errors.New(e.ErrorGettingToken))

	handler := Handler{RooftopApi: mockRooftopApi}

	_, err := handler.HandleRequest()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorGettingToken, err.Error())
}

func Test_WhenGetDataFails_ErrorIsReturned(t *testing.T) {
	mockRooftopApi := new(MockRooftopApi)

	mockRooftopApi.On(getTokenMethodName).Return("", nil)
	mockRooftopApi.On(getDataMethodName, mock.Anything).Return(nil, errors.New(e.ErrorGettingData))

	handler := Handler{RooftopApi: mockRooftopApi}

	_, err := handler.HandleRequest()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorGettingData, err.Error())
}

func Test_WhenCheckBlockFails_ErrorIsReturned(t *testing.T) {
	mockRooftopApi := new(MockRooftopApi)

	mockRooftopApi.On(getTokenMethodName).Return("", nil)
	mockRooftopApi.On(getDataMethodName, mock.Anything).Return(getData(), nil)
	mockRooftopApi.On(checkBlockMethodName, mock.Anything, mock.Anything).Return(nil, errors.New(e.ErrorCheckingBlock))

	handler := Handler{RooftopApi: mockRooftopApi}

	_, err := handler.HandleRequest()

	assert.NotNil(t, err)
	assert.Equal(t, e.ErrorCheckingBlock, err.Error())
}

func Test_GivenAnUnorderedSlice_ReturnsAnOrderedSlice(t *testing.T) {
	handler := Handler{RooftopApi: api.RooftopApiImpl{}, Mocked: func() string {
		return ""
	}}

	result, err := handler.check(getData(), "")

	expected := []string{blockA, blockC, blockB}

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_GivenAnUnorderedSliceuu_ReturnsAnOrderedSlice(t *testing.T) {
	handler := Handler{RooftopApi: api.RooftopApiImpl{}, Mocked: func() string {
		return ""
	}}

	result := handler.concatenateBlocks(getData())

	expected := blockA + blockB + blockC

	assert.Equal(t, expected, result)
}

func getData() []string {
	return []string{blockA, blockB, blockC}
}

type MockRooftopApi struct{ mock.Mock }

func (m *MockRooftopApi) GetToken() (string, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return "", args.Error(1)
	}

	return args.Get(0).(string), nil
}

func (m *MockRooftopApi) GetData(token string) ([]string, error) {
	args := m.Called(token)

	if args.Get(0) == nil {
		return []string{}, args.Error(1)
	}

	return args.Get(0).([]string), nil
}

func (m *MockRooftopApi) CheckBlock(request models.CheckRequest, token string) (bool, error) {
	args := m.Called(request, token)

	if args.Get(0) == nil {
		return true, args.Error(1)
	}

	return args.Get(0).(bool), nil
}

func (m *MockRooftopApi) MockCheckBlock(request models.CheckRequest) (bool, error) {
	args := m.Called(request)

	if args.Get(0) == nil {
		return true, args.Error(1)
	}

	return args.Get(0).(bool), nil
}
