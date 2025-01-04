# \ChangeProducerClaimRepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostChangeProducerClaimRepositories**](ChangeProducerClaimRepositoryAPI.md#PostChangeProducerClaimRepositories) | **Post** /api/change-producer-claim-repository | 



## PostChangeProducerClaimRepositories

> ResponseWithGenericOfRepository PostChangeProducerClaimRepositories(ctx).RepositoryChangeProducerClaimRequest(repositoryChangeProducerClaimRequest).Execute()



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
	repositoryChangeProducerClaimRequest := *openapiclient.NewRepositoryChangeProducerClaimRequest() // RepositoryChangeProducerClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChangeProducerClaimRepositoryAPI.PostChangeProducerClaimRepositories(context.Background()).RepositoryChangeProducerClaimRequest(repositoryChangeProducerClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChangeProducerClaimRepositoryAPI.PostChangeProducerClaimRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostChangeProducerClaimRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `ChangeProducerClaimRepositoryAPI.PostChangeProducerClaimRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostChangeProducerClaimRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryChangeProducerClaimRequest** | [**RepositoryChangeProducerClaimRequest**](RepositoryChangeProducerClaimRequest.md) |  | 

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

