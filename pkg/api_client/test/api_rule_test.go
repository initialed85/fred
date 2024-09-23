/*
Djangolang

Testing RuleAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_client

import (
	"context"
	"testing"

	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_client_RuleAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RuleAPIService DeleteRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		httpRes, err := apiClient.RuleAPI.DeleteRule(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RuleAPIService GetRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.RuleAPI.GetRule(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RuleAPIService GetRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RuleAPI.GetRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RuleAPIService PatchRule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.RuleAPI.PatchRule(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RuleAPIService PostRules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RuleAPI.PostRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
