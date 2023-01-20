package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	gintracing "gitlab.com/ricardo-public/tracing/pkg/gin"
	"gitlab.com/ricardo134/link-service/internal/driving/http/link"
	"gitlab.com/ricardo134/link-service/internal/driving/http/party"
	"log"
	"net/http"
)

var (
	router *gin.Engine

	extLinkURL string
)

func initRoutes() {
	router.Use(gintracing.TraceRequest)

	// Ready route
	router.GET("/", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})

	linkController := link.NewController(linkService, []byte(accessSecret))
	partyController := party.NewController(partyService)
	tokenMiddleware := tokens.NewJwtAuthMiddleware([]byte(accessSecret))
	magicMiddleware := link.ValidateMiddleware(linkService, "link")

	linkGroup := router.Group("/link")
	linkGroup.GET("", tokenMiddleware.Authorize, linkController.Get)
	linkGroup.GET("/:link_id", tokenMiddleware.Authorize, linkController.GetOne)
	linkGroup.GET("/party/:party_id", tokenMiddleware.Authorize, linkController.GetAllForParty)
	linkGroup.POST("", tokenMiddleware.Authorize, linkController.Create)
	linkGroup.PATCH("/:link_id", tokenMiddleware.Authorize, linkController.Update)
	linkGroup.DELETE("/:link_id", tokenMiddleware.Authorize, linkController.Delete)

	router.GET("/see/:link", magicMiddleware, partyController.See)
	router.POST("/join/:link", magicMiddleware, tokenMiddleware.Authorize, partyController.Join)
}

func ServeHTTP() {
	router = gin.Default()
	_ = router.SetTrustedProxies(nil)

	initRoutes()

	appURL := fmt.Sprintf("%s:%s", url, port)
	extLinkURL = fmt.Sprintf("%s/see/", appURL)
	log.Printf("Launching server on %s...\n", appURL)

	log.Fatalln(router.Run(appURL))

	// TODO: go func and etc
	//log.Println("HTTP server stopped, exiting...")
}
