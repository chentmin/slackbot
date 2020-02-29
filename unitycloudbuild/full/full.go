package main

import (
	"context"
	"flag"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/chentmin/slackbot/slackbot"
	"github.com/gin-gonic/gin"
	"os"
)

var ginLambda *ginadapter.GinLambda

var (
	local = flag.String("local", "", "if set, local mode. should be bind addr like :8080")
)

func newGinRouter() *gin.Engine{
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	botManager := slackbot.New(os.Getenv("SLACK__TOKEN"), os.Getenv("SLACK_VERIFICATION_TOKEN"))

	botManager.RegisterMentionCommand(buildCommand, processBuildCommand)
	botManager.RegisterMentionCommand(pingCommand, processPingCommand)
	botManager.RegisterMentionCommand(installCommand, processInstallCommand)

	botManager.RegisterCallback("cancel_build", processCancelBuild)

	r.POST("/message", botManager.HandleMessageEvent)
	r.POST("/interact", botManager.HandleCallbackEvent)
	r.GET("/install", handleRedirectManifest)
	r.GET("/manifest/:tag/:build/manifest.plist", handleInstallManifest)
	r.GET("/redirect-download/:tag/:build/build.ipa", handleRedirectDownload)

	return r
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	if ginLambda == nil{
		ginLambda = ginadapter.New(newGinRouter())
	}
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	flag.Parse()

	if *local == ""{
		lambda.Start(Handler)
	} else{
		if err := newGinRouter().Run(*local); err != nil{
			panic(err)
		}
	}
}