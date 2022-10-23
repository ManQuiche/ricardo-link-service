package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/link-service/internal/driving/http/link"
	"gitlab.com/ricardo134/link-service/internal/driving/http/party"
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

	linkController := link.NewController(linkService, []byte(accessSecret))
	partyController := party.NewController(partyService)
	tokenMiddleware := tokens.NewJwtAuthMiddleware([]byte(accessSecret))

	linkGroup := router.Group("/link")
	linkGroup.GET("/:link_id", tokenMiddleware.Authorize, linkController.GetOne)
	linkGroup.GET("/party/:party_id", tokenMiddleware.Authorize, linkController.GetAllForParty)
	linkGroup.POST("", tokenMiddleware.Authorize, linkController.Create)
	linkGroup.PATCH("/:link_id", tokenMiddleware.Authorize, linkController.Update)
	linkGroup.DELETE("/:link_id", tokenMiddleware.Authorize, linkController.Delete)

	router.GET("/see/:link", partyController.See)
	router.POST("/join/:link", tokenMiddleware.Authorize, partyController.Join)
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
