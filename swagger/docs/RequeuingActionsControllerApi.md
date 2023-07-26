# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRequeuingAction**](RequeuingActionsControllerApi.md#CancelRequeuingAction) | **Post** /requeuing-actions/queuing/{migrationId}/{actionId}/cancel | Remove requeue action from migration
[**GetAllRequeuingActions**](RequeuingActionsControllerApi.md#GetAllRequeuingActions) | **Get** /requeuing-actions/queuing | Get all requeuing actions
[**GetAllRequeuingActionsStats**](RequeuingActionsControllerApi.md#GetAllRequeuingActionsStats) | **Get** /requeuing-actions/stats | Get all requeuing actions stats
[**GetMigrationRequeuingActions**](RequeuingActionsControllerApi.md#GetMigrationRequeuingActions) | **Get** /requeuing-actions/queuing/{migrationId} | Get all requeuing actions for specific migration
[**GetMigrationRequeuingActionsStats**](RequeuingActionsControllerApi.md#GetMigrationRequeuingActionsStats) | **Get** /requeuing-actions/stats/{migrationId} | Get all requeuing actions stats for specific migration
[**GetRequeuingAction**](RequeuingActionsControllerApi.md#GetRequeuingAction) | **Get** /requeuing-actions/queuing/{migrationId}/{actionId} | Get requeue action by migration ID and action ID

# **CancelRequeuingAction**
> CancelRequeuingAction(ctx, migrationId, actionId)
Remove requeue action from migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **actionId** | **int32**| Unique identifier of the action object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRequeuingActions**
> []RequeuingAction GetAllRequeuingActions(ctx, optional)
Get all requeuing actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequeuingActionsControllerApiGetAllRequeuingActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequeuingActionsControllerApiGetAllRequeuingActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Limit number of actions returned. | [default to 1000]

### Return type

[**[]RequeuingAction**](RequeuingAction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRequeuingActionsStats**
> RequeuingActionsStats GetAllRequeuingActionsStats(ctx, )
Get all requeuing actions stats

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RequeuingActionsStats**](RequeuingActionsStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationRequeuingActions**
> []RequeuingAction GetMigrationRequeuingActions(ctx, migrationId, optional)
Get all requeuing actions for specific migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
 **optional** | ***RequeuingActionsControllerApiGetMigrationRequeuingActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequeuingActionsControllerApiGetMigrationRequeuingActionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Limit number of actions returned. | [default to 1000]

### Return type

[**[]RequeuingAction**](RequeuingAction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMigrationRequeuingActionsStats**
> RequeuingActionsStats GetMigrationRequeuingActionsStats(ctx, migrationId)
Get all requeuing actions stats for specific migration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 

### Return type

[**RequeuingActionsStats**](RequeuingActionsStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequeuingAction**
> RequeuingAction GetRequeuingAction(ctx, migrationId, actionId)
Get requeue action by migration ID and action ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **migrationId** | **string**| Unique identifier of the migration object | 
  **actionId** | **int32**| Unique identifier of the action object | 

### Return type

[**RequeuingAction**](RequeuingAction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

