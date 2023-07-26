# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](FileSystemControllerApi.md#CreateFile) | **Post** /fs/{role}/{fsId}/createFile | Create file on the file system
[**DeleteByPath**](FileSystemControllerApi.md#DeleteByPath) | **Post** /fs/{role}/{fsId}/deleteByPath | Delete file from the file system
[**GetFileSystemURI**](FileSystemControllerApi.md#GetFileSystemURI) | **Get** /fs/{role}/{fsId}/fileSystemURI | Get file system URI for given filesystem id
[**ListPaths**](FileSystemControllerApi.md#ListPaths) | **Get** /fs/{role}/{fsId}/listPaths | Get files list from the file system
[**Mkdir**](FileSystemControllerApi.md#Mkdir) | **Post** /fs/{role}/{fsId}/mkdir | Create directory in the file system
[**MoveByPath**](FileSystemControllerApi.md#MoveByPath) | **Post** /fs/{role}/{fsId}/moveByPath | Move file or directory on the file system
[**Path**](FileSystemControllerApi.md#Path) | **Get** /fs/{role}/{fsId}/pathExists | Check if file or directory exists
[**ReadFile**](FileSystemControllerApi.md#ReadFile) | **Get** /fs/{role}/{fsId}/readFile | Read file from the file system
[**WriteFile**](FileSystemControllerApi.md#WriteFile) | **Post** /fs/{role}/{fsId}/writeFile | Write file into the file system

# **CreateFile**
> bool CreateFile(ctx, fsId, role, optional)
Create file on the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
 **optional** | ***FileSystemControllerApiCreateFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileSystemControllerApiCreateFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **optional.**|  | 
 **fileBody** | **optional.Interface of *os.File****optional.**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteByPath**
> bool DeleteByPath(ctx, body, fsId, role, path)
Delete file from the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FileDeleteRequest**](FileDeleteRequest.md)| File delete parameters | 
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file or directory | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileSystemURI**
> FileSystemUri GetFileSystemURI(ctx, fsId, role)
Get file system URI for given filesystem id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 

### Return type

[**FileSystemUri**](FileSystemURI.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaths**
> []FileStatus ListPaths(ctx, fsId, role, path)
Get files list from the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file or directory | 

### Return type

[**[]FileStatus**](FileStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Mkdir**
> bool Mkdir(ctx, fsId, role, path)
Create directory in the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the directory | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveByPath**
> bool MoveByPath(ctx, fsId, role, path, destPath)
Move file or directory on the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file or directory | 
  **destPath** | **string**| Destination path to the file or directory | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Path**
> bool Path(ctx, fsId, role, path)
Check if file or directory exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file or directory | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFile**
> *os.File ReadFile(ctx, fsId, role, path)
Read file from the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WriteFile**
> bool WriteFile(ctx, body, fsId, role, path)
Write file into the file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FileWriteRequest**](FileWriteRequest.md)| File write parameters | 
  **fsId** | **string**|  | 
  **role** | **string**|  | 
  **path** | **string**| Path to the file | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

