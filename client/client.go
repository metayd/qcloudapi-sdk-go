package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/dbdd4us/qcloudapi-sdk-go/types"
)

type Credential struct {
	SecretId  []byte
	SecretKey []byte
}

type Opts struct {
	Region string
}

type Client struct {
	client *http.Client

	region    string
	secretId  []byte
	secretKey []byte
}

func NewCredential(secretId []byte, secretKey []byte) Credential {
	return Credential{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
}

func NewClient(credential Credential, opts Opts) (*Client, error) {
	return &Client{
		client:    &http.Client{},
		region:    opts.Region,
		secretId:  credential.SecretId,
		secretKey: credential.SecretKey,
	}, nil
}

func (cli *Client) Invoke(host string, action string, args interface{}, response interface{}) error {
	commonArgs := types.CommonArgs{}
	commonArgs.Action = action
	commonArgs.Region = cli.region
	commonArgs.Timestamp = uint(time.Now().Unix())
	commonArgs.SecretId = string(cli.secretId)
	commonArgs.Nonce = uint(rand.Int())
	commonArgs.SignatureMethod = "HmacSHA256"
	commonArgValues, _ := query.Values(commonArgs)
	requestArgValues, _ := query.Values(args)
	signature := cli.sign("GET", host, &commonArgValues, &requestArgValues)
	requestArgValues.Set("Signature", string(signature))
	reqArgs := requestArgValues.Encode()
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s/v2/index.php?%s", host, reqArgs),
		nil,
	)
	if err != nil {
		return err
	}
	resp, err := cli.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	fmt.Println(resp.StatusCode)
	return nil
}

func (cli *Client) sign(method string, host string, commonArgs *url.Values, requestArgs *url.Values) types.RequestSignature {
	for k, _ := range *commonArgs {
		requestArgs.Set(k, commonArgs.Get(k))
	}
	keys := make([]string, 0)
	for k, _ := range *requestArgs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	kvs := make([]string, 0, len(keys))
	for _, k := range keys {
		kvs = append(kvs, fmt.Sprintf("%s=%s", k, requestArgs.Get(k)))
	}
	reqStr := strings.Join(kvs, "&")
	reqStr = fmt.Sprintf("%s%s/v2/index.php?%s", method, host, reqStr)
	mac := hmac.New(sha256.New, cli.secretKey)
	mac.Write([]byte(reqStr))
	signature := mac.Sum(nil)
	b64Encoded := base64.StdEncoding.EncodeToString(signature)
	return types.RequestSignature(b64Encoded)
}