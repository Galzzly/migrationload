# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearTargetFileSystems**](TargetFileSystemControllerApi.md#ClearTargetFileSystems) | **Delete** /fs/targets | Delete all target file systems. NOTE this endpoint is currently for test purposes only.
[**DeleteTargetFileSystem**](TargetFileSystemControllerApi.md#DeleteTargetFileSystem) | **Delete** /fs/targets/{targetId} | Delete target file system for given ID
[**GetAvailableFileSystems**](TargetFileSystemControllerApi.md#GetAvailableFileSystems) | **Get** /fs/targets/availableFileSystems | List all available file systems
[**GetFileSystem**](TargetFileSystemControllerApi.md#GetFileSystem) | **Get** /fs/targets/{targetId} | Get target file system for given ID
[**ListTargetFileSystems**](TargetFileSystemControllerApi.md#ListTargetFileSystems) | **Get** /fs/targets | List all target file systems

# **ClearTargetFileSystems**
> ClearTargetFileSystems(ctx, )
Delete all target file systems. NOTE this endpoint is currently for test purposes only.

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

# **DeleteTargetFileSystem**
> DeleteTargetFileSystem(ctx, targetId)
Delete target file system for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetId** | **string**| Unique identifier of the target file system Object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableFileSystems**
> []FsType GetAvailableFileSystems(ctx, )
List all available file systems

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FsType**](FsType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileSystem**
> FileSystemConfiguration GetFileSystem(ctx, targetId, optional)
Get target file system for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetId** | **string**| Unique identifier of the target file system Object | 
 **optional** | ***TargetFileSystemControllerApiGetFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetFileSystemControllerApiGetFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verbose** | **optional.Bool**| Optional parameter to return all filesystem properties | 

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargetFileSystems**
> []FileSystemConfiguration ListTargetFileSystems(ctx, optional)
List all target file systems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TargetFileSystemControllerApiListTargetFileSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetFileSystemControllerApiListTargetFileSystemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verbose** | **optional.Bool**| Optional parameter to return all filesystem properties | 

### Return type

[**[]FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

