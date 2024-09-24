# \ExecutionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExecution**](ExecutionAPI.md#DeleteExecution) | **Delete** /api/executions/{primaryKey} | 
[**GetExecution**](ExecutionAPI.md#GetExecution) | **Get** /api/executions/{primaryKey} | 
[**GetExecutions**](ExecutionAPI.md#GetExecutions) | **Get** /api/executions | 
[**PatchExecution**](ExecutionAPI.md#PatchExecution) | **Patch** /api/executions/{primaryKey} | 
[**PostExecutions**](ExecutionAPI.md#PostExecutions) | **Post** /api/executions | 



## DeleteExecution

> DeleteExecution(ctx, primaryKey).Depth(depth).Execute()



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
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExecutionAPI.DeleteExecution(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.DeleteExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> ResponseWithGenericOfExecution GetExecution(ctx, primaryKey).Depth(depth).Execute()



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
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecution(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecution`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutions

> ResponseWithGenericOfExecution GetExecutions(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).TriggerIdEq(triggerIdEq).TriggerIdNe(triggerIdNe).TriggerIdGt(triggerIdGt).TriggerIdGte(triggerIdGte).TriggerIdLt(triggerIdLt).TriggerIdLte(triggerIdLte).TriggerIdIn(triggerIdIn).TriggerIdNotin(triggerIdNotin).TriggerIdLike(triggerIdLike).TriggerIdNotlike(triggerIdNotlike).TriggerIdIlike(triggerIdIlike).TriggerIdNotilike(triggerIdNotilike).TriggerIdDesc(triggerIdDesc).TriggerIdAsc(triggerIdAsc).TriggerIdObjectDesc(triggerIdObjectDesc).TriggerIdObjectAsc(triggerIdObjectAsc).BuildOutputIdEq(buildOutputIdEq).BuildOutputIdNe(buildOutputIdNe).BuildOutputIdGt(buildOutputIdGt).BuildOutputIdGte(buildOutputIdGte).BuildOutputIdLt(buildOutputIdLt).BuildOutputIdLte(buildOutputIdLte).BuildOutputIdIn(buildOutputIdIn).BuildOutputIdNotin(buildOutputIdNotin).BuildOutputIdLike(buildOutputIdLike).BuildOutputIdNotlike(buildOutputIdNotlike).BuildOutputIdIlike(buildOutputIdIlike).BuildOutputIdNotilike(buildOutputIdNotilike).BuildOutputIdDesc(buildOutputIdDesc).BuildOutputIdAsc(buildOutputIdAsc).BuildOutputIdObjectDesc(buildOutputIdObjectDesc).BuildOutputIdObjectAsc(buildOutputIdObjectAsc).TestOutputIdEq(testOutputIdEq).TestOutputIdNe(testOutputIdNe).TestOutputIdGt(testOutputIdGt).TestOutputIdGte(testOutputIdGte).TestOutputIdLt(testOutputIdLt).TestOutputIdLte(testOutputIdLte).TestOutputIdIn(testOutputIdIn).TestOutputIdNotin(testOutputIdNotin).TestOutputIdLike(testOutputIdLike).TestOutputIdNotlike(testOutputIdNotlike).TestOutputIdIlike(testOutputIdIlike).TestOutputIdNotilike(testOutputIdNotilike).TestOutputIdDesc(testOutputIdDesc).TestOutputIdAsc(testOutputIdAsc).TestOutputIdObjectDesc(testOutputIdObjectDesc).TestOutputIdObjectAsc(testOutputIdObjectAsc).PublishOutputIdEq(publishOutputIdEq).PublishOutputIdNe(publishOutputIdNe).PublishOutputIdGt(publishOutputIdGt).PublishOutputIdGte(publishOutputIdGte).PublishOutputIdLt(publishOutputIdLt).PublishOutputIdLte(publishOutputIdLte).PublishOutputIdIn(publishOutputIdIn).PublishOutputIdNotin(publishOutputIdNotin).PublishOutputIdLike(publishOutputIdLike).PublishOutputIdNotlike(publishOutputIdNotlike).PublishOutputIdIlike(publishOutputIdIlike).PublishOutputIdNotilike(publishOutputIdNotilike).PublishOutputIdDesc(publishOutputIdDesc).PublishOutputIdAsc(publishOutputIdAsc).PublishOutputIdObjectDesc(publishOutputIdObjectDesc).PublishOutputIdObjectAsc(publishOutputIdObjectAsc).DeployOutputIdEq(deployOutputIdEq).DeployOutputIdNe(deployOutputIdNe).DeployOutputIdGt(deployOutputIdGt).DeployOutputIdGte(deployOutputIdGte).DeployOutputIdLt(deployOutputIdLt).DeployOutputIdLte(deployOutputIdLte).DeployOutputIdIn(deployOutputIdIn).DeployOutputIdNotin(deployOutputIdNotin).DeployOutputIdLike(deployOutputIdLike).DeployOutputIdNotlike(deployOutputIdNotlike).DeployOutputIdIlike(deployOutputIdIlike).DeployOutputIdNotilike(deployOutputIdNotilike).DeployOutputIdDesc(deployOutputIdDesc).DeployOutputIdAsc(deployOutputIdAsc).DeployOutputIdObjectDesc(deployOutputIdObjectDesc).DeployOutputIdObjectAsc(deployOutputIdObjectAsc).ValidateOutputIdEq(validateOutputIdEq).ValidateOutputIdNe(validateOutputIdNe).ValidateOutputIdGt(validateOutputIdGt).ValidateOutputIdGte(validateOutputIdGte).ValidateOutputIdLt(validateOutputIdLt).ValidateOutputIdLte(validateOutputIdLte).ValidateOutputIdIn(validateOutputIdIn).ValidateOutputIdNotin(validateOutputIdNotin).ValidateOutputIdLike(validateOutputIdLike).ValidateOutputIdNotlike(validateOutputIdNotlike).ValidateOutputIdIlike(validateOutputIdIlike).ValidateOutputIdNotilike(validateOutputIdNotilike).ValidateOutputIdDesc(validateOutputIdDesc).ValidateOutputIdAsc(validateOutputIdAsc).ValidateOutputIdObjectDesc(validateOutputIdObjectDesc).ValidateOutputIdObjectAsc(validateOutputIdObjectAsc).JobIdEq(jobIdEq).JobIdNe(jobIdNe).JobIdGt(jobIdGt).JobIdGte(jobIdGte).JobIdLt(jobIdLt).JobIdLte(jobIdLte).JobIdIn(jobIdIn).JobIdNotin(jobIdNotin).JobIdLike(jobIdLike).JobIdNotlike(jobIdNotlike).JobIdIlike(jobIdIlike).JobIdNotilike(jobIdNotilike).JobIdDesc(jobIdDesc).JobIdAsc(jobIdAsc).JobIdObjectDesc(jobIdObjectDesc).JobIdObjectAsc(jobIdObjectAsc).ReferencedByTriggerHasExecutionExecutionIdObjectsDesc(referencedByTriggerHasExecutionExecutionIdObjectsDesc).ReferencedByTriggerHasExecutionExecutionIdObjectsAsc(referencedByTriggerHasExecutionExecutionIdObjectsAsc).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := int32(56) // int32 | SQL LIMIT operator (optional)
	offset := int32(56) // int32 | SQL OFFSET operator (optional)
	depth := int32(56) // int32 | Max recursion depth for loading foreign objects; default = 1  (0 = recurse until graph cycle detected, 1 = this object only, 2 = this object + neighbours, 3 = this object + neighbours + their neighbours... etc) (optional)
	idEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	idNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	idGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	idGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	idLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	idLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	idIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	idNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	idLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	idDesc := "idDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	idAsc := "idAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	createdAtEq := time.Now() // time.Time | SQL = comparison (optional)
	createdAtNe := time.Now() // time.Time | SQL != comparison (optional)
	createdAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	createdAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	createdAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	createdAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	createdAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	createdAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	createdAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	createdAtDesc := "createdAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	createdAtAsc := "createdAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	updatedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	updatedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	updatedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	updatedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	updatedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	updatedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	updatedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	updatedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	updatedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	updatedAtDesc := "updatedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	updatedAtAsc := "updatedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deletedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	deletedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	deletedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	deletedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	deletedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	deletedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	deletedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	deletedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	deletedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deletedAtDesc := "deletedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deletedAtAsc := "deletedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	statusEq := "statusEq_example" // string | SQL = comparison (optional)
	statusNe := "statusNe_example" // string | SQL != comparison (optional)
	statusGt := "statusGt_example" // string | SQL > comparison, may not work with all column types (optional)
	statusGte := "statusGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	statusLt := "statusLt_example" // string | SQL < comparison, may not work with all column types (optional)
	statusLte := "statusLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	statusIn := "statusIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	statusNotin := "statusNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	statusLike := "statusLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotlike := "statusNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusIlike := "statusIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotilike := "statusNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusDesc := "statusDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	statusAsc := "statusAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	triggerIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	triggerIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	triggerIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	triggerIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	triggerIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	triggerIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	triggerIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	triggerIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	triggerIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	triggerIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	triggerIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	triggerIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	triggerIdDesc := "triggerIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	triggerIdAsc := "triggerIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	triggerIdObjectDesc := "triggerIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	triggerIdObjectAsc := "triggerIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	buildOutputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	buildOutputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	buildOutputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	buildOutputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	buildOutputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	buildOutputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	buildOutputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	buildOutputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	buildOutputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildOutputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildOutputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildOutputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildOutputIdDesc := "buildOutputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	buildOutputIdAsc := "buildOutputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	buildOutputIdObjectDesc := "buildOutputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	buildOutputIdObjectAsc := "buildOutputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	testOutputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	testOutputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	testOutputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	testOutputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	testOutputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	testOutputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	testOutputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	testOutputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	testOutputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testOutputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testOutputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testOutputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testOutputIdDesc := "testOutputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	testOutputIdAsc := "testOutputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	testOutputIdObjectDesc := "testOutputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	testOutputIdObjectAsc := "testOutputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	publishOutputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	publishOutputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	publishOutputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	publishOutputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	publishOutputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	publishOutputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	publishOutputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	publishOutputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	publishOutputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishOutputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishOutputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishOutputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishOutputIdDesc := "publishOutputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	publishOutputIdAsc := "publishOutputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	publishOutputIdObjectDesc := "publishOutputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	publishOutputIdObjectAsc := "publishOutputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deployOutputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	deployOutputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	deployOutputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	deployOutputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	deployOutputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	deployOutputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	deployOutputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	deployOutputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	deployOutputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployOutputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployOutputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployOutputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployOutputIdDesc := "deployOutputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deployOutputIdAsc := "deployOutputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deployOutputIdObjectDesc := "deployOutputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deployOutputIdObjectAsc := "deployOutputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	validateOutputIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	validateOutputIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	validateOutputIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	validateOutputIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	validateOutputIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	validateOutputIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	validateOutputIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	validateOutputIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	validateOutputIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateOutputIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateOutputIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateOutputIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateOutputIdDesc := "validateOutputIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	validateOutputIdAsc := "validateOutputIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	validateOutputIdObjectDesc := "validateOutputIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	validateOutputIdObjectAsc := "validateOutputIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	jobIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	jobIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	jobIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	jobIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	jobIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	jobIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	jobIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	jobIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	jobIdDesc := "jobIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdAsc := "jobIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdObjectDesc := "jobIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	jobIdObjectAsc := "jobIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerHasExecutionExecutionIdObjectsDesc := "referencedByTriggerHasExecutionExecutionIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTriggerHasExecutionExecutionIdObjectsAsc := "referencedByTriggerHasExecutionExecutionIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.GetExecutions(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).TriggerIdEq(triggerIdEq).TriggerIdNe(triggerIdNe).TriggerIdGt(triggerIdGt).TriggerIdGte(triggerIdGte).TriggerIdLt(triggerIdLt).TriggerIdLte(triggerIdLte).TriggerIdIn(triggerIdIn).TriggerIdNotin(triggerIdNotin).TriggerIdLike(triggerIdLike).TriggerIdNotlike(triggerIdNotlike).TriggerIdIlike(triggerIdIlike).TriggerIdNotilike(triggerIdNotilike).TriggerIdDesc(triggerIdDesc).TriggerIdAsc(triggerIdAsc).TriggerIdObjectDesc(triggerIdObjectDesc).TriggerIdObjectAsc(triggerIdObjectAsc).BuildOutputIdEq(buildOutputIdEq).BuildOutputIdNe(buildOutputIdNe).BuildOutputIdGt(buildOutputIdGt).BuildOutputIdGte(buildOutputIdGte).BuildOutputIdLt(buildOutputIdLt).BuildOutputIdLte(buildOutputIdLte).BuildOutputIdIn(buildOutputIdIn).BuildOutputIdNotin(buildOutputIdNotin).BuildOutputIdLike(buildOutputIdLike).BuildOutputIdNotlike(buildOutputIdNotlike).BuildOutputIdIlike(buildOutputIdIlike).BuildOutputIdNotilike(buildOutputIdNotilike).BuildOutputIdDesc(buildOutputIdDesc).BuildOutputIdAsc(buildOutputIdAsc).BuildOutputIdObjectDesc(buildOutputIdObjectDesc).BuildOutputIdObjectAsc(buildOutputIdObjectAsc).TestOutputIdEq(testOutputIdEq).TestOutputIdNe(testOutputIdNe).TestOutputIdGt(testOutputIdGt).TestOutputIdGte(testOutputIdGte).TestOutputIdLt(testOutputIdLt).TestOutputIdLte(testOutputIdLte).TestOutputIdIn(testOutputIdIn).TestOutputIdNotin(testOutputIdNotin).TestOutputIdLike(testOutputIdLike).TestOutputIdNotlike(testOutputIdNotlike).TestOutputIdIlike(testOutputIdIlike).TestOutputIdNotilike(testOutputIdNotilike).TestOutputIdDesc(testOutputIdDesc).TestOutputIdAsc(testOutputIdAsc).TestOutputIdObjectDesc(testOutputIdObjectDesc).TestOutputIdObjectAsc(testOutputIdObjectAsc).PublishOutputIdEq(publishOutputIdEq).PublishOutputIdNe(publishOutputIdNe).PublishOutputIdGt(publishOutputIdGt).PublishOutputIdGte(publishOutputIdGte).PublishOutputIdLt(publishOutputIdLt).PublishOutputIdLte(publishOutputIdLte).PublishOutputIdIn(publishOutputIdIn).PublishOutputIdNotin(publishOutputIdNotin).PublishOutputIdLike(publishOutputIdLike).PublishOutputIdNotlike(publishOutputIdNotlike).PublishOutputIdIlike(publishOutputIdIlike).PublishOutputIdNotilike(publishOutputIdNotilike).PublishOutputIdDesc(publishOutputIdDesc).PublishOutputIdAsc(publishOutputIdAsc).PublishOutputIdObjectDesc(publishOutputIdObjectDesc).PublishOutputIdObjectAsc(publishOutputIdObjectAsc).DeployOutputIdEq(deployOutputIdEq).DeployOutputIdNe(deployOutputIdNe).DeployOutputIdGt(deployOutputIdGt).DeployOutputIdGte(deployOutputIdGte).DeployOutputIdLt(deployOutputIdLt).DeployOutputIdLte(deployOutputIdLte).DeployOutputIdIn(deployOutputIdIn).DeployOutputIdNotin(deployOutputIdNotin).DeployOutputIdLike(deployOutputIdLike).DeployOutputIdNotlike(deployOutputIdNotlike).DeployOutputIdIlike(deployOutputIdIlike).DeployOutputIdNotilike(deployOutputIdNotilike).DeployOutputIdDesc(deployOutputIdDesc).DeployOutputIdAsc(deployOutputIdAsc).DeployOutputIdObjectDesc(deployOutputIdObjectDesc).DeployOutputIdObjectAsc(deployOutputIdObjectAsc).ValidateOutputIdEq(validateOutputIdEq).ValidateOutputIdNe(validateOutputIdNe).ValidateOutputIdGt(validateOutputIdGt).ValidateOutputIdGte(validateOutputIdGte).ValidateOutputIdLt(validateOutputIdLt).ValidateOutputIdLte(validateOutputIdLte).ValidateOutputIdIn(validateOutputIdIn).ValidateOutputIdNotin(validateOutputIdNotin).ValidateOutputIdLike(validateOutputIdLike).ValidateOutputIdNotlike(validateOutputIdNotlike).ValidateOutputIdIlike(validateOutputIdIlike).ValidateOutputIdNotilike(validateOutputIdNotilike).ValidateOutputIdDesc(validateOutputIdDesc).ValidateOutputIdAsc(validateOutputIdAsc).ValidateOutputIdObjectDesc(validateOutputIdObjectDesc).ValidateOutputIdObjectAsc(validateOutputIdObjectAsc).JobIdEq(jobIdEq).JobIdNe(jobIdNe).JobIdGt(jobIdGt).JobIdGte(jobIdGte).JobIdLt(jobIdLt).JobIdLte(jobIdLte).JobIdIn(jobIdIn).JobIdNotin(jobIdNotin).JobIdLike(jobIdLike).JobIdNotlike(jobIdNotlike).JobIdIlike(jobIdIlike).JobIdNotilike(jobIdNotilike).JobIdDesc(jobIdDesc).JobIdAsc(jobIdAsc).JobIdObjectDesc(jobIdObjectDesc).JobIdObjectAsc(jobIdObjectAsc).ReferencedByTriggerHasExecutionExecutionIdObjectsDesc(referencedByTriggerHasExecutionExecutionIdObjectsDesc).ReferencedByTriggerHasExecutionExecutionIdObjectsAsc(referencedByTriggerHasExecutionExecutionIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.GetExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExecutions`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.GetExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
 **idEq** | **string** | SQL &#x3D; comparison | 
 **idNe** | **string** | SQL !&#x3D; comparison | 
 **idGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **idGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **idLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **idLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **idIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **idNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **idLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **idDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **idAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **createdAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **createdAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **createdAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **createdAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **createdAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **createdAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **createdAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **createdAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **createdAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **createdAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **createdAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **updatedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **updatedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **updatedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **updatedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **updatedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **updatedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **updatedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **updatedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **updatedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **updatedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **updatedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deletedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **deletedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **deletedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **deletedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **deletedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **deletedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **deletedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **deletedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **deletedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deletedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deletedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **statusEq** | **string** | SQL &#x3D; comparison | 
 **statusNe** | **string** | SQL !&#x3D; comparison | 
 **statusGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **statusGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **statusLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **statusLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **statusIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **statusNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **statusLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **statusAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **triggerIdEq** | **string** | SQL &#x3D; comparison | 
 **triggerIdNe** | **string** | SQL !&#x3D; comparison | 
 **triggerIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **triggerIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **triggerIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **triggerIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **triggerIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **triggerIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **triggerIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **triggerIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **triggerIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **triggerIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **triggerIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **triggerIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **triggerIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **triggerIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **buildOutputIdEq** | **string** | SQL &#x3D; comparison | 
 **buildOutputIdNe** | **string** | SQL !&#x3D; comparison | 
 **buildOutputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **buildOutputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **buildOutputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **buildOutputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **buildOutputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **buildOutputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **buildOutputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildOutputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildOutputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildOutputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildOutputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **buildOutputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **buildOutputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **buildOutputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **testOutputIdEq** | **string** | SQL &#x3D; comparison | 
 **testOutputIdNe** | **string** | SQL !&#x3D; comparison | 
 **testOutputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **testOutputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **testOutputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **testOutputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **testOutputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **testOutputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **testOutputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testOutputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testOutputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testOutputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testOutputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **testOutputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **testOutputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **testOutputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **publishOutputIdEq** | **string** | SQL &#x3D; comparison | 
 **publishOutputIdNe** | **string** | SQL !&#x3D; comparison | 
 **publishOutputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **publishOutputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **publishOutputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **publishOutputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **publishOutputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **publishOutputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **publishOutputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishOutputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishOutputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishOutputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishOutputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **publishOutputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **publishOutputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **publishOutputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deployOutputIdEq** | **string** | SQL &#x3D; comparison | 
 **deployOutputIdNe** | **string** | SQL !&#x3D; comparison | 
 **deployOutputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **deployOutputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **deployOutputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **deployOutputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **deployOutputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **deployOutputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **deployOutputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployOutputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployOutputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployOutputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployOutputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deployOutputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deployOutputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deployOutputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **validateOutputIdEq** | **string** | SQL &#x3D; comparison | 
 **validateOutputIdNe** | **string** | SQL !&#x3D; comparison | 
 **validateOutputIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **validateOutputIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **validateOutputIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **validateOutputIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **validateOutputIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **validateOutputIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **validateOutputIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateOutputIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateOutputIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateOutputIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateOutputIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **validateOutputIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **validateOutputIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **validateOutputIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **jobIdEq** | **string** | SQL &#x3D; comparison | 
 **jobIdNe** | **string** | SQL !&#x3D; comparison | 
 **jobIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **jobIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **jobIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **jobIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **jobIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **jobIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **jobIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **jobIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **jobIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **jobIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerHasExecutionExecutionIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTriggerHasExecutionExecutionIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfExecution**](ResponseWithGenericOfExecution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchExecution

> ResponseWithGenericOfExecution PatchExecution(ctx, primaryKey).Execution(execution).Depth(depth).Execute()



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
	primaryKey := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Path parameter primaryKey
	execution := *openapiclient.NewExecution() // Execution | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.PatchExecution(context.Background(), primaryKey).Execution(execution).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.PatchExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchExecution`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.PatchExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | [**Execution**](Execution.md) |  | 
 **depth** | **int64** | Query parameter depth | 

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


## PostExecutions

> ResponseWithGenericOfExecution PostExecutions(ctx).Execution(execution).Depth(depth).Execute()



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
	execution := []openapiclient.Execution{*openapiclient.NewExecution()} // []Execution | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExecutionAPI.PostExecutions(context.Background()).Execution(execution).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExecutionAPI.PostExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostExecutions`: ResponseWithGenericOfExecution
	fmt.Fprintf(os.Stdout, "Response from `ExecutionAPI.PostExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **execution** | [**[]Execution**](Execution.md) |  | 
 **depth** | **int64** | Query parameter depth | 

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

