/**
 * file: web/handler/error_json.go
 * description: file responsible for the error layer of the application.
 * data: 01/28/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package handler

import "encoding/json"

func jsonError(message string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		message,
	}

	result, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}

	return result
}