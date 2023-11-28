package stub

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {

	type info struct {
		expression string
		statusCode int
		body       string
	}

	var input info
	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

			expression := req.URL.Query().Get("expr")
			if expression != input.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("invalid expression"))
			}
			rw.WriteHeader(input.statusCode)
			rw.Write([]byte(input.body))
		}))

	defer server.Close()
	rs := RemoteSolver{
		url:    server.URL,
		client: server.Client(),
	}

	testCases := []struct {
		name   string
		input  info
		result int
		errMsg string
	}{
		{"case1", info{"2+2*10", http.StatusOK, "22"}, 22, ""},
		{"case2", info{"(2+2*10", http.StatusBadRequest, "invalid expression"}, 0, "invalid expression"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			input = testCase.input
			result, err := rs.Resolve(testCase.input.expression)
			if result != testCase.result {
				t.Errorf("expected %d, got %d", testCase.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != testCase.errMsg {
				t.Errorf("io error %s, got %s", testCase.errMsg, errMsg)
			}
		})
	}

}
