package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/OJTServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

//	@Summary		delete an ojt
//	@Description	delete ojt information
//	@Tags			OJTS
//	@Accept			json
//	@Produce		json
//	@Param			ojt-id	path		string					true	"OJT ID"
//	@Success		200		{object}	utils.success_response	"Successful delete"
//	@Failure		400		{object}	utils.error_response	"delete failure"
//	@Router			/ojts/{ojt-id} [delete]
func handleDeleteOJT(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		ojt_id := mux.Vars(req)["ojt-id"]
		if ojt_id == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing account ID", fmt.Errorf("missing account ID")))
			return
		}

		store := repository.NewOjtStore(db)
		biz := business.NewDeleteOJTBiz(store)
		if err := biz.DeleteOJTBiz(ojt_id); err != nil {
			if strings.Contains(err.Error(), "invalid") || strings.Contains(err.Error(), "not exist") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_NoContent("Account deleted successfully!"))
	}
}
