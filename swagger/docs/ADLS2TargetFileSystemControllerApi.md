# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdls2FileSystem**](ADLS2TargetFileSystemControllerApi.md#AddAdls2FileSystem) | **Post** /fs/targets/adls2/shared-key/{fileSystemId} | Add a new target ADLS2 file system with Shared Key authentication
[**AddAdls2FileSystemServicePrincipal**](ADLS2TargetFileSystemControllerApi.md#AddAdls2FileSystemServicePrincipal) | **Post** /fs/targets/adls2/oauth2/{fileSystemId} | Add a new target ADLS2 file system with service principal and OAuth 2.0 authentication
[**UpdateAdls2FileSystem**](ADLS2TargetFileSystemControllerApi.md#UpdateAdls2FileSystem) | **Patch** /fs/targets/adls2/shared-key/{fileSystemId} | Update a target ADLS2 file system with Shared Key authentication
[**UpdateAdls2FileSystemServicePrincipal**](ADLS2TargetFileSystemControllerApi.md#UpdateAdls2FileSystemServicePrincipal) | **Patch** /fs/targets/adls2/oauth2/{fileSystemId} | Update a target ADLS2 file system with service principal and OAuth 2.0 authentication

# **AddAdls2FileSystem**
> FileSystemConfiguration AddAdls2FileSystem(ctx, body, fileSystemId, optional)
Add a new target ADLS2 file system with Shared Key authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2SharedKeyFilesystemProperties**](Adls2SharedKeyFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***ADLS2TargetFileSystemControllerApiAddAdls2FileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2TargetFileSystemControllerApiAddAdls2FileSystemOpts struct
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

# **AddAdls2FileSystemServicePrincipal**
> FileSystemConfiguration AddAdls2FileSystemServicePrincipal(ctx, body, fileSystemId, optional)
Add a new target ADLS2 file system with service principal and OAuth 2.0 authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2ServicePrincipalFilesystemProperties**](Adls2ServicePrincipalFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***ADLS2TargetFileSystemControllerApiAddAdls2FileSystemServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2TargetFileSystemControllerApiAddAdls2FileSystemServicePrincipalOpts struct
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

# **UpdateAdls2FileSystem**
> FileSystemConfiguration UpdateAdls2FileSystem(ctx, body, fileSystemId, optional)
Update a target ADLS2 file system with Shared Key authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2SharedKeyFilesystemProperties**](Adls2SharedKeyFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***ADLS2TargetFileSystemControllerApiUpdateAdls2FileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2TargetFileSystemControllerApiUpdateAdls2FileSystemOpts struct
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

# **UpdateAdls2FileSystemServicePrincipal**
> FileSystemConfiguration UpdateAdls2FileSystemServicePrincipal(ctx, body, fileSystemId, optional)
Update a target ADLS2 file system with service principal and OAuth 2.0 authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2ServicePrincipalFilesystemProperties**](Adls2ServicePrincipalFilesystemProperties.md)| Target file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the target file system | 
 **optional** | ***ADLS2TargetFileSystemControllerApiUpdateAdls2FileSystemServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2TargetFileSystemControllerApiUpdateAdls2FileSystemServicePrincipalOpts struct
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

