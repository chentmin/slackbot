# \OrgsApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBillingPlans**](OrgsApi.md#GetBillingPlans) | **Get** /orgs/{orgid}/billingplan | Get billing plan
[**GetSSHKey**](OrgsApi.md#GetSSHKey) | **Get** /orgs/{orgid}/sshkey | Get SSH Key
[**RegenerateSSHKey**](OrgsApi.md#RegenerateSSHKey) | **Post** /orgs/{orgid}/sshkey | Regenerate SSH Key



## GetBillingPlans

> map[string]interface{} GetBillingPlans(ctx, orgid)
Get billing plan

Get the billing plan for the specified organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHKey

> map[string]interface{} GetSSHKey(ctx, orgid)
Get SSH Key

Get the ssh public key for the specified org

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateSSHKey

> map[string]interface{} RegenerateSSHKey(ctx, orgid)
Regenerate SSH Key

Regenerate the ssh key for the specified org *WARNING* this is a destructive operation that will permanently remove your current SSH key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

