package deleter_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"project/internal/http-server/handlers/deleter"
	"project/internal/http-server/handlers/deleter/mocks"
	"project/internal/lib/logger/handlers/slogdiscard"
)

func TestDeleteHandler(t *testing.T) {
	cases := []struct {
		name           string
		alias          string
		mockError      error
		expectedStatus int
		expectedBody   interface{} // string for non-JSON, map[string]string for JSON
	}{
		{
			name:           "Success",
			alias:          "testalias",
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "OK"},
		},
		{
			name:           "Missing alias",
			alias:          "",
			mockError:      nil,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "404 page not found\n",
		},
		{
			name:           "Delete error",
			alias:          "testalias",
			mockError:      assert.AnError,
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"error": "url deleter: assert.AnError general error for testing", "status": "Error"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.NewDeleter(t)

			if tc.alias != "" {
				mock.On("DeleteURL", tc.alias).Return(tc.mockError).Once()
			}

			r := chi.NewRouter()
			r.Delete("/{alias}", deleter.New(slogdiscard.NewDiscardLogger(), mock))

			ts := httptest.NewServer(r)
			defer ts.Close()

			url := ts.URL + "/"
			if tc.alias != "" {
				url += tc.alias
			}

			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			body := make([]byte, resp.ContentLength)
			resp.Body.Read(body)

			switch expected := tc.expectedBody.(type) {
			case string:
				assert.Equal(t, expected, string(body))
			case map[string]string:
				var actual map[string]string
				err := json.Unmarshal(body, &actual)
				require.NoError(t, err)
				assert.Equal(t, expected, actual)
			default:
				t.Fatalf("unsupported expectedBody type: %T", expected)
			}
		})
	}
}
