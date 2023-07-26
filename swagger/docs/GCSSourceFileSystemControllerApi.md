# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceGcsFileSystemDefault**](GCSSourceFileSystemControllerApi.md#AddSourceGcsFileSystemDefault) | **Post** /fs/sources/gcs/{fileSystemId} | Add a new source GCS file system
[**AddSourceGcsFileSystemJsonKeyFile**](GCSSourceFileSystemControllerApi.md#AddSourceGcsFileSystemJsonKeyFile) | **Post** /fs/sources/gcs/keyfiles/json/{fileSystemId} | Add a new source GCS file system with credentials uploaded with JSON key file
[**AddSourceGcsFileSystemP12KeyFile**](GCSSourceFileSystemControllerApi.md#AddSourceGcsFileSystemP12KeyFile) | **Post** /fs/sources/gcs/keyfiles/p12/{fileSystemId} | Add a new source GCS file system with credentials uploaded via P12 key file
[**UpdateSourceGcsFileSystemDefault**](GCSSourceFileSystemControllerApi.md#UpdateSourceGcsFileSystemDefault) | **Patch** /fs/sources/gcs/{fileSystemId} | Update a source GCS file system
[**UpdateSourceGcsFileSystemJsonKeyFile**](GCSSourceFileSystemControllerApi.md#UpdateSourceGcsFileSystemJsonKeyFile) | **Patch** /fs/sources/gcs/keyfiles/json/{fileSystemId} | Update a source GCS file system with credentials uploaded with JSON key file
[**UpdateSourceGcsFileSystemP12KeyFile**](GCSSourceFileSystemControllerApi.md#UpdateSourceGcsFileSystemP12KeyFile) | **Patch** /fs/sources/gcs/keyfiles/p12/{fileSystemId} | Update a source GCS file system with credentials uploaded via P12 key file

# **AddSourceGcsFileSystemDefault**
> FileSystemConfiguration AddSourceGcsFileSystemDefault(ctx, body, fileSystemId, optional)
Add a new source GCS file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GcsFilesystemProperties**](GcsFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiAddSourceGcsFileSystemDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiAddSourceGcsFileSystemDefaultOpts struct
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

# **AddSourceGcsFileSystemJsonKeyFile**
> FileSystemConfiguration AddSourceGcsFileSystemJsonKeyFile(ctx, fileSystemId, optional)
Add a new source GCS file system with credentials uploaded with JSON key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiAddSourceGcsFileSystemJsonKeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiAddSourceGcsFileSystemJsonKeyFileOpts struct
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

# **AddSourceGcsFileSystemP12KeyFile**
> FileSystemConfiguration AddSourceGcsFileSystemP12KeyFile(ctx, fileSystemId, optional)
Add a new source GCS file system with credentials uploaded via P12 key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiAddSourceGcsFileSystemP12KeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiAddSourceGcsFileSystemP12KeyFileOpts struct
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

# **UpdateSourceGcsFileSystemDefault**
> FileSystemConfiguration UpdateSourceGcsFileSystemDefault(ctx, body, fileSystemId, optional)
Update a source GCS file system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GcsFilesystemProperties**](GcsFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemDefaultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemDefaultOpts struct
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

# **UpdateSourceGcsFileSystemJsonKeyFile**
> FileSystemConfiguration UpdateSourceGcsFileSystemJsonKeyFile(ctx, fileSystemId, optional)
Update a source GCS file system with credentials uploaded with JSON key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemJsonKeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemJsonKeyFileOpts struct
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

# **UpdateSourceGcsFileSystemP12KeyFile**
> FileSystemConfiguration UpdateSourceGcsFileSystemP12KeyFile(ctx, fileSystemId, optional)
Update a source GCS file system with credentials uploaded via P12 key file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemP12KeyFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GCSSourceFileSystemControllerApiUpdateSourceGcsFileSystemP12KeyFileOpts struct
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

