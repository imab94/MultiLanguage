package services

import log "rApp/tsrApp/driver/tsrAppLogger"

var nearRtRicInfos nearRtRicInfoList
var trainedModelsJson trainedModels
var Update_interval int64

var Logger, _ = log.GetrappLogger()

//storing all the services
var Service_list_Arr = make(map[string]string)

type trainedModels struct {
	Models []ModelInfo
}
type ModelInfo struct {
	PolicyType           string
	ModelID              int64
	ModelName            string
	PlaybookName         string
	PlaybookID           int64
	TrainingDatasource   string
	LiveDatasource       string
	PredictionDatasource string
}

//Predication Data
type PredicationDataList struct {
	PredicationInstances []PredicationData
}

var PredictionDataList []struct {
	PredicationInstances PredicationData
}

type PredicationData struct {
	Cellid string
	Tag    string
}

type nearRtRicInfoList struct {
	Rics []RicInfo
}
type RicInfo struct {
	ManagedElementIds []string `json:"managed_element_ids"`
	PolicytypeIds     []string `json:"policytype_ids"`
	RicID             string   `json:"ric_id"`
	State             string   `json:"state"`
}

func SetRicInfo(nearRtRics nearRtRicInfoList) {
	Logger.RappDebug("Old NearRtRicInfoList: ", nearRtRicInfos)
	nearRtRicInfos = nearRtRics
	Logger.RappDebug("New NearRtRicInfoList: ", nearRtRicInfos)
}

func GetRicInfo() nearRtRicInfoList {
	Logger.RappDebug("Returning NearRtRicInfoList: ", nearRtRicInfos)
	return nearRtRicInfos
}

func SetModelsInfo(ipJson trainedModels) {
	Logger.RappDebug("Old trainedModelsJson: ", trainedModelsJson)
	trainedModelsJson = ipJson
	Logger.RappDebug("New trainedModelsJson: ", trainedModelsJson)
}

func GetModelsInfo() trainedModels {
	Logger.RappDebug("Returning trainedModelsJson: ", trainedModelsJson)
	return trainedModelsJson
}

func GetUpdateInterval() int64 {
	return Update_interval
}

func SetUpdateInterval(interval_value int64) {
	Update_interval = interval_value
}
