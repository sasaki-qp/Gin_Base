package helper

import()

type ErrorType struct {
	statusCode int "json:statusCode"
	message string "json:message"
}

func ReturnErrorReponse(statusCode int) interface{} {
	var message string
	switch statusCode {
		case 400:
			message = "Bad request"

		case 401:
			message = "Unauthorized"

		case 403:
			message = "Forbidden"

		case 404:
			message = "Not found" 

		case 500:
			message = "Internal server error"
	}
	res := map[string] interface{} {
		"statusCode": statusCode,
		"message": message,
	}
	return res
}