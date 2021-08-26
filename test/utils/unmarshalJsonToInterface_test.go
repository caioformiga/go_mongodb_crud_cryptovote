package test

import (
	"encoding/json"
	"testing"

	"github.com/caioformiga/go_mongodb_crud_cryptovote/model"
	"github.com/caioformiga/go_mongodb_crud_cryptovote/utils"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJsonToCryptoVoteInvalidData(t *testing.T) {
	var ptr *[]model.CryptoVote
	err := json.Unmarshal(utils.JsonBadData, &ptr)
	assert.NotNil(t, err, "err should not be nil")
}

func TestUnmarshalJsonToCryptoVoteValidData(t *testing.T) {
	var ptr *[]model.CryptoVote
	err := json.Unmarshal(utils.JsonInData, &ptr)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, ptr, "ptr should not nil")
}
