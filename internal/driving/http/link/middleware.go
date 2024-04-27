package link

import (
	"fmt"
	"github.com/gin-gonic/gin"
	errorsext "gitlab.com/ricardo-public/errors/v2/pkg/errors"
	"gitlab.com/ricardo134/link-service/internal/core/app/link"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

func ValidateMiddleware(linkService link.Service, linkParam string) func(gtx *gin.Context) {
	return func(gtx *gin.Context) {
		linkString := gtx.Param(linkParam)
		magicLink, err := entities.NewMagicLinkFromString(linkString)
		if err != nil {
			_ = errorsext.GinErrorHandler(gtx, fmt.Errorf("%s: %w", err, errorsext.ErrBadRequest))
			gtx.Abort()
			return
		}

		valid, err := linkService.IsValid(gtx.Request.Context(), magicLink)
		if err != nil || !valid {
			_ = errorsext.GinErrorHandler(gtx, fmt.Errorf("%s: %w", "link not valid", errorsext.ErrBadRequest))
			gtx.Abort()
			return
		}

		gtx.Next()
	}
}
