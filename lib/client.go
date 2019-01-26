package lib

import (
	"github.com/zelenin/go-tdlib/client"
)

// CreateClient
func CreateClient() (*client.Client, error) {
	client.SetLogVerbosityLevel(1)
	client.SetLogFilePath("/dev/stderr")

	// client authorizer
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	const (
		apiID   = 2496
		apiHash = "8da85b0d5bfe62527e5b244c209159c3"
	)

	authorizer.TdlibParameters <- &client.TdlibParameters{
		UseTestDc:              false,
		DatabaseDirectory:      "./.tdlib/database",
		FilesDirectory:         "./.tdlib/files",
		UseFileDatabase:        false,
		UseChatInfoDatabase:    false,
		UseMessageDatabase:     false,
		UseSecretChats:         false,
		ApiId:                  apiID,
		ApiHash:                apiHash,
		SystemLanguageCode:     "en",
		DeviceModel:            "Server",
		SystemVersion:          "1.0.0",
		ApplicationVersion:     "1.0.0",
		EnableStorageOptimizer: false,
		IgnoreFileNames:        false,
	}

	return client.NewClient(authorizer)
}
