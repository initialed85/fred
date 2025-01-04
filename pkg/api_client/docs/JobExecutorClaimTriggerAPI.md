# \JobExecutorClaimTriggerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostJobExecutorClaimTriggers**](JobExecutorClaimTriggerAPI.md#PostJobExecutorClaimTriggers) | **Post** /api/job-executor-claim-trigger | 



## PostJobExecutorClaimTriggers

> ResponseWithGenericOfTrigger PostJobExecutorClaimTriggers(ctx).TriggerJobExecutorClaimRequest(triggerJobExecutorClaimRequest).Execute()



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
	triggerJobExecutorClaimRequest := *openapiclient.NewTriggerJobExecutorClaimRequest() // TriggerJobExecutorClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobExecutorClaimTriggerAPI.PostJobExecutorClaimTriggers(context.Background()).TriggerJobExecutorClaimRequest(triggerJobExecutorClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobExecutorClaimTriggerAPI.PostJobExecutorClaimTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostJobExecutorClaimTriggers`: ResponseWithGenericOfTrigger
	fmt.Fprintf(os.Stdout, "Response from `JobExecutorClaimTriggerAPI.PostJobExecutorClaimTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobExecutorClaimTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerJobExecutorClaimRequest** | [**TriggerJobExecutorClaimRequest**](TriggerJobExecutorClaimRequest.md) |  | 

### Return type

[**ResponseWithGenericOfTrigger**](ResponseWithGenericOfTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

