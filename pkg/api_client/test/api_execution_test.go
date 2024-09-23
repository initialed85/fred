/*
Djangolang

Testing ExecutionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_api_client_ExecutionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExecutionAPIService DeleteExecution", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		httpRes, err := apiClient.ExecutionAPI.DeleteExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionAPIService GetExecution", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.ExecutionAPI.GetExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionAPIService GetExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExecutionAPI.GetExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionAPIService PatchExecution", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.ExecutionAPI.PatchExecution(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExecutionAPIService PostExecutions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExecutionAPI.PostExecutions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}