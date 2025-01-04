# \JobExecutorClaimExecutionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostJobExecutorClaimExecutions**](JobExecutorClaimExecutionAPI.md#PostJobExecutorClaimExecutions) | **Post** /api/job-executor-claim-execution | 



## PostJobExecutorClaimExecutions

> ResponseWithGenericOfExecution PostJobExecutorClaimExecutions(ctx).ExecutionJobExecutorClaimRequest(executionJobExecutorClaimRequest).Execute()



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
	executionJobExecutorClaimRequest := *openapiclient.NewExecutionJobExecutorClaimRequest() // ExecutionJobExecutorClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobExecutorClaimExecutionAPI.PostJobExecutorClaimExecutions(context.Background()).ExecutionJobExecutorClaimRequest(executionJobExecutorClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobExecutorClaimExecutionAPI.PostJobExecutorClaimExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostJobExecutorClaimExecutions`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `JobExecutorClaimExecutionAPI.PostJobExecutorClaimExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobExecutorClaimExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionJobExecutorClaimRequest** | [**ExecutionJobExecutorClaimRequest**](ExecutionJobExecutorClaimRequest.md) |  | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

