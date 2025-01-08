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

> ResponseWithGenericOfJob GetJobs(ctx).Limit(limit).Offset(offset).Depth(depth).RepositoryLoad(repositoryLoad).ReferencedByDependsOnLoad(referencedByDependsOnLoad).ReferencedByExecutionLoad(referencedByExecutionLoad).ReferencedByTaskLoad(referencedByTaskLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameContains(nameContains).NameNotcontains(nameNotcontains).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).BranchesEq(branchesEq).BranchesNe(branchesNe).BranchesGt(branchesGt).BranchesGte(branchesGte).BranchesLt(branchesLt).BranchesLte(branchesLte).BranchesIn(branchesIn).BranchesNotin(branchesNotin).BranchesContains(branchesContains).BranchesNotcontains(branchesNotcontains).BranchesLike(branchesLike).BranchesNotlike(branchesNotlike).BranchesIlike(branchesIlike).BranchesNotilike(branchesNotilike).BranchesDesc(branchesDesc).BranchesAsc(branchesAsc).TagsEq(tagsEq).TagsNe(tagsNe).TagsGt(tagsGt).TagsGte(tagsGte).TagsLt(tagsLt).TagsLte(tagsLte).TagsIn(tagsIn).TagsNotin(tagsNotin).TagsContains(tagsContains).TagsNotcontains(tagsNotcontains).TagsLike(tagsLike).TagsNotlike(tagsNotlike).TagsIlike(tagsIlike).TagsNotilike(tagsNotilike).TagsDesc(tagsDesc).TagsAsc(tagsAsc).RepositoryIdEq(repositoryIdEq).RepositoryIdNe(repositoryIdNe).RepositoryIdGt(repositoryIdGt).RepositoryIdGte(repositoryIdGte).RepositoryIdLt(repositoryIdLt).RepositoryIdLte(repositoryIdLte).RepositoryIdIn(repositoryIdIn).RepositoryIdNotin(repositoryIdNotin).RepositoryIdContains(repositoryIdContains).RepositoryIdNotcontains(repositoryIdNotcontains).RepositoryIdLike(repositoryIdLike).RepositoryIdNotlike(repositoryIdNotlike).RepositoryIdIlike(repositoryIdIlike).RepositoryIdNotilike(repositoryIdNotilike).RepositoryIdDesc(repositoryIdDesc).RepositoryIdAsc(repositoryIdAsc).RepositoryIdObjectContains(repositoryIdObjectContains).RepositoryIdObjectNotcontains(repositoryIdObjectNotcontains).RepositoryIdObjectDesc(repositoryIdObjectDesc).RepositoryIdObjectAsc(repositoryIdObjectAsc).ReferencedByDependsOnSourceJobIdObjectsContains(referencedByDependsOnSourceJobIdObjectsContains).ReferencedByDependsOnSourceJobIdObjectsNotcontains(referencedByDependsOnSourceJobIdObjectsNotcontains).ReferencedByDependsOnSourceJobIdObjectsDesc(referencedByDependsOnSourceJobIdObjectsDesc).ReferencedByDependsOnSourceJobIdObjectsAsc(referencedByDependsOnSourceJobIdObjectsAsc).ReferencedByDependsOnSinkJobIdObjectsContains(referencedByDependsOnSinkJobIdObjectsContains).ReferencedByDependsOnSinkJobIdObjectsNotcontains(referencedByDependsOnSinkJobIdObjectsNotcontains).ReferencedByDependsOnSinkJobIdObjectsDesc(referencedByDependsOnSinkJobIdObjectsDesc).ReferencedByDependsOnSinkJobIdObjectsAsc(referencedByDependsOnSinkJobIdObjectsAsc).ReferencedByExecutionJobIdObjectsContains(referencedByExecutionJobIdObjectsContains).ReferencedByExecutionJobIdObjectsNotcontains(referencedByExecutionJobIdObjectsNotcontains).ReferencedByExecutionJobIdObjectsDesc(referencedByExecutionJobIdObjectsDesc).ReferencedByExecutionJobIdObjectsAsc(referencedByExecutionJobIdObjectsAsc).ReferencedByTaskJobIdObjectsContains(referencedByTaskJobIdObjectsContains).ReferencedByTaskJobIdObjectsNotcontains(referencedByTaskJobIdObjectsNotcontains).ReferencedByTaskJobIdObjectsDesc(referencedByTaskJobIdObjectsDesc).ReferencedByTaskJobIdObjectsAsc(referencedByTaskJobIdObjectsAsc).Execute()



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
	repositoryLoad := "repositoryLoad_example" // string | load the given directly related object, value is ignored (presence of key is sufficient) (optional)
	referencedByDependsOnLoad := "referencedByDependsOnLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionLoad := "referencedByExecutionLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
	referencedByTaskLoad := "referencedByTaskLoad_example" // string | load the given indirectly related objects, value is ignored (presence of key is sufficient) (optional)
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
	nameEq := "nameEq_example" // string | SQL = comparison (optional)
	nameNe := "nameNe_example" // string | SQL != comparison (optional)
	nameGt := "nameGt_example" // string | SQL > comparison, may not work with all column types (optional)
	nameGte := "nameGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	nameLt := "nameLt_example" // string | SQL < comparison, may not work with all column types (optional)
	nameLte := "nameLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	nameIn := "nameIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	nameNotin := "nameNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	nameContains := "nameContains_example" // string | SQL @> comparison (optional)
	nameNotcontains := "nameNotcontains_example" // string | SQL NOT @> comparison (optional)
	nameLike := "nameLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotlike := "nameNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameIlike := "nameIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameNotilike := "nameNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	nameDesc := "nameDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	nameAsc := "nameAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	branchesEq := "branchesEq_example" // string | SQL = comparison (optional)
	branchesNe := "branchesNe_example" // string | SQL != comparison (optional)
	branchesGt := "branchesGt_example" // string | SQL > comparison, may not work with all column types (optional)
	branchesGte := "branchesGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	branchesLt := "branchesLt_example" // string | SQL < comparison, may not work with all column types (optional)
	branchesLte := "branchesLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	branchesIn := "branchesIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	branchesNotin := "branchesNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	branchesContains := "branchesContains_example" // string | SQL @> comparison (optional)
	branchesNotcontains := "branchesNotcontains_example" // string | SQL NOT @> comparison (optional)
	branchesLike := "branchesLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchesNotlike := "branchesNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchesIlike := "branchesIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchesNotilike := "branchesNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	branchesDesc := "branchesDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	branchesAsc := "branchesAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	tagsEq := "tagsEq_example" // string | SQL = comparison (optional)
	tagsNe := "tagsNe_example" // string | SQL != comparison (optional)
	tagsGt := "tagsGt_example" // string | SQL > comparison, may not work with all column types (optional)
	tagsGte := "tagsGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	tagsLt := "tagsLt_example" // string | SQL < comparison, may not work with all column types (optional)
	tagsLte := "tagsLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	tagsIn := "tagsIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	tagsNotin := "tagsNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	tagsContains := "tagsContains_example" // string | SQL @> comparison (optional)
	tagsNotcontains := "tagsNotcontains_example" // string | SQL NOT @> comparison (optional)
	tagsLike := "tagsLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	tagsNotlike := "tagsNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	tagsIlike := "tagsIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	tagsNotilike := "tagsNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	tagsDesc := "tagsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	tagsAsc := "tagsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdEq := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL = comparison (optional)
	repositoryIdNe := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL != comparison (optional)
	repositoryIdGt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL > comparison, may not work with all column types (optional)
	repositoryIdGte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL >= comparison, may not work with all column types (optional)
	repositoryIdLt := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL < comparison, may not work with all column types (optional)
	repositoryIdLte := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL <= comparison, may not work with all column types (optional)
	repositoryIdIn := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL IN comparison, permits comma-separated values (optional)
	repositoryIdNotin := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	repositoryIdContains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL @> comparison (optional)
	repositoryIdNotcontains := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT @> comparison (optional)
	repositoryIdLike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdNotlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdIlike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdNotilike := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	repositoryIdDesc := "repositoryIdDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdAsc := "repositoryIdAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdObjectContains := TODO // interface{} | SQL @> comparison (optional)
	repositoryIdObjectNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	repositoryIdObjectDesc := "repositoryIdObjectDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	repositoryIdObjectAsc := "repositoryIdObjectAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByDependsOnSourceJobIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByDependsOnSourceJobIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByDependsOnSourceJobIdObjectsDesc := "referencedByDependsOnSourceJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByDependsOnSourceJobIdObjectsAsc := "referencedByDependsOnSourceJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByDependsOnSinkJobIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByDependsOnSinkJobIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByDependsOnSinkJobIdObjectsDesc := "referencedByDependsOnSinkJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByDependsOnSinkJobIdObjectsAsc := "referencedByDependsOnSinkJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionJobIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByExecutionJobIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByExecutionJobIdObjectsDesc := "referencedByExecutionJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByExecutionJobIdObjectsAsc := "referencedByExecutionJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTaskJobIdObjectsContains := TODO // interface{} | SQL @> comparison (optional)
	referencedByTaskJobIdObjectsNotcontains := TODO // interface{} | SQL NOT @> comparison (optional)
	referencedByTaskJobIdObjectsDesc := "referencedByTaskJobIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByTaskJobIdObjectsAsc := "referencedByTaskJobIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JobAPI.GetJobs(context.Background()).Limit(limit).Offset(offset).Depth(depth).RepositoryLoad(repositoryLoad).ReferencedByDependsOnLoad(referencedByDependsOnLoad).ReferencedByExecutionLoad(referencedByExecutionLoad).ReferencedByTaskLoad(referencedByTaskLoad).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdContains(idContains).IdNotcontains(idNotcontains).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtContains(createdAtContains).CreatedAtNotcontains(createdAtNotcontains).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtContains(updatedAtContains).UpdatedAtNotcontains(updatedAtNotcontains).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtContains(deletedAtContains).DeletedAtNotcontains(deletedAtNotcontains).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).NameEq(nameEq).NameNe(nameNe).NameGt(nameGt).NameGte(nameGte).NameLt(nameLt).NameLte(nameLte).NameIn(nameIn).NameNotin(nameNotin).NameContains(nameContains).NameNotcontains(nameNotcontains).NameLike(nameLike).NameNotlike(nameNotlike).NameIlike(nameIlike).NameNotilike(nameNotilike).NameDesc(nameDesc).NameAsc(nameAsc).BranchesEq(branchesEq).BranchesNe(branchesNe).BranchesGt(branchesGt).BranchesGte(branchesGte).BranchesLt(branchesLt).BranchesLte(branchesLte).BranchesIn(branchesIn).BranchesNotin(branchesNotin).BranchesContains(branchesContains).BranchesNotcontains(branchesNotcontains).BranchesLike(branchesLike).BranchesNotlike(branchesNotlike).BranchesIlike(branchesIlike).BranchesNotilike(branchesNotilike).BranchesDesc(branchesDesc).BranchesAsc(branchesAsc).TagsEq(tagsEq).TagsNe(tagsNe).TagsGt(tagsGt).TagsGte(tagsGte).TagsLt(tagsLt).TagsLte(tagsLte).TagsIn(tagsIn).TagsNotin(tagsNotin).TagsContains(tagsContains).TagsNotcontains(tagsNotcontains).TagsLike(tagsLike).TagsNotlike(tagsNotlike).TagsIlike(tagsIlike).TagsNotilike(tagsNotilike).TagsDesc(tagsDesc).TagsAsc(tagsAsc).RepositoryIdEq(repositoryIdEq).RepositoryIdNe(repositoryIdNe).RepositoryIdGt(repositoryIdGt).RepositoryIdGte(repositoryIdGte).RepositoryIdLt(repositoryIdLt).RepositoryIdLte(repositoryIdLte).RepositoryIdIn(repositoryIdIn).RepositoryIdNotin(repositoryIdNotin).RepositoryIdContains(repositoryIdContains).RepositoryIdNotcontains(repositoryIdNotcontains).RepositoryIdLike(repositoryIdLike).RepositoryIdNotlike(repositoryIdNotlike).RepositoryIdIlike(repositoryIdIlike).RepositoryIdNotilike(repositoryIdNotilike).RepositoryIdDesc(repositoryIdDesc).RepositoryIdAsc(repositoryIdAsc).RepositoryIdObjectContains(repositoryIdObjectContains).RepositoryIdObjectNotcontains(repositoryIdObjectNotcontains).RepositoryIdObjectDesc(repositoryIdObjectDesc).RepositoryIdObjectAsc(repositoryIdObjectAsc).ReferencedByDependsOnSourceJobIdObjectsContains(referencedByDependsOnSourceJobIdObjectsContains).ReferencedByDependsOnSourceJobIdObjectsNotcontains(referencedByDependsOnSourceJobIdObjectsNotcontains).ReferencedByDependsOnSourceJobIdObjectsDesc(referencedByDependsOnSourceJobIdObjectsDesc).ReferencedByDependsOnSourceJobIdObjectsAsc(referencedByDependsOnSourceJobIdObjectsAsc).ReferencedByDependsOnSinkJobIdObjectsContains(referencedByDependsOnSinkJobIdObjectsContains).ReferencedByDependsOnSinkJobIdObjectsNotcontains(referencedByDependsOnSinkJobIdObjectsNotcontains).ReferencedByDependsOnSinkJobIdObjectsDesc(referencedByDependsOnSinkJobIdObjectsDesc).ReferencedByDependsOnSinkJobIdObjectsAsc(referencedByDependsOnSinkJobIdObjectsAsc).ReferencedByExecutionJobIdObjectsContains(referencedByExecutionJobIdObjectsContains).ReferencedByExecutionJobIdObjectsNotcontains(referencedByExecutionJobIdObjectsNotcontains).ReferencedByExecutionJobIdObjectsDesc(referencedByExecutionJobIdObjectsDesc).ReferencedByExecutionJobIdObjectsAsc(referencedByExecutionJobIdObjectsAsc).ReferencedByTaskJobIdObjectsContains(referencedByTaskJobIdObjectsContains).ReferencedByTaskJobIdObjectsNotcontains(referencedByTaskJobIdObjectsNotcontains).ReferencedByTaskJobIdObjectsDesc(referencedByTaskJobIdObjectsDesc).ReferencedByTaskJobIdObjectsAsc(referencedByTaskJobIdObjectsAsc).Execute()
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
 **repositoryLoad** | **string** | load the given directly related object, value is ignored (presence of key is sufficient) | 
 **referencedByDependsOnLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
 **referencedByTaskLoad** | **string** | load the given indirectly related objects, value is ignored (presence of key is sufficient) | 
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
 **nameEq** | **string** | SQL &#x3D; comparison | 
 **nameNe** | **string** | SQL !&#x3D; comparison | 
 **nameGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **nameGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **nameLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **nameLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **nameIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **nameNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **nameContains** | **string** | SQL @&gt; comparison | 
 **nameNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **nameLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **nameDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **nameAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **branchesEq** | **string** | SQL &#x3D; comparison | 
 **branchesNe** | **string** | SQL !&#x3D; comparison | 
 **branchesGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **branchesGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **branchesLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **branchesLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **branchesIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **branchesNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **branchesContains** | **string** | SQL @&gt; comparison | 
 **branchesNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **branchesLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchesNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchesIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchesNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **branchesDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **branchesAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **tagsEq** | **string** | SQL &#x3D; comparison | 
 **tagsNe** | **string** | SQL !&#x3D; comparison | 
 **tagsGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **tagsGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **tagsLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **tagsLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **tagsIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **tagsNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **tagsContains** | **string** | SQL @&gt; comparison | 
 **tagsNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **tagsLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **tagsNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **tagsIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **tagsNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **tagsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **tagsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdEq** | **string** | SQL &#x3D; comparison | 
 **repositoryIdNe** | **string** | SQL !&#x3D; comparison | 
 **repositoryIdGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **repositoryIdGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **repositoryIdLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **repositoryIdLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **repositoryIdIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **repositoryIdNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **repositoryIdContains** | **string** | SQL @&gt; comparison | 
 **repositoryIdNotcontains** | **string** | SQL NOT @&gt; comparison | 
 **repositoryIdLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **repositoryIdDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdObjectContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **repositoryIdObjectNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **repositoryIdObjectDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **repositoryIdObjectAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByDependsOnSourceJobIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByDependsOnSourceJobIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByDependsOnSourceJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByDependsOnSourceJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByDependsOnSinkJobIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByDependsOnSinkJobIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByDependsOnSinkJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByDependsOnSinkJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionJobIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByExecutionJobIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByExecutionJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByExecutionJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTaskJobIdObjectsContains** | [**interface{}**](interface{}.md) | SQL @&gt; comparison | 
 **referencedByTaskJobIdObjectsNotcontains** | [**interface{}**](interface{}.md) | SQL NOT @&gt; comparison | 
 **referencedByTaskJobIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByTaskJobIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

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

