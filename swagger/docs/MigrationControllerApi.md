# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddExclusionToMigration**](MigrationControllerApi.md#AddExclusionToMigration) | **Put** /migrations/{migrationId}/exclusions/{exclusionId} | Add exclusion to migration
[**AddPendingRegionToMigration**](MigrationControllerApi.md#AddPendingRegionToMigration) | **Post** /migrations/{migrationId}/addPendingRegion | Add a path as a pending region to migration, so LDM will rescan that path.
[**CheckMigrationPathMappings**](MigrationControllerApi.md#CheckMigrationPathMappings) | **Post** /migrations/{migrationId}/checkMigrationPathMappings | Check migration path mappings
[**ClearMigrations**](MigrationControllerApi.md#ClearMigrations) | **Delete** /migrations | Delete all migrations
[**DeleteExclusionFromMigration**](MigrationControllerApi.md#DeleteExclusionFromMigration) | **Delete** /migrations/{migrationId}/exclusions/{exclusionId} | Remove exclusion from migration
[**DeleteMigration**](MigrationControllerApi.md#DeleteMigration) | **Delete** /migrations/{migrationId} | Delete migration for given migration ID
[**DeleteMigrationWithAssets**](MigrationControllerApi.md#DeleteMigrationWithAssets) | **Delete** /migrations/{migrationId}/withAssets | Delete migration for given migration ID along with its assets
[**FetchMigration**](MigrationControllerApi.md#FetchMigration) | **Get** /migrations/{migrationId} | Get migration for given migration ID
[**GetOldestIncompleteWorkItem**](MigrationControllerApi.md#GetOldestIncompleteWorkItem) | **Get** /migrations/{migrationId}/oldestIncompleteWorkItem | Get the oldest incomplete work item for a Migration.
[**GetOldestIncompleteWorkItems**](MigrationControllerApi.md#GetOldestIncompleteWorkItems) | **Get** /migrations/oldestIncompleteWorkItems | Get a list of oldest incomplete work items across all migrations.
[**GetPendingRegionsForMigration**](MigrationControllerApi.md#GetPendingRegionsForMigration) | **Get** /migrations/{migrationId}/pendingRegions | Get the PendingRegions for a Migration.
[**ListAutoSourceCleanupQueuedPaths**](MigrationControllerApi.md#ListAutoSourceCleanupQueuedPaths) | **Get** /migrations/{migrationId}/autoSourceCleanupQueuedPaths | List the top 10000 Auto-source Cleanup Paths Queued for deletion.
[**ListMigration**](MigrationControllerApi.md#ListMigration) | **Get** /migrations | List all migrations
[**ListMigrationErrors**](MigrationControllerApi.md#ListMigrationErrors) | **Get** /migrations/{migrationId}/failures | List the last 10000 files which failed to migrate
[**NewMigrationWithId**](MigrationControllerApi.md#NewMigrationWithId) | **Put** /migrations/{migrationId} | Add a new migration
[**ResetMigration**](MigrationControllerApi.md#ResetMigration) | **Post** /migrations/{migrationId}/reset | Reset migration for given migration ID
[**ResumeMigration**](MigrationControllerApi.md#ResumeMigration) | **Post** /migrations/{migrationId}/resume | Resume migration for given migration ID
[**StartMigration**](MigrationControllerApi.md#StartMigration) | **Post** /migrations/{migrationId}/start | Start migration for given migration ID
[**StopAllMigrations**](MigrationControllerApi.md#StopAllMigrations) | **Post** /migrations/stop | Stop all migrations
[**StopMigration**](MigrationControllerApi.md#StopMigration) | **Post** /migrations/{migrationId}/stop | Stop migration for given migration ID
[**ToggleMigrationDeletionMode**](MigrationControllerApi.md#ToggleMigrationDeletionMode) | **Put** /migrations/{migrationId}/toggleDeletionMode | Toggle the migration deletion mode
[**UpdateMigrationMaxPendingRegions**](MigrationControllerApi.md#UpdateMigrationMaxPendingRegions) | **Put** /migrations/{migrationId}/maxPendingRegions/{maxPendingRegions} | Update the maximum number of pending regions for a migration.
[**UpdateMigrationMaxUnscheduledEvents**](MigrationControllerApi.md#UpdateMigrationMaxUnscheduledEvents) | **Put** /migrations/{migrationId}/maxStoppedUnscheduledEventBuffer/{maxEvents} | Update the maximum number of unscheduled events to store for a migration if it is stopped.
[**UpdateMigrationProperties**](MigrationControllerApi.md#UpdateMigrationProperties) | **Patch** /migrations/{migrationId} | Update a migration&#x27;s additional properties

# **AddExclusionToMigration**
> AddExclusionToMigration(ctx, migrationId, exclusionId)
Add exclusion to migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **exclusionId** | **string**| Unique identifier of the exclusion object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPendingRegionToMigration**
> AddPendingRegionToMigration(ctx, migrationId, path, optional)
Add a path as a pending region to migration, so LDM will rescan that path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **path** | **string**| Path of the new pending region | 
 **optional** | ***MigrationControllerApiAddPendingRegionToMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiAddPendingRegionToMigrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionPolicy** | **optional.String**| Overwrite policy for existing files on target | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckMigrationPathMappings**
> MigrationMappingCheckResult CheckMigrationPathMappings(ctx, migrationId, path, target, optional)
Check migration path mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique Identifier for the migration | 
  **path** | **string**| Migration path of the source file system | 
  **target** | **string**| ID of the target file system | 
 **optional** | ***MigrationControllerApiCheckMigrationPathMappingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiCheckMigrationPathMappingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **source** | **optional.String**| ID of the source file system | 

### Return type

[**MigrationMappingCheckResult**](MigrationMappingCheckResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearMigrations**
> ClearMigrations(ctx, )
Delete all migrations

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExclusionFromMigration**
> DeleteExclusionFromMigration(ctx, migrationId, exclusionId)
Remove exclusion from migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **exclusionId** | **string**| Unique identifier of the exclusion object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMigration**
> DeleteMigration(ctx, migrationId)
Delete migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMigrationWithAssets**
> map[string]AssetDeletionStats DeleteMigrationWithAssets(ctx, migrationId)
Delete migration for given migration ID along with its assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**map[string]AssetDeletionStats**](AssetDeletionStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMigration**
> Migration FetchMigration(ctx, migrationId)
Get migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOldestIncompleteWorkItem**
> OldestIncompleteWorkItem GetOldestIncompleteWorkItem(ctx, migrationId)
Get the oldest incomplete work item for a Migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**OldestIncompleteWorkItem**](OldestIncompleteWorkItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOldestIncompleteWorkItems**
> []OldestIncompleteWorkItem GetOldestIncompleteWorkItems(ctx, )
Get a list of oldest incomplete work items across all migrations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OldestIncompleteWorkItem**](OldestIncompleteWorkItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPendingRegionsForMigration**
> []PendingRegion GetPendingRegionsForMigration(ctx, migrationId)
Get the PendingRegions for a Migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**[]PendingRegion**](PendingRegion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAutoSourceCleanupQueuedPaths**
> []SourceCleanupQueuedPath ListAutoSourceCleanupQueuedPaths(ctx, migrationId)
List the top 10000 Auto-source Cleanup Paths Queued for deletion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**[]SourceCleanupQueuedPath**](SourceCleanupQueuedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMigration**
> []Migration ListMigration(ctx, )
List all migrations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMigrationErrors**
> MigrationFailures ListMigrationErrors(ctx, migrationId, optional)
List the last 10000 files which failed to migrate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
 **optional** | ***MigrationControllerApiListMigrationErrorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiListMigrationErrorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **optional.Time**| If a failure is earlier than this, it will be excluded from the list of failed paths. If no value is provided, all known failed paths will be returned. | 

### Return type

[**MigrationFailures**](MigrationFailures.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NewMigrationWithId**
> Migration NewMigrationWithId(ctx, migrationId, path, target, optional)
Add a new migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique Identifier for the migration | 
  **path** | **string**| Migration path of the source file system | 
  **target** | **string**| ID of the target file system | 
 **optional** | ***MigrationControllerApiNewMigrationWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiNewMigrationWithIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **source** | **optional.String**| ID of the source file system | 
 **exclusions** | [**optional.Interface of []string**](string.md)| List of exclusions to apply | 
 **actionPolicy** | **optional.String**| Overwrite policy for existing files on target | 
 **autoStart** | **optional.Bool**| Flag to start migration immediately | 
 **scanOnlyMigration** | **optional.Bool**| Flag to run scan-only migration | 
 **additionalProperties** | [**optional.Interface of map[string]string**](string.md)| Migration additional properties | [default to {}]
 **deletionMode** | **optional.String**| Migration Deletion Mode | 
 **delayedDeletionPeriod** | **optional.Int32**| Delayed Deletion Period in seconds | 
 **autoSourceCleanupEnabled** | **optional.Bool**| Flag to enable auto source clean up actions on the migration | 
 **recurringMigration** | **optional.Bool**| Flag for Migration Recurring Mode | 
 **recurringPeriod** | **optional.Int32**| Recurring Period in seconds | 
 **migrationPriority** | **optional.String**| Migration priority | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetMigration**
> Migration ResetMigration(ctx, migrationId, optional)
Reset migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
 **optional** | ***MigrationControllerApiResetMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiResetMigrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionPolicy** | **optional.String**| Override overwrite policy for existing files on target | 
 **reloadMappings** | **optional.Bool**| Reload path mappings | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeMigration**
> Migration ResumeMigration(ctx, migrationId)
Resume migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartMigration**
> Migration StartMigration(ctx, migrationId)
Start migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopAllMigrations**
> StopAllMigrations(ctx, optional)
Stop all migrations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MigrationControllerApiStopAllMigrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiStopAllMigrationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resumable** | **optional.Bool**| If the migration may be resumed later | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopMigration**
> StopMigration(ctx, migrationId, optional)
Stop migration for given migration ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
 **optional** | ***MigrationControllerApiStopMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiStopMigrationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resumable** | **optional.Bool**| If the migration may be resumed later | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleMigrationDeletionMode**
> Migration ToggleMigrationDeletionMode(ctx, migrationId, deletionMode, actionPolicy, optional)
Toggle the migration deletion mode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **deletionMode** | **string**| Desired Deletion mode to toggle for migration | 
  **actionPolicy** | **string**| Overwrite policy for existing files on target | [default to com.wandisco.livemigrator2.migration.OverwriteActionPolicy]
 **optional** | ***MigrationControllerApiToggleMigrationDeletionModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiToggleMigrationDeletionModeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **delayedDeletionPeriod** | **optional.Int32**| Delayed Deletion Period in seconds | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigrationMaxPendingRegions**
> UpdateMigrationMaxPendingRegions(ctx, migrationId, maxPendingRegions)
Update the maximum number of pending regions for a migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **maxPendingRegions** | **int32**| Maximum number of pending regions | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigrationMaxUnscheduledEvents**
> UpdateMigrationMaxUnscheduledEvents(ctx, migrationId, maxEvents)
Update the maximum number of unscheduled events to store for a migration if it is stopped.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **maxEvents** | **int32**| Maximum number of unscheduled events to store | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMigrationProperties**
> UpdateMigrationProperties(ctx, migrationId, optional)
Update a migration's additional properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
 **optional** | ***MigrationControllerApiUpdateMigrationPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MigrationControllerApiUpdateMigrationPropertiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MigrationPatch**](MigrationPatch.md)| Migration configuration | 
 **additionalProperties** | [**optional.Interface of map[string]string**](string.md)| Migration additional properties | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

