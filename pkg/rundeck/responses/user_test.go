package responses

import (
	"encoding/json"
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck/responses/testdata"
	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestUserInfoResponse(t *testing.T) {
	obj := &UserProfileResponse{}
	data, dataErr := testdata.GetBytes(UserProfileResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}

func TestUsersInfoResponse(t *testing.T) {
	obj := &ListUsersResponse{}
	data, dataErr := testdata.GetBytes(ListUsersResponseTestFile)
	if dataErr != nil {
		t.Fatalf(dataErr.Error())
	}
	placeholder := make(map[string]interface{})
	_ = json.Unmarshal(data, &placeholder)
	config := newMSDecoderConfig()
	config.Result = obj
	decoder, newErr := mapstructure.NewDecoder(config)
	assert.NoError(t, newErr)
	dErr := decoder.Decode(placeholder)
	assert.NoError(t, dErr)
	assert.Implements(t, (*VersionedResponse)(nil), obj)
}
