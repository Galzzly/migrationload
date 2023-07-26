# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdls2SourceFileSystemServicePrincipal**](ADLS2SourceFileSystemControllerApi.md#AddAdls2SourceFileSystemServicePrincipal) | **Post** /fs/sources/adls2/oauth2/{fileSystemId} | Add a new source ADLS2 file system with service principal and OAuth 2.0 authentication
[**AddSourceAdls2FileSystem**](ADLS2SourceFileSystemControllerApi.md#AddSourceAdls2FileSystem) | **Post** /fs/sources/adls2/shared-key/{fileSystemId} | Add a Adls2 file system With shared key authentication
[**UpdateAdls2SourceFileSystemServicePrincipal**](ADLS2SourceFileSystemControllerApi.md#UpdateAdls2SourceFileSystemServicePrincipal) | **Patch** /fs/sources/adls2/oauth2/{fileSystemId} | Update a source ADLS2 file system with service principal and OAuth 2.0 authentication
[**UpdateSourceAdls2FileSystem**](ADLS2SourceFileSystemControllerApi.md#UpdateSourceAdls2FileSystem) | **Patch** /fs/sources/adls2/shared-key/{fileSystemId} | Update a Adls2 file system With shared key authentication

# **AddAdls2SourceFileSystemServicePrincipal**
> FileSystemConfiguration AddAdls2SourceFileSystemServicePrincipal(ctx, body, fileSystemId, optional)
Add a new source ADLS2 file system with service principal and OAuth 2.0 authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2ServicePrincipalFilesystemProperties**](Adls2ServicePrincipalFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***ADLS2SourceFileSystemControllerApiAddAdls2SourceFileSystemServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2SourceFileSystemControllerApiAddAdls2SourceFileSystemServicePrincipalOpts struct
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

# **AddSourceAdls2FileSystem**
> FileSystemConfiguration AddSourceAdls2FileSystem(ctx, body, fileSystemId, optional)
Add a Adls2 file system With shared key authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2SharedKeyFilesystemProperties**](Adls2SharedKeyFilesystemProperties.md)|  | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***ADLS2SourceFileSystemControllerApiAddSourceAdls2FileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2SourceFileSystemControllerApiAddSourceAdls2FileSystemOpts struct
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

# **UpdateAdls2SourceFileSystemServicePrincipal**
> FileSystemConfiguration UpdateAdls2SourceFileSystemServicePrincipal(ctx, body, fileSystemId, optional)
Update a source ADLS2 file system with service principal and OAuth 2.0 authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2ServicePrincipalFilesystemProperties**](Adls2ServicePrincipalFilesystemProperties.md)| Source file system properties | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***ADLS2SourceFileSystemControllerApiUpdateAdls2SourceFileSystemServicePrincipalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2SourceFileSystemControllerApiUpdateAdls2SourceFileSystemServicePrincipalOpts struct
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

# **UpdateSourceAdls2FileSystem**
> FileSystemConfiguration UpdateSourceAdls2FileSystem(ctx, body, fileSystemId, optional)
Update a Adls2 file system With shared key authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Adls2SharedKeyFilesystemProperties**](Adls2SharedKeyFilesystemProperties.md)|  | 
  **fileSystemId** | **string**| Unique Identifier for the source file system | 
 **optional** | ***ADLS2SourceFileSystemControllerApiUpdateSourceAdls2FileSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ADLS2SourceFileSystemControllerApiUpdateSourceAdls2FileSystemOpts struct
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

