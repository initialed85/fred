/*
Djangolang

Testing OutputAPIService

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

func Test_api_client_OutputAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OutputAPIService DeleteOutput", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		httpRes, err := apiClient.OutputAPI.DeleteOutput(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputAPIService GetOutput", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.OutputAPI.GetOutput(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputAPIService GetOutputs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OutputAPI.GetOutputs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputAPIService PatchOutput", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.OutputAPI.PatchOutput(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OutputAPIService PostOutputs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OutputAPI.PostOutputs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
