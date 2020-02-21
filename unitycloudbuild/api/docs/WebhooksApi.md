# \WebhooksApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddHookForOrg**](WebhooksApi.md#AddHookForOrg) | **Post** /orgs/{orgid}/hooks | Add hook for organization
[**AddHookForProject**](WebhooksApi.md#AddHookForProject) | **Post** /orgs/{orgid}/projects/{projectid}/hooks | Add hook for project
[**DeleteHook**](WebhooksApi.md#DeleteHook) | **Delete** /orgs/{orgid}/hooks/{id} | Delete organization hook
[**GetHook**](WebhooksApi.md#GetHook) | **Get** /orgs/{orgid}/hooks/{id} | Get organization hook details
[**GetProjectsHook**](WebhooksApi.md#GetProjectsHook) | **Get** /orgs/{orgid}/projects/{projectid}/hooks/{id} | Get project hook details
[**ListHooksForOrg**](WebhooksApi.md#ListHooksForOrg) | **Get** /orgs/{orgid}/hooks | List hooks for organization
[**ListHooksForProject**](WebhooksApi.md#ListHooksForProject) | **Get** /orgs/{orgid}/projects/{projectid}/hooks | List hooks for project
[**OrgsOrgidProjectsProjectidHooksIdDelete**](WebhooksApi.md#OrgsOrgidProjectsProjectidHooksIdDelete) | **Delete** /orgs/{orgid}/projects/{projectid}/hooks/{id} | Delete project hook
[**OrgsOrgidProjectsProjectidHooksIdPingPost**](WebhooksApi.md#OrgsOrgidProjectsProjectidHooksIdPingPost) | **Post** /orgs/{orgid}/projects/{projectid}/hooks/{id}/ping | Ping a project hook
[**OrgsOrgidProjectsProjectidHooksIdPut**](WebhooksApi.md#OrgsOrgidProjectsProjectidHooksIdPut) | **Put** /orgs/{orgid}/projects/{projectid}/hooks/{id} | Update hook for project
[**PingHook**](WebhooksApi.md#PingHook) | **Post** /orgs/{orgid}/hooks/{id}/ping | Ping an org hook
[**UpdateHook**](WebhooksApi.md#UpdateHook) | **Put** /orgs/{orgid}/hooks/{id} | Update hook for organization



## AddHookForOrg

> map[string]interface{} AddHookForOrg(ctx, orgid, optional)
Add hook for organization

Adds a new organization level hook. An organization level hook is triggered by events from all projects belonging to the organziation. NOTE: you must be a manager in the organization to add new hooks. <h4>Hook Type Configuration Parameters</h4> <div class=\"webhook-tag-desc\"> <table> <tr><th>Type</th><th>Configuration Options</th></tr> <tr>    <td><code>web</code>    <td>       <table>          <tr><th>url</th><td>Endpoint to submit POST request</td></tr>          <tr><th>encoding</th><td>Either <code>json</code> (default) or <code>form</code></td></tr>          <tr><th>sslVerify</th><td>Verify SSL certificates of HTTPS endpoint</td></tr>          <tr><th>secret</th><td>Used to compute the SHA256 HMAC signature of the hook body and adds          a <code>X-UnityCloudBuild-Signature</code> header to the payload</td></tr>       </table>    </td> </tr> <tr>    <td><code>slack</code>    <td>       <table>          <tr><th>url</th><td>Slack incoming webhook URL. Learn more at https://api.slack.com/incoming-webhooks</td></tr>       </table>    </td> </tr> </table> </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
 **optional** | ***AddHookForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddHookForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **options** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddHookForProject

> map[string]interface{} AddHookForProject(ctx, orgid, projectid, optional)
Add hook for project

Adds a new project level hook. A project level hook is only triggered by events from the specific project. NOTE: you must be a manager in the organization to add new hooks. <h4>Hook Type Configuration Parameters</h4> <div class=\"webhook-tag-desc\"> <table> <tr><th>Type</th><th>Configuration Options</th></tr> <tr>    <td><code>web</code>    <td>       <table>          <tr><th>url</th><td>Endpoint to submit POST request</td></tr>          <tr><th>encoding</th><td>Either <code>json</code> (default) or <code>form</code></td></tr>          <tr><th>sslVerify</th><td>Verify SSL certificates of HTTPS endpoint</td></tr>          <tr><th>secret</th><td>Used to compute the SHA256 HMAC signature of the hook body and adds          a <code>X-UnityCloudBuild-Signature</code> header to the payload</td></tr>       </table>    </td> </tr> <tr>    <td><code>slack</code>    <td>       <table>          <tr><th>url</th><td>Slack incoming webhook URL. Learn more at https://api.slack.com/incoming-webhooks</td></tr>       </table>    </td> </tr> </table> </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
 **optional** | ***AddHookForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddHookForProjectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **options** | [**optional.Interface of InlineObject4**](InlineObject4.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHook

> string DeleteHook(ctx, orgid, id)
Delete organization hook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**id** | **string**| Hook record identifier | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHook

> map[string]interface{} GetHook(ctx, orgid, id)
Get organization hook details

Get details of a hook by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**id** | **string**| Hook record identifier | 

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


## GetProjectsHook

> map[string]interface{} GetProjectsHook(ctx, orgid, projectid, id)
Get project hook details

Get details of a hook by id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**id** | **string**| Hook record identifier | 

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


## ListHooksForOrg

> []InlineResponse2004 ListHooksForOrg(ctx, orgid)
List hooks for organization

List all hooks configured for the specified organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHooksForProject

> []InlineResponse2004 ListHooksForProject(ctx, orgid, projectid)
List hooks for project

List all hooks configured for the specified project

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidHooksIdDelete

> string OrgsOrgidProjectsProjectidHooksIdDelete(ctx, orgid, projectid, id)
Delete project hook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**id** | **string**| Hook record identifier | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidHooksIdPingPost

> string OrgsOrgidProjectsProjectidHooksIdPingPost(ctx, orgid, projectid, id)
Ping a project hook

Send a ping event to a project hook. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**id** | **string**| Hook record identifier | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgsOrgidProjectsProjectidHooksIdPut

> map[string]interface{} OrgsOrgidProjectsProjectidHooksIdPut(ctx, orgid, projectid, id, optional)
Update hook for project

Update an existing hook. NOTE: you must be a manager of the project to update hooks. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**id** | **string**| Hook record identifier | 
 **optional** | ***OrgsOrgidProjectsProjectidHooksIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a OrgsOrgidProjectsProjectidHooksIdPutOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **options** | [**optional.Interface of InlineObject5**](InlineObject5.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingHook

> string PingHook(ctx, orgid, id)
Ping an org hook

Send a ping event to an org hook. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**id** | **string**| Hook record identifier | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHook

> map[string]interface{} UpdateHook(ctx, orgid, id, optional)
Update hook for organization

Update a new hook. NOTE: you must be a manager in the organization to update hooks. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**id** | **string**| Hook record identifier | 
 **optional** | ***UpdateHookOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateHookOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **options** | [**optional.Interface of InlineObject3**](InlineObject3.md)|  | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

