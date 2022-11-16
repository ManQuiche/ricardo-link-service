package link

import "github.com/gin-gonic/gin"

// See
// @Summary See a magiclink
// @Description See the party behind the magiclink
// @Param magic_link path string true "Link id"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/see/{magic_link} [GET]
func (c controller) See(gtx *gin.Context) {}

// Join
// @Summary Join a party
// @Description Join the party behind the magiclink
// @Param magic_link path string true "Link id"
// @Success 200 {object} entities.Link
// @Failure 400 {object} errorsext.RicardoError
// @Failure 404 {object} errorsext.RicardoError
// @Router /link/join/{magic_link} [GET]
func (c controller) Join(gtx *gin.Context) {}
