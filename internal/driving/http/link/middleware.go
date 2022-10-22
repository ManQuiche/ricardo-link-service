package link

import (
	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/link-service/internal/core/app"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type ginMiddleware func(gtx *gin.Context)

func validateMiddleware(linkService app.LinkService, linkParam string) func(gtx *gin.Context) {
	return func(gtx *gin.Context) {
		linkString := gtx.Param(linkParam)
		magicLink, err := entities.NewMagicLinkFromString(linkString)
		if err != nil {
			_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
			return
		}

		valid, err := linkService.IsValid(gtx.Request.Context(), magicLink)
		if err != nil || !valid {
			_ = errorsext.GinErrorHandler(gtx, errorsext.New(errorsext.ErrBadRequest, err.Error()))
			return
		}

		gtx.Next()
	}
}
