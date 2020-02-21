# OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityAssetBundles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildBundles** | **bool** | enable asset bundle builds for this target | [optional] 
**BasePath** | **string** | base path relative to Assets folder where asset bundles are output. Default is &#39;AssetBundles&#39; | [optional] 
**BuildAssetBundleOptions** | **string** | comma separated list of flags from BuildAssetBundleOptions. see https://docs.unity3d.com/ScriptReference/BuildAssetBundleOptions.html | [optional] 
**CopyToStreamingAssets** | **bool** | copy bundles to streaming assets folder, which will be packaged into the exported player. | [optional] 
**CopyBundlePatterns** | **[]string** | array of patterns to match (C# Regular Expressions) when copying asset bundle files. By default, all bundles will be copied. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


