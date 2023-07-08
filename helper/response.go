package helper

import "crud-simple-api/models"

func Response(status int, message string, data interface{}) *models.Response {
	return &models.Response{
		Status: 	status,
		Message: 	message,
		Data: 		data,
	}
}

// func FailedResponse(status int, data interface{}) *models.Response {
// 	return &models.Response{
// 		Status: 	status,
// 		Message: 	"Failed",
// 		Data: 		data,
// 	}
// }

