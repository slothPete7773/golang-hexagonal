package apihandler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{
		accService: accService,
	}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customer_id"])

	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	reqAcc := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&reqAcc)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	response, err := h.accService.NewAccount(customerId, reqAcc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customer_id"])

	responses, err := h.accService.GetAccounts(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
