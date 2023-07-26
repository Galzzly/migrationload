# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllDiagnosticsByAgentName**](DataAgentDiagnosticsControllerApi.md#FetchAllDiagnosticsByAgentName) | **Get** /diagnostics/dta/{agentName} | Get all Diagnostics snapshots for all data agents over time. \&quot;kind\&quot; can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.
[**FetchLatestDiagnosticsByAgentName**](DataAgentDiagnosticsControllerApi.md#FetchLatestDiagnosticsByAgentName) | **Get** /diagnostics/dta/lastCollected/{agentName} | Get last collected snapshot of Diagnostics for specified agent by agent name. Will return the last collected snapshot of Diagnostics. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

# **FetchAllDiagnosticsByAgentName**
> []DataAgentDiagnosticsSnapshot FetchAllDiagnosticsByAgentName(ctx, agentName, optional)
Get all Diagnostics snapshots for all data agents over time. \"kind\" can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| Data agent name | 
 **optional** | ***DataAgentDiagnosticsControllerApiFetchAllDiagnosticsByAgentNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataAgentDiagnosticsControllerApiFetchAllDiagnosticsByAgentNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | **optional.String**| Filter based on Diagnostic kind. | 
 **count** | **optional.Int32**| Limit number of Diagnostic snapshots returned. Higher counts may cause large responses. It is advised to use curl or another tool in this case. | 

### Return type

[**[]DataAgentDiagnosticsSnapshot**](DataAgentDiagnosticsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLatestDiagnosticsByAgentName**
> DataAgentDiagnosticsSnapshot FetchLatestDiagnosticsByAgentName(ctx, agentName, optional)
Get last collected snapshot of Diagnostics for specified agent by agent name. Will return the last collected snapshot of Diagnostics. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **agentName** | **string**| Data agent name | 
 **optional** | ***DataAgentDiagnosticsControllerApiFetchLatestDiagnosticsByAgentNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataAgentDiagnosticsControllerApiFetchLatestDiagnosticsByAgentNameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kind** | **optional.String**| Filter based on Diagnostic kind | 

### Return type

[**DataAgentDiagnosticsSnapshot**](DataAgentDiagnosticsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

