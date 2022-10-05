package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/link-service/internal/driving/http/invite"
	"log"
	"net/http"
)

var (
	router *gin.Engine
)

func initRoutes() {
	// Ready route
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	inviteController := invite.NewController(inviteService, []byte(accessSecret))
	tokenMiddleware := tokens.NewJwtAuthMiddleware([]byte(accessSecret))

	partyGroup := router.Group("/invites")
	partyGroup.GET("", tokenMiddleware.Authorize, inviteController.Get)
	partyGroup.GET("/user/:user_id", tokenMiddleware.Authorize, inviteController.GetAllForUser)
	partyGroup.GET("/party/:party_id", tokenMiddleware.Authorize, inviteController.GetOne)
	partyGroup.POST("", tokenMiddleware.Authorize, inviteController.Create)
	partyGroup.PATCH("", tokenMiddleware.Authorize, inviteController.Update)
	partyGroup.DELETE("", tokenMiddleware.Authorize, inviteController.Delete)
}

func ServeHTTP() {
	router = gin.Default()

	initRoutes()

	appURL := fmt.Sprintf("%s:%s", url, port)
	log.Printf("Launching server on %s...\n", appURL)

	log.Fatalln(router.Run(appURL))

	// TODO: go func and etc
	//log.Println("HTTP server stopped, exiting...")
}
