package cls

import (
	"net/http"
	"strings"
	"fmt"
	"net/url"
	"crypto/sha1"
	"crypto/hmac"
	//"encoding/json"
	"encoding/hex"
	"time"
	"io/ioutil"
	"encoding/json"
)

const (
	QSignAlgorithm = "sha1"
)

type ApiClient struct {
	client *http.Client

	host string

	credential TmpCredential
}

type TmpCredential struct {
	Token     string
	SecretId  string
	SecretKey string
}

func NewApiClient(cred TmpCredential, host string) (*ApiClient, error) {
	return &ApiClient{
		client:     http.DefaultClient,
		credential: cred,
		host:       host,
	}, nil
}

type ErrorResponse struct {
	ErrorCode    string `json:"errorcode"`
	ErrorMessage string `json:"errormessage"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("api error: code %s, message %s", err.ErrorCode, err.ErrorMessage)
}

type LogSet struct {
	LogsetId   string `json:"logset_id"`
	LogSetName string `json:"logset_name"`
	Period     int    `json:"period"`
	CreateTime string `json:"create_time"`
}

type GetLogSetResponse LogSet

type ListLogSetsResponse struct {
	LogSets []LogSet `json:"logsets"`
}

type LogTopic struct {
	LogsetId     string       `json:"logset_id"`
	TopicId      string       `json:"topic_id"`
	TopicName    string       `json:"topic_name"`
	Path         string       `json:"path"`
	Collection   bool         `json:"collection"`
	Index        bool         `json:"index"`
	LogType      string       `json:"log_type"`
	ExtractRule  ExtractRule  `json:"extract_rule"`
	MachineGroup MachineGroup `json:"machine_group"`
	CreateTime   string       `json:"create_time"`
}

type ExtractRule struct {
	TimeKey     string   `json:"time_key"`
	TimeFormat  string   `json:"time_format"`
	Delimiter   string   `json:"delimiter"`
	Keys        []string `json:"keys"`
	FilterKeys  []string `json:"filter_keys"`
	FilterRegex []string `json:"filter_regex"`
}

type MachineGroup struct {
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
}

type GetLogTopicResponse LogTopic

type ListLogTopicsResponse struct {
	Topics []LogTopic `json:"topics"`
}

func (api *ApiClient) GetLogSet(logsetId string) (*LogSet, error) {
	reqUrl := url.URL{}
	reqUrl.Scheme = "http"
	reqUrl.Host = api.host
	reqUrl.Path = "/logset"
	query := reqUrl.Query()
	query.Set("logset_id", logsetId)
	reqUrl.RawQuery = query.Encode()
	request, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	request.URL.Query().Set("logset_id", logsetId)

	response := GetLogSetResponse{}
	err = api.do(request, &response)
	if err != nil {
		return nil, err
	}

	logset := LogSet(response)

	return &logset, nil
}

func (api *ApiClient) ListLogSets() (*ListLogSetsResponse, error) {
	reqUrl := url.URL{}
	reqUrl.Scheme = "http"
	reqUrl.Host = api.host
	reqUrl.Path = "/logsets"
	query := reqUrl.Query()
	reqUrl.RawQuery = query.Encode()
	request, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	response := ListLogSetsResponse{}
	err = api.do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *ApiClient) GetLogTopic(topicId string) (*LogTopic, error) {
	reqUrl := url.URL{}
	reqUrl.Scheme = "http"
	reqUrl.Host = api.host
	reqUrl.Path = "/topic"
	query := reqUrl.Query()
	query.Set("topic_id", topicId)
	reqUrl.RawQuery = query.Encode()
	request, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	response := GetLogTopicResponse{}
	err = api.do(request, &response)
	if err != nil {
		return nil, err
	}

	topic := LogTopic(response)

	return &topic, nil
}

func (api *ApiClient) ListLogTopics(logsetId string) (*ListLogTopicsResponse, error) {
	reqUrl := url.URL{}
	reqUrl.Scheme = "http"
	reqUrl.Host = api.host
	reqUrl.Path = "/topics"
	query := reqUrl.Query()
	query.Set("logset_id", logsetId)
	reqUrl.RawQuery = query.Encode()
	request, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	request.URL.Query().Set("logset_id", logsetId)

	response := ListLogTopicsResponse{}
	err = api.do(request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (api *ApiClient) do(req *http.Request, response interface{}) error {
	api.sign(req)

	resp, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if 400 <= resp.StatusCode && resp.StatusCode <= 500 {
		errResponse := &ErrorResponse{}
		//dec := json.NewDecoder(resp.Body)
		//if err := dec.Decode(errResponse); err != nil {
		if err := json.Unmarshal(rawBody, errResponse); err != nil {
			return err
		}
		return errResponse
	}

	//dec := json.NewDecoder(resp.Body)
	//if err := dec.Decode(response); err != nil {
	if err := json.Unmarshal(rawBody, response); err != nil {
		return err
	}

	return nil
}

func (api *ApiClient) sign(req *http.Request) error {

	var queryKey []string
	queryC := make(map[string][]string)
	queryCopy := url.Values(queryC)
	for k, v := range req.URL.Query() {
		if len(v) <= 0 {
			continue
		}
		queryKey = append(queryKey, strings.ToLower(k))
		queryCopy.Set(strings.ToLower(k), v[0])
	}
	var headerKey []string
	headerC := make(map[string][]string)
	headerCopy := url.Values(headerC)
	for k, v := range req.Header {
		if len(v) <= 0 {
			continue
		}
		headerKey = append(headerKey, strings.ToLower(k))
		headerCopy.Set(strings.ToLower(k), v[0])
	}
	headerCopy.Set("host", api.host)
	headerKey = append(headerKey, "host")

	lowerMethod := strings.ToLower(req.Method)

	httpRequestInfo := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n",
		lowerMethod,
		req.URL.Path,
		queryCopy.Encode(),
		headerCopy.Encode(),
	)

	httpRequestInfoSha1 := sha1.New()
	httpRequestInfoSha1.Write([]byte(httpRequestInfo))
	httpRequestInfoSha1String := hex.EncodeToString(httpRequestInfoSha1.Sum(nil))

	now := time.Now()

	qKeyTime := fmt.Sprintf("%d;%d", now.Unix(), now.Add(time.Second * 3600).Unix())
	stringToSign := fmt.Sprintf("%s\n%s\n%s\n", QSignAlgorithm, qKeyTime, httpRequestInfoSha1String)

	signKeyHmac := hmac.New(sha1.New, []byte(api.credential.SecretKey))
	signKeyHmac.Write([]byte(qKeyTime))
	signKey := hex.EncodeToString(signKeyHmac.Sum(nil))

	signatureHmac := hmac.New(sha1.New, []byte(signKey))
	signatureHmac.Write([]byte(stringToSign))
	signature := hex.EncodeToString(signatureHmac.Sum(nil))

	authorizationHeaderMap := map[string]string{
		"q-sign-algorithm": "sha1",
		"q-ak":             api.credential.SecretId,
		"q-sign-time":      qKeyTime,
		"q-key-time":       qKeyTime,
		"q-signature":      signature,
		"q-header-list":    strings.Join(headerKey, ";"),
		"q-url-param-list": strings.Join(queryKey, ";"),
	}

	var authorizationHeader []string

	for k, v := range authorizationHeaderMap {
		authorizationHeader = append(authorizationHeader, fmt.Sprintf("%s=%s", k, v))
	}

	req.Header.Set("Authorization", strings.Join(authorizationHeader, "&"))
	req.Header.Set("x-cls-Token", api.credential.Token)

	return nil
}
