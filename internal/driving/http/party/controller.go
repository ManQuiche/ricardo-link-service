package link

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type Controller interface {
	Join(gtx *gin.Context)
	See(gtx *gin.Context)
}

type controller struct {
	service      app.PartyService
	accessSecret []byte
}

func (c controller) See(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

}

func (c controller) Join(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

}
