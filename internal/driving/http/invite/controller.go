package invite

import (
	"errors"
	"github.com/gin-gonic/gin"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Controller interface {
	Create(gtx *gin.Context)
	Update(gtx *gin.Context)
	Get(gtx *gin.Context)
	GetAllForUser(gtx *gin.Context)
	GetAllForParty(gtx *gin.Context)
	GetOne(gtx *gin.Context)
	Delete(gtx *gin.Context)
}

type controller struct {
	service      app.InviteService
	accessSecret []byte
}

func NewController(service app.InviteService, accessSecret []byte) Controller {
	return controller{service: service, accessSecret: accessSecret}
}

func (c controller) Create(gtx *gin.Context) {
	var cir entities.CreateInviteRequest
	err := gtx.ShouldBindJSON(&cir)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
		return
	}

	i := entities.Invite{
		PartyID: cir.PartyID,
		UserID:  cir.UserID,
	}
	invite, err := c.service.Save(gtx.Request.Context(), i)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, invite)
}

func (c controller) Update(gtx *gin.Context) {
	var uir entities.UpdateInviteRequest
	err := gtx.ShouldBindJSON(&uir)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, ""))
		return
	}

	_, err = c.canUpdateOrDelete(gtx, uir.ID)
	if err != nil {
		return
	}

	i := entities.Invite{
		Model: gorm.Model{
			ID: uir.ID,
		},
		Answered:    uir.Answered,
		Accepted:    uir.Accepted,
		Explanation: uir.Explanation,
	}

	invite, err := c.service.Save(gtx.Request.Context(), i)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, invite)
}

func (c controller) Get(gtx *gin.Context) {
	invites, err := c.service.GetAll(gtx.Request.Context())
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, ""))
		return
	}

	gtx.JSON(http.StatusOK, invites)
}

func (c controller) GetAllForUser(gtx *gin.Context) {
	userID, err := strconv.ParseUint(gtx.Param("user_id"), 10, 64)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
	}

	invites, err := c.service.GetAllForUser(gtx.Request.Context(), uint(userID))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, invites)
}

func (c controller) GetAllForParty(gtx *gin.Context) {
	partyID, err := strconv.ParseUint(gtx.Param("party_id"), 10, 64)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, err.Error()))
	}

	invites, err := c.service.GetAllForParty(gtx.Request.Context(), uint(partyID))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, invites)
}

func (c controller) GetOne(gtx *gin.Context) {
	inviteId, err := strconv.ParseUint(gtx.Param("invite_id"), 10, 64)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, ""))
		return
	}

	invites, err := c.service.Get(gtx.Request.Context(), uint(inviteId))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.JSON(http.StatusOK, invites)
}

func (c controller) Delete(gtx *gin.Context) {
	var dir entities.DeleteInviteRequest
	err := gtx.ShouldBindJSON(&dir)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, ""))
		return
	}

	_, err = c.canUpdateOrDelete(gtx, dir.ID)
	if err != nil {
		return
	}

	err = c.service.Delete(gtx.Request.Context(), dir.ID)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return
	}

	gtx.Status(http.StatusOK)
}

func (c controller) canUpdateOrDelete(gtx *gin.Context, inviteID uint) (bool, error) {
	i, err := c.service.Get(gtx.Request.Context(), inviteID)
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, ricardoErr.New(ricardoErr.ErrBadRequest, ""))
		return false, err
	}

	strToken, err := tokens.ExtractTokenFromHeader(gtx.GetHeader(tokens.AuthorizationHeader))
	if err != nil {
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return false, err
	}

	pToken, err := tokens.Parse(strToken, c.accessSecret)
	claims := pToken.Claims.(tokens.RicardoClaims)
	userID, err := strconv.ParseUint(claims.Subject, 10, 64)

	if uint(userID) != i.UserID && claims.Role != tokens.AdminRole {
		err = errors.New("unauthorized to update or delete")
		_ = ricardoErr.GinErrorHandler(gtx, err)
		return false, err
	}

	return true, nil
}
