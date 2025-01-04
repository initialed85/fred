# \TriggerProducerClaimChangeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostTriggerProducerClaimChanges**](TriggerProducerClaimChangeAPI.md#PostTriggerProducerClaimChanges) | **Post** /api/trigger-producer-claim-change | 



## PostTriggerProducerClaimChanges

> ResponseWithGenericOfChange PostTriggerProducerClaimChanges(ctx).ChangeTriggerProducerClaimRequest(changeTriggerProducerClaimRequest).Execute()



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
	changeTriggerProducerClaimRequest := *openapiclient.NewChangeTriggerProducerClaimRequest() // ChangeTriggerProducerClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TriggerProducerClaimChangeAPI.PostTriggerProducerClaimChanges(context.Background()).ChangeTriggerProducerClaimRequest(changeTriggerProducerClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TriggerProducerClaimChangeAPI.PostTriggerProducerClaimChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTriggerProducerClaimChanges`: ResponseWithGenericOfChange
	fmt.Fprintf(os.Stdout, "Response from `TriggerProducerClaimChangeAPI.PostTriggerProducerClaimChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTriggerProducerClaimChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeTriggerProducerClaimRequest** | [**ChangeTriggerProducerClaimRequest**](ChangeTriggerProducerClaimRequest.md) |  | 

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

