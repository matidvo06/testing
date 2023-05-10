/*
Siguiendo con la metodología TDD, testear los siguientes casos antes del código, en el paquete tests:
Se configura la presa para que sea más veloz que el tiburón, y al simular la caza logra escaparse.
Se configura el tiburón más rápido que la presa, pero se encuentra demasiado lejos y no logra cazarla.
El tiburón y la presa se configuran de modo que el tiburón logra cazarla luego de 24 segundos (tener en cuenta el algoritmo que usa el 
simulador).
Testear, para todos los endpoints de configuración, casos donde los tipos de los campos no son esperados.
Tener en cuenta los structs creados en el archivo handlers.go. Es importante que utilicemos los métodos creados en el archivo utils.go 
del paquete tests.
*/
package tests

import (
	"encoding/json"
	"net/http"
	"testing"
	"bytes"
	"errors"
	"functional/cmd/server"
	"functional/prey"
	"functional/shark"
	"net/http/httptest"
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

func TestSharkCantCatchPrey(t *testing.T) {
	sim := &mockCatchSimulator{}
	sim.distance = 1000
	whiteShark := shark.CreateWhiteShark(sim)
	
	tuna := prey.CreateTuna()
	tuna.SetSpeed(25)
	
	handler := server.NewHandler(whiteShark, tuna)
	srv := server.NewServer(handler, nil)
	
	reqBody := map[string]float64{
		"x_position": 250,
		"y_position": 250,
		"speed": 35.4,
	}
	reqBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/v1/shark", bytes.NewReader(reqBytes))
	
	w := httptest.NewRecorder()
	srv.Engine().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
	
	reqBody = map[string]float64{
		"speed": 25,
	}
	reqBytes, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("PUT", "/v1/prey", bytes.NewReader(reqBytes))
	
	w = httptest.NewRecorder()
	srv.Engine().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
	
	req, _ = http.NewRequest("POST", "/v1/simulate", nil)
	
	w = httptest.NewRecorder()
	srv.Engine().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", w.Code)
	}
	var respBody map[string]interface{}
	json.NewDecoder(w.Body).Decode(&respBody)
	if !respBody["success"].(bool) {
		t.Errorf("expected success true; got %v", respBody["success"])
	}
	if respBody["time"].(float64) != 0 {
		t.Errorf("expected time 0; got %v", respBody["time"])
	}
	if respBody["message"].(string) != "Shark could not catch prey" {
		t.Errorf(expected message 'Shark could not catch prey'; got %v", respBody["message"])
	}
}
			 
func TestHunt(t *testing.T) {
