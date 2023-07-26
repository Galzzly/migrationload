# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceS3aFileSystem**](S3aSourceFileSystemControllerApi.md#AddSourceS3aFileSystem) | **Post** /fs/sources/s3a/{fileSystemId} | Add a new source S3a file system
[**UpdateSourceS3aFileSystem**](S3aSourceFileSystemControllerApi.md#UpdateSourceS3aFileSystem) | **Patch** /fs/sources/s3a/{fileSystemId} | Update a source S3a file system

# **AddSourceS3aFileSystem**
> FileSystemConfiguration AddSourceS3aFileSystem(ctx, body, fileSystemId, optional)
Add a new source S3a file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3aFilesystemProperties**](S3aFilesystemProperties.md)| S3A file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***S3aSourceFileSystemControllerApiAddSourceS3aFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3aSourceFileSystemControllerApiAddSourceS3aFileSystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]
 **scanOnly** | **optional.**| Scan-only will not pick up any changes to the underlying filesystem during migrations | 

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSourceS3aFileSystem**
> FileSystemConfiguration UpdateSourceS3aFileSystem(ctx, body, fileSystemId, optional)
Update a source S3a file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**S3aFilesystemProperties**](S3aFilesystemProperties.md)| S3A file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***S3aSourceFileSystemControllerApiUpdateSourceS3aFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3aSourceFileSystemControllerApiUpdateSourceS3aFileSystemOpts struct
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

