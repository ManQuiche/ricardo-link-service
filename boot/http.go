package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/link-service/internal/driving/http/link"
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
	tokenMiddleware := tokens.NewJwtAuthMiddleware([]byte(accessSecret))

	linkGroup := router.Group("/link")
	linkGroup.GET("/:link", tokenMiddleware.Authorize, linkController.GetOne)
	linkGroup.GET("/party/:party_id", tokenMiddleware.Authorize, linkController.GetAllForParty)
	linkGroup.POST("", tokenMiddleware.Authorize, linkController.Create)
	linkGroup.PATCH("/:link", tokenMiddleware.Authorize, linkController.Update)
	linkGroup.DELETE("/:link", tokenMiddleware.Authorize, linkController.Delete)

	joinGroup := router.Group("/join")
	joinGroup.POST("/:link", tokenMiddleware.Authorize, func(context *gin.Context) {
		// TODO: add logic for joining a party
		panic("implement me")
	})
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
