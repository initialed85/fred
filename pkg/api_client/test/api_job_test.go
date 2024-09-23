/*
Djangolang

Testing JobAPIService

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

func Test_api_client_JobAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JobAPIService DeleteJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		httpRes, err := apiClient.JobAPI.DeleteJob(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.JobAPI.GetJob(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.JobAPI.GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService PatchJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.JobAPI.PatchJob(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JobAPIService PostJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.JobAPI.PostJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
