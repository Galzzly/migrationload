# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMapping**](PathMappingControllerApi.md#CreateMapping) | **Put** /transformations/mappings | Create a Mapping
[**Delete**](PathMappingControllerApi.md#Delete) | **Delete** /transformations/mappings/{id} | Delete a Mapping
[**FetchAll**](PathMappingControllerApi.md#FetchAll) | **Get** /transformations/mappings | Get all Mappings
[**FetchMapping**](PathMappingControllerApi.md#FetchMapping) | **Get** /transformations/mappings/{id} | Fetch a Mapping

# **CreateMapping**
> PathMapping CreateMapping(ctx, sourceFileSystem, sourcePath, targetFileSystem, targetPath, optional)
Create a Mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceFileSystem** | **string**| Name of source file system | 
  **sourcePath** | **string**| Path of source file system | 
  **targetFileSystem** | **string**| Name of target file system | 
  **targetPath** | **string**| Path of target file system | 
 **optional** | ***PathMappingControllerApiCreateMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PathMappingControllerApiCreateMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pathMappingId** | **optional.String**| Mapping id | 
 **description** | **optional.String**| Mapping description | 

### Return type

[**PathMapping**](PathMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, id)
Delete a Mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier for mapping | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchAll**
> []PathMapping FetchAll(ctx, optional)
Get all Mappings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PathMappingControllerApiFetchAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PathMappingControllerApiFetchAllOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceFileSystemName** | **optional.String**| Name of the source file system. | 
 **targetFileSystemName** | **optional.String**| Name of the target file system. | 

### Return type

[**[]PathMapping**](PathMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchMapping**
> PathMapping FetchMapping(ctx, id)
Fetch a Mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier for mapping | 

### Return type

[**PathMapping**](PathMapping.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

