package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/business"
	"github.com/PhuPhuoc/hrm_nextbean_api/services/AccountServices/repository"
	"github.com/PhuPhuoc/hrm_nextbean_api/utils"
	"github.com/gorilla/mux"
)

// @Summary		delete an account
// @Description	delete account information
// @Tags			Accounts
// @Accept			json
// @Produce		json
// @Param			account-id	path		string					true	"Account ID"
// @Success		200			{object}	utils.success_response	"Successful delete"
// @Failure		400			{object}	utils.error_response	"delete failure"
// @Router			/accounts/{account-id} [delete]
func handleDeleteAccount(db *sql.DB) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		accountID := mux.Vars(req)["account-id"]
		if accountID == "" {
			utils.WriteJSON(rw, utils.ErrorResponse_BadRequest("Missing account ID", fmt.Errorf("missing account ID")))
			return
		}

		store := repository.NewAccountStore(db)
		biz := business.NewDeleteAccountBusiness(store)
		if err := biz.DeleteAccountBiz(accountID); err != nil {
			if strings.Contains(err.Error(), "invalid-request") {
				utils.WriteJSON(rw, utils.ErrorResponse_BadRequest(err.Error(), err))
			} else {
				utils.WriteJSON(rw, utils.ErrorResponse_DB(err))
			}
			return
		}
		utils.WriteJSON(rw, utils.SuccessResponse_NoContent("Account deleted successfully!"))
	}
}
