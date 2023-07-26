# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceHdfsFileSystem**](HDFSSourceFileSystemControllerApi.md#AddSourceHdfsFileSystem) | **Post** /fs/sources/hdfs/{fileSystemId} | Add a new Hadoop source file system
[**UpdateSourceHdfsFileSystem**](HDFSSourceFileSystemControllerApi.md#UpdateSourceHdfsFileSystem) | **Patch** /fs/sources/hdfs/{fileSystemId} | Update a Hadoop source file system

# **AddSourceHdfsFileSystem**
> FileSystemConfiguration AddSourceHdfsFileSystem(ctx, body, fileSystemId, optional)
Add a new Hadoop source file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HdfsFilesystemProperties**](HdfsFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***HDFSSourceFileSystemControllerApiAddSourceHdfsFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HDFSSourceFileSystemControllerApiAddSourceHdfsFileSystemOpts struct
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

# **UpdateSourceHdfsFileSystem**
> FileSystemConfiguration UpdateSourceHdfsFileSystem(ctx, body, fileSystemId, optional)
Update a Hadoop source file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HdfsFilesystemProperties**](HdfsFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***HDFSSourceFileSystemControllerApiUpdateSourceHdfsFileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HDFSSourceFileSystemControllerApiUpdateSourceHdfsFileSystemOpts struct
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

