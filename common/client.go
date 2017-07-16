package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	RequestMethodGET  = "GET"
	RequestMethodPOST = "POST"

	SignatureMethodHMacSha256 = "HmacSHA256"
)

type Client struct {
	*http.Client

	credential Credential
	opts       Opts
}

type Opts struct {
	Method          string
	Region          string
	Host            string
	Path            string
	SignatureMethod string
	Schema          string
}

type Credential struct {
	SecretId  string
	SecretKey string
}

func NewClient(credential Credential, opts Opts) (*Client, error) {
	if opts.Method == "" {
		opts.Method = RequestMethodGET
	}
	if opts.SignatureMethod == "" {
		opts.SignatureMethod = SignatureMethodHMacSha256
	}
	if opts.Schema == "" {
		opts.Schema = "https"
	}
	return &Client{
		&http.Client{},
		credential,
		opts,
	}, nil
}

func (client *Client) Invoke(action string, args interface{}, response interface{}) error {
	switch client.opts.Method {
	case "GET":
		return client.InvokeWithGET(action, args, response)
	default:
		return client.InvokeWithPOST(action, args, response)
	}
}

func (client *Client) initCommonArgs(args *url.Values) {
	args.Set("Region", client.opts.Region)
	args.Set("Timestamp", fmt.Sprint(uint(time.Now().Unix())))
	args.Set("SecretId", client.credential.SecretId)
	args.Set("Nonce", fmt.Sprint(uint(rand.Int())))
	args.Set("SignatureMethod", client.opts.SignatureMethod)
}

func (client *Client) signGetRequest(values *url.Values) string {
	keys := make([]string, 0, len(*values))
	for k := range *values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	kvs := make([]string, 0, len(keys))
	for _, k := range keys {
		kvs = append(kvs, fmt.Sprintf("%s=%s", k, values.Get(k)))
	}
	queryStr := strings.Join(kvs, "&")
	reqStr := fmt.Sprintf("GET%s%s?%s", client.opts.Host, client.opts.Path, queryStr)

	mac := hmac.New(sha256.New, []byte(client.credential.SecretKey))
	mac.Write([]byte(reqStr))
	signature := mac.Sum(nil)

	b64Encoded := base64.StdEncoding.EncodeToString(signature)

	return b64Encoded
}

func (client *Client) InvokeWithGET(action string, args interface{}, response interface{}) error {
	reqValues := url.Values{}
	err := EncodeStruct(args, &reqValues)
	if err != nil {
		return err
	}
	reqValues.Set("Action", action)
	client.initCommonArgs(&reqValues)
	signature := client.signGetRequest(&reqValues)
	reqValues.Set("Signature", signature)

	reqQuery := reqValues.Encode()

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s://%s%s?%s", client.opts.Schema, client.opts.Host, client.opts.Path, reqQuery),
		nil,
	)

	if err != nil {
		return makeClientError(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return makeClientError(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return makeClientError(err)
	}

	errorResponse := ErrorResponse{}

	if err = json.Unmarshal(body, &errorResponse); err != nil {
		return makeClientError(err)
	}

	if (errorResponse.Code != NoErr) || (errorResponse.CodeDesc != "" && errorResponse.CodeDesc != NoErrCodeDesc) {
		return Error{
			ErrorResponse: errorResponse,
			StatusCode:    resp.StatusCode,
		}
	}

	if err = json.Unmarshal(body, response); err != nil {
		return makeClientError(err)
	}

	return nil
}

func (client *Client) InvokeWithPOST(action string, args interface{}, response interface{}) error {
	return nil
}
