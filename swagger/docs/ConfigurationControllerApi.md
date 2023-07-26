# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigurationProperties**](ConfigurationControllerApi.md#GetConfigurationProperties) | **Get** /configuration/properties | Get all configuration properties
[**GetConfigurationProperty**](ConfigurationControllerApi.md#GetConfigurationProperty) | **Get** /configuration/property | Get value for given configuration property
[**SetConfigurationProperty**](ConfigurationControllerApi.md#SetConfigurationProperty) | **Put** /configuration/property | Set key/value configuration property

# **GetConfigurationProperties**
> []Configuration GetConfigurationProperties(ctx, )
Get all configuration properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Configuration**](Configuration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigurationProperty**
> Configuration GetConfigurationProperty(ctx, key)
Get value for given configuration property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Key - property name to be get | 

### Return type

[**Configuration**](Configuration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetConfigurationProperty**
> SetConfigurationProperty(ctx, key, value)
Set key/value configuration property

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Key - property name to be set | 
  **value** | **string**| Value - value to set corresponding property to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

