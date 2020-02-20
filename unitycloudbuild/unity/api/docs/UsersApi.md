# \UsersApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserApiKey**](UsersApi.md#GetUserApiKey) | **Get** /users/me/apikey | Get current user&#39;s API key
[**GetUserSelf**](UsersApi.md#GetUserSelf) | **Get** /users/me | Get current user
[**RegenApiKey**](UsersApi.md#RegenApiKey) | **Post** /users/me/apikey | Regenerate API Key
[**UpdateUserSelf**](UsersApi.md#UpdateUserSelf) | **Put** /users/me | Update current user



## GetUserApiKey

> map[string]interface{} GetUserApiKey(ctx, )
Get current user's API key

Get the currently authenticated user's API key.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSelf

> map[string]interface{} GetUserSelf(ctx, optional)
Get current user

Get the currently authenticated user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserSelfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserSelfOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **optional.String**| Extra fields to include in the response | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenApiKey

> InlineResponse2002 RegenApiKey(ctx, )
Regenerate API Key

Remove current API key and generate a new one. *WARNING* you will need to use the returned API key in all subsequent calls.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserSelf

> map[string]interface{} UpdateUserSelf(ctx, optional)
Update current user

You can update a few fields on the current user. Each field is optional and you do not need to specify all fields on update.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateUserSelfOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserSelfOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **options** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

