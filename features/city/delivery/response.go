package delivery

import "capstone-alta1/features/city"

type CityResponse struct {
	ID       uint   `json:"id"`
	CityName string `json:"city_name"`
}

func fromCore(dataCore city.Core) CityResponse {
	return CityResponse{
		ID:       dataCore.ID,
		CityName: dataCore.CityName,
	}
}

func fromCoreList(dataCore []city.Core) []CityResponse {
	var dataResponse []CityResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
