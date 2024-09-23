/*
Djangolang

Testing TriggerHasExecutionAPIService

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

func Test_api_client_TriggerHasExecutionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TriggerHasExecutionAPIService DeleteTriggerHasExecution", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		httpRes, err := apiClient.TriggerHasExecutionAPI.DeleteTriggerHasExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerHasExecutionAPIService GetTriggerHasExecution", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.TriggerHasExecutionAPI.GetTriggerHasExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerHasExecutionAPIService GetTriggerHasExecutions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TriggerHasExecutionAPI.GetTriggerHasExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerHasExecutionAPIService PatchTriggerHasExecution", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.TriggerHasExecutionAPI.PatchTriggerHasExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TriggerHasExecutionAPIService PostTriggerHasExecutions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TriggerHasExecutionAPI.PostTriggerHasExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
