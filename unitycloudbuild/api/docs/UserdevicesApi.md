# \UserdevicesApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](UserdevicesApi.md#CreateDevice) | **Post** /users/me/devices | Create iOS device profile
[**ListDevicesForUser**](UserdevicesApi.md#ListDevicesForUser) | **Get** /users/me/devices | List iOS device profiles



## CreateDevice

> map[string]interface{} CreateDevice(ctx, options)
Create iOS device profile

Create iOS device profile for the current user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**options** | [**InlineObject1**](InlineObject1.md)|  | 

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


## ListDevicesForUser

> []InlineResponse2003 ListDevicesForUser(ctx, )
List iOS device profiles

List all iOS device profiles for the current user

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

[apikey](../README.md#apikey), [filetoken](../README.md#filetoken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

