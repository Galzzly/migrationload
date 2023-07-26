# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDtaStats**](DataAgentStatisticsControllerApi.md#GetDtaStats) | **Get** /stats/dta/{agentName} | Get full stats over for given Data Agent
[**GetDtaStatsList**](DataAgentStatisticsControllerApi.md#GetDtaStatsList) | **Get** /stats/dta/ | Get full stats over all time for all Data Agents
[**GetDtaStatsSummary**](DataAgentStatisticsControllerApi.md#GetDtaStatsSummary) | **Get** /stats/dta/{agentName}/summary | Get summary stats for given Data Agent

# **GetDtaStats**
> DataAgentStatsHistory GetDtaStats(ctx, agentName)
Get full stats over for given Data Agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| User-defined name of the Data Agent | 

### Return type

[**DataAgentStatsHistory**](DataAgentStatsHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDtaStatsList**
> OverallDataAgentStatsHistory GetDtaStatsList(ctx, )
Get full stats over all time for all Data Agents

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OverallDataAgentStatsHistory**](OverallDataAgentStatsHistory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDtaStatsSummary**
> DataAgentStatsSummary GetDtaStatsSummary(ctx, agentName)
Get summary stats for given Data Agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| User-defined name of the Data Agent | 

### Return type

[**DataAgentStatsSummary**](DataAgentStatsSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

