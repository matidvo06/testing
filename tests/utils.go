// Agrego las últimas 3 funciones adicionales para ayudar a realizar las pruebas
package tests

import (
	"bytes"
	"functional/cmd/server"
	"functional/prey"
	"functional/shark"
	"functional/simulator"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func createServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	g := r.Group("/v1")

	g.PUT("/shark", handler.ConfigureShark())
	g.PUT("/prey", handler.ConfigurePrey())
	g.POST("/simulate", handler.SimulateHunt())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func performRequest(handler http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	bodyReader := strings.NewReader("")
	if body != nil {
		bodyBytes, _ := json.Marshal(body)
		bodyReader = strings.NewReader(string(bodyBytes))
	}
	
	req, _ := http.NewRequest(method, path, bodyReader)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	
	return w
}

func assertStatusCode(t *testing.T, response *httptest.ResponseRecorder, expectedStatusCode int) {
	if response.Body.String() != expectedStatusCode {
		t.Fatalf("Expected status code %d but got %d\n", expectedStatusCode, response.Code)
	}
}

func assertJSONResponse(t *testing.T, response *httptest.ResponseRecorder, expectedBody string) {
	if response.Body.String() != expectedBody {
		t.Fatalf("Expected body %s but got %s\n", expectedBody, response.Body.String())
	}
}
