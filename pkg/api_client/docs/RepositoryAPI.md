# \RepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRepository**](RepositoryAPI.md#DeleteRepository) | **Delete** /api/repositories/{primaryKey} | 
[**GetRepositories**](RepositoryAPI.md#GetRepositories) | **Get** /api/repositories | 
[**GetRepository**](RepositoryAPI.md#GetRepository) | **Get** /api/repositories/{primaryKey} | 
[**PatchRepository**](RepositoryAPI.md#PatchRepository) | **Patch** /api/repositories/{primaryKey} | 
[**PostRepositories**](RepositoryAPI.md#PostRepositories) | **Post** /api/repositories | 



## DeleteRepository

> DeleteRepository(ctx, primaryKey).Depth(depth).Execute()



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
	r, err := apiClient.RepositoryAPI.DeleteRepository(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.DeleteRepository``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


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


## GetRepositories

> ResponseWithGenericOfRepository GetRepositories(ctx).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).LastSyncedEq(lastSyncedEq).LastSyncedNe(lastSyncedNe).LastSyncedGt(lastSyncedGt).LastSyncedGte(lastSyncedGte).LastSyncedLt(lastSyncedLt).LastSyncedLte(lastSyncedLte).LastSyncedIn(lastSyncedIn).LastSyncedNotin(lastSyncedNotin).LastSyncedLike(lastSyncedLike).LastSyncedNotlike(lastSyncedNotlike).LastSyncedIlike(lastSyncedIlike).LastSyncedNotilike(lastSyncedNotilike).LastSyncedDesc(lastSyncedDesc).LastSyncedAsc(lastSyncedAsc).UrlEq(urlEq).UrlNe(urlNe).UrlGt(urlGt).UrlGte(urlGte).UrlLt(urlLt).UrlLte(urlLte).UrlIn(urlIn).UrlNotin(urlNotin).UrlLike(urlLike).UrlNotlike(urlNotlike).UrlIlike(urlIlike).UrlNotilike(urlNotilike).UrlDesc(urlDesc).UrlAsc(urlAsc).UsernameEq(usernameEq).UsernameNe(usernameNe).UsernameGt(usernameGt).UsernameGte(usernameGte).UsernameLt(usernameLt).UsernameLte(usernameLte).UsernameIn(usernameIn).UsernameNotin(usernameNotin).UsernameLike(usernameLike).UsernameNotlike(usernameNotlike).UsernameIlike(usernameIlike).UsernameNotilike(usernameNotilike).UsernameDesc(usernameDesc).UsernameAsc(usernameAsc).PasswordEq(passwordEq).PasswordNe(passwordNe).PasswordGt(passwordGt).PasswordGte(passwordGte).PasswordLt(passwordLt).PasswordLte(passwordLte).PasswordIn(passwordIn).PasswordNotin(passwordNotin).PasswordLike(passwordLike).PasswordNotlike(passwordNotlike).PasswordIlike(passwordIlike).PasswordNotilike(passwordNotilike).PasswordDesc(passwordDesc).PasswordAsc(passwordAsc).SshKeyEq(sshKeyEq).SshKeyNe(sshKeyNe).SshKeyGt(sshKeyGt).SshKeyGte(sshKeyGte).SshKeyLt(sshKeyLt).SshKeyLte(sshKeyLte).SshKeyIn(sshKeyIn).SshKeyNotin(sshKeyNotin).SshKeyLike(sshKeyLike).SshKeyNotlike(sshKeyNotlike).SshKeyIlike(sshKeyIlike).SshKeyNotilike(sshKeyNotilike).SshKeyDesc(sshKeyDesc).SshKeyAsc(sshKeyAsc).ReferencedByChangeRepositoryIdObjectsDesc(referencedByChangeRepositoryIdObjectsDesc).ReferencedByChangeRepositoryIdObjectsAsc(referencedByChangeRepositoryIdObjectsAsc).ReferencedByRuleRepositoryIdObjectsDesc(referencedByRuleRepositoryIdObjectsDesc).ReferencedByRuleRepositoryIdObjectsAsc(referencedByRuleRepositoryIdObjectsAsc).Execute()



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
	lastSyncedEq := time.Now() // time.Time | SQL = comparison (optional)
	lastSyncedNe := time.Now() // time.Time | SQL != comparison (optional)
	lastSyncedGt := time.Now() // time.Time | SQL > comparison, may not work with all column types (optional)
	lastSyncedGte := time.Now() // time.Time | SQL >= comparison, may not work with all column types (optional)
	lastSyncedLt := time.Now() // time.Time | SQL < comparison, may not work with all column types (optional)
	lastSyncedLte := time.Now() // time.Time | SQL <= comparison, may not work with all column types (optional)
	lastSyncedIn := time.Now() // time.Time | SQL IN comparison, permits comma-separated values (optional)
	lastSyncedNotin := time.Now() // time.Time | SQL NOT IN comparison, permits comma-separated values (optional)
	lastSyncedLike := time.Now() // time.Time | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	lastSyncedNotlike := time.Now() // time.Time | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	lastSyncedIlike := time.Now() // time.Time | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	lastSyncedNotilike := time.Now() // time.Time | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	lastSyncedDesc := "lastSyncedDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	lastSyncedAsc := "lastSyncedAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	urlEq := "urlEq_example" // string | SQL = comparison (optional)
	urlNe := "urlNe_example" // string | SQL != comparison (optional)
	urlGt := "urlGt_example" // string | SQL > comparison, may not work with all column types (optional)
	urlGte := "urlGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	urlLt := "urlLt_example" // string | SQL < comparison, may not work with all column types (optional)
	urlLte := "urlLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	urlIn := "urlIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	urlNotin := "urlNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	urlLike := "urlLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlNotlike := "urlNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlIlike := "urlIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlNotilike := "urlNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	urlDesc := "urlDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	urlAsc := "urlAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	usernameEq := "usernameEq_example" // string | SQL = comparison (optional)
	usernameNe := "usernameNe_example" // string | SQL != comparison (optional)
	usernameGt := "usernameGt_example" // string | SQL > comparison, may not work with all column types (optional)
	usernameGte := "usernameGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	usernameLt := "usernameLt_example" // string | SQL < comparison, may not work with all column types (optional)
	usernameLte := "usernameLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	usernameIn := "usernameIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	usernameNotin := "usernameNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	usernameLike := "usernameLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	usernameNotlike := "usernameNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	usernameIlike := "usernameIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	usernameNotilike := "usernameNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	usernameDesc := "usernameDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	usernameAsc := "usernameAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	passwordEq := "passwordEq_example" // string | SQL = comparison (optional)
	passwordNe := "passwordNe_example" // string | SQL != comparison (optional)
	passwordGt := "passwordGt_example" // string | SQL > comparison, may not work with all column types (optional)
	passwordGte := "passwordGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	passwordLt := "passwordLt_example" // string | SQL < comparison, may not work with all column types (optional)
	passwordLte := "passwordLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	passwordIn := "passwordIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	passwordNotin := "passwordNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	passwordLike := "passwordLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	passwordNotlike := "passwordNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	passwordIlike := "passwordIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	passwordNotilike := "passwordNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	passwordDesc := "passwordDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	passwordAsc := "passwordAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	sshKeyEq := "sshKeyEq_example" // string | SQL = comparison (optional)
	sshKeyNe := "sshKeyNe_example" // string | SQL != comparison (optional)
	sshKeyGt := "sshKeyGt_example" // string | SQL > comparison, may not work with all column types (optional)
	sshKeyGte := "sshKeyGte_example" // string | SQL >= comparison, may not work with all column types (optional)
	sshKeyLt := "sshKeyLt_example" // string | SQL < comparison, may not work with all column types (optional)
	sshKeyLte := "sshKeyLte_example" // string | SQL <= comparison, may not work with all column types (optional)
	sshKeyIn := "sshKeyIn_example" // string | SQL IN comparison, permits comma-separated values (optional)
	sshKeyNotin := "sshKeyNotin_example" // string | SQL NOT IN comparison, permits comma-separated values (optional)
	sshKeyLike := "sshKeyLike_example" // string | SQL LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	sshKeyNotlike := "sshKeyNotlike_example" // string | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	sshKeyIlike := "sshKeyIlike_example" // string | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	sshKeyNotilike := "sshKeyNotilike_example" // string | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % (optional)
	sshKeyDesc := "sshKeyDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	sshKeyAsc := "sshKeyAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByChangeRepositoryIdObjectsDesc := "referencedByChangeRepositoryIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByChangeRepositoryIdObjectsAsc := "referencedByChangeRepositoryIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByRuleRepositoryIdObjectsDesc := "referencedByRuleRepositoryIdObjectsDesc_example" // string | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) (optional)
	referencedByRuleRepositoryIdObjectsAsc := "referencedByRuleRepositoryIdObjectsAsc_example" // string | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.GetRepositories(context.Background()).Limit(limit).Offset(offset).Depth(depth).IdEq(idEq).IdNe(idNe).IdGt(idGt).IdGte(idGte).IdLt(idLt).IdLte(idLte).IdIn(idIn).IdNotin(idNotin).IdLike(idLike).IdNotlike(idNotlike).IdIlike(idIlike).IdNotilike(idNotilike).IdDesc(idDesc).IdAsc(idAsc).CreatedAtEq(createdAtEq).CreatedAtNe(createdAtNe).CreatedAtGt(createdAtGt).CreatedAtGte(createdAtGte).CreatedAtLt(createdAtLt).CreatedAtLte(createdAtLte).CreatedAtIn(createdAtIn).CreatedAtNotin(createdAtNotin).CreatedAtLike(createdAtLike).CreatedAtNotlike(createdAtNotlike).CreatedAtIlike(createdAtIlike).CreatedAtNotilike(createdAtNotilike).CreatedAtDesc(createdAtDesc).CreatedAtAsc(createdAtAsc).UpdatedAtEq(updatedAtEq).UpdatedAtNe(updatedAtNe).UpdatedAtGt(updatedAtGt).UpdatedAtGte(updatedAtGte).UpdatedAtLt(updatedAtLt).UpdatedAtLte(updatedAtLte).UpdatedAtIn(updatedAtIn).UpdatedAtNotin(updatedAtNotin).UpdatedAtLike(updatedAtLike).UpdatedAtNotlike(updatedAtNotlike).UpdatedAtIlike(updatedAtIlike).UpdatedAtNotilike(updatedAtNotilike).UpdatedAtDesc(updatedAtDesc).UpdatedAtAsc(updatedAtAsc).DeletedAtEq(deletedAtEq).DeletedAtNe(deletedAtNe).DeletedAtGt(deletedAtGt).DeletedAtGte(deletedAtGte).DeletedAtLt(deletedAtLt).DeletedAtLte(deletedAtLte).DeletedAtIn(deletedAtIn).DeletedAtNotin(deletedAtNotin).DeletedAtLike(deletedAtLike).DeletedAtNotlike(deletedAtNotlike).DeletedAtIlike(deletedAtIlike).DeletedAtNotilike(deletedAtNotilike).DeletedAtDesc(deletedAtDesc).DeletedAtAsc(deletedAtAsc).LastSyncedEq(lastSyncedEq).LastSyncedNe(lastSyncedNe).LastSyncedGt(lastSyncedGt).LastSyncedGte(lastSyncedGte).LastSyncedLt(lastSyncedLt).LastSyncedLte(lastSyncedLte).LastSyncedIn(lastSyncedIn).LastSyncedNotin(lastSyncedNotin).LastSyncedLike(lastSyncedLike).LastSyncedNotlike(lastSyncedNotlike).LastSyncedIlike(lastSyncedIlike).LastSyncedNotilike(lastSyncedNotilike).LastSyncedDesc(lastSyncedDesc).LastSyncedAsc(lastSyncedAsc).UrlEq(urlEq).UrlNe(urlNe).UrlGt(urlGt).UrlGte(urlGte).UrlLt(urlLt).UrlLte(urlLte).UrlIn(urlIn).UrlNotin(urlNotin).UrlLike(urlLike).UrlNotlike(urlNotlike).UrlIlike(urlIlike).UrlNotilike(urlNotilike).UrlDesc(urlDesc).UrlAsc(urlAsc).UsernameEq(usernameEq).UsernameNe(usernameNe).UsernameGt(usernameGt).UsernameGte(usernameGte).UsernameLt(usernameLt).UsernameLte(usernameLte).UsernameIn(usernameIn).UsernameNotin(usernameNotin).UsernameLike(usernameLike).UsernameNotlike(usernameNotlike).UsernameIlike(usernameIlike).UsernameNotilike(usernameNotilike).UsernameDesc(usernameDesc).UsernameAsc(usernameAsc).PasswordEq(passwordEq).PasswordNe(passwordNe).PasswordGt(passwordGt).PasswordGte(passwordGte).PasswordLt(passwordLt).PasswordLte(passwordLte).PasswordIn(passwordIn).PasswordNotin(passwordNotin).PasswordLike(passwordLike).PasswordNotlike(passwordNotlike).PasswordIlike(passwordIlike).PasswordNotilike(passwordNotilike).PasswordDesc(passwordDesc).PasswordAsc(passwordAsc).SshKeyEq(sshKeyEq).SshKeyNe(sshKeyNe).SshKeyGt(sshKeyGt).SshKeyGte(sshKeyGte).SshKeyLt(sshKeyLt).SshKeyLte(sshKeyLte).SshKeyIn(sshKeyIn).SshKeyNotin(sshKeyNotin).SshKeyLike(sshKeyLike).SshKeyNotlike(sshKeyNotlike).SshKeyIlike(sshKeyIlike).SshKeyNotilike(sshKeyNotilike).SshKeyDesc(sshKeyDesc).SshKeyAsc(sshKeyAsc).ReferencedByChangeRepositoryIdObjectsDesc(referencedByChangeRepositoryIdObjectsDesc).ReferencedByChangeRepositoryIdObjectsAsc(referencedByChangeRepositoryIdObjectsAsc).ReferencedByRuleRepositoryIdObjectsDesc(referencedByRuleRepositoryIdObjectsDesc).ReferencedByRuleRepositoryIdObjectsAsc(referencedByRuleRepositoryIdObjectsAsc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


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
 **lastSyncedEq** | **time.Time** | SQL &#x3D; comparison | 
 **lastSyncedNe** | **time.Time** | SQL !&#x3D; comparison | 
 **lastSyncedGt** | **time.Time** | SQL &gt; comparison, may not work with all column types | 
 **lastSyncedGte** | **time.Time** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **lastSyncedLt** | **time.Time** | SQL &lt; comparison, may not work with all column types | 
 **lastSyncedLte** | **time.Time** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **lastSyncedIn** | **time.Time** | SQL IN comparison, permits comma-separated values | 
 **lastSyncedNotin** | **time.Time** | SQL NOT IN comparison, permits comma-separated values | 
 **lastSyncedLike** | **time.Time** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **lastSyncedNotlike** | **time.Time** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **lastSyncedIlike** | **time.Time** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **lastSyncedNotilike** | **time.Time** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **lastSyncedDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **lastSyncedAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **urlEq** | **string** | SQL &#x3D; comparison | 
 **urlNe** | **string** | SQL !&#x3D; comparison | 
 **urlGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **urlGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **urlLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **urlLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **urlIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **urlNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **urlLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **urlDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **urlAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **usernameEq** | **string** | SQL &#x3D; comparison | 
 **usernameNe** | **string** | SQL !&#x3D; comparison | 
 **usernameGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **usernameGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **usernameLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **usernameLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **usernameIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **usernameNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **usernameLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **usernameNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **usernameIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **usernameNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **usernameDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **usernameAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **passwordEq** | **string** | SQL &#x3D; comparison | 
 **passwordNe** | **string** | SQL !&#x3D; comparison | 
 **passwordGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **passwordGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **passwordLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **passwordLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **passwordIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **passwordNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **passwordLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **passwordNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **passwordIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **passwordNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **passwordDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **passwordAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **sshKeyEq** | **string** | SQL &#x3D; comparison | 
 **sshKeyNe** | **string** | SQL !&#x3D; comparison | 
 **sshKeyGt** | **string** | SQL &gt; comparison, may not work with all column types | 
 **sshKeyGte** | **string** | SQL &gt;&#x3D; comparison, may not work with all column types | 
 **sshKeyLt** | **string** | SQL &lt; comparison, may not work with all column types | 
 **sshKeyLte** | **string** | SQL &lt;&#x3D; comparison, may not work with all column types | 
 **sshKeyIn** | **string** | SQL IN comparison, permits comma-separated values | 
 **sshKeyNotin** | **string** | SQL NOT IN comparison, permits comma-separated values | 
 **sshKeyLike** | **string** | SQL LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **sshKeyNotlike** | **string** | SQL NOT LIKE comparison, value is implicitly prefixed and suffixed with % | 
 **sshKeyIlike** | **string** | SQL ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **sshKeyNotilike** | **string** | SQL NOT ILIKE comparison, value is implicitly prefixed and suffixed with % | 
 **sshKeyDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **sshKeyAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByChangeRepositoryIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByChangeRepositoryIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 
 **referencedByRuleRepositoryIdObjectsDesc** | **string** | SQL ORDER BY _ DESC clause, value is ignored (presence of key is sufficient) | 
 **referencedByRuleRepositoryIdObjectsAsc** | **string** | SQL ORDER BY _ ASC clause, value is ignored (presence of key is sufficient) | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> ResponseWithGenericOfRepository GetRepository(ctx, primaryKey).Depth(depth).Execute()



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
	resp, r, err := apiClient.RepositoryAPI.GetRepository(context.Background(), primaryKey).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **depth** | **int64** | Query parameter depth | 

### Return type

[**ResponseWithGenericOfRepository**](ResponseWithGenericOfRepository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRepository

> ResponseWithGenericOfRepository PatchRepository(ctx, primaryKey).Repository(repository).Depth(depth).Execute()



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
	repository := *openapiclient.NewRepository() // Repository | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.PatchRepository(context.Background(), primaryKey).Repository(repository).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.PatchRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchRepository`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.PatchRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**primaryKey** | **string** | Path parameter primaryKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | [**Repository**](Repository.md) |  | 
 **depth** | **int64** | Query parameter depth | 

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


## PostRepositories

> ResponseWithGenericOfRepository PostRepositories(ctx).Repository(repository).Depth(depth).Execute()



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
	repository := []openapiclient.Repository{*openapiclient.NewRepository()} // []Repository | 
	depth := int64(789) // int64 | Query parameter depth (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryAPI.PostRepositories(context.Background()).Repository(repository).Depth(depth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryAPI.PostRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRepositories`: ResponseWithGenericOfRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryAPI.PostRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | [**[]Repository**](Repository.md) |  | 
 **depth** | **int64** | Query parameter depth | 

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

