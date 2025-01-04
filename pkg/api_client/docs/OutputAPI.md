# \OutputAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOutput**](OutputAPI.md#DeleteOutput) | **Delete** /api/outputs/{primaryKey} | 
[**GetOutput**](OutputAPI.md#GetOutput) | **Get** /api/outputs/{primaryKey} | 
[**GetOutputs**](OutputAPI.md#GetOutputs) | **Get** /api/outputs | 
[**PatchOutput**](OutputAPI.md#PatchOutput) | **Patch** /api/outputs/{primaryKey} | 
[**PostOutputs**](OutputAPI.md#PostOutputs) | **Post** /api/outputs | 



## DeleteOutput

> DeleteOutput(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.OutputAPI.DeleteOutput(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputAPI.DeleteOutput``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOutputRequest struct via the builder pattern


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


## GetOutput

> ResponseWithGenericOfOutput GetOutput(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.OutputAPI.GetOutput(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputAPI.GetOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutput`: ResponseWithGenericOfOutput
	fmt.Fprintf(os.Stdout, "Response from `OutputAPI.GetOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfOutput**](ResponseWithGenericOfOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutputs

> ResponseWithGenericOfOutput GetOutputs(ctx).Limit(limit).Offset(offset).Depth(depth).ExecutionLoad(executionLoad).TaskLoad(taskLoad).LogLoad(logLoad).ReferencedByLogLoad(referencedByLogLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusContains(statusContains).StatusNotcontains(statusNotcontains).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).StartedAtEq(startedAtEq).StartedAtNe(startedAtNe).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtIn(startedAtIn).StartedAtNotin(startedAtNotin).StartedAtContains(startedAtContains).StartedAtNotcontains(startedAtNotcontains).StartedAtLike(startedAtLike).StartedAtNotlike(startedAtNotlike).StartedAtIlike(startedAtIlike).StartedAtNotilike(startedAtNotilike).StartedAtDesc(startedAtDesc).StartedAtAsc(startedAtAsc).EndedAtEq(endedAtEq).EndedAtNe(endedAtNe).EndedAtGt(endedAtGt).EndedAtGte(endedAtGte).EndedAtLt(endedAtLt).EndedAtLte(endedAtLte).EndedAtIn(endedAtIn).EndedAtNotin(endedAtNotin).EndedAtContains(endedAtContains).EndedAtNotcontains(endedAtNotcontains).EndedAtLike(endedAtLike).EndedAtNotlike(endedAtNotlike).EndedAtIlike(endedAtIlike).EndedAtNotilike(endedAtNotilike).EndedAtDesc(endedAtDesc).EndedAtAsc(endedAtAsc).ExitStatusEq(exitStatusEq).ExitStatusNe(exitStatusNe).ExitStatusGt(exitStatusGt).ExitStatusGte(exitStatusGte).ExitStatusLt(exitStatusLt).ExitStatusLte(exitStatusLte).ExitStatusIn(exitStatusIn).ExitStatusNotin(exitStatusNotin).ExitStatusContains(exitStatusContains).ExitStatusNotcontains(exitStatusNotcontains).ExitStatusDesc(exitStatusDesc).ExitStatusAsc(exitStatusAsc).ErrorEq(errorEq).ErrorNe(errorNe).ErrorGt(errorGt).ErrorGte(errorGte).ErrorLt(errorLt).ErrorLte(errorLte).ErrorIn(errorIn).ErrorNotin(errorNotin).ErrorContains(errorContains).ErrorNotcontains(errorNotcontains).ErrorLike(errorLike).ErrorNotlike(errorNotlike).ErrorIlike(errorIlike).ErrorNotilike(errorNotilike).ErrorDesc(errorDesc).ErrorAsc(errorAsc).ExecutionIdEq(executionIdEq).ExecutionIdNe(executionIdNe).ExecutionIdGt(executionIdGt).ExecutionIdGte(executionIdGte).ExecutionIdLt(executionIdLt).ExecutionIdLte(executionIdLte).ExecutionIdIn(executionIdIn).ExecutionIdNotin(executionIdNotin).ExecutionIdContains(executionIdContains).ExecutionIdNotcontains(executionIdNotcontains).ExecutionIdLike(executionIdLike).ExecutionIdNotlike(executionIdNotlike).ExecutionIdIlike(executionIdIlike).ExecutionIdNotilike(executionIdNotilike).ExecutionIdDesc(executionIdDesc).ExecutionIdAsc(executionIdAsc).ExecutionIdObjectContains(executionIdObjectContains).ExecutionIdObjectNotcontains(executionIdObjectNotcontains).ExecutionIdObjectDesc(executionIdObjectDesc).ExecutionIdObjectAsc(executionIdObjectAsc).TaskIdEq(taskIdEq).TaskIdNe(taskIdNe).TaskIdGt(taskIdGt).TaskIdGte(taskIdGte).TaskIdLt(taskIdLt).TaskIdLte(taskIdLte).TaskIdIn(taskIdIn).TaskIdNotin(taskIdNotin).TaskIdContains(taskIdContains).TaskIdNotcontains(taskIdNotcontains).TaskIdLike(taskIdLike).TaskIdNotlike(taskIdNotlike).TaskIdIlike(taskIdIlike).TaskIdNotilike(taskIdNotilike).TaskIdDesc(taskIdDesc).TaskIdAsc(taskIdAsc).TaskIdObjectContains(taskIdObjectContains).TaskIdObjectNotcontains(taskIdObjectNotcontains).TaskIdObjectDesc(taskIdObjectDesc).TaskIdObjectAsc(taskIdObjectAsc).LogIdEq(logIdEq).LogIdNe(logIdNe).LogIdGt(logIdGt).LogIdGte(logIdGte).LogIdLt(logIdLt).LogIdLte(logIdLte).LogIdIn(logIdIn).LogIdNotin(logIdNotin).LogIdContains(logIdContains).LogIdNotcontains(logIdNotcontains).LogIdLike(logIdLike).LogIdNotlike(logIdNotlike).LogIdIlike(logIdIlike).LogIdNotilike(logIdNotilike).LogIdDesc(logIdDesc).LogIdAsc(logIdAsc).LogIdObjectContains(logIdObjectContains).LogIdObjectNotcontains(logIdObjectNotcontains).LogIdObjectDesc(logIdObjectDesc).LogIdObjectAsc(logIdObjectAsc).ReferencedByLogOutputIdObjectsContains(referencedByLogOutputIdObjectsContains).ReferencedByLogOutputIdObjectsNotcontains(referencedByLogOutputIdObjectsNotcontains).ReferencedByLogOutputIdObjectsDesc(referencedByLogOutputIdObjectsDesc).ReferencedByLogOutputIdObjectsAsc(referencedByLogOutputIdObjectsAsc).Execute()



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
	executionLoad := "executionLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	taskLoad := "taskLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	logLoad := "logLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	referencedByLogLoad := "referencedByLogLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
	idEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	idNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	idGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	idGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	idLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	idLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	idIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	idNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	idContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	idNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
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
	createdAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	createdAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	updatedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	updatedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	deletedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	deletedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
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
	statusContains := "statusContains_example" // string | SQL @> comparison (optional)
	statusNotcontains := "statusNotcontains_example" // string | SQL NOT @> comparison (optional)
	statusLike := "statusLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotlike := "statusNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusIlike := "statusIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusNotilike := "statusNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	statusDesc := "statusDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	statusAsc := "statusAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	startedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	startedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	startedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	startedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	startedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	startedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	startedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	startedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	startedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	startedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
	startedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	startedAtDesc := "startedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	startedAtAsc := "startedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	endedAtEq := time.Now() // time.Time | SQL = comparison (optional)
	endedAtNe := time.Now() // time.Time | SQL != comparison (optional)
	endedAtGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	endedAtGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	endedAtLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	endedAtLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	endedAtIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	endedAtNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	endedAtContains := time.Now() // time.Time | SQL @> comparison (optional)
	endedAtNotcontains := time.Now() // time.Time | SQL NOT @> comparison (optional)
	endedAtLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	endedAtDesc := "endedAtDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	endedAtAsc := "endedAtAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	exitStatusEq := int64(789) // int64 | SQL = comparison (optional)
	exitStatusNe := int64(789) // int64 | SQL != comparison (optional)
	exitStatusGt := int64(789) // int64 | SQL > comparison, may not work with all column types (optional)
	exitStatusGte := int64(789) // int64 | SQL >= comparison, may not work with all column types (optional)
	exitStatusLt := int64(789) // int64 | SQL < comparison, may not work with all column types (optional)
	exitStatusLte := int64(789) // int64 | SQL <= comparison, may not work with all column types (optional)
	exitStatusIn := int64(789) // int64 | SQL IN comparison, permits comma-separated values (optional)
	exitStatusNotin := int64(789) // int64 | SQL NOT IN comparison, permits comma-separated values (optional)
	exitStatusContains := int64(789) // int64 | SQL @> comparison (optional)
	exitStatusNotcontains := int64(789) // int64 | SQL NOT @> comparison (optional)
	exitStatusDesc := "exitStatusDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	exitStatusAsc := "exitStatusAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	errorEq := "errorEq_example" // string | SQL = comparison (optional)
	errorNe := "errorNe_example" // string | SQL != comparison (optional)
	errorGt := "errorGt_example" // string | SQL > comparison, may not work with all column types (optional)
	errorGte := "errorGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	errorLt := "errorLt_example" // string | SQL < comparison, may not work with all column types (optional)
	errorLte := "errorLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	errorIn := "errorIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	errorNotin := "errorNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	errorContains := "errorContains_example" // string | SQL @> comparison (optional)
	errorNotcontains := "errorNotcontains_example" // string | SQL NOT @> comparison (optional)
	errorLike := "errorLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	errorNotlike := "errorNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	errorIlike := "errorIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	errorNotilike := "errorNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	errorDesc := "errorDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	errorAsc := "errorAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	executionIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	executionIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	executionIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	executionIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	executionIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	executionIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	executionIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	executionIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	executionIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	executionIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	executionIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	executionIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	executionIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	executionIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	executionIdDesc := "executionIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	executionIdAsc := "executionIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	executionIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	executionIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	executionIdObjectDesc := "executionIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	executionIdObjectAsc := "executionIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	taskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	taskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	taskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	taskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	taskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	taskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	taskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	taskIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	taskIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	taskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	taskIdDesc := "taskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdAsc := "taskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	taskIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	taskIdObjectDesc := "taskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	taskIdObjectAsc := "taskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	logIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	logIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	logIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	logIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	logIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	logIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	logIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	logIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	logIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	logIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	logIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	logIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	logIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	logIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	logIdDesc := "logIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	logIdAsc := "logIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	logIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	logIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	logIdObjectDesc := "logIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	logIdObjectAsc := "logIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByLogOutputIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByLogOutputIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByLogOutputIdObjectsDesc := "referencedByLogOutputIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByLogOutputIdObjectsAsc := "referencedByLogOutputIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputAPI.GetOutputs(context.Background()).Limit(limit).Offset(offset).Depth(depth).ExecutionLoad(executionLoad).TaskLoad(taskLoad).LogLoad(logLoad).ReferencedByLogLoad(referencedByLogLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).StatusEq(statusEq).StatusNe(statusNe).StatusGt(statusGt).StatusGte(statusGte).StatusLt(statusLt).StatusLte(statusLte).StatusIn(statusIn).StatusNotin(statusNotin).StatusContains(statusContains).StatusNotcontains(statusNotcontains).StatusLike(statusLike).StatusNotlike(statusNotlike).StatusIlike(statusIlike).StatusNotilike(statusNotilike).StatusDesc(statusDesc).StatusAsc(statusAsc).StartedAtEq(startedAtEq).StartedAtNe(startedAtNe).StartedAtGt(startedAtGt).StartedAtGte(startedAtGte).StartedAtLt(startedAtLt).StartedAtLte(startedAtLte).StartedAtIn(startedAtIn).StartedAtNotin(startedAtNotin).StartedAtContains(startedAtContains).StartedAtNotcontains(startedAtNotcontains).StartedAtLike(startedAtLike).StartedAtNotlike(startedAtNotlike).StartedAtIlike(startedAtIlike).StartedAtNotilike(startedAtNotilike).StartedAtDesc(startedAtDesc).StartedAtAsc(startedAtAsc).EndedAtEq(endedAtEq).EndedAtNe(endedAtNe).EndedAtGt(endedAtGt).EndedAtGte(endedAtGte).EndedAtLt(endedAtLt).EndedAtLte(endedAtLte).EndedAtIn(endedAtIn).EndedAtNotin(endedAtNotin).EndedAtContains(endedAtContains).EndedAtNotcontains(endedAtNotcontains).EndedAtLike(endedAtLike).EndedAtNotlike(endedAtNotlike).EndedAtIlike(endedAtIlike).EndedAtNotilike(endedAtNotilike).EndedAtDesc(endedAtDesc).EndedAtAsc(endedAtAsc).ExitStatusEq(exitStatusEq).ExitStatusNe(exitStatusNe).ExitStatusGt(exitStatusGt).ExitStatusGte(exitStatusGte).ExitStatusLt(exitStatusLt).ExitStatusLte(exitStatusLte).ExitStatusIn(exitStatusIn).ExitStatusNotin(exitStatusNotin).ExitStatusContains(exitStatusContains).ExitStatusNotcontains(exitStatusNotcontains).ExitStatusDesc(exitStatusDesc).ExitStatusAsc(exitStatusAsc).ErrorEq(errorEq).ErrorNe(errorNe).ErrorGt(errorGt).ErrorGte(errorGte).ErrorLt(errorLt).ErrorLte(errorLte).ErrorIn(errorIn).ErrorNotin(errorNotin).ErrorContains(errorContains).ErrorNotcontains(errorNotcontains).ErrorLike(errorLike).ErrorNotlike(errorNotlike).ErrorIlike(errorIlike).ErrorNotilike(errorNotilike).ErrorDesc(errorDesc).ErrorAsc(errorAsc).ExecutionIdEq(executionIdEq).ExecutionIdNe(executionIdNe).ExecutionIdGt(executionIdGt).ExecutionIdGte(executionIdGte).ExecutionIdLt(executionIdLt).ExecutionIdLte(executionIdLte).ExecutionIdIn(executionIdIn).ExecutionIdNotin(executionIdNotin).ExecutionIdContains(executionIdContains).ExecutionIdNotcontains(executionIdNotcontains).ExecutionIdLike(executionIdLike).ExecutionIdNotlike(executionIdNotlike).ExecutionIdIlike(executionIdIlike).ExecutionIdNotilike(executionIdNotilike).ExecutionIdDesc(executionIdDesc).ExecutionIdAsc(executionIdAsc).ExecutionIdObjectContains(executionIdObjectContains).ExecutionIdObjectNotcontains(executionIdObjectNotcontains).ExecutionIdObjectDesc(executionIdObjectDesc).ExecutionIdObjectAsc(executionIdObjectAsc).TaskIdEq(taskIdEq).TaskIdNe(taskIdNe).TaskIdGt(taskIdGt).TaskIdGte(taskIdGte).TaskIdLt(taskIdLt).TaskIdLte(taskIdLte).TaskIdIn(taskIdIn).TaskIdNotin(taskIdNotin).TaskIdContains(taskIdContains).TaskIdNotcontains(taskIdNotcontains).TaskIdLike(taskIdLike).TaskIdNotlike(taskIdNotlike).TaskIdIlike(taskIdIlike).TaskIdNotilike(taskIdNotilike).TaskIdDesc(taskIdDesc).TaskIdAsc(taskIdAsc).TaskIdObjectContains(taskIdObjectContains).TaskIdObjectNotcontains(taskIdObjectNotcontains).TaskIdObjectDesc(taskIdObjectDesc).TaskIdObjectAsc(taskIdObjectAsc).LogIdEq(logIdEq).LogIdNe(logIdNe).LogIdGt(logIdGt).LogIdGte(logIdGte).LogIdLt(logIdLt).LogIdLte(logIdLte).LogIdIn(logIdIn).LogIdNotin(logIdNotin).LogIdContains(logIdContains).LogIdNotcontains(logIdNotcontains).LogIdLike(logIdLike).LogIdNotlike(logIdNotlike).LogIdIlike(logIdIlike).LogIdNotilike(logIdNotilike).LogIdDesc(logIdDesc).LogIdAsc(logIdAsc).LogIdObjectContains(logIdObjectContains).LogIdObjectNotcontains(logIdObjectNotcontains).LogIdObjectDesc(logIdObjectDesc).LogIdObjectAsc(logIdObjectAsc).ReferencedByLogOutputIdObjectsContains(referencedByLogOutputIdObjectsContains).ReferencedByLogOutputIdObjectsNotcontains(referencedByLogOutputIdObjectsNotcontains).ReferencedByLogOutputIdObjectsDesc(referencedByLogOutputIdObjectsDesc).ReferencedByLogOutputIdObjectsAsc(referencedByLogOutputIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputAPI.GetOutputs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutputs`: ResponseWithGenericOfOutput
	fmt.Fprintf(os.Stdout, "Response from `OutputAPI.GetOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | SQL LIMIT operator | 
 **offset** | **int32** | SQL OFFSET operator | 
 **depth** | **int32** | Max recursion depth for loading foreign objects; default &#x3D; 1  (0 &#x3D; recurse until graph cycle detected, 1 &#x3D; this object only, 2 &#x3D; this object + neighbours, 3 &#x3D; this object + neighbours + their neighbours... etc) | 
 **executionLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **taskLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **logLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **referencedByLogLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
 **idEq** | **string** | SQL &#x3D; comparison | 
 **idNe** | **string** | SQL !&#x3D; comparison | 
 **idGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **idGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **idLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **idLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **idIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **idNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **idContains** | **string** | SQL @&gt; comparison | 
 **idNotcontains** | **string** | SQL NOT @&gt; comparison | 
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
 **createdAtContains** | **time.Time** | SQL @&gt; comparison | 
 **createdAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **updatedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **updatedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **deletedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **deletedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
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
 **statusContains** | **string** | SQL @&gt; comparison | 
 **statusNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **statusLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **statusDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **statusAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **startedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **startedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **startedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **startedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **startedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **startedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **startedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **startedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **startedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **startedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
 **startedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **startedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **startedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **endedAtEq** | **time.Time** | SQL &#x3D; comparison | 
 **endedAtNe** | **time.Time** | SQL !&#x3D; comparison | 
 **endedAtGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **endedAtGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **endedAtLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **endedAtLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **endedAtIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **endedAtNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **endedAtContains** | **time.Time** | SQL @&gt; comparison | 
 **endedAtNotcontains** | **time.Time** | SQL NOT @&gt; comparison | 
 **endedAtLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **endedAtDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **endedAtAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **exitStatusEq** | **int64** | SQL &#x3D; comparison | 
 **exitStatusNe** | **int64** | SQL !&#x3D; comparison | 
 **exitStatusGt** | **int64** | SQL &gt; comparison, may not work with all column types | 
 **exitStatusGte** | **int64** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **exitStatusLt** | **int64** | SQL &lt; comparison, may not work with all column types | 
 **exitStatusLte** | **int64** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **exitStatusIn** | **int64** | SQL IN comparison, permits comma-separated values | 
 **exitStatusNotin** | **int64** | SQL NOT IN comparison, permits comma-separated values | 
 **exitStatusContains** | **int64** | SQL @&gt; comparison | 
 **exitStatusNotcontains** | **int64** | SQL NOT @&gt; comparison | 
 **exitStatusDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **exitStatusAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **errorEq** | **string** | SQL &#x3D; comparison | 
 **errorNe** | **string** | SQL !&#x3D; comparison | 
 **errorGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **errorGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **errorLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **errorLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **errorIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **errorNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **errorContains** | **string** | SQL @&gt; comparison | 
 **errorNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **errorLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **errorNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **errorIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **errorNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **errorDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **errorAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **executionIdEq** | **string** | SQL &#x3D; comparison | 
 **executionIdNe** | **string** | SQL !&#x3D; comparison | 
 **executionIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **executionIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **executionIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **executionIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **executionIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **executionIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **executionIdContains** | **string** | SQL @&gt; comparison | 
 **executionIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **executionIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **executionIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **executionIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **executionIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **executionIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **executionIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **executionIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **executionIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **executionIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **executionIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **taskIdEq** | **string** | SQL &#x3D; comparison | 
 **taskIdNe** | **string** | SQL !&#x3D; comparison | 
 **taskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **taskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **taskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **taskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **taskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **taskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **taskIdContains** | **string** | SQL @&gt; comparison | 
 **taskIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **taskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **taskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **taskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **taskIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **taskIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **taskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **taskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **logIdEq** | **string** | SQL &#x3D; comparison | 
 **logIdNe** | **string** | SQL !&#x3D; comparison | 
 **logIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **logIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **logIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **logIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **logIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **logIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **logIdContains** | **string** | SQL @&gt; comparison | 
 **logIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **logIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **logIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **logIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **logIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **logIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **logIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **logIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **logIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **logIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **logIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByLogOutputIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByLogOutputIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByLogOutputIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByLogOutputIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfOutput**](ResponseWithGenericOfOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOutput

> ResponseWithGenericOfOutput PatchOutput(ctx, primaryKey).Output(output).Depth(depth).Execute()



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
	output := *openapiclient.NewOutput() // Output | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputAPI.PatchOutput(context.Background(), primaryKey).Output(output).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputAPI.PatchOutput``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchOutput`: ResponseWithGenericOfOutput
	fmt.Fprintf(os.Stdout, "Response from `OutputAPI.PatchOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **output** | [**Output**](Output.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfOutput**](ResponseWithGenericOfOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOutputs

> ResponseWithGenericOfOutput PostOutputs(ctx).Output(output).Depth(depth).Execute()



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
	output := []openapiclient.Output{*openapiclient.NewOutput()} // []Output | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutputAPI.PostOutputs(context.Background()).Output(output).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutputAPI.PostOutputs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOutputs`: ResponseWithGenericOfOutput
	fmt.Fprintf(os.Stdout, "Response from `OutputAPI.PostOutputs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOutputsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **output** | [**[]Output**](Output.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfOutput**](ResponseWithGenericOfOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

