# \TaskAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTask**](TaskAPI.md#DeleteTask) | **Delete** /api/tasks/{primaryKey} | 
[**GetTask**](TaskAPI.md#GetTask) | **Get** /api/tasks/{primaryKey} | 
[**GetTasks**](TaskAPI.md#GetTasks) | **Get** /api/tasks | 
[**PatchTask**](TaskAPI.md#PatchTask) | **Patch** /api/tasks/{primaryKey} | 
[**PostTasks**](TaskAPI.md#PostTasks) | **Post** /api/tasks | 



## DeleteTask

> DeleteTask(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.TaskAPI.DeleteTask(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.DeleteTask``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


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


## GetTask

> ResponseWithGenericOfTask GetTask(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.TaskAPI.GetTask(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTask`: ResponseWithGenericOfTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTask**](ResponseWithGenericOfTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasks

> ResponseWithGenericOfTask GetTasks(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).PlatformEq(platformEq).PlatformNe(platformNe).PlatformGt(platformGt).PlatformGte(platformGte).PlatformLt(platformLt).PlatformLte(platformLte).PlatformIn(platformIn).PlatformNotin(platformNotin).PlatformLike(platformLike).PlatformNotlike(platformNotlike).PlatformIlike(platformIlike).PlatformNotilike(platformNotilike).PlatformDesc(platformDesc).PlatformAsc(platformAsc).ImageEq(imageEq).ImageNe(imageNe).ImageGt(imageGt).ImageGte(imageGte).ImageLt(imageLt).ImageLte(imageLte).ImageIn(imageIn).ImageNotin(imageNotin).ImageLike(imageLike).ImageNotlike(imageNotlike).ImageIlike(imageIlike).ImageNotilike(imageNotilike).ImageDesc(imageDesc).ImageAsc(imageAsc).ScriptEq(scriptEq).ScriptNe(scriptNe).ScriptGt(scriptGt).ScriptGte(scriptGte).ScriptLt(scriptLt).ScriptLte(scriptLte).ScriptIn(scriptIn).ScriptNotin(scriptNotin).ScriptLike(scriptLike).ScriptNotlike(scriptNotlike).ScriptIlike(scriptIlike).ScriptNotilike(scriptNotilike).ScriptDesc(scriptDesc).ScriptAsc(scriptAsc).ReferencedByOutputTaskIdObjectsDesc(referencedByOutputTaskIdObjectsDesc).ReferencedByOutputTaskIdObjectsAsc(referencedByOutputTaskIdObjectsAsc).ReferencedByJobBuildTaskIdObjectsDesc(referencedByJobBuildTaskIdObjectsDesc).ReferencedByJobBuildTaskIdObjectsAsc(referencedByJobBuildTaskIdObjectsAsc).ReferencedByJobTestTaskIdObjectsDesc(referencedByJobTestTaskIdObjectsDesc).ReferencedByJobTestTaskIdObjectsAsc(referencedByJobTestTaskIdObjectsAsc).ReferencedByJobPublishTaskIdObjectsDesc(referencedByJobPublishTaskIdObjectsDesc).ReferencedByJobPublishTaskIdObjectsAsc(referencedByJobPublishTaskIdObjectsAsc).ReferencedByJobDeployTaskIdObjectsDesc(referencedByJobDeployTaskIdObjectsDesc).ReferencedByJobDeployTaskIdObjectsAsc(referencedByJobDeployTaskIdObjectsAsc).ReferencedByJobValidateTaskIdObjectsDesc(referencedByJobValidateTaskIdObjectsDesc).ReferencedByJobValidateTaskIdObjectsAsc(referencedByJobValidateTaskIdObjectsAsc).Execute()



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
	nameEq := "nameEq_example" // string | SQL = comparison (optional)
	nameNe := "nameNe_example" // string | SQL != comparison (optional)
	nameGt := "nameGt_example" // string | SQL > comparison, may not work with all column types (optional)
	nameGte := "nameGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	nameLt := "nameLt_example" // string | SQL < comparison, may not work with all column types (optional)
	nameLte := "nameLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	nameIn := "nameIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	nameNotin := "nameNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	nameLike := "nameLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotlike := "nameNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameIlike := "nameIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotilike := "nameNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameDesc := "nameDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	nameAsc := "nameAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	platformEq := "platformEq_example" // string | SQL = comparison (optional)
	platformNe := "platformNe_example" // string | SQL != comparison (optional)
	platformGt := "platformGt_example" // string | SQL > comparison, may not work with all column types (optional)
	platformGte := "platformGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	platformLt := "platformLt_example" // string | SQL < comparison, may not work with all column types (optional)
	platformLte := "platformLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	platformIn := "platformIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	platformNotin := "platformNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	platformLike := "platformLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	platformNotlike := "platformNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	platformIlike := "platformIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	platformNotilike := "platformNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	platformDesc := "platformDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	platformAsc := "platformAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	imageEq := "imageEq_example" // string | SQL = comparison (optional)
	imageNe := "imageNe_example" // string | SQL != comparison (optional)
	imageGt := "imageGt_example" // string | SQL > comparison, may not work with all column types (optional)
	imageGte := "imageGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	imageLt := "imageLt_example" // string | SQL < comparison, may not work with all column types (optional)
	imageLte := "imageLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	imageIn := "imageIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	imageNotin := "imageNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	imageLike := "imageLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	imageNotlike := "imageNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	imageIlike := "imageIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	imageNotilike := "imageNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	imageDesc := "imageDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	imageAsc := "imageAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	scriptEq := "scriptEq_example" // string | SQL = comparison (optional)
	scriptNe := "scriptNe_example" // string | SQL != comparison (optional)
	scriptGt := "scriptGt_example" // string | SQL > comparison, may not work with all column types (optional)
	scriptGte := "scriptGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	scriptLt := "scriptLt_example" // string | SQL < comparison, may not work with all column types (optional)
	scriptLte := "scriptLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	scriptIn := "scriptIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	scriptNotin := "scriptNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	scriptLike := "scriptLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	scriptNotlike := "scriptNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	scriptIlike := "scriptIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	scriptNotilike := "scriptNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	scriptDesc := "scriptDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	scriptAsc := "scriptAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByOutputTaskIdObjectsDesc := "referencedByOutputTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByOutputTaskIdObjectsAsc := "referencedByOutputTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobBuildTaskIdObjectsDesc := "referencedByJobBuildTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobBuildTaskIdObjectsAsc := "referencedByJobBuildTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobTestTaskIdObjectsDesc := "referencedByJobTestTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobTestTaskIdObjectsAsc := "referencedByJobTestTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobPublishTaskIdObjectsDesc := "referencedByJobPublishTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobPublishTaskIdObjectsAsc := "referencedByJobPublishTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobDeployTaskIdObjectsDesc := "referencedByJobDeployTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobDeployTaskIdObjectsAsc := "referencedByJobDeployTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobValidateTaskIdObjectsDesc := "referencedByJobValidateTaskIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByJobValidateTaskIdObjectsAsc := "referencedByJobValidateTaskIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.GetTasks(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).PlatformEq(platformEq).PlatformNe(platformNe).PlatformGt(platformGt).PlatformGte(platformGte).PlatformLt(platformLt).PlatformLte(platformLte).PlatformIn(platformIn).PlatformNotin(platformNotin).PlatformLike(platformLike).PlatformNotlike(platformNotlike).PlatformIlike(platformIlike).PlatformNotilike(platformNotilike).PlatformDesc(platformDesc).PlatformAsc(platformAsc).ImageEq(imageEq).ImageNe(imageNe).ImageGt(imageGt).ImageGte(imageGte).ImageLt(imageLt).ImageLte(imageLte).ImageIn(imageIn).ImageNotin(imageNotin).ImageLike(imageLike).ImageNotlike(imageNotlike).ImageIlike(imageIlike).ImageNotilike(imageNotilike).ImageDesc(imageDesc).ImageAsc(imageAsc).ScriptEq(scriptEq).ScriptNe(scriptNe).ScriptGt(scriptGt).ScriptGte(scriptGte).ScriptLt(scriptLt).ScriptLte(scriptLte).ScriptIn(scriptIn).ScriptNotin(scriptNotin).ScriptLike(scriptLike).ScriptNotlike(scriptNotlike).ScriptIlike(scriptIlike).ScriptNotilike(scriptNotilike).ScriptDesc(scriptDesc).ScriptAsc(scriptAsc).ReferencedByOutputTaskIdObjectsDesc(referencedByOutputTaskIdObjectsDesc).ReferencedByOutputTaskIdObjectsAsc(referencedByOutputTaskIdObjectsAsc).ReferencedByJobBuildTaskIdObjectsDesc(referencedByJobBuildTaskIdObjectsDesc).ReferencedByJobBuildTaskIdObjectsAsc(referencedByJobBuildTaskIdObjectsAsc).ReferencedByJobTestTaskIdObjectsDesc(referencedByJobTestTaskIdObjectsDesc).ReferencedByJobTestTaskIdObjectsAsc(referencedByJobTestTaskIdObjectsAsc).ReferencedByJobPublishTaskIdObjectsDesc(referencedByJobPublishTaskIdObjectsDesc).ReferencedByJobPublishTaskIdObjectsAsc(referencedByJobPublishTaskIdObjectsAsc).ReferencedByJobDeployTaskIdObjectsDesc(referencedByJobDeployTaskIdObjectsDesc).ReferencedByJobDeployTaskIdObjectsAsc(referencedByJobDeployTaskIdObjectsAsc).ReferencedByJobValidateTaskIdObjectsDesc(referencedByJobValidateTaskIdObjectsDesc).ReferencedByJobValidateTaskIdObjectsAsc(referencedByJobValidateTaskIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTasks`: ResponseWithGenericOfTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksRequest struct via the builder pattern


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
 **nameEq** | **string** | SQL &#x3D; comparison | 
 **nameNe** | **string** | SQL !&#x3D; comparison | 
 **nameGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **nameGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **nameLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **nameLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **nameIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **nameNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **nameLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **nameAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **platformEq** | **string** | SQL &#x3D; comparison | 
 **platformNe** | **string** | SQL !&#x3D; comparison | 
 **platformGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **platformGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **platformLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **platformLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **platformIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **platformNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **platformLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **platformNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **platformIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **platformNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **platformDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **platformAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **imageEq** | **string** | SQL &#x3D; comparison | 
 **imageNe** | **string** | SQL !&#x3D; comparison | 
 **imageGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **imageGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **imageLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **imageLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **imageIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **imageNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **imageLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **imageNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **imageIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **imageNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **imageDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **imageAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **scriptEq** | **string** | SQL &#x3D; comparison | 
 **scriptNe** | **string** | SQL !&#x3D; comparison | 
 **scriptGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **scriptGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **scriptLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **scriptLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **scriptIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **scriptNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **scriptLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **scriptNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **scriptIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **scriptNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **scriptDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **scriptAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByOutputTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByOutputTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobBuildTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobBuildTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobTestTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobTestTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobPublishTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobPublishTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobDeployTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobDeployTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobValidateTaskIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByJobValidateTaskIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfTask**](ResponseWithGenericOfTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTask

> ResponseWithGenericOfTask PatchTask(ctx, primaryKey).Task(task).Depth(depth).Execute()



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
	task := *openapiclient.NewTask() // Task | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PatchTask(context.Background(), primaryKey).Task(task).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PatchTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTask`: ResponseWithGenericOfTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PatchTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | [**Task**](Task.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTask**](ResponseWithGenericOfTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTasks

> ResponseWithGenericOfTask PostTasks(ctx).Task(task).Depth(depth).Execute()



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
	task := []openapiclient.Task{*openapiclient.NewTask()} // []Task | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskAPI.PostTasks(context.Background()).Task(task).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.PostTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTasks`: ResponseWithGenericOfTask
	fmt.Fprintf(os.Stdout, "Response from `TaskAPI.PostTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | [**[]Task**](Task.md) |  | 
 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfTask**](ResponseWithGenericOfTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

