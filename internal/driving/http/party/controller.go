package party

import (
	"fmt"
	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/v2/pkg/errors"
	"gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo134/link-service/internal/core/app/party"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"net/http"
	"strconv"
)

type Controller interface {
	Join(gtx *gin.Context)
	See(gtx *gin.Context)
}

type controller struct {
	service party.PartyService
}

func NewController(service party.PartyService) Controller {
	return controller{service: service}
}

// See
// @Summary See a magiclink
// @Description See the party behind the magiclink
// @Param magic_link path string true "Link id"
// @Success 200 {object} entities.Link
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/see/{magic_link} [GET]
func (c controller) See(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

	party, err := c.service.Request(gtx.Request.Context(), magicLink.PartyID)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, fmt.Errorf("%s: %w", err, errorsext.ErrBadRequest))
		return
	}

	gtx.JSON(http.StatusOK, party)
}

// Join
// @Summary Join a party
// @Description Join the party behind the magiclink
// @Param magic_link path string true "Link id"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/join/{magic_link} [POST]
func (c controller) Join(gtx *gin.Context) {
	linkString := gtx.Param("link")
	// Already checked by middleware
	magicLink, _ := entities.NewMagicLinkFromString(linkString)

	userID, _ := strconv.Atoi(gtx.Param(token.UserIDKey))

	err := c.service.Joined(gtx.Request.Context(), magicLink.PartyID, uint(userID))
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, fmt.Errorf("%s: %w", err, errorsext.ErrBadRequest))
		return
	}

	gtx.Status(http.StatusOK)
}
