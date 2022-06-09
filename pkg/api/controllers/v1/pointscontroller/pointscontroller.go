package pointscontroller

import (
	"encoding/json"
	"fmt"
	"github.com/iitheo/theofetchrewards/pkg/app/config/httpresponses"
	"github.com/iitheo/theofetchrewards/pkg/app/dto/pointsdto"
	"github.com/iitheo/theofetchrewards/pkg/app/helper"
	"github.com/iitheo/theofetchrewards/pkg/repositories/pointsrepo"
	"net/http"
)

func GetAllPoints(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	listOfPointsDTO := make([]pointsdto.GetAllPointsDTO, 0)

	listOfPointsFromRepo := pointsrepo.GetAllPointsFromRepo()
	for _, v := range listOfPointsFromRepo {
		singleItem := pointsdto.GetAllPointsDTO{
			Payer:     v.Payer,
			Points:    v.Points,
			Timestamp: v.Timestamp,
		}
		listOfPointsDTO = append(listOfPointsDTO, singleItem)
	}

	var msg string
	var responseHeader int

	if len(listOfPointsDTO) == 0 {
		msg = fmt.Sprintf("No record found")
		responseHeader = http.StatusNotFound
	} else {
		msg = fmt.Sprintf("%d record(s) successfully fetched", len(listOfPointsDTO))
		responseHeader = http.StatusOK
	}

	response := &httpresponses.HttpResponse{
		Success: true,
		Data:    listOfPointsDTO,
		Message: msg,
	}
	res.WriteHeader(responseHeader)
	_ = json.NewEncoder(res).Encode(response)

}

func AddPoints(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data := &pointsdto.AddPointsDTO{}
	var resp httpresponses.HttpResponse

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Error = fmt.Sprintf("error decoding add points data - %s", err.Error())
		resp.Message = fmt.Sprintf("error decoding add points data - %s", err.Error())
		_ = json.NewEncoder(res).Encode(resp)
		return
	}

	dataToInsertToRepo, err := helper.IsAddPointsValid(data)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Error = fmt.Sprintf("invalid add points data - %s", err.Error())
		resp.Message = fmt.Sprintf("invalid add points data - %s", err.Error())
		_ = json.NewEncoder(res).Encode(resp)
		return
	}

	pointsrepo.SavePoints(dataToInsertToRepo)
	response := &httpresponses.HttpResponse{
		Success: true,
		Data:    data,
		Message: "points successfully added.",
	}
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(response)

}

func SpendPoints(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	pointsToSpend := &pointsdto.SpendPointsRequestDTO{}
	remainingPoints := make([]pointsdto.SpendPointsResponseDTO, 0)
	var resp httpresponses.HttpResponse

	err := json.NewDecoder(req.Body).Decode(&pointsToSpend)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Error = fmt.Sprintf("error decoding points to spend - %s", err.Error())
		resp.Message = fmt.Sprintf("error decoding points to spend - %s", err.Error())
		_ = json.NewEncoder(res).Encode(resp)
		return
	}
	if pointsToSpend.Points <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		resp.Success = false
		resp.Error = fmt.Sprintf("points to spend must be greater than 0 - %s", err.Error())
		resp.Message = fmt.Sprintf("points to spend must be greater than 0 - %s", err.Error())
		_ = json.NewEncoder(res).Encode(resp)
		return
	}

	dataFromRepo := pointsrepo.SpendPoints(pointsToSpend.Points)

	for _, v := range dataFromRepo {
		singlePoint := &pointsdto.SpendPointsResponseDTO{
			Payer:  v.Payer,
			Points: v.Points,
		}
		remainingPoints = append(remainingPoints, *singlePoint)
	}

	response := &httpresponses.HttpResponse{
		Success: true,
		Data:    remainingPoints,
		Message: "points successfully fetched.",
	}
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(response)

}
