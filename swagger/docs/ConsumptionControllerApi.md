# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtendLicense**](ConsumptionControllerApi.md#ExtendLicense) | **Put** /consumption/extendTime | Extend consumption license expiry time 
[**SwitchToConsumption**](ConsumptionControllerApi.md#SwitchToConsumption) | **Put** /consumption/switchToConsumption | Change license type to consumption 

# **ExtendLicense**
> bool ExtendLicense(ctx, token)
Extend consumption license expiry time 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Token used to extend license | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchToConsumption**
> SwitchToConsumption(ctx, )
Change license type to consumption 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

