# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetS3aFileSystem**](S3aTargetFileSystemControllerApi.md#AddTargetS3aFileSystem) | **Post** /fs/targets/s3a/{fileSystemId} | Add a new target S3a file system
[**UpdateTargetS3aFileSystem**](S3aTargetFileSystemControllerApi.md#UpdateTargetS3aFileSystem) | **Patch** /fs/targets/s3a/{fileSystemId} | Update a target S3a file system

# **AddTargetS3aFileSystem**
> FileSystemConfiguration AddTargetS3aFileSystem(ctx, body, fileSystemId, optional)
Add a new target S3a file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3aFilesystemProperties**](S3aFilesystemProperties.md)| S3A file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***S3aTargetFileSystemControllerApiAddTargetS3aFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3aTargetFileSystemControllerApiAddTargetS3aFileSystemOpts struct
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

# **UpdateTargetS3aFileSystem**
> FileSystemConfiguration UpdateTargetS3aFileSystem(ctx, body, fileSystemId, optional)
Update a target S3a file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3aFilesystemProperties**](S3aFilesystemProperties.md)| S3A file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***S3aTargetFileSystemControllerApiUpdateTargetS3aFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3aTargetFileSystemControllerApiUpdateTargetS3aFileSystemOpts struct
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

