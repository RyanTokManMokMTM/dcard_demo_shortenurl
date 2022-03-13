package main

import (
	"fmt"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/global"
	"github.com/RyanTokManMokMTM/dcard_demo_shortenurl/pkg/errCode"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//

func httpPerformance(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	res := httptest.NewRecorder() //test response

	r.ServeHTTP(res, req)
	return res
}

//func TestHello(t *testing.T) {
//	//open the server and listen
//	engine := setUpServer()
//
//	res := httpPerformance(engine, "GET", "/")
//	if res.Code != http.StatusOK {
//		t.Error("response error")
//	}
//
//	t.Log("Hello unit test")
//}

func TestUploadRoute(t *testing.T) {
	//set up the server
	engine := setUpServer()
	testDatas := []struct {
		URL            string
		ExpiredTime    string
		ExpectedStatus int
	}{
		//testing succeed cases
		{
			URL:            "https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals",
			ExpiredTime:    "2022-03-14T09:00:10Z",
			ExpectedStatus: errCode.Success.StatusCode(),
		},
		{
			URL:            "https://www.youtube.com/watch?v=FXE5xzR0LuA&list=RDMMFXE5xzR0LuA&start_radio=1",
			ExpiredTime:    "2022-03-15T07:00:10Z",
			ExpectedStatus: errCode.Success.StatusCode(),
		},
		{
			URL:            "https://github.com/RyanTokManMokMTM/dcard_demo_shortenurl",
			ExpiredTime:    "2022-03-18T10:00:10Z",
			ExpectedStatus: errCode.Success.StatusCode(),
		},

		//testing invalid param withs both url and expired time
		{
			//missing url and expired time
			URL:            "",
			ExpiredTime:    "",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},
		{
			//missing expired time
			URL:            "https://www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},
		{
			//missing url
			URL:            "",
			ExpiredTime:    "2022-03-18T10:00:10Z",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},

		//testing expired time passed (less than now time + 24hour - we set the expired time must greater than the current time/current day)
		{
			URL:            "https://github.com/RyanTokManMokMTM/dcard_demo_shortenurl",
			ExpiredTime:    "2022-03-12T10:00:10Z",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},

		//
		//testing invalid param with url only - must contain http:// => ://
		{
			//missing format error - without http
			URL:            "www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "2022-03-18T10:00:10Z",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},
		{
			//missing format error - missing semicolon
			URL:            "https//www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "2022-03-18T10:00:10Z",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},

		//testing invalid param with expired time only
		//just allow UTC format
		{
			//NOT UTC Format
			URL:            "https//www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "2022-03-18 10:00:10",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},
		{
			//NOT UTC Format
			URL:            "https//www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "2022-03-18",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},
		{
			//NOT UTC Format
			URL:            "https//www.youtube.com/watch?v=-7EZdOLWN18&list=RDTyWuxwqc_uY&index=2",
			ExpiredTime:    "2022-03",
			ExpectedStatus: errCode.InvalidParams.StatusCode(),
		},

		////testing url existed
		{
			URL:            "https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals",
			ExpiredTime:    "2022-03-14T09:00:10Z",
			ExpectedStatus: errCode.ErrorCreateShortenURL.StatusCode(),
		},
		{
			URL:            "https://www.youtube.com/watch?v=FXE5xzR0LuA&list=RDMMFXE5xzR0LuA&start_radio=1",
			ExpiredTime:    "2022-03-15T07:00:10Z",
			ExpectedStatus: errCode.ErrorCreateShortenURL.StatusCode(),
		},
		{
			URL:            "https://github.com/RyanTokManMokMTM/dcard_demo_shortenurl",
			ExpiredTime:    "2022-03-18T10:00:10Z",
			ExpectedStatus: errCode.ErrorCreateShortenURL.StatusCode(),
		},
	}

	apiHost := fmt.Sprintf("http://%s:%s", global.ServerSetting.Host, global.ServerSetting.Port)
	uri := "/api/v1/urls"

	url, _ := url.Parse(apiHost) //return an url object with apiHost
	url.Path = uri
	apiStr := url.String()
	for _, info := range testDatas {
		t.Run(
			fmt.Sprintf("upload url services with info : {URL:%s, ExpiredTime:%s}",
				info.URL,
				info.ExpiredTime,
			), func(t *testing.T) {
				w := uploadPerformance(info.URL, info.ExpiredTime, apiStr, engine)
				//check the response is same as expectedCode
				if w.Code != info.ExpectedStatus {
					t.Error(
						fmt.Sprintf(
							"response code:%d,expected code: %d,errors message:%s",
							w.Code,
							info.ExpectedStatus,
							w.Body))
				}

			})
	}

}

func uploadPerformance(originalURL, time, uri string, s http.Handler) *httptest.ResponseRecorder {
	//sending the request
	data := url.Values{} //key : value
	data.Add("URL", originalURL)
	data.Add("ExpiredTime", time)

	req, _ := http.NewRequest("POST", uri, strings.NewReader(data.Encode())) //end code the key-pair to string
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()
	//adding the data to request
	s.ServeHTTP(res, req)
	return res
}

func TestGetUrl(t *testing.T) {
	engine := setUpServer()
	testCase := []struct {
		UrlCode        string
		ExpectedStatus int
	}{
		//redirected
		{
			UrlCode:        "5HjtN1",
			ExpectedStatus: errCode.PermanentlyRedirect.StatusCode(), //301
		},
		{
			UrlCode:        "6HjtN1",
			ExpectedStatus: errCode.PermanentlyRedirect.StatusCode(), //301
		},
		{
			UrlCode:        "7HjtN1",
			ExpectedStatus: errCode.PermanentlyRedirect.StatusCode(), //301
		},
	}

	apiHost := fmt.Sprintf("http://%s:%s", global.ServerSetting.Host, global.ServerSetting.Port)
	uri := "/"
	url, _ := url.Parse(apiHost)
	url.Path = uri
	apiStr := url.String()
	for _, caseInfo := range testCase {
		t.Run(fmt.Sprintf("Redirect to longest url"), func(t *testing.T) {
			w := redirectPerformance(caseInfo.UrlCode, apiStr, engine)
			if w.Code != caseInfo.ExpectedStatus {
				t.Error(
					fmt.Sprintf(
						"response code:%d,expected code: %d,errors message:%s",
						w.Code,
						caseInfo.ExpectedStatus,
						w.Body))
			}
		})
	}

}

func redirectPerformance(shortenCode, path string, handler http.Handler) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", path+shortenCode, nil)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	return res
}
