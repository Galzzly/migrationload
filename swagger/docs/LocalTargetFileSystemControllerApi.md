# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetLocalFileSystem**](LocalTargetFileSystemControllerApi.md#AddTargetLocalFileSystem) | **Post** /fs/targets/local/{fileSystemId} | Add a new target Local file system
[**UpdateTargetLocalFileSystem**](LocalTargetFileSystemControllerApi.md#UpdateTargetLocalFileSystem) | **Patch** /fs/targets/local/{fileSystemId} | Update a target Local file system

# **AddTargetLocalFileSystem**
> FileSystemConfiguration AddTargetLocalFileSystem(ctx, body, fileSystemId, optional)
Add a new target Local file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LocalFilesystemProperties**](LocalFilesystemProperties.md)| Local file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***LocalTargetFileSystemControllerApiAddTargetLocalFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalTargetFileSystemControllerApiAddTargetLocalFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTargetLocalFileSystem**
> FileSystemConfiguration UpdateTargetLocalFileSystem(ctx, body, fileSystemId, optional)
Update a target Local file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LocalFilesystemProperties**](LocalFilesystemProperties.md)| Local file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***LocalTargetFileSystemControllerApiUpdateTargetLocalFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalTargetFileSystemControllerApiUpdateTargetLocalFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

