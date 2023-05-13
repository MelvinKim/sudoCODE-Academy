package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/MelvinKim/users/application/common/dto"
	"github.com/MelvinKim/users/presentation"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/imroc/req"
)

var srv *http.Server
var baseURL string
var serverErr error

func startTestServer(ctx context.Context) (*http.Server, string, error) {
	// prepare the server
	port := randomPort()
	srv := presentation.PrepareServer(ctx, port)
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	fmt.Println("base url: ", baseURL)
	if srv == nil {
		return nil, "", fmt.Errorf("nil test server")
	}

	// set up the TCP listener
	// this is done early so that we are sure we can connect to the port in
	// the tests; backlogs will be sent to the listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, "", fmt.Errorf("unable to listen on port %d: %w", port, err)
	}
	if l == nil {
		return nil, "", fmt.Errorf("nil test server listener")
	}
	log.Printf("LISTENING on port %d", port)

	// start serving
	go func() {
		err := srv.Serve(l)
		if err != nil {
			log.Printf("serve error: %s", err)
		}
	}()

	// the cleanup of this server (deferred shutdown) needs to occur in the
	// acceptance test that will use this
	return srv, baseURL, nil
}

func TestMain(m *testing.M) {
	// setup
	ctx := context.Background()
	srv, baseURL, serverErr = startTestServer(ctx) // set the globals
	if serverErr != nil {
		log.Printf("unable to start test server: %s", serverErr)
	}

	// run the tests
	code := m.Run()

	// cleanup here
	defer func() {
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("test server shutdown error: %s", err)
		}
	}()
	os.Exit(code)
}

func randomPort() int {
	rand.Seed(time.Now().Unix())
	min := 32768
	max := 60999
	port := rand.Intn(max-min+1) + min
	return port
}

func TestHandlersInterfacesImpl_CreateStudent(t *testing.T) {
	client := http.Client{}
	headers := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	payload := dto.StudentCreationPayload{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}

	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("failed to marshall payload: %v", err)
		return
	}
	validPayload := bytes.NewBuffer(marshalled)

	type args struct {
		url        string
		httpMethod string
		headers    map[string]string
		body       io.Reader
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "Happy Case: Valid payload",
			args: args{
				url:        fmt.Sprintf("%s/api/v1/users", baseURL),
				httpMethod: http.MethodPost,
				headers:    headers,
				body:       validPayload,
			},
			wantStatus: http.StatusCreated,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			for k, v := range tt.args.headers {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf(
					"expected status %d, got %d and response %s",
					tt.wantStatus,
					resp.StatusCode,
					string(data),
				)
				return
			}
		})
	}

}

func TestHandlersInterfacesImpl_GetStudent(t *testing.T) {
	client := http.Client{}
	headers := req.Header{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	payload := dto.GetStudentPayload{
		Email: "test@test.com",
	}

	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("failed to marshall payload: %v", err)
		return
	}
	validPayload := bytes.NewBuffer(marshalled)

	type args struct {
		url        string
		httpMethod string
		headers    map[string]string
		body       io.Reader
	}

	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "Happy Case: Valid payload",
			args: args{
				url:        fmt.Sprintf("%s/api/v1/user", baseURL),
				httpMethod: http.MethodGet,
				headers:    headers,
				body:       validPayload,
			},
			wantStatus: http.StatusOK,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(
				tt.args.httpMethod,
				tt.args.url,
				tt.args.body,
			)
			if err != nil {
				t.Errorf("can't create new request: %v", err)
				return
			}

			for k, v := range tt.args.headers {
				r.Header.Add(k, v)
			}

			resp, err := client.Do(r)
			if err != nil {
				t.Errorf("HTTP error: %v", err)
				return
			}

			if resp == nil {
				t.Errorf("unexpected nil response (did not expect an error)")
				return
			}

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read response body: %v", err)
				return
			}

			if data == nil {
				t.Errorf("nil response body data")
				return
			}

			if tt.wantStatus != resp.StatusCode {
				t.Errorf(
					"expected status %d, got %d and response %s",
					tt.wantStatus,
					resp.StatusCode,
					string(data),
				)
				return
			}
		})
	}

}
