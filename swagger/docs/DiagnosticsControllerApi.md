# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DumpActionStoreForMigration**](DiagnosticsControllerApi.md#DumpActionStoreForMigration) | **Get** /diagnostics/{migrationId}/actionStore | To be used for debugging purposes only - dumps the ActionStore for a Migration to file.
[**FetchAllDiagnostics**](DiagnosticsControllerApi.md#FetchAllDiagnostics) | **Get** /diagnostics | Get all Diagnostics sets over time. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.
[**FetchCountTcpStates**](DiagnosticsControllerApi.md#FetchCountTcpStates) | **Get** /diagnostics/netstat/tcpStatesCount | Get a map of total connections to their TCP state.
[**FetchCurrentDiagnostics**](DiagnosticsControllerApi.md#FetchCurrentDiagnostics) | **Get** /diagnostics/collect | Trigger Diagnostics Collection. Will trigger the collection of Diagnostics and return the new set of Diagnostics generated. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.
[**FetchDiagnosticsDescriptions**](DiagnosticsControllerApi.md#FetchDiagnosticsDescriptions) | **Get** /diagnostics/descriptions | Get Diagnostics Descriptions. Return the set of available Diagnostics and a description of the information they return.
[**FetchDiagnosticsSummary**](DiagnosticsControllerApi.md#FetchDiagnosticsSummary) | **Get** /diagnostics/summary | Get Diagnostic Summary
[**FetchDiagnosticsSummaryTxt**](DiagnosticsControllerApi.md#FetchDiagnosticsSummaryTxt) | **Get** /diagnostics/summary.txt | Get Diagnostic Summary Plain Text
[**FetchLatestDiagnostics**](DiagnosticsControllerApi.md#FetchLatestDiagnostics) | **Get** /diagnostics/lastCollected | Get last collected set of Diagnostics. Will return the last collected set of Diagnostics. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.
[**FetchNetworkConnections**](DiagnosticsControllerApi.md#FetchNetworkConnections) | **Get** /diagnostics/netstat/connections | Get Established Connections.
[**FetchNetworkStatus**](DiagnosticsControllerApi.md#FetchNetworkStatus) | **Get** /diagnostics/netstat | Get Netstat Status.
[**FetchNetworkStatusTotals**](DiagnosticsControllerApi.md#FetchNetworkStatusTotals) | **Get** /diagnostics/netstat/totals | Get Connections Total Queue Sizes and ReTransmit Counts.

# **DumpActionStoreForMigration**
> DumpActionStoreForMigration(ctx, migrationId)
To be used for debugging purposes only - dumps the ActionStore for a Migration to file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchAllDiagnostics**
> []LiveMigratorDiagnosticsSnapshot FetchAllDiagnostics(ctx, optional)
Get all Diagnostics sets over time. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiagnosticsControllerApiFetchAllDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiagnosticsControllerApiFetchAllDiagnosticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **optional.String**| Filter based on Diagnostic kind. | 
 **count** | **optional.Int32**| Limit number of Diagnostic sets returned. Higher counts may cause large responses. It is advised to use curl or another tool in this case. | 

### Return type

[**[]LiveMigratorDiagnosticsSnapshot**](LiveMigratorDiagnosticsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCountTcpStates**
> map[string]int32 FetchCountTcpStates(ctx, )
Get a map of total connections to their TCP state.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchCurrentDiagnostics**
> LiveMigratorDiagnosticsSnapshot FetchCurrentDiagnostics(ctx, optional)
Trigger Diagnostics Collection. Will trigger the collection of Diagnostics and return the new set of Diagnostics generated. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiagnosticsControllerApiFetchCurrentDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiagnosticsControllerApiFetchCurrentDiagnosticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **optional.String**| Filter based on Diagnostic kind | 

### Return type

[**LiveMigratorDiagnosticsSnapshot**](LiveMigratorDiagnosticsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsDescriptions**
> []DiagnosticDescription FetchDiagnosticsDescriptions(ctx, )
Get Diagnostics Descriptions. Return the set of available Diagnostics and a description of the information they return.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DiagnosticDescription**](DiagnosticDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsSummary**
> DiagnosticSummary FetchDiagnosticsSummary(ctx, )
Get Diagnostic Summary

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiagnosticSummary**](DiagnosticSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchDiagnosticsSummaryTxt**
> string FetchDiagnosticsSummaryTxt(ctx, )
Get Diagnostic Summary Plain Text

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchLatestDiagnostics**
> LiveMigratorDiagnosticsSnapshot FetchLatestDiagnostics(ctx, optional)
Get last collected set of Diagnostics. Will return the last collected set of Diagnostics. kind can be used to filter the diagnostics returned, it is a comma separated list of Diagnostics kinds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DiagnosticsControllerApiFetchLatestDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiagnosticsControllerApiFetchLatestDiagnosticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **optional.String**| Filter based on Diagnostic kind | 

### Return type

[**LiveMigratorDiagnosticsSnapshot**](LiveMigratorDiagnosticsSnapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNetworkConnections**
> []NetworkConnection FetchNetworkConnections(ctx, )
Get Established Connections.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]NetworkConnection**](NetworkConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNetworkStatus**
> string FetchNetworkStatus(ctx, )
Get Netstat Status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNetworkStatusTotals**
> map[string]RxTxTotals FetchNetworkStatusTotals(ctx, )
Get Connections Total Queue Sizes and ReTransmit Counts.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]RxTxTotals**](RxTxTotals.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

