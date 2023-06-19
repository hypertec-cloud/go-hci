package hci

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hypertec-cloud/go-hci/api"
	"github.com/hypertec-cloud/go-hci/mocks"
	"github.com/hypertec-cloud/go-hci/mocks/services_mocks"
	"github.com/stretchr/testify/assert"
)

const (
	TEST_BAREMETAL_ID                    = "test_baremetal_id"
	TEST_BAREMETAL_NAME                  = "test_baremetal"
	TEST_BAREMETAL_STATE                 = "test_baremetal_state"
	TEST_BAREMETAL_TEMPLATE_ID           = "test_baremetal_template_id"
	TEST_BAREMETAL_TEMPLATE_NAME         = "test_baremetal_template_name"
	TEST_BAREMETAL_IS_PASSWORD_ENABLED   = true
	TEST_BAREMETAL_IS_SSH_KEY_ENABLED    = false
	TEST_BAREMETAL_USERNAME              = "test_baremetal_username"
	TEST_BAREMETAL_COMPUTE_OFFERING_ID   = "test_baremetal_compute_offering_id"
	TEST_BAREMETAL_COMPUTE_OFFERING_NAME = "test_baremetal_compute_offering_name"
	TEST_BAREMETAL_ZONE_ID               = "test_baremetal_zone_id"
	TEST_BAREMETAL_ZONE_NAME             = "test_baremetal_zone_name"
	TEST_BAREMETAL_PROJECT_ID            = "test_baremetal_project_id"
	TEST_BAREMETAL_NETWORK_ID            = "test_baremetal_network_id"
	TEST_BAREMETAL_NETWORK_NAME          = "test_baremetal_network_name"
	TEST_BAREMETAL_VPC_ID                = "test_baremetal_vpc_id"
	TEST_BAREMETAL_VPC_NAME              = "test_baremetal_vpc_name"
	TEST_BAREMETAL_MAC_ADDRESS           = "test_baremetal_mac_address"
	TEST_BAREMETAL_IP_ADDRESS            = "test_baremetal_ip_address"
	TEST_BAREMETAL_VOLUME_ID_TO_ATTACH   = "test_volume_id_to_attach"
	TEST_BAREMETAL_USER_DATA             = "test_baremetal_user_data"
	TEST_BAREMETAL_PUBLIC_KEY            = "test_baremetal_public_key"
)

func buildTestBaremetalJsonResponse(baremetal *Baremetal) []byte {
	return []byte(`{"id": "` + baremetal.Id + `", ` +
		`"name":"` + baremetal.Name + `", ` +
		`"state":"` + baremetal.State + `", ` +
		`"templateId":"` + baremetal.TemplateId + `", ` +
		`"templateName":"` + baremetal.TemplateName + `", ` +
		`"isPasswordEnabled":` + strconv.FormatBool(baremetal.IsPasswordEnabled) + `, ` +
		`"isSshKeyEnabled":` + strconv.FormatBool(baremetal.IsSSHKeyEnabled) + `, ` +
		`"username":"` + baremetal.Username + `", ` +
		`"computeOfferingId":"` + baremetal.ComputeOfferingId + `", ` +
		`"computeOfferingName":"` + baremetal.ComputeOfferingName + `", ` +
		`"zoneId":"` + baremetal.ZoneId + `", ` +
		`"zoneName":"` + baremetal.ZoneName + `", ` +
		`"projectId":"` + baremetal.ProjectId + `", ` +
		`"networkId":"` + baremetal.NetworkId + `", ` +
		`"networkName":"` + baremetal.NetworkName + `", ` +
		`"vpcId":"` + baremetal.VpcId + `", ` +
		`"vpcName":"` + baremetal.VpcName + `", ` +
		`"macAddress":"` + baremetal.MacAddress + `", ` +
		`"ipAddress":"` + baremetal.IpAddress + `", ` +
		`"volumeIdToAttach":"` + baremetal.VolumeIdToAttach + `", ` +
		`"publicKey":"` + baremetal.PublicKey + `", ` +
		`"userData":"` + baremetal.UserData + `"}`)
}

func buildListTestBaremetalJsonResponse(baremetals []Baremetal) []byte {
	resp := `[`
	for i, inst := range baremetals {
		resp += string(buildTestBaremetalJsonResponse(&inst))
		if i != len(baremetals)-1 {
			resp += `,`
		}
	}
	resp += `]`
	return []byte(resp)
}

func TestGetBaremetalReturnBaremetalIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	expectedBaremetal := Baremetal{Id: TEST_BAREMETAL_ID,
		Name:                TEST_BAREMETAL_NAME,
		State:               TEST_BAREMETAL_STATE,
		TemplateId:          TEST_BAREMETAL_TEMPLATE_ID,
		TemplateName:        TEST_BAREMETAL_TEMPLATE_NAME,
		IsPasswordEnabled:   TEST_BAREMETAL_IS_PASSWORD_ENABLED,
		IsSSHKeyEnabled:     TEST_BAREMETAL_IS_SSH_KEY_ENABLED,
		Username:            TEST_BAREMETAL_USERNAME,
		ComputeOfferingId:   TEST_BAREMETAL_COMPUTE_OFFERING_ID,
		ComputeOfferingName: TEST_BAREMETAL_COMPUTE_OFFERING_NAME,
		ZoneId:              TEST_BAREMETAL_ZONE_ID,
		ZoneName:            TEST_BAREMETAL_ZONE_NAME,
		ProjectId:           TEST_BAREMETAL_PROJECT_ID,
		NetworkId:           TEST_BAREMETAL_NETWORK_ID,
		NetworkName:         TEST_BAREMETAL_NETWORK_NAME,
		VpcId:               TEST_BAREMETAL_VPC_ID,
		VpcName:             TEST_BAREMETAL_VPC_NAME,
		MacAddress:          TEST_BAREMETAL_MAC_ADDRESS,
		IpAddress:           TEST_BAREMETAL_IP_ADDRESS,
		VolumeIdToAttach:    TEST_BAREMETAL_VOLUME_ID_TO_ATTACH,
		PublicKey:           TEST_BAREMETAL_PUBLIC_KEY,
		UserData:            TEST_BAREMETAL_USER_DATA}

	mockEntityService.EXPECT().Get(TEST_BAREMETAL_ID, gomock.Any()).Return(buildTestBaremetalJsonResponse(&expectedBaremetal), nil)

	//when
	baremetal, _ := baremetalService.Get(TEST_BAREMETAL_ID)

	//then
	if assert.NotNil(t, baremetal) {
		assert.Equal(t, expectedBaremetal, *baremetal)
	}
}

func TestGetBaremetalReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_get_error"}

	mockEntityService.EXPECT().Get(TEST_BAREMETAL_ID, gomock.Any()).Return(nil, mockError)

	//when
	baremetal, err := baremetalService.Get(TEST_BAREMETAL_ID)

	//then
	assert.Nil(t, baremetal)
	assert.Equal(t, mockError, err)

}

func TestListBaremetalReturnBaremetalsIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	expectedBaremetal1 := Baremetal{Id: "list_id_1",
		Name:                "list_name_1",
		State:               "list_state_1",
		TemplateId:          "list_template_id_1",
		TemplateName:        "list_template_name_1",
		IsPasswordEnabled:   false,
		IsSSHKeyEnabled:     true,
		Username:            "list_username_1",
		ComputeOfferingId:   "list_compute_offering_id_1",
		ComputeOfferingName: "list_compute_offering_name_1",
		ZoneId:              "list_zone_id_1",
		ZoneName:            "list_zone_name_1",
		ProjectId:           "list_project_id_1",
		NetworkId:           "list_network_id_1",
		NetworkName:         "list_network_name_1",
		VpcId:               "list_vpc_id_1",
		VpcName:             "list_vpc_name_1",
		MacAddress:          "list_mac_address_1",
		VolumeIdToAttach:    "list_volume_id_to_attach_1",
		IpAddress:           "list_ip_address_1",
		PublicKey:           "list_public_key_1",
		UserData:            "list_user_data_1"}

	expectedBaremetals := []Baremetal{expectedBaremetal1}

	mockEntityService.EXPECT().List(gomock.Any()).Return(buildListTestBaremetalJsonResponse(expectedBaremetals), nil)

	//when
	baremetals, _ := baremetalService.List()

	//then
	if assert.NotNil(t, baremetals) {
		assert.Equal(t, expectedBaremetals, baremetals)
	}
}

func TestListBaremetalReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_list_error"}

	mockEntityService.EXPECT().List(gomock.Any()).Return(nil, mockError)

	//when
	baremetals, err := baremetalService.List()

	//then
	assert.Nil(t, baremetals)
	assert.Equal(t, mockError, err)

}

