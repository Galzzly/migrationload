# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceLocalFileSystem**](LocalSourceFileSystemControllerApi.md#AddSourceLocalFileSystem) | **Post** /fs/sources/local/{fileSystemId} | Add a new Local source file system
[**UpdateSourceLocalFileSystem**](LocalSourceFileSystemControllerApi.md#UpdateSourceLocalFileSystem) | **Patch** /fs/sources/local/{fileSystemId} | Update a Local source file system

# **AddSourceLocalFileSystem**
> FileSystemConfiguration AddSourceLocalFileSystem(ctx, body, fileSystemId, optional)
Add a new Local source file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LocalFilesystemProperties**](LocalFilesystemProperties.md)| Local file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***LocalSourceFileSystemControllerApiAddSourceLocalFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalSourceFileSystemControllerApiAddSourceLocalFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scanOnly** | **optional.**| Scan-only is the only configuration supported for this type of source. If a value of &#x27;false&#x27; is supplied for this parameter, it will be ignored. | [default to true]
 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSourceLocalFileSystem**
> FileSystemConfiguration UpdateSourceLocalFileSystem(ctx, body, fileSystemId, optional)
Update a Local source file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LocalFilesystemProperties**](LocalFilesystemProperties.md)| Local file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***LocalSourceFileSystemControllerApiUpdateSourceLocalFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalSourceFileSystemControllerApiUpdateSourceLocalFileSystemOpts struct
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

