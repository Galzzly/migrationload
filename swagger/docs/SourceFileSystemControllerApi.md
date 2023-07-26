# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSourceFileSystems**](SourceFileSystemControllerApi.md#ClearSourceFileSystems) | **Delete** /fs/sources | Delete all source file systems
[**DeleteSourceFileSystem**](SourceFileSystemControllerApi.md#DeleteSourceFileSystem) | **Delete** /fs/sources/{sourceId} | Delete source file system for given ID
[**GetSourceAutoDiscoverMessage**](SourceFileSystemControllerApi.md#GetSourceAutoDiscoverMessage) | **Get** /fs/sources/auto-discover-message | Get auto-discover file system message
[**GetSourceAutoDiscoverStatus**](SourceFileSystemControllerApi.md#GetSourceAutoDiscoverStatus) | **Get** /fs/sources/auto-discover-status | Get auto-discover file system status object
[**GetSourceFileSystem**](SourceFileSystemControllerApi.md#GetSourceFileSystem) | **Get** /fs/sources/{sourceId} | Get source file system for given ID
[**ListSourceFileSystems**](SourceFileSystemControllerApi.md#ListSourceFileSystems) | **Get** /fs/sources | List all source filesystems

# **ClearSourceFileSystems**
> ClearSourceFileSystems(ctx, optional)
Delete all source file systems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SourceFileSystemControllerApiClearSourceFileSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceFileSystemControllerApiClearSourceFileSystemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteListener** | **optional.Bool**| Optional parameter to delete an event listener on source system while deleting the source from LDM | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSourceFileSystem**
> DeleteSourceFileSystem(ctx, sourceId, optional)
Delete source file system for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**| Unique identifier of the source filesystem | 
 **optional** | ***SourceFileSystemControllerApiDeleteSourceFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceFileSystemControllerApiDeleteSourceFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteListener** | **optional.Bool**| Optional parameter to delete an event listener on source system while deleting the source from LDM | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceAutoDiscoverMessage**
> string GetSourceAutoDiscoverMessage(ctx, )
Get auto-discover file system message

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceAutoDiscoverStatus**
> AutoDiscoverFsStatus GetSourceAutoDiscoverStatus(ctx, )
Get auto-discover file system status object

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AutoDiscoverFsStatus**](AutoDiscoverFSStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSourceFileSystem**
> FileSystemConfiguration GetSourceFileSystem(ctx, sourceId, optional)
Get source file system for given ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceId** | **string**| Unique identifier of the source filesystem | 
 **optional** | ***SourceFileSystemControllerApiGetSourceFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceFileSystemControllerApiGetSourceFileSystemOpts struct
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

# **ListSourceFileSystems**
> []FileSystemConfiguration ListSourceFileSystems(ctx, optional)
List all source filesystems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SourceFileSystemControllerApiListSourceFileSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SourceFileSystemControllerApiListSourceFileSystemsOpts struct
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

