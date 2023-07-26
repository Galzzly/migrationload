# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataAgent**](DataAgentsControllerApi.md#AddDataAgent) | **Post** /scaling/dataagents/ | Add new Data Transfer Agent
[**DeleteDataAgent**](DataAgentsControllerApi.md#DeleteDataAgent) | **Delete** /scaling/dataagents/{agentName} | Delete Data Transfer Agent from LM2
[**GetDataAgent**](DataAgentsControllerApi.md#GetDataAgent) | **Get** /scaling/dataagents/{agentName} | Get Data Agent with given name
[**GetDataAgentList**](DataAgentsControllerApi.md#GetDataAgentList) | **Get** /scaling/dataagents/ | Get Data Agent list

# **AddDataAgent**
> DataAgent AddDataAgent(ctx, optional)
Add new Data Transfer Agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataAgentsControllerApiAddDataAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataAgentsControllerApiAddDataAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DataAgentAddRequest**](DataAgentAddRequest.md)|  | 

### Return type

[**DataAgent**](DataAgent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDataAgent**
> DeleteDataAgent(ctx, agentName)
Delete Data Transfer Agent from LM2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| Data Agent user-defined name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataAgent**
> DataAgent GetDataAgent(ctx, agentName)
Get Data Agent with given name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| Data Agent user-defined name. | 

### Return type

[**DataAgent**](DataAgent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDataAgentList**
> []DataAgent GetDataAgentList(ctx, )
Get Data Agent list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DataAgent**](DataAgent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

