# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoDiscoverSourceFilesystem**](HDFSSourceFileSystemAutodiscoveryControllerApi.md#AutoDiscoverSourceFilesystem) | **Post** /fs/autodiscover/sources/hdfs | Automatically discover the local hadoop source file system

# **AutoDiscoverSourceFilesystem**
> FileSystemConfiguration AutoDiscoverSourceFilesystem(ctx, body, optional)
Automatically discover the local hadoop source file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HdfsFilesystemProperties**](HdfsFilesystemProperties.md)| Source file system properties | 
 **optional** | ***HDFSSourceFileSystemAutodiscoveryControllerApiAutoDiscoverSourceFilesystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HDFSSourceFileSystemAutodiscoveryControllerApiAutoDiscoverSourceFilesystemOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scanOnly** | **optional.**| Scan-only will not pick up any changes to the underlying filesystem during migrations | 

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

