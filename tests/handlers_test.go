package tests

import (
	"encoding/json"
	"net/http"
	"testing"
)

// Utilizamos la funcion ConfigureShark para configurar la velocidad del tiburón y la ConfigurePrey para configurar la velocidad de la 
// presa. Luego, podemos llamar a la función SimulateHunt para simular la caza y verificar que la presa logra escaparse.
func TestSharkFasterThanPrey(t *testing.T) {
	srv := NewTestServer()
	defer srv.Close()
	
	sharkBody := map[string]interface{}{
		"x_position": 0.0,
		"y_position": 0.0,
		"speed": 10.0,
	}
	performRequest(srv.Handler, http.MethodPut, "/v1/shark", sharkBody)
	
	preyBody := map[string]interface{}{
		"speed": 5.0,
	}
	performRequest(srv.Handler, http.MethodPut, "/v1/prey", preyBody)
	
	resp := performRequest(srv.Handler, http.MethodPost, "/v1/simulate", nil)
	
	assertStatusCode(t, resp, http.StatusOK)
	
	var responseJSON map[string]interface{}
	json.Unmarshal([]byte(resp.Body.String()), &responseJSON)
	
	if responsJSON["success"] != false {
		t.Fatalf("Expected success to be false but got %v\n", responseJSON["success"])
	}
}
