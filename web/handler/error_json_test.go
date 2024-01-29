/**
 * file: web/handler/error_json.go
 * description: file responsible for the error layer of the application.
 * data: 01/28/2024
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	message := "Hello Json Error"
	result := jsonError(message)

	require.Equal(t, []byte(`{"message":"Hello Json Error"}`), result)
}
