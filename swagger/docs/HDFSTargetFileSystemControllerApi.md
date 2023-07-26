# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetHdfsFileSystem**](HDFSTargetFileSystemControllerApi.md#AddTargetHdfsFileSystem) | **Post** /fs/targets/hdfs/{fileSystemId} | Add a new target Hadoop target file system
[**UpdateTargetHdfsFileSystem**](HDFSTargetFileSystemControllerApi.md#UpdateTargetHdfsFileSystem) | **Patch** /fs/targets/hdfs/{fileSystemId} | Update a target Hadoop target file system

# **AddTargetHdfsFileSystem**
> FileSystemConfiguration AddTargetHdfsFileSystem(ctx, body, fileSystemId, optional)
Add a new target Hadoop target file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HdfsFilesystemProperties**](HdfsFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***HDFSTargetFileSystemControllerApiAddTargetHdfsFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HDFSTargetFileSystemControllerApiAddTargetHdfsFileSystemOpts struct
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

# **UpdateTargetHdfsFileSystem**
> FileSystemConfiguration UpdateTargetHdfsFileSystem(ctx, body, fileSystemId, optional)
Update a target Hadoop target file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HdfsFilesystemProperties**](HdfsFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***HDFSTargetFileSystemControllerApiUpdateTargetHdfsFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HDFSTargetFileSystemControllerApiUpdateTargetHdfsFileSystemOpts struct
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

