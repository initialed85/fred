# \JobCoordinatorClaimChangeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostJobCoordinatorClaimChanges**](JobCoordinatorClaimChangeAPI.md#PostJobCoordinatorClaimChanges) | **Post** /api/job-coordinator-claim-change | 



## PostJobCoordinatorClaimChanges

> ResponseWithGenericOfChange PostJobCoordinatorClaimChanges(ctx).ChangeJobCoordinatorClaimRequest(changeJobCoordinatorClaimRequest).Execute()



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
	changeJobCoordinatorClaimRequest := *openapiclient.NewChangeJobCoordinatorClaimRequest() // ChangeJobCoordinatorClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobCoordinatorClaimChangeAPI.PostJobCoordinatorClaimChanges(context.Background()).ChangeJobCoordinatorClaimRequest(changeJobCoordinatorClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobCoordinatorClaimChangeAPI.PostJobCoordinatorClaimChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostJobCoordinatorClaimChanges`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `JobCoordinatorClaimChangeAPI.PostJobCoordinatorClaimChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobCoordinatorClaimChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeJobCoordinatorClaimRequest** | [**ChangeJobCoordinatorClaimRequest**](ChangeJobCoordinatorClaimRequest.md) |  | 

### Return type

[**ResponseWithGenericOfChange**](ResponseWithGenericOfChange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

