package service

import (
	"encoding/json"
	"net/http"

	"github.com/gautamrege/packt/sweatbead/sweatmgr/db"
	"github.com/gautamrege/packt/sweatbead/sweatmgr/logger"
)

// @Title createSweatHandler
// @Description create sweat entry for given user
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
// @Router /sweat [post]
func createSweatHandler(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var s db.Sweat
	err := decoder.Decode(&s)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	req = WithUserContext(req)
	err = s.Create(req.Context())
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}

func getSweatSamplesHandler(rw http.ResponseWriter, req *http.Request) {
	sweats, err := db.ListAllSweat()
	if err != nil {
		logger.Get().Info("Error fetching data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(sweats)
	if err != nil {
		logger.Get().Info("Error marshaling data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(respBytes)

}

func getSweatByUserIdHandler(rw http.ResponseWriter, req *http.Request) {
	req = WithUserContext(req)

	sweats, err := db.ListUserSweat(req.Context())
	if err != nil {
		logger.Get().Info("Error fetching data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBytes, err := json.Marshal(sweats)
	if err != nil {
		logger.Get().Info("Error marshaling data", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(respBytes)
}