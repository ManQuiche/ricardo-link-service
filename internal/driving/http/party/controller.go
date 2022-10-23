package party

import (
	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"net/http"
	"strconv"
)

type Controller interface {
	Join(gtx *gin.Context)
	See(gtx *gin.Context)
}

type controller struct {
	service app.PartyService
}

func NewController(service app.PartyService) Controller {
	return controller{service: service}
}

func (c controller) See(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

	party, err := c.service.Request(gtx.Request.Context(), magicLink.PartyID)
	if err != nil {
		// TODO: how to handle this ?
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
	}

	gtx.JSON(http.StatusOK, party)
}

func (c controller) Join(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

	userID, _ := strconv.Atoi(gtx.Param(token.UserIDKey))

	err := c.service.Joined(gtx.Request.Context(), magicLink.PartyID, uint(userID))
	if err != nil {
		// TODO: how to handle this ?
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
	}

	gtx.Status(http.StatusOK)
}
