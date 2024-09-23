/*
Djangolang

Testing TriggerAPIService

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

func Test_api_client_TriggerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TriggerAPIService DeleteTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		httpRes, err := apiClient.TriggerAPI.DeleteTrigger(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerAPIService GetTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.TriggerAPI.GetTrigger(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerAPIService GetTriggers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TriggerAPI.GetTriggers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerAPIService PatchTrigger", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.TriggerAPI.PatchTrigger(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerAPIService PostTriggers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TriggerAPI.PostTriggers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
