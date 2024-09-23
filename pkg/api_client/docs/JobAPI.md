# \JobAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJob**](JobAPI.md#DeleteJob) | **Delete** /api/jobs/{primaryKey} | 
[**GetJob**](JobAPI.md#GetJob) | **Get** /api/jobs/{primaryKey} | 
[**GetJobs**](JobAPI.md#GetJobs) | **Get** /api/jobs | 
[**PatchJob**](JobAPI.md#PatchJob) | **Patch** /api/jobs/{primaryKey} | 
[**PostJobs**](JobAPI.md#PostJobs) | **Post** /api/jobs | 



## DeleteJob

> DeleteJob(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.JobAPI.DeleteJob(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.DeleteJob``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


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


## GetJob

> ResponseWithGenericOfJob GetJob(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.JobAPI.GetJob(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: ResponseWithGenericOfJob
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfJob**](ResponseWithGenericOfJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> ResponseWithGenericOfJob GetJobs(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).BuildTaskIdEq(buildTaskIdEq).BuildTaskIdNe(buildTaskIdNe).BuildTaskIdGt(buildTaskIdGt).BuildTaskIdGte(buildTaskIdGte).BuildTaskIdLt(buildTaskIdLt).BuildTaskIdLte(buildTaskIdLte).BuildTaskIdIn(buildTaskIdIn).BuildTaskIdNotin(buildTaskIdNotin).BuildTaskIdLike(buildTaskIdLike).BuildTaskIdNotlike(buildTaskIdNotlike).BuildTaskIdIlike(buildTaskIdIlike).BuildTaskIdNotilike(buildTaskIdNotilike).BuildTaskIdDesc(buildTaskIdDesc).BuildTaskIdAsc(buildTaskIdAsc).BuildTaskIdObjectDesc(buildTaskIdObjectDesc).BuildTaskIdObjectAsc(buildTaskIdObjectAsc).TestTaskIdEq(testTaskIdEq).TestTaskIdNe(testTaskIdNe).TestTaskIdGt(testTaskIdGt).TestTaskIdGte(testTaskIdGte).TestTaskIdLt(testTaskIdLt).TestTaskIdLte(testTaskIdLte).TestTaskIdIn(testTaskIdIn).TestTaskIdNotin(testTaskIdNotin).TestTaskIdLike(testTaskIdLike).TestTaskIdNotlike(testTaskIdNotlike).TestTaskIdIlike(testTaskIdIlike).TestTaskIdNotilike(testTaskIdNotilike).TestTaskIdDesc(testTaskIdDesc).TestTaskIdAsc(testTaskIdAsc).TestTaskIdObjectDesc(testTaskIdObjectDesc).TestTaskIdObjectAsc(testTaskIdObjectAsc).PublishTaskIdEq(publishTaskIdEq).PublishTaskIdNe(publishTaskIdNe).PublishTaskIdGt(publishTaskIdGt).PublishTaskIdGte(publishTaskIdGte).PublishTaskIdLt(publishTaskIdLt).PublishTaskIdLte(publishTaskIdLte).PublishTaskIdIn(publishTaskIdIn).PublishTaskIdNotin(publishTaskIdNotin).PublishTaskIdLike(publishTaskIdLike).PublishTaskIdNotlike(publishTaskIdNotlike).PublishTaskIdIlike(publishTaskIdIlike).PublishTaskIdNotilike(publishTaskIdNotilike).PublishTaskIdDesc(publishTaskIdDesc).PublishTaskIdAsc(publishTaskIdAsc).PublishTaskIdObjectDesc(publishTaskIdObjectDesc).PublishTaskIdObjectAsc(publishTaskIdObjectAsc).DeployTaskIdEq(deployTaskIdEq).DeployTaskIdNe(deployTaskIdNe).DeployTaskIdGt(deployTaskIdGt).DeployTaskIdGte(deployTaskIdGte).DeployTaskIdLt(deployTaskIdLt).DeployTaskIdLte(deployTaskIdLte).DeployTaskIdIn(deployTaskIdIn).DeployTaskIdNotin(deployTaskIdNotin).DeployTaskIdLike(deployTaskIdLike).DeployTaskIdNotlike(deployTaskIdNotlike).DeployTaskIdIlike(deployTaskIdIlike).DeployTaskIdNotilike(deployTaskIdNotilike).DeployTaskIdDesc(deployTaskIdDesc).DeployTaskIdAsc(deployTaskIdAsc).DeployTaskIdObjectDesc(deployTaskIdObjectDesc).DeployTaskIdObjectAsc(deployTaskIdObjectAsc).ValidateTaskIdEq(validateTaskIdEq).ValidateTaskIdNe(validateTaskIdNe).ValidateTaskIdGt(validateTaskIdGt).ValidateTaskIdGte(validateTaskIdGte).ValidateTaskIdLt(validateTaskIdLt).ValidateTaskIdLte(validateTaskIdLte).ValidateTaskIdIn(validateTaskIdIn).ValidateTaskIdNotin(validateTaskIdNotin).ValidateTaskIdLike(validateTaskIdLike).ValidateTaskIdNotlike(validateTaskIdNotlike).ValidateTaskIdIlike(validateTaskIdIlike).ValidateTaskIdNotilike(validateTaskIdNotilike).ValidateTaskIdDesc(validateTaskIdDesc).ValidateTaskIdAsc(validateTaskIdAsc).ValidateTaskIdObjectDesc(validateTaskIdObjectDesc).ValidateTaskIdObjectAsc(validateTaskIdObjectAsc).ReferencedByExecutionJobIdObjectsDesc(referencedByExecutionJobIdObjectsDesc).ReferencedByExecutionJobIdObjectsAsc(referencedByExecutionJobIdObjectsAsc).ReferencedByRuleRequiresJobJobIdObjectsDesc(referencedByRuleRequiresJobJobIdObjectsDesc).ReferencedByRuleRequiresJobJobIdObjectsAsc(referencedByRuleRequiresJobJobIdObjectsAsc).Execute()



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
	ruleIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	ruleIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	ruleIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	ruleIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	ruleIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	ruleIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	ruleIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	ruleIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	ruleIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	ruleIdDesc := "ruleIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdAsc := "ruleIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectDesc := "ruleIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	ruleIdObjectAsc := "ruleIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	buildTaskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	buildTaskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	buildTaskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	buildTaskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	buildTaskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	buildTaskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	buildTaskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	buildTaskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	buildTaskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildTaskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildTaskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildTaskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	buildTaskIdDesc := "buildTaskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	buildTaskIdAsc := "buildTaskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	buildTaskIdObjectDesc := "buildTaskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	buildTaskIdObjectAsc := "buildTaskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	testTaskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	testTaskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	testTaskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	testTaskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	testTaskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	testTaskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	testTaskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	testTaskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	testTaskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testTaskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testTaskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testTaskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	testTaskIdDesc := "testTaskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	testTaskIdAsc := "testTaskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	testTaskIdObjectDesc := "testTaskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	testTaskIdObjectAsc := "testTaskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	publishTaskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	publishTaskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	publishTaskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	publishTaskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	publishTaskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	publishTaskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	publishTaskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	publishTaskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	publishTaskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishTaskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishTaskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishTaskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	publishTaskIdDesc := "publishTaskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	publishTaskIdAsc := "publishTaskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	publishTaskIdObjectDesc := "publishTaskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	publishTaskIdObjectAsc := "publishTaskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deployTaskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	deployTaskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	deployTaskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	deployTaskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	deployTaskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	deployTaskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	deployTaskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	deployTaskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	deployTaskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployTaskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployTaskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployTaskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	deployTaskIdDesc := "deployTaskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deployTaskIdAsc := "deployTaskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	deployTaskIdObjectDesc := "deployTaskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	deployTaskIdObjectAsc := "deployTaskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	validateTaskIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	validateTaskIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	validateTaskIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	validateTaskIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	validateTaskIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	validateTaskIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	validateTaskIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	validateTaskIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	validateTaskIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateTaskIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateTaskIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateTaskIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	validateTaskIdDesc := "validateTaskIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	validateTaskIdAsc := "validateTaskIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	validateTaskIdObjectDesc := "validateTaskIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	validateTaskIdObjectAsc := "validateTaskIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionJobIdObjectsDesc := "referencedByExecutionJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionJobIdObjectsAsc := "referencedByExecutionJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByRuleRequiresJobJobIdObjectsDesc := "referencedByRuleRequiresJobJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByRuleRequiresJobJobIdObjectsAsc := "referencedByRuleRequiresJobJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobs(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).RuleIdEq(ruleIdEq).RuleIdNe(ruleIdNe).RuleIdGt(ruleIdGt).RuleIdGte(ruleIdGte).RuleIdLt(ruleIdLt).RuleIdLte(ruleIdLte).RuleIdIn(ruleIdIn).RuleIdNotin(ruleIdNotin).RuleIdLike(ruleIdLike).RuleIdNotlike(ruleIdNotlike).RuleIdIlike(ruleIdIlike).RuleIdNotilike(ruleIdNotilike).RuleIdDesc(ruleIdDesc).RuleIdAsc(ruleIdAsc).RuleIdObjectDesc(ruleIdObjectDesc).RuleIdObjectAsc(ruleIdObjectAsc).BuildTaskIdEq(buildTaskIdEq).BuildTaskIdNe(buildTaskIdNe).BuildTaskIdGt(buildTaskIdGt).BuildTaskIdGte(buildTaskIdGte).BuildTaskIdLt(buildTaskIdLt).BuildTaskIdLte(buildTaskIdLte).BuildTaskIdIn(buildTaskIdIn).BuildTaskIdNotin(buildTaskIdNotin).BuildTaskIdLike(buildTaskIdLike).BuildTaskIdNotlike(buildTaskIdNotlike).BuildTaskIdIlike(buildTaskIdIlike).BuildTaskIdNotilike(buildTaskIdNotilike).BuildTaskIdDesc(buildTaskIdDesc).BuildTaskIdAsc(buildTaskIdAsc).BuildTaskIdObjectDesc(buildTaskIdObjectDesc).BuildTaskIdObjectAsc(buildTaskIdObjectAsc).TestTaskIdEq(testTaskIdEq).TestTaskIdNe(testTaskIdNe).TestTaskIdGt(testTaskIdGt).TestTaskIdGte(testTaskIdGte).TestTaskIdLt(testTaskIdLt).TestTaskIdLte(testTaskIdLte).TestTaskIdIn(testTaskIdIn).TestTaskIdNotin(testTaskIdNotin).TestTaskIdLike(testTaskIdLike).TestTaskIdNotlike(testTaskIdNotlike).TestTaskIdIlike(testTaskIdIlike).TestTaskIdNotilike(testTaskIdNotilike).TestTaskIdDesc(testTaskIdDesc).TestTaskIdAsc(testTaskIdAsc).TestTaskIdObjectDesc(testTaskIdObjectDesc).TestTaskIdObjectAsc(testTaskIdObjectAsc).PublishTaskIdEq(publishTaskIdEq).PublishTaskIdNe(publishTaskIdNe).PublishTaskIdGt(publishTaskIdGt).PublishTaskIdGte(publishTaskIdGte).PublishTaskIdLt(publishTaskIdLt).PublishTaskIdLte(publishTaskIdLte).PublishTaskIdIn(publishTaskIdIn).PublishTaskIdNotin(publishTaskIdNotin).PublishTaskIdLike(publishTaskIdLike).PublishTaskIdNotlike(publishTaskIdNotlike).PublishTaskIdIlike(publishTaskIdIlike).PublishTaskIdNotilike(publishTaskIdNotilike).PublishTaskIdDesc(publishTaskIdDesc).PublishTaskIdAsc(publishTaskIdAsc).PublishTaskIdObjectDesc(publishTaskIdObjectDesc).PublishTaskIdObjectAsc(publishTaskIdObjectAsc).DeployTaskIdEq(deployTaskIdEq).DeployTaskIdNe(deployTaskIdNe).DeployTaskIdGt(deployTaskIdGt).DeployTaskIdGte(deployTaskIdGte).DeployTaskIdLt(deployTaskIdLt).DeployTaskIdLte(deployTaskIdLte).DeployTaskIdIn(deployTaskIdIn).DeployTaskIdNotin(deployTaskIdNotin).DeployTaskIdLike(deployTaskIdLike).DeployTaskIdNotlike(deployTaskIdNotlike).DeployTaskIdIlike(deployTaskIdIlike).DeployTaskIdNotilike(deployTaskIdNotilike).DeployTaskIdDesc(deployTaskIdDesc).DeployTaskIdAsc(deployTaskIdAsc).DeployTaskIdObjectDesc(deployTaskIdObjectDesc).DeployTaskIdObjectAsc(deployTaskIdObjectAsc).ValidateTaskIdEq(validateTaskIdEq).ValidateTaskIdNe(validateTaskIdNe).ValidateTaskIdGt(validateTaskIdGt).ValidateTaskIdGte(validateTaskIdGte).ValidateTaskIdLt(validateTaskIdLt).ValidateTaskIdLte(validateTaskIdLte).ValidateTaskIdIn(validateTaskIdIn).ValidateTaskIdNotin(validateTaskIdNotin).ValidateTaskIdLike(validateTaskIdLike).ValidateTaskIdNotlike(validateTaskIdNotlike).ValidateTaskIdIlike(validateTaskIdIlike).ValidateTaskIdNotilike(validateTaskIdNotilike).ValidateTaskIdDesc(validateTaskIdDesc).ValidateTaskIdAsc(validateTaskIdAsc).ValidateTaskIdObjectDesc(validateTaskIdObjectDesc).ValidateTaskIdObjectAsc(validateTaskIdObjectAsc).ReferencedByExecutionJobIdObjectsDesc(referencedByExecutionJobIdObjectsDesc).ReferencedByExecutionJobIdObjectsAsc(referencedByExecutionJobIdObjectsAsc).ReferencedByRuleRequiresJobJobIdObjectsDesc(referencedByRuleRequiresJobJobIdObjectsDesc).ReferencedByRuleRequiresJobJobIdObjectsAsc(referencedByRuleRequiresJobJobIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobs`: ResponseWithGenericOfJob
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


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
 **ruleIdEq** | **string** | SQL &#x3D; comparison | 
 **ruleIdNe** | **string** | SQL !&#x3D; comparison | 
 **ruleIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **ruleIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **ruleIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **ruleIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **ruleIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **ruleIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **ruleIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **ruleIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **ruleIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **buildTaskIdEq** | **string** | SQL &#x3D; comparison | 
 **buildTaskIdNe** | **string** | SQL !&#x3D; comparison | 
 **buildTaskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **buildTaskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **buildTaskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **buildTaskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **buildTaskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **buildTaskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **buildTaskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildTaskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildTaskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildTaskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **buildTaskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **buildTaskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **buildTaskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **buildTaskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **testTaskIdEq** | **string** | SQL &#x3D; comparison | 
 **testTaskIdNe** | **string** | SQL !&#x3D; comparison | 
 **testTaskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **testTaskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **testTaskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **testTaskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **testTaskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **testTaskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **testTaskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testTaskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testTaskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testTaskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **testTaskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **testTaskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **testTaskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **testTaskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **publishTaskIdEq** | **string** | SQL &#x3D; comparison | 
 **publishTaskIdNe** | **string** | SQL !&#x3D; comparison | 
 **publishTaskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **publishTaskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **publishTaskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **publishTaskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **publishTaskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **publishTaskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **publishTaskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishTaskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishTaskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishTaskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **publishTaskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **publishTaskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **publishTaskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **publishTaskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deployTaskIdEq** | **string** | SQL &#x3D; comparison | 
 **deployTaskIdNe** | **string** | SQL !&#x3D; comparison | 
 **deployTaskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **deployTaskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **deployTaskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **deployTaskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **deployTaskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **deployTaskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **deployTaskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployTaskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployTaskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployTaskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **deployTaskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deployTaskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **deployTaskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **deployTaskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **validateTaskIdEq** | **string** | SQL &#x3D; comparison | 
 **validateTaskIdNe** | **string** | SQL !&#x3D; comparison | 
 **validateTaskIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **validateTaskIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **validateTaskIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **validateTaskIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **validateTaskIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **validateTaskIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **validateTaskIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateTaskIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateTaskIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateTaskIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **validateTaskIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **validateTaskIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **validateTaskIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **validateTaskIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByRuleRequiresJobJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByRuleRequiresJobJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfJob**](ResponseWithGenericOfJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchJob

> ResponseWithGenericOfJob PatchJob(ctx, primaryKey).Job(job).Depth(depth).Execute()



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
	job := *openapiclient.NewJob() // Job | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.PatchJob(context.Background(), primaryKey).Job(job).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.PatchJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchJob`: ResponseWithGenericOfJob
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.PatchJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **job** | [**Job**](Job.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfJob**](ResponseWithGenericOfJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobs

> ResponseWithGenericOfJob PostJobs(ctx).Job(job).Depth(depth).Execute()



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
	job := []openapiclient.Job{*openapiclient.NewJob()} // []Job | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.PostJobs(context.Background()).Job(job).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JobAPI.PostJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostJobs`: ResponseWithGenericOfJob
	fmt.Fprintf(os.Stdout, "Response from `JobAPI.PostJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job** | [**[]Job**](Job.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfJob**](ResponseWithGenericOfJob.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