func TestCreateBaremetalReturnCreatedBaremetalIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	baremetalToCreate := Baremetal{Id: "new_id",
		Name:              "new_name",
		TemplateId:        "templateId",
		ComputeOfferingId: "computeOfferingId",
		NetworkId:         "networkId"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return([]byte(`{"id":"new_id", "password": "new_password"}`), nil)

	//when
	createdBaremetal, _ := baremetalService.Create(baremetalToCreate)

	//then
	if assert.NotNil(t, createdBaremetal) {
		assert.Equal(t, "new_password", createdBaremetal.Password)
	}
}

func TestCreateBaremetalReturnNilWithErrorIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_create_baremetal_error"}

	mockEntityService.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, mockError)

	baremetalToCreate := Baremetal{Name: "new_name",
		TemplateId:        "templateId",
		ComputeOfferingId: "computeOfferingId",
		NetworkId:         "networkId"}

	//when
	createdBaremetal, err := baremetalService.Create(baremetalToCreate)

	//then
	assert.Nil(t, createdBaremetal)
	assert.Equal(t, mockError, err)

}

func TestStartBaremetalReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_START_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.Start(TEST_BAREMETAL_ID)

	//then
	assert.True(t, success)
}

func TestStartBaremetalReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_start_baremetal_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_START_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.Start(TEST_BAREMETAL_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestStopBaremetalReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_STOP_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.Stop(TEST_BAREMETAL_ID)

	//then
	assert.True(t, success)
}

func TestStopBaremetalReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_stop_baremetal_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_STOP_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.Stop(TEST_BAREMETAL_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)

}

func TestDestroyBaremetalReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)
	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_PURGE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.Destroy(TEST_BAREMETAL_ID)

	//then
	assert.True(t, success)
}

func TestDestroyBaremetalReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_destroy_baremetal_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_PURGE_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.Destroy(TEST_BAREMETAL_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestRecoverBaremetalReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_RECOVER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.Recover(TEST_BAREMETAL_ID)

	//then
	assert.True(t, success)
}

func TestRecoverBaremetalReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_recover_baremetal_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_RECOVER_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.Recover(TEST_BAREMETAL_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestRebootBaremetalReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_REBOOT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.Reboot(TEST_BAREMETAL_ID)

	//then
	assert.True(t, success)
}

func TestRebootBaremetalReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_reboot_baremetal_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_REBOOT_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.Reboot(TEST_BAREMETAL_ID)

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func BMTestAssociateSSHKeyReturnTrueIfSuccess(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_ASSOCIATE_SSH_KEY_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), nil)

	//when
	success, _ := baremetalService.AssociateSSHKey(TEST_BAREMETAL_ID, "new_ssh_key")

	//then
	assert.True(t, success)
}

func BMTestAssociateSSHKeyReturnFalseIfError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_associate_ssh_key_error"}
	mockEntityService.EXPECT().Execute(TEST_BAREMETAL_ID, BAREMETAL_ASSOCIATE_SSH_KEY_OPERATION, gomock.Any(), gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	success, err := baremetalService.AssociateSSHKey(TEST_BAREMETAL_ID, "new_ssh_key")

	//then
	assert.False(t, success)
	assert.Equal(t, mockError, err)
}

func TestExistsReturnTrueIfBaremetalExists(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockEntityService.EXPECT().Get(TEST_BAREMETAL_ID, gomock.Any()).Return([]byte(`{"id": "foo"}`), nil)

	//when
	exists, _ := baremetalService.Exists(TEST_BAREMETAL_ID)

	//then
	assert.True(t, exists)
}

func TestExistsReturnFalseIfBaremetalDoesntExist(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockApiError := api.HciErrorResponse(api.HciResponse{StatusCode: api.NOT_FOUND})
	mockEntityService.EXPECT().Get(TEST_BAREMETAL_ID, gomock.Any()).Return([]byte(`{}`), mockApiError)

	//when
	exists, err := baremetalService.Exists(TEST_BAREMETAL_ID)

	//then
	assert.Nil(t, err)
	assert.False(t, exists)
}

func BMTestExistsReturnErrorIfUnexpectedError(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEntityService := services_mocks.NewMockEntityService(ctrl)

	baremetalService := BaremetalApi{
		entityService: mockEntityService,
	}

	mockError := mocks.MockError{"some_exists_error"}
	mockEntityService.EXPECT().Get(TEST_BAREMETAL_ID, gomock.Any()).Return([]byte(`{}`), mockError)

	//when
	_, err := baremetalService.Exists(TEST_BAREMETAL_ID)

	//then
	assert.Equal(t, mockError, err)
}
