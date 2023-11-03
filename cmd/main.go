package main

import (
	"github.com/Tristan-Muggridge/custom-oauth-provider/internal/server"

	"github.com/Tristan-Muggridge/custom-oauth-provider/internal/oauth"
)

func main() {
	server.StartServer()
	oauth.WelcomeFromHandlers()
}
