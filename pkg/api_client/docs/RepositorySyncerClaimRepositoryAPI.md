# \RepositorySyncerClaimRepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostRepositorySyncerClaimRepositories**](RepositorySyncerClaimRepositoryAPI.md#PostRepositorySyncerClaimRepositories) | **Post** /api/repository-syncer-claim-repository | 



## PostRepositorySyncerClaimRepositories

> ResponseWithGenericOfRepository PostRepositorySyncerClaimRepositories(ctx).RepositoryRepositorySyncerClaimRequest(repositoryRepositorySyncerClaimRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	repositoryRepositorySyncerClaimRequest := *openapiclient.NewRepositoryRepositorySyncerClaimRequest() // RepositoryRepositorySyncerClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositorySyncerClaimRepositoryAPI.PostRepositorySyncerClaimRepositories(context.Background()).RepositoryRepositorySyncerClaimRequest(repositoryRepositorySyncerClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositorySyncerClaimRepositoryAPI.PostRepositorySyncerClaimRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRepositorySyncerClaimRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositorySyncerClaimRepositoryAPI.PostRepositorySyncerClaimRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRepositorySyncerClaimRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryRepositorySyncerClaimRequest** | [**RepositoryRepositorySyncerClaimRequest**](RepositoryRepositorySyncerClaimRequest.md) |  | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

