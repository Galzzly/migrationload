# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Atom**](NotificationControllerApi.md#Atom) | **Get** /notifications/atom | Get Atom feed for Notifications
[**FetchNotification**](NotificationControllerApi.md#FetchNotification) | **Get** /notifications/{notificationId} | Get a Notification
[**FetchNotifications**](NotificationControllerApi.md#FetchNotifications) | **Get** /notifications | Notifications

# **Atom**
> interface{} Atom(ctx, )
Get Atom feed for Notifications

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/atom+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNotification**
> Notification FetchNotification(ctx, notificationId)
Get a Notification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **notificationId** | **string**| Identifier for Notification | 

### Return type

[**Notification**](Notification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNotifications**
> []Notification FetchNotifications(ctx, optional)
Notifications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NotificationControllerApiFetchNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationControllerApiFetchNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Limit number of Notifications returned. | 
 **since** | **optional.Time**| If a notification&#x27;s last modification time is earlier than this, it will be excluded. + \&quot;Expects the ISO format e.g. \\\&quot;2020-02-03T10:15:05-05:00\\\&quot;. | 
 **type_** | **optional.String**| If this parameter is supplied only notifications of the given type will be returned. | 
 **excludeResolved** | **optional.Bool**| Exclude notifications which have been marked resolved (default false) | [default to false]
 **level** | [**optional.Interface of NotificationLevel**](.md)| If this parameter is supplied only notifications of the given level will be returned. | 

### Return type

[**[]Notification**](Notification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

