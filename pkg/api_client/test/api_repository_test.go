/*
Djangolang

Testing RepositoryAPIService

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

func Test_api_client_RepositoryAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RepositoryAPIService DeleteRepository", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		httpRes, err := apiClient.RepositoryAPI.DeleteRepository(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService GetRepositories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoryAPI.GetRepositories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService GetRepository", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.RepositoryAPI.GetRepository(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService PatchRepository", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var primaryKey string

		resp, httpRes, err := apiClient.RepositoryAPI.PatchRepository(context.Background(), primaryKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RepositoryAPIService PostRepositories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RepositoryAPI.PostRepositories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}