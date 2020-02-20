# OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreExportMethod** | **string** | The fully-qualified name of a public static method you want us to call before we start the Unity build process. For example: ClassName.NeatMethod or NameSpace.ClassName.NeatMethod. No trailing parenthesis, and it can&#39;t have the same name as your Post-Export method! | [optional] 
**PostExportMethod** | **string** | The fully-qualified name of a public static method you want us to call after we finish the Unity build process (but before Xcode). For example: ClassName.CoolMethod or NameSpace.ClassName.CoolMethod. No trailing parenthesis, and it can&#39;t have the same name as your Post-Export method! This method must accept a string parameter, which will receive the path to the exported Unity player (or Xcode project in the case of iOS). | [optional] 
**PreBuildScript** | **string** | Relative path to the script that should be run before the build process starts. | [optional] 
**PostBuildScript** | **string** | Relative path to the script that should be run after the build process finishes. | [optional] 
**ScriptingDefineSymbols** | **string** | Enter the names of the symbols you want to define for iOS. These symbols can then be used as the conditions for #if directives just like the built-in ones. (i.e. #IF MYDEFINE or #IF AMAZON) | [optional] 
**PlayerExporter** | [**OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityPlayerExporter**](_orgs_orgid_projects_projectid_buildtargets_settings_advanced_unity_playerExporter.md) |  | [optional] 
**PlayerSettings** | [**OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityPlayerSettings**](_orgs_orgid_projects_projectid_buildtargets_settings_advanced_unity_playerSettings.md) |  | [optional] 
**EditorUserBuildSettings** | [**OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityEditorUserBuildSettings**](_orgs_orgid_projects_projectid_buildtargets_settings_advanced_unity_editorUserBuildSettings.md) |  | [optional] 
**AssetBundles** | [**OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityAssetBundles**](_orgs_orgid_projects_projectid_buildtargets_settings_advanced_unity_assetBundles.md) |  | [optional] 
**Addressables** | [**OrgsOrgidProjectsProjectidBuildtargetsSettingsAdvancedUnityAddressables**](_orgs_orgid_projects_projectid_buildtargets_settings_advanced_unity_addressables.md) |  | [optional] 
**RunUnitTests** | **bool** | Run any unit tests your project has when a build happens. | [optional] 
**RunEditModeTests** | **bool** | Should Edit Mode unit tests be run? NOTE: requires runUnitTests to be true and building with Unity 5.6 or newer. | [optional] 
**RunPlayModeTests** | **bool** | Should Play Mode unit tests be run? NOTE: requires runUnitTests to be true and building with Unity 5.6 or newer. | [optional] 
**FailedUnitTestFailsBuild** | **bool** | Mark builds as failed if the unit tests do not pass. | [optional] 
**UnitTestMethod** | **string** | LEGACY - The Unity method to call when running unit tests (only supported in Unity 5.2 and lower). | [optional] 
**EnableLightBake** | **bool** | Enable lightmap baking (disabled by default since it is very slow and usually unnecessary) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


