# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCheckClosedGates**](MigrationGateControllerApi.md#BulkCheckClosedGates) | **Post** /gates/gate/bulk/check | Bulk check for closed Gates by their ID&#x27;s
[**CheckClosedGates**](MigrationGateControllerApi.md#CheckClosedGates) | **Get** /gates/gate/check | Check for closed Gates in bulk
[**CloseMigrationGate**](MigrationGateControllerApi.md#CloseMigrationGate) | **Put** /gates/gate/close | Close a Gate by its ID
[**CreateMigrationGate**](MigrationGateControllerApi.md#CreateMigrationGate) | **Put** /gates/gate | Request a gate on a given migration by path
[**DeleteAllMigrationGates**](MigrationGateControllerApi.md#DeleteAllMigrationGates) | **Delete** /gates/gate/deleteAll | Delete all Migration Gates
[**DeleteBulkMigrationGates**](MigrationGateControllerApi.md#DeleteBulkMigrationGates) | **Post** /gates/gate/bulk/delete | Bulk delete migration gates by their ID&#x27;s
[**DeleteMigrationGates**](MigrationGateControllerApi.md#DeleteMigrationGates) | **Delete** /gates/gate | Delete a Gate by its ID
[**GetBulkMigrationGates**](MigrationGateControllerApi.md#GetBulkMigrationGates) | **Post** /gates/gate/bulk | Check status of a number of gates by their ID&#x27;s
[**GetMigrationGates**](MigrationGateControllerApi.md#GetMigrationGates) | **Get** /gates/gate | Check status of a gate by its ID

# **BulkCheckClosedGates**
> []string BulkCheckClosedGates(ctx, body)
Bulk check for closed Gates by their ID's

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationGateIds**](MigrationGateIds.md)| Gate IDs | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckClosedGates**
> []string CheckClosedGates(ctx, gateIds)
Check for closed Gates in bulk

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateIds** | [**[]string**](string.md)| Gate ID | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseMigrationGate**
> CloseMigrationGate(ctx, gateId)
Close a Gate by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateId** | **string**| Gate ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMigrationGate**
> MigrationGate CreateMigrationGate(ctx, path, sourceFilesystemId, targetFilesystemId, gateId)
Request a gate on a given migration by path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Path of the Migration gate | 
  **sourceFilesystemId** | **string**| ID of the source Filesystem | 
  **targetFilesystemId** | **string**| ID of the target Filesystem | 
  **gateId** | **string**| Gate ID | 

### Return type

[**MigrationGate**](MigrationGate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllMigrationGates**
> DeleteAllMigrationGates(ctx, )
Delete all Migration Gates

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBulkMigrationGates**
> DeleteBulkMigrationGates(ctx, body)
Bulk delete migration gates by their ID's

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationGateIds**](MigrationGateIds.md)| Gate IDs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMigrationGates**
> DeleteMigrationGates(ctx, gateIds)
Delete a Gate by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateIds** | [**[]string**](string.md)| Gate IDs | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulkMigrationGates**
> []MigrationGate GetBulkMigrationGates(ctx, body)
Check status of a number of gates by their ID's

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationGateIds**](MigrationGateIds.md)| Gate IDs | 

### Return type

[**[]MigrationGate**](MigrationGate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationGates**
> []MigrationGate GetMigrationGates(ctx, gateIds)
Check status of a gate by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gateIds** | [**[]string**](string.md)| Gate ID | 

### Return type

[**[]MigrationGate**](MigrationGate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

