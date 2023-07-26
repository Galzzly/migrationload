# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTargetGcsFileSystemDefault**](GCSTargetFileSystemControllerApi.md#AddTargetGcsFileSystemDefault) | **Post** /fs/targets/gcs/{fileSystemId} | Add a new target GCS file system
[**AddTargetGcsFileSystemJsonKeyFile**](GCSTargetFileSystemControllerApi.md#AddTargetGcsFileSystemJsonKeyFile) | **Post** /fs/targets/gcs/keyfiles/json/{fileSystemId} | Add a new target GCS file system with credentials uploaded with JSON key file
[**AddTargetGcsFileSystemP12KeyFile**](GCSTargetFileSystemControllerApi.md#AddTargetGcsFileSystemP12KeyFile) | **Post** /fs/targets/gcs/keyfiles/p12/{fileSystemId} | Add a new target GCS file system with credentials uploaded via P12 key file
[**UpdateTargetGcsFileSystemDefault**](GCSTargetFileSystemControllerApi.md#UpdateTargetGcsFileSystemDefault) | **Patch** /fs/targets/gcs/{fileSystemId} | Update a target GCS file system
[**UpdateTargetGcsFileSystemJsonKeyFile**](GCSTargetFileSystemControllerApi.md#UpdateTargetGcsFileSystemJsonKeyFile) | **Patch** /fs/targets/gcs/keyfiles/json/{fileSystemId} | Update a target GCS file system with credentials uploaded with JSON key file
[**UpdateTargetGcsFileSystemP12KeyFile**](GCSTargetFileSystemControllerApi.md#UpdateTargetGcsFileSystemP12KeyFile) | **Patch** /fs/targets/gcs/keyfiles/p12/{fileSystemId} | Update a target GCS file system with credentials uploaded via P12 key file

# **AddTargetGcsFileSystemDefault**
> FileSystemConfiguration AddTargetGcsFileSystemDefault(ctx, body, fileSystemId, optional)
Add a new target GCS file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GcsFilesystemProperties**](GcsFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiAddTargetGcsFileSystemDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiAddTargetGcsFileSystemDefaultOpts struct
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

# **AddTargetGcsFileSystemJsonKeyFile**
> FileSystemConfiguration AddTargetGcsFileSystemJsonKeyFile(ctx, fileSystemId, optional)
Add a new target GCS file system with credentials uploaded with JSON key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiAddTargetGcsFileSystemJsonKeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiAddTargetGcsFileSystemJsonKeyFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fsProperties** | [**optional.Interface of GcsFilesystemProperties**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTargetGcsFileSystemP12KeyFile**
> FileSystemConfiguration AddTargetGcsFileSystemP12KeyFile(ctx, fileSystemId, optional)
Add a new target GCS file system with credentials uploaded via P12 key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiAddTargetGcsFileSystemP12KeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiAddTargetGcsFileSystemP12KeyFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fsProperties** | [**optional.Interface of GcsFilesystemProperties**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTargetGcsFileSystemDefault**
> FileSystemConfiguration UpdateTargetGcsFileSystemDefault(ctx, body, fileSystemId, optional)
Update a target GCS file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GcsFilesystemProperties**](GcsFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemDefaultOpts struct
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

# **UpdateTargetGcsFileSystemJsonKeyFile**
> FileSystemConfiguration UpdateTargetGcsFileSystemJsonKeyFile(ctx, fileSystemId, optional)
Update a target GCS file system with credentials uploaded with JSON key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemJsonKeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemJsonKeyFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fsProperties** | [**optional.Interface of GcsFilesystemProperties**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTargetGcsFileSystemP12KeyFile**
> FileSystemConfiguration UpdateTargetGcsFileSystemP12KeyFile(ctx, fileSystemId, optional)
Update a target GCS file system with credentials uploaded via P12 key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemP12KeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSTargetFileSystemControllerApiUpdateTargetGcsFileSystemP12KeyFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fsProperties** | [**optional.Interface of GcsFilesystemProperties**](.md)|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **propertiesFiles** | [**optional.Interface of []string**](string.md)| Load configuration from files that are accessible from server | [default to []]

### Return type

[**FileSystemConfiguration**](FileSystemConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

