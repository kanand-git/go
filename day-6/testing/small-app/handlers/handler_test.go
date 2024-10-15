package handlers

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"small-app/models"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	tt := []struct {
		name             string // Name of the test case
		query            string // query thar our server would receive
		expectedStatus   int    // Expected status of the response
		expectedResponse string // Expected response body
	}{
		{
			name:             "OK",
			query:            "user_id=123",
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"f_name":"Bob","l_name":"abc","email":"bob@email.com"}`,
		},
		{
			name:             "Not a number",
			query:            "user_id=abc",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"Message":"Please provide a valid user id"}`,
		},
	}
	ms := models.NewService("postgres")
	c, err := NewController(ms)
	//if err != nil {
	//	t.Fatal(err)
	//}
	require.NoError(t, err) // require is like fatalf

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// NewRecorder gives us implementation of response writer interface
			rec := httptest.NewRecorder()

			// constructing the new http request
			r, err := http.NewRequest(http.MethodGet, "localhost:8080/user?"+tc.query, nil)
			require.NoError(t, err, "problem in creating request")

			c.GetUser(rec, r)

			//getting the response body
			b := rec.Body.String()

			//TrimSpace returns a slice of the string s, with all leading and trailing white space removed
			gotResp := strings.TrimSpace(b)

			require.Equal(t, tc.expectedStatus, rec.Code, "status code")
			require.Equal(t, tc.expectedResponse, gotResp, "response body")

		})
	}
}
