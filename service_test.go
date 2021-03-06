package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/borischen0203/Get2Json/dto"
	"github.com/borischen0203/Get2Json/services"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *dto.HeadResponse
	expectedError error
}

func TestInputValidURL1(t *testing.T) {
	//Create a Tests struct and Mock GET response:(the expected response url will be assign later)
	test := Tests{
		name: "Should return the correct output when input is valid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently) //mock response
			w.Write([]byte(`
			<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.16.1</center>
</body>
</html>
			`))
		})),
		response: &dto.HeadResponse{ //expected value
			Url:           "", //assign later
			StatusCode:    301,
			ContentLength: 169,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		//Assign the expected response url
		test.response.Url = test.server.URL
		defer test.server.Close()

		/**
		 *  For the mock purposes, the GetHeadResponse should be input test.server.URL but not a real URL
		 *  test.server.URL would be http://127.0.0.1:59019 and the port would change everytime.
		 */

		//input the test URL
		resp := services.GetHeadResponse(test.server.URL)

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL2(t *testing.T) {
	//Create a Tests struct and Mock GET response:(the expected response url will be assign later)
	test := Tests{
		name: "Should return the correct output when input is valid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`
			<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx</center>
</body>
</html>
			`))
		})),
		response: &dto.HeadResponse{
			Url:           "", //assign later
			StatusCode:    301,
			ContentLength: 162,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		//Assign the expected response url
		test.response.Url = test.server.URL
		defer test.server.Close()

		//input the test URL
		resp := services.GetHeadResponse(test.server.URL)

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL3(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the correct output when input is ending with multiple slashes",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`
			<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.16.1</center>
</body>
</html>
			`))
		})),
		response: &dto.HeadResponse{
			Url:           "", //assign later
			StatusCode:    301,
			ContentLength: 169,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		//Assign the expected response url
		test.response.Url = test.server.URL + "///////"
		defer test.server.Close()

		//input the test URL
		resp := services.GetHeadResponse(test.server.URL + "///////")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL4(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the correct output when input is ending with multiple space",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`
			<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.16.1</center>
</body>
</html>
			`))
		})),
		response: &dto.HeadResponse{
			Url:           "", //assign later
			StatusCode:    301,
			ContentLength: 169,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		//Assign the expected response url
		test.response.Url = test.server.URL
		defer test.server.Close()

		//input the test URL
		resp := services.GetHeadResponse(test.server.URL + "               ")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputValidURL5(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the correct output when input is starting with multiple space",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMovedPermanently)
			w.Write([]byte(`
			<html>
<head><title>301 Moved Permanently</title></head>
<body>
<center><h1>301 Moved Permanently</h1></center>
<hr><center>nginx/1.16.1</center>
</body>
</html>
			`))
		})),
		response: &dto.HeadResponse{
			Url:           "", //assign later
			StatusCode:    301,
			ContentLength: 169,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		//Assign the expected response url
		test.response.Url = test.server.URL
		defer test.server.Close()

		//input the test URL
		resp := services.GetHeadResponse("          " + test.server.URL)

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInvalidURL1(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the output with fields:0 when input is invalid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(``))
		})),
		response: &dto.HeadResponse{
			Url:           "HelloWorld",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := services.GetHeadResponse("HelloWorld")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInvalidURL2(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the output with fields:0 when input is invalid URL",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(``))
		})),
		response: &dto.HeadResponse{
			Url:           "https://bbc.",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := services.GetHeadResponse("https://bbc.")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInvalidURL3(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the output with fields:0 when input is empty",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(``))
		})),
		response: &dto.HeadResponse{
			Url:           "",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := services.GetHeadResponse("")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

func TestInputInvalidURL4(t *testing.T) {
	//Mock GET response:
	test := Tests{
		name: "Should return the output with fields:0 when input are symbols",
		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(0)
			w.Write([]byte(""))
		})),
		response: &dto.HeadResponse{
			Url:           "~!@#$%^\u0026*()_+{}|:\"\u003c\u003e?,./';\\[]=????????",
			StatusCode:    0,
			ContentLength: 0,
		},
		expectedError: nil,
	}

	//Test service
	t.Run(test.name, func(t *testing.T) {
		defer test.server.Close()

		resp := services.GetHeadResponse("~!@#$%^\u0026*()_+{}|:\"\u003c\u003e?,./';\\[]=????????")

		if !reflect.DeepEqual(resp, test.response) {
			t.Errorf("FAILED: expected %v, got %v\n", test.response, resp)
		}
	})
}

// // func TestInputTimeOutURL(t *testing.T) {
// //google.com:81
// // }
