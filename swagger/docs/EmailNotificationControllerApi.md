# {{classname}}

All URIs are relative to *http://liamg01-vm5.bdfrem.wandisco.com:18080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchEmailNotificationTypes**](EmailNotificationControllerApi.md#FetchEmailNotificationTypes) | **Get** /notifications/config/email/types | Get email notifications types
[**FetchEmailRegistrations**](EmailNotificationControllerApi.md#FetchEmailRegistrations) | **Get** /notifications/config/email/registration | Get registered email address for receiving push notifications about specific notification types
[**FetchSmtpConfig**](EmailNotificationControllerApi.md#FetchSmtpConfig) | **Get** /notifications/config/email | Get the current SMTP adaptor config
[**SendTestEmail**](EmailNotificationControllerApi.md#SendTestEmail) | **Post** /notifications/config/email/test | Send test e-mail to the specified list of recipient
[**SetEmailRegistrations**](EmailNotificationControllerApi.md#SetEmailRegistrations) | **Put** /notifications/config/email/registration | Update set of emails registered for receiving push notifications about specific notification types
[**SetSmtpConfig**](EmailNotificationControllerApi.md#SetSmtpConfig) | **Put** /notifications/config/email | Update the SMTP adaptor config

# **FetchEmailNotificationTypes**
> []string FetchEmailNotificationTypes(ctx, )
Get email notifications types

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchEmailRegistrations**
> EmailRegistrations FetchEmailRegistrations(ctx, )
Get registered email address for receiving push notifications about specific notification types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EmailRegistrations**](EmailRegistrations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchSmtpConfig**
> SmtpConfiguration FetchSmtpConfig(ctx, )
Get the current SMTP adaptor config

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SmtpConfiguration**](SmtpConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmail**
> SendTestEmail(ctx, addresses)
Send test e-mail to the specified list of recipient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addresses** | [**[]string**](string.md)| Emails of receivers | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetEmailRegistrations**
> EmailRegistrations SetEmailRegistrations(ctx, optional)
Update set of emails registered for receiving push notifications about specific notification types

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailNotificationControllerApiSetEmailRegistrationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailNotificationControllerApiSetEmailRegistrationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | [**optional.Interface of []string**](string.md)| Types of notifications | [default to []]
 **addresses** | [**optional.Interface of []string**](string.md)| Emails of receivers | [default to []]

### Return type

[**EmailRegistrations**](EmailRegistrations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSmtpConfig**
> SmtpConfiguration SetSmtpConfig(ctx, host, port, securityType, email, optional)
Update the SMTP adaptor config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **host** | **string**| Host of the SMTP server | 
  **port** | **int32**| Port of the SMTP server | 
  **securityType** | [**SmtpSecurityType**](.md)| Security type of the SMTP server | 
  **email** | **string**| Email that will be used by notification sender | 
 **optional** | ***EmailNotificationControllerApiSetSmtpConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailNotificationControllerApiSetSmtpConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **authEnabled** | **optional.Bool**| Authentication is enabled | [default to false]
 **login** | **optional.String**| Login | 
 **password** | **optional.String**| Password | 
 **subjectPrefix** | **optional.String**| E-Mail subject prefix | 

### Return type

[**SmtpConfiguration**](SmtpConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

