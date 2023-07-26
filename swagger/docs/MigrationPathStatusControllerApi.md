# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPathStatus**](MigrationPathStatusControllerApi.md#GetPathStatus) | **Get** /migration-path-status | Displays information on a path in a current migration.

# **GetPathStatus**
> MigrationPathStatus GetPathStatus(ctx, path, sourceId)
Displays information on a path in a current migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| Path on source filesystem. | 
  **sourceId** | **string**|  | 

### Return type

[**MigrationPathStatus**](MigrationPathStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

