# \CredentialsApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredentialsAndroid**](CredentialsApi.md#AddCredentialsAndroid) | **Post** /orgs/{orgid}/projects/{projectid}/credentials/signing/android | Upload Android Credentials
[**AddCredentialsAndroidForOrg**](CredentialsApi.md#AddCredentialsAndroidForOrg) | **Post** /orgs/{orgid}/credentials/signing/android | Upload Android Credentials
[**AddCredentialsIos**](CredentialsApi.md#AddCredentialsIos) | **Post** /orgs/{orgid}/projects/{projectid}/credentials/signing/ios | Upload iOS Credentials
[**AddCredentialsIosForOrg**](CredentialsApi.md#AddCredentialsIosForOrg) | **Post** /orgs/{orgid}/credentials/signing/ios | Upload iOS Credentials for organization
[**DeleteAndroid**](CredentialsApi.md#DeleteAndroid) | **Delete** /orgs/{orgid}/projects/{projectid}/credentials/signing/android/{credentialid} | Delete Android Credentials
[**DeleteAndroidForOrg**](CredentialsApi.md#DeleteAndroidForOrg) | **Delete** /orgs/{orgid}/credentials/signing/android/{credentialid} | Delete Android Credentials for organization
[**DeleteIos**](CredentialsApi.md#DeleteIos) | **Delete** /orgs/{orgid}/projects/{projectid}/credentials/signing/ios/{credentialid} | Delete iOS Credentials
[**DeleteIosForOrg**](CredentialsApi.md#DeleteIosForOrg) | **Delete** /orgs/{orgid}/credentials/signing/ios/{credentialid} | Delete iOS Credentials for organization
[**GetAllAndroid**](CredentialsApi.md#GetAllAndroid) | **Get** /orgs/{orgid}/projects/{projectid}/credentials/signing/android | Get All Android Credentials
[**GetAllAndroidForOrg**](CredentialsApi.md#GetAllAndroidForOrg) | **Get** /orgs/{orgid}/credentials/signing/android | Get All Android Credentials for an organization
[**GetAllIos**](CredentialsApi.md#GetAllIos) | **Get** /orgs/{orgid}/projects/{projectid}/credentials/signing/ios | Get All iOS Credentials
[**GetAllIosForOrg**](CredentialsApi.md#GetAllIosForOrg) | **Get** /orgs/{orgid}/credentials/signing/ios | Get All iOS Credentials for an oganization
[**GetOneAndroid**](CredentialsApi.md#GetOneAndroid) | **Get** /orgs/{orgid}/projects/{projectid}/credentials/signing/android/{credentialid} | Get Android Credential Details
[**GetOneAndroidForOrg**](CredentialsApi.md#GetOneAndroidForOrg) | **Get** /orgs/{orgid}/credentials/signing/android/{credentialid} | Get Android Credential Details for organization
[**GetOneIos**](CredentialsApi.md#GetOneIos) | **Get** /orgs/{orgid}/projects/{projectid}/credentials/signing/ios/{credentialid} | Get iOS Credential Details
[**GetOneIosForOrg**](CredentialsApi.md#GetOneIosForOrg) | **Get** /orgs/{orgid}/credentials/signing/ios/{credentialid} | Get iOS Credential Details for organization
[**UpdateAndroid**](CredentialsApi.md#UpdateAndroid) | **Put** /orgs/{orgid}/projects/{projectid}/credentials/signing/android/{credentialid} | Update Android Credentials
[**UpdateAndroidForOrg**](CredentialsApi.md#UpdateAndroidForOrg) | **Put** /orgs/{orgid}/credentials/signing/android/{credentialid} | Update Android Credentials for organization
[**UpdateIos**](CredentialsApi.md#UpdateIos) | **Put** /orgs/{orgid}/projects/{projectid}/credentials/signing/ios/{credentialid} | Update iOS Credentials
[**UpdateIosForOrg**](CredentialsApi.md#UpdateIosForOrg) | **Put** /orgs/{orgid}/credentials/signing/ios/{credentialid} | Update iOS Credentials for organization



## AddCredentialsAndroid

> map[string]interface{} AddCredentialsAndroid(ctx, orgid, projectid, label, file, alias, keypass, storepass)
Upload Android Credentials

Upload a new android keystore for the project. NOTE: you must be a manager in the project's organization to add new credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**label** | **string**| Label for the uploaded credential | 
**file** | ***os.File*****os.File**| Keystore file | 
**alias** | **string**| Keystore alias | 
**keypass** | **string**| Keystore keypass | 
**storepass** | **string**| Keystore storepass | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCredentialsAndroidForOrg

> map[string]interface{} AddCredentialsAndroidForOrg(ctx, orgid, label, file, alias, keypass, storepass)
Upload Android Credentials

Upload a new android keystore for an organization. NOTE: you must be a manager in the organization to add new credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**label** | **string**| Label for the uploaded credential | 
**file** | ***os.File*****os.File**| Keystore file | 
**alias** | **string**| Keystore alias | 
**keypass** | **string**| Keystore keypass | 
**storepass** | **string**| Keystore storepass | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCredentialsIos

> map[string]interface{} AddCredentialsIos(ctx, orgid, projectid, label, fileCertificate, fileProvisioningProfile, optional)
Upload iOS Credentials

Upload a new iOS certificate and provisioning profile for the project. NOTE: you must be a manager in the project's organization to add new credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**label** | **string**| Label for the uploaded credentials | 
**fileCertificate** | ***os.File*****os.File**| Certificate file (.p12) | 
**fileProvisioningProfile** | ***os.File*****os.File**| Provisioning Profile (.mobileprovision) | 
 **optional** | ***AddCredentialsIosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCredentialsIosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **certificatePass** | **optional.String**| Certificate (.p12) password | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddCredentialsIosForOrg

> map[string]interface{} AddCredentialsIosForOrg(ctx, orgid, label, fileCertificate, fileProvisioningProfile, optional)
Upload iOS Credentials for organization

Upload a new iOS certificate and provisioning profile for the organization. NOTE: you must be a manager in the organization to add new credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**label** | **string**| Label for the uploaded credentials | 
**fileCertificate** | ***os.File*****os.File**| Certificate file (.p12) | 
**fileProvisioningProfile** | ***os.File*****os.File**| Provisioning Profile (.mobileprovision) | 
 **optional** | ***AddCredentialsIosForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCredentialsIosForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **certificatePass** | **optional.String**| Certificate (.p12) password | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAndroid

> string DeleteAndroid(ctx, orgid, projectid, credentialid)
Delete Android Credentials

Delete specific android credentials for a project. NOTE: you must be a manager in the project's organization to delete credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## DeleteAndroidForOrg

> string DeleteAndroidForOrg(ctx, orgid, credentialid)
Delete Android Credentials for organization

Delete specific android credentials for an organization. NOTE: you must be a manager in the organization to delete credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## DeleteIos

> string DeleteIos(ctx, orgid, projectid, credentialid)
Delete iOS Credentials

Delete specific ios credentials for a project. NOTE: you must be a manager in the project's organization to delete credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## DeleteIosForOrg

> string DeleteIosForOrg(ctx, orgid, credentialid)
Delete iOS Credentials for organization

Delete specific ios credentials. NOTE: you must be a manager in the project's organization to delete credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## GetAllAndroid

> []InlineResponse2007 GetAllAndroid(ctx, orgid, projectid)
Get All Android Credentials

Get all credentials available for the project. A user in the projects org will see all credentials uploaded for any project within the org, whereas a user with project-level permissions will only see credentials assigned to the specific project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAndroidForOrg

> []InlineResponse2007 GetAllAndroidForOrg(ctx, orgid)
Get All Android Credentials for an organization

Get all credentials available for the organization. A list of projects using a credential is included in the links element. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIos

> []InlineResponse2008 GetAllIos(ctx, orgid, projectid)
Get All iOS Credentials

Get all credentials available for the project. A user in the projects org will see all credentials uploaded for any project within the org, whereas a user with project-level permissions will only see credentials assigned to the specific project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 

### Return type

[**[]InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIosForOrg

> []InlineResponse2008 GetAllIosForOrg(ctx, orgid)
Get All iOS Credentials for an oganization

Get all credentials available for the project. A user in the projects org will see all credentials uploaded for any project within the org, whereas a user with project-level permissions will only see credentials assigned to the specific project. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 

### Return type

[**[]InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOneAndroid

> map[string]interface{} GetOneAndroid(ctx, orgid, projectid, credentialid)
Get Android Credential Details

Get specific android credential details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## GetOneAndroidForOrg

> map[string]interface{} GetOneAndroidForOrg(ctx, orgid, credentialid)
Get Android Credential Details for organization

Get specific organization android credential details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## GetOneIos

> map[string]interface{} GetOneIos(ctx, orgid, projectid, credentialid)
Get iOS Credential Details

Get specific iOS credential details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## GetOneIosForOrg

> map[string]interface{} GetOneIosForOrg(ctx, orgid, credentialid)
Get iOS Credential Details for organization

Get specific iOS credential details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 

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


## UpdateAndroid

> map[string]interface{} UpdateAndroid(ctx, orgid, projectid, credentialid, optional)
Update Android Credentials

Update an android keystore for the project. NOTE: you must be a manager in the project's organization to add new credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 
 **optional** | ***UpdateAndroidOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAndroidOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **label** | **optional.String**| Label for the uploaded credential | 
 **file** | **optional.Interface of *os.File****optional.*os.File**| Keystore file | 
 **alias** | **optional.String**| Keystore alias | 
 **keypass** | **optional.String**| Keystore keypass | 
 **storepass** | **optional.String**| Keystore storepass | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAndroidForOrg

> map[string]interface{} UpdateAndroidForOrg(ctx, orgid, credentialid, optional)
Update Android Credentials for organization

Update an android keystore for the organization. NOTE: you must be a manager in the organization to update credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 
 **optional** | ***UpdateAndroidForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAndroidForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | **optional.String**| Label for the uploaded credential | 
 **file** | **optional.Interface of *os.File****optional.*os.File**| Keystore file | 
 **alias** | **optional.String**| Keystore alias | 
 **keypass** | **optional.String**| Keystore keypass | 
 **storepass** | **optional.String**| Keystore storepass | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIos

> map[string]interface{} UpdateIos(ctx, orgid, projectid, credentialid, optional)
Update iOS Credentials

Update an iOS certificate / provisioning profile for the project. NOTE: you must be a manager in the project's organization to update credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**projectid** | **string**| Project identifier | 
**credentialid** | **string**| Credential Identifier | 
 **optional** | ***UpdateIosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **label** | **optional.String**| Label for the updated credentials | 
 **fileCertificate** | **optional.Interface of *os.File****optional.*os.File**| Certificate file (.p12) | 
 **fileProvisioningProfile** | **optional.Interface of *os.File****optional.*os.File**| Provisioning Profile (.mobileprovision) | 
 **certificatePass** | **optional.String**| Certificate (.p12) password | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIosForOrg

> map[string]interface{} UpdateIosForOrg(ctx, orgid, credentialid, optional)
Update iOS Credentials for organization

Update an iOS certificate / provisioning profile. NOTE: you must be a manager in the project's organization to update credentials. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgid** | **string**| Organization identifier | 
**credentialid** | **string**| Credential Identifier | 
 **optional** | ***UpdateIosForOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIosForOrgOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **label** | **optional.String**| Label for the updated credentials | 
 **fileCertificate** | **optional.Interface of *os.File****optional.*os.File**| Certificate file (.p12) | 
 **fileProvisioningProfile** | **optional.Interface of *os.File****optional.*os.File**| Provisioning Profile (.mobileprovision) | 
 **certificatePass** | **optional.String**| Certificate (.p12) password | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[apikey](../README.md#apikey), [permissions](../README.md#permissions)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, text/plain, text/html, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

