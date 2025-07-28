package controllersAcc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DevOps-Group-D/YouToFy-API/models"
	servicesAcc "github.com/DevOps-Group-D/YouToFy-API/services"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding invalid account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	res, err := servicesAcc.Register(account.Username, account.Password)
	if err != nil {
		errMsg := fmt.Sprintf("Error registering account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(202)

	json.NewEncoder(w).Encode(res)
}
