package controller

// Delete deletes article by given id
// @Tags News (Admin)
// @Summary Delete article
// @Description This endpoint is for deleting a single article with all localizations.
// @Param id path string true "Article ID"
// @Param platform header string true "The client platform" Enums(Cms, Android, Ios, Web)
// @Param session header string true "JWT"
// @Produce json
// @Success 200 {object} api.swagStatusOk
// @Router /news/admin/{id} [delete]
func Delete() {

}
