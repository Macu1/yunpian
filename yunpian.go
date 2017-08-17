package yunpian

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const BlackWordURL = "https://sms.yunpian.com/v2/sms/get_black_word.json"
const PullStatusUrl = "https://sms.yunpian.com/v2/sms/pull_status.json"
const GetRecordURL = "https://sms.yunpian.com/v2/sms/get_record.json"
const SMSURL = "https://sms.yunpian.com/v2/sms/single_send.json"

//const Apikey = "123"

type params map[string]string

type ypSMS interface {
	//	encode() string
	GetReqUrl() string
	GetReqBody(apikey string) io.Reader
}

type BlackWordReq struct {
	Content string
}

type PullStatus struct {
	Page_size string
}

type SMSData struct {
	Mobile string
	Text   string
}

type GetRecord struct {
	Mobile     string
	Start_time string
	End_time   string
}

func NewApiParams(apiKey string) *url.Values {
	params := &url.Values{}
	params.Add("apikey", apiKey)
	return params
}

func (w *BlackWordReq) GetReqBody(apiKey string) io.Reader {
	params := NewApiParams(apiKey)
	params.Add("text", w.Content)
	return strings.NewReader(params.Encode())
}

func (w *BlackWordReq) GetReqUrl() string {
	return BlackWordURL
}

func (w *PullStatus) GetReqBody(apiKey string) io.Reader {
	params := NewApiParams(apiKey)
	params.Add("page_size", w.Page_size)
	return strings.NewReader(params.Encode())
}

func (w *PullStatus) GetReqUrl() string {
	return PullStatusUrl
}

func (w *GetRecord) GetReqBody(apiKey string) io.Reader {
	params := NewApiParams(apiKey)
	params.Add("mobile", w.Mobile)
	params.Add("start_time", w.Start_time)
	params.Add("end_time", w.End_time)
	return strings.NewReader(params.Encode())
}

func (w *GetRecord) GetReqUrl() string {
	return GetRecordURL
}

func (w *SMSData) GetReqBody(apiKey string) io.Reader {
	params := NewApiParams(apiKey)
	params.Add("text", w.Text)
	params.Add("mobile", w.Mobile)
	return strings.NewReader(params.Encode())
}

func (w *SMSData) GetReqUrl() string {
	return SMSURL
}

func DoRequest(v ypSMS, apiKey string) (body []byte, code int, ec error) {
	client := &http.Client{}
	rbody := v.GetReqBody(apiKey)
	url := v.GetReqUrl()
	request, err := http.NewRequest("POST", url, rbody)
	response, _ := client.Do(request)
	body, _ = ioutil.ReadAll(response.Body)
	return body, response.StatusCode, err
}

type GetRecordResult struct {
	Sid               int
	Mobile            string
	Send_time         string
	Text              string
	Send_status       string
	Report_status     string
	Fee               int
	User_receive_time string
	Error_msg         string
	Uid               string
}


