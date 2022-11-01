package services

import (
	"github.com/golang/mock/gomock"
	"github.com/hypertec-cloud/go-hci/api"
	"github.com/hypertec-cloud/go-hci/mocks"
	"github.com/hypertec-cloud/go-hci/mocks/api_mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TEST_TASK_ID = "test_task_id"
)

func TestGetTaskReturnTaskIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHciClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockHciClient,
	}

	expectedTask := Task{
		Id:      TEST_TASK_ID,
		Status:  "SUCCESS",
		Created: "2015-07-07",
		Result:  []byte(`{"key": "value"}`),
	}

	mockHciClient.EXPECT().Do(api.HciRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(&api.HciResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"SUCCESS", "created":"2015-07-07", "result":{"key": "value"}}`),
	}, nil)

	//when
	task, _ := taskService.Get(TEST_TASK_ID)

	//then
	assert.Equal(t, expectedTask, *task)
}

func TestGetTaskReturnErrorIfHasHciErrors(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHciClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockHciClient,
	}

	hciResponse := api.HciResponse{
		StatusCode: 400,
		Errors:     []api.HciError{{}},
	}
	mockHciClient.EXPECT().Do(api.HciRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(&hciResponse, nil)

	//when
	task, err := taskService.Get(TEST_TASK_ID)

	//then
	assert.Nil(t, task)
	assert.Equal(t, api.HciErrorResponse(hciResponse), err)
}

func TestGetTaskReturnErrorIfHasUnexpectedErrors(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHciClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockHciClient,
	}

	mockError := mocks.MockError{"some_get_task_error"}

	mockHciClient.EXPECT().Do(api.HciRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}).Return(nil, mockError)

	//when
	task, err := taskService.Get(TEST_TASK_ID)

	//then
	assert.Nil(t, task)
	assert.Equal(t, mockError, err)
}

func TestPollingReturnTaskResultOnSuccessfulComplete(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHciClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockHciClient,
	}

	request := api.HciRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}

	expectedResult := []byte(`{"foo":"bar"}`)

	pendingResponse := &api.HciResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"PENDING", "created":"2015-07-07"}`),
	}
	successResponse := &api.HciResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"SUCCESS", "created":"2015-07-07", "result":` + string(expectedResult) + `}`),
	}
	gomock.InOrder(
		mockHciClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockHciClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockHciClient.EXPECT().Do(request).Return(successResponse, nil),
	)

	//when
	result, _ := taskService.Poll(TEST_TASK_ID, 10)

	//then
	if assert.NotNil(t, result) {
		assert.Equal(t, expectedResult, result)
	}
}

func TestPollingGetErrorOnTaskFailure(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHciClient := api_mocks.NewMockApiClient(ctrl)

	taskService := TaskApi{
		apiClient: mockHciClient,
	}

	request := api.HciRequest{
		Method:   api.GET,
		Endpoint: "tasks/" + TEST_TASK_ID,
	}

	expectedResult := []byte(`{"foo":"bar"}`)

	pendingResponse := &api.HciResponse{
		StatusCode: 200,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"PENDING", "created":"2015-07-07"}`),
	}
	failedResponse := &api.HciResponse{
		StatusCode: 400,
		Data:       []byte(`{"id":"` + TEST_TASK_ID + `", "status":"FAILED", "created":"2015-07-07", "result":` + string(expectedResult) + `}`),
	}
	gomock.InOrder(
		mockHciClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockHciClient.EXPECT().Do(request).Return(pendingResponse, nil),
		mockHciClient.EXPECT().Do(request).Return(failedResponse, nil),
	)

	//when
	_, err := taskService.Poll(TEST_TASK_ID, 10)

	//then
	assert.NotNil(t, err)

}
