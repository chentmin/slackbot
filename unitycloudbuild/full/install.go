package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"text/template"
)

func handleRedirectManifest(c *gin.Context){
	tag := c.Query("tag")
	buildNumber := c.Query("build")

	if tag == "" || buildNumber == ""{
		c.String(http.StatusBadRequest, "tag或build不存在")
		return
	}

	//c.HTML(http.StatusTemporaryRedirect, "redirect", gin.H{
	//	"Url": os.Getenv("SELF_URL"),
	//	"Tag": tag,
	//	"Build": buildNumber,
	//})


	url := fmt.Sprintf(`itms-services://?action=download-manifest&url=%s/manifest/%s/%s/manifest.plist`, os.Getenv("SELF_URL"), tag, buildNumber)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleRedirectDownload(c *gin.Context){
	tag := c.Param("tag")
	buildNumber := c.Param("build")

	if tag == "" || buildNumber == ""{
		c.String(http.StatusBadRequest, "tag或build不存在")
		return
	}

	result, _, err := unityClient().BuildsApi.GetBuild(c, UNITY_ORG, UNITY_PROJECT, tag, buildNumber, nil)
	if err != nil{
		fmt.Printf("unity返回错误: %s\n", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("unity返回错误: %s\n", err))
		return
	}

	switch result.BuildStatus {
	case "success":
		fmt.Printf("result: %+v\n", result)

		url := result.Links.DownloadUrl.Href

		c.Redirect(http.StatusTemporaryRedirect, url)

	default:
		fmt.Printf("当前不是success状态: %+v", result)
		c.String(http.StatusOK, "当前不是success状态")
	}
}

func handleInstallManifest(c *gin.Context){
	tag := c.Param("tag")
	buildNumber := c.Param("build")

	if tag == "" || buildNumber == ""{
		c.String(http.StatusBadRequest, "tag或build不存在")
		return
	}

	result, _, err := unityClient().BuildsApi.GetBuild(c, UNITY_ORG, UNITY_PROJECT, tag, buildNumber, nil)
	if err != nil{
		fmt.Printf("unity返回错误: %s\n", err)
		c.String(http.StatusBadRequest, fmt.Sprintf("unity返回错误: %s\n", err))
		return
	}

	switch result.BuildStatus {
	case "success":
		fmt.Printf("result: %+v\n", result)

		url := fmt.Sprintf("%s/redirect-download/%s/%s/build.ipa", os.Getenv("SELF_URL"), tag, buildNumber)
		bundleId := result.ProjectVersion.BundleId
		name := result.ProjectVersion.Name
		buildNum := result.Build

		info := &InstallInfo{
			Url:        url,
			Identifier: bundleId,
			Bundle:     fmt.Sprintf("0.0.%v", buildNum),
			GameName:   name,
		}

		b := &bytes.Buffer{}

		if err := intmp.Execute(b, info); err != nil{
			fmt.Printf("template 出错: %s\n", err)
			return
		}

		c.Data(http.StatusOK, "text/xml", b.Bytes())

	default:
		fmt.Printf("当前不是success状态: %+v", result)
		c.String(http.StatusOK, "当前不是success状态")
	}
}

type InstallInfo struct{
	Url string
	Identifier string
	Bundle string
	GameName string
}

var intmp = template.Must(template.New("in").Parse(installTemplate))

var installTemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>items</key>
    <array>
        <dict>
            <key>assets</key>
            <array>
                <dict>
                    <key>kind</key>
                    <string>software-package</string>
                    <key>url</key>
                    <string>{{.Url}}</string>
                </dict>
            </array>
            <key>metadata</key>
            <dict>
                <key>bundle-identifier</key>
                <string>{{.Identifier}}</string>
                <key>bundle-version</key>
                <string>{{.Bundle}}</string>
                <key>kind</key>
                <string>software</string>
                <key>title</key>
                <string>{{.GameName}}</string>
            </dict>
        </dict>
    </array>
</dict>
</plist>
`
