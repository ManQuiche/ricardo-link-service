package link

import (
	"errors"
	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"net/http"
	"strconv"
)

type Controller interface {
	Create(gtx *gin.Context)
	Update(gtx *gin.Context)
	Get(gtx *gin.Context)
	GetAllForParty(gtx *gin.Context)
	GetOne(gtx *gin.Context)
	Delete(gtx *gin.Context)
}

type controller struct {
	service      app.LinkService
	accessSecret []byte
}

func NewController(service app.LinkService, accessSecret []byte) Controller {
	return controller{service: service, accessSecret: accessSecret}
}

// Create
// @Summary Create a link
// @Description Create a link
// @Param link_id path int true "Link id"
// @Param link body entities.CreateLinkRequest true "Created link info"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/{link_id} [POST]
func (c controller) Create(gtx *gin.Context) {
	var cir entities.CreateLinkRequest
	err := gtx.ShouldBindJSON(&cir)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
		return
	}

	i := entities.Link{
		PartyID:    cir.PartyID,
		Expiration: cir.Expiration,
	}
	link, err := c.service.Save(gtx.Request.Context(), i)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, link)
}

// Update
// @Summary Update a link
// @Description Update a link
// @Param link_id path int true "Link id"
// @Param link body entities.UpdateLinkRequest true "Updated link info"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/{link_id} [PATCH]
func (c controller) Update(gtx *gin.Context) {
	linkID, err := strconv.ParseUint(gtx.Param("linkID"), 10, 32)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, "invalid ID format"))
		return
	}
	uintLinkId := uint(linkID)

	var ulr entities.UpdateLinkRequest
	err = gtx.ShouldBindJSON(&ulr)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
		return
	}

	_, err = c.canUpdateOrDelete(gtx, uintLinkId)
	if err != nil {
		return
	}

	l := entities.Link{
		ID:         uintLinkId,
		Expiration: ulr.Expiration,
	}

	link, err := c.service.Save(gtx.Request.Context(), l)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, link)
}

// Get
// @Summary Get all links
// @Description Get all links
// @Success 200 {object} []entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link [GET]
func (c controller) Get(gtx *gin.Context) {
	links, err := c.service.GetAll(gtx.Request.Context())
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
		return
	}

	gtx.JSON(http.StatusOK, links)
}

// GetAllForParty
// @Summary Get all links for a party
// @Description Get all links for a party
// @Param party_id path int true "Party id"
// @Success 200 {object} []entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/party/{link_id} [GET]
func (c controller) GetAllForParty(gtx *gin.Context) {
	partyID, err := strconv.ParseUint(gtx.Param("party_id"), 10, 64)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
	}

	links, err := c.service.GetAllForParty(gtx.Request.Context(), uint(partyID))
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, links)
}

// GetOne
// @Summary Get a link
// @Description Get a link
// @Param link_id path int true "Link id"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/{link_id} [GET]
func (c controller) GetOne(gtx *gin.Context) {
	linkId, err := strconv.ParseUint(gtx.Param("link_id"), 10, 64)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
		return
	}

	links, err := c.service.Get(gtx.Request.Context(), uint(linkId))
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, links)
}

// Delete
// @Summary Delete a link
// @Description Delete a link
// @Param link_id path int true "Link id"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/{link_id} [DELETE]
func (c controller) Delete(gtx *gin.Context) {
	linkID, err := strconv.ParseUint(gtx.Param("linkID"), 10, 32)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, "invalid ID format"))
		return
	}
	uintLinkId := uint(linkID)

	_, err = c.canUpdateOrDelete(gtx, uintLinkId)
	if err != nil {
		return
	}

	err = c.service.Delete(gtx.Request.Context(), uintLinkId)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return
	}

	gtx.Status(http.StatusOK)
}

func (c controller) canUpdateOrDelete(gtx *gin.Context, linkID uint) (bool, error) {
	l, err := c.service.Get(gtx.Request.Context(), linkID)
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, ""))
		return false, err
	}

	strToken, err := tokens.ExtractTokenFromHeader(gtx.GetHeader(tokens.AuthorizationHeader))
	if err != nil {
		_ = errorsext.GinErrorHandler(gtx, err)
		return false, err
	}

	pToken, err := tokens.Parse(strToken, c.accessSecret)
	claims := pToken.Claims.(tokens.RicardoClaims)
	userID, err := strconv.ParseUint(claims.Subject, 10, 64)

	if uint(userID) != l.CreatorID && claims.Role != tokens.AdminRole {
		err = errors.New("unauthorized to update or delete")
		_ = errorsext.GinErrorHandler(gtx, err)
		return false, err
	}

	return true, nil
}
