package cmd

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *GetHeadResponse
	expectedError error
}

func TestInputValidURL1(t *testing.T) {
	//Mock GET response:
	test := Tests{

		name: "Should return the correct output when input is valid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("137.22.165.138"))
		})),
		response: &GetHeadResponse{
			Url:           "http://checkip.amazonaws.com",
			StatusCode:    200,
			ContentLength: 15,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := GetHeadResponseService("http://checkip.amazonaws.com")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL2(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the correct output when input is valid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`<html>
			<head><title>301 Moved Permanently</title></head>
			<body>
			<center><h1>301 Moved Permanently</h1></center>
			<hr><center>nginx/1.16.1</center>
			</body>
			</html>`))
		})),
		response: &GetHeadResponse{
			Url:           "http://www.bbc.co.uk/iplayer",
			StatusCode:    301,
			ContentLength: 169,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := GetHeadResponseService("http://www.bbc.co.uk/iplayer")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL3(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the correct output when input is valid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`<html>
			<head><title>301 Moved Permanently</title></head>
			<body>
			<center><h1>301 Moved Permanently</h1></center>
			<hr><center>nginx/1.16.1</center>
			</body>
			</html>`))
		})),
		response: &GetHeadResponse{
			Url:           "http://www.bbc.co.uk/missing/thing",
			StatusCode:    301,
			ContentLength: 162,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := GetHeadResponseService("http://www.bbc.co.uk/missing/thing")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInValidURL1(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the fields:0 output when input is invalid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(``))
		})),
		response: &GetHeadResponse{
			Url:           "HelloWorld",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := GetHeadResponseService("HelloWorld")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInValidURL2(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the fields:0 output when input is invalid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(""))
		})),
		response: &GetHeadResponse{
			Url:           "https://bbc.",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := GetHeadResponseService("https://bbc.")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}
