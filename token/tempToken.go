package token

import (
	"encoding/json"
	"github.com/dbdd4us/qcloudapi-sdk-go/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
	"strings"
	"errors"
	"bufio"
)

var tokenMutex = &sync.Mutex{}

const (
	ENV_TOKEN_SESSION_TOKEN = "BM_CREDENTIAL_SESSION_TOKEN"
	ENV_TOKEN_SECRET_ID     = "BM_CREDENTIAL_SECRET_ID"
	ENV_TOKEN_SECRET_KEY    = "BM_CREDENTIAL_SECRET_KEY"
	ENV_TOKEN_EXPIRE_TIME   = "BM_CREDENTIAL_EXPIRE_TIME"

	BM_CCS_NOMR_SERVER_URL = "http://127.0.0.1:9090/bmccs/tmpToken"
	BM_TMEP_TOKEN_PATH     = "/opt/ccs_agent/temptoken"
	BM_METADATA_PATH       = "/etc/cpminfo"
)

type TempTokenSource struct {
	Uin string     //用户帐号
	AppId string 
	Region string  

	SecretId  string //永久密钥
	SecretKey string
}

func SetTokenToEnv(token *Token) error {
	os.Setenv(ENV_TOKEN_SESSION_TOKEN, token.SessionToken)
	os.Setenv(ENV_TOKEN_SECRET_ID, token.SecretId)
	os.Setenv(ENV_TOKEN_SECRET_KEY, token.SecretKey)

	if false == token.Expiry.IsZero() {
		os.Setenv(ENV_TOKEN_EXPIRE_TIME, token.Expiry.Format(time.RFC3339Nano))
	}
	return nil
}

func GetTokenFromEnv() (*Token, error) {
	token := &Token{
		SecretId:     os.Getenv(ENV_TOKEN_SECRET_ID),
		SecretKey:    os.Getenv(ENV_TOKEN_SECRET_KEY),
		SessionToken: os.Getenv(ENV_TOKEN_SESSION_TOKEN),
	}
	expireTime := os.Getenv(ENV_TOKEN_EXPIRE_TIME)

	//环境变量未设置
	if token.SecretId == "" || token.SecretKey == "" || expireTime == "" {
		return nil, nil
	}

	expiry, err := time.Parse(time.RFC3339Nano, expireTime)
	if err != nil {
		return nil, err
	}
	token.Expiry = expiry

	return token, nil
}

type AppConfigProperties map[string]string

func ReadPropertiesFile(filename string) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func GetInstanceIdFromMetadata(metadataFile string)(string,error){
    appConfigMap, err := ReadPropertiesFile(metadataFile)
	if err != nil {
		return "", err
	}

	instanceId,ok := appConfigMap["instanceId"]
	if !ok {
		return "",errors.New("get instance from /etc/cpminfo failed")
	}

	return instanceId,nil 
}


func GetTokenFromFile(tokenFile string )(*Token, error) {
	appConfigMap, err := ReadPropertiesFile(tokenFile)
	if err != nil {
		return nil, err
	}

	var ok bool
	token := &Token{}

	if token.SecretId, ok = appConfigMap["secretId"]; !ok {
		return nil ,errors.New("secretId not found in temptoken file")
	}

	if token.SecretKey, ok = appConfigMap["secretKey"]; !ok {
		return nil ,errors.New("secretKey not found in temptoken file")
	}

	if token.SessionToken, ok = appConfigMap["sessionToken"]; !ok {
		return nil ,errors.New("sessionToken not found in temptoken file")
	}

	expireTime,ok :=  appConfigMap["expireTime"]
	if !ok {
		return nil ,errors.New("expireTime not found in temptoken file")
	}

	if token.Expiry, err = time.Parse(time.RFC3339Nano, expireTime); err != nil {
		return nil, err
	}

	return token,nil 
}

type CamResponse struct {
	Code int32       `json:"return_code"`
	Msg  string      `json:"msg"`
}
//先调用接口，接口调用成功后，再从文件中获取
func GetNewTokenFromNormService(uin,appId,region,instanceId string) (*Token, error) {
	v := url.Values{}
	v.Set("uin",uin)	
	v.Set("region",region)
	v.Set("instanceId",instanceId)
	v.Set("appId",appId)
	resp, err := http.PostForm(BM_CCS_NOMR_SERVER_URL,v)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rsp := CamResponse{}
	if err = json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}

	if  rsp.Code != 0 {
		return nil, errors.New(rsp.Msg)
	}

	return GetTokenFromFile(BM_TMEP_TOKEN_PATH)
}

func (tokenSource *TempTokenSource) Token() (*Token, error) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	
	//如果有永久secretId和secretKey
	if tokenSource.SecretId != "" && tokenSource.SecretKey != "" {
		token := &Token{
			SecretId:  tokenSource.SecretId,
			SecretKey: tokenSource.SecretKey,
		}
		return token, nil
	}
	//从环境变量中读取,如果能读到就用环境变量
	token, err := GetTokenFromEnv()
	if nil != token && false == token.expired() {
		return token, nil
	}

	//调用API获取,获取成功后放入环境变量中
	instanceId,err := GetInstanceIdFromMetadata(BM_METADATA_PATH)
	if err != nil {
		return nil ,err
	}
	
	if token, err = GetNewTokenFromNormService(tokenSource.Uin,tokenSource.AppId,tokenSource.Region,instanceId); err != nil {
		return nil, err
	}

	SetTokenToEnv(token)
	return token, nil
}

//实现credential接口
func (tokenSource *TempTokenSource) GetSecretId() (string, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.SecretId, nil
}

func (tokenSource *TempTokenSource) GetSecretKey() (string, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return "", err
	}
	return token.SecretKey, nil
}

func (tokenSource *TempTokenSource) Values() (common.CredentialValues, error) {
	values := common.CredentialValues{}
	token, err := tokenSource.Token()

	if err != nil {
		return values, err
	}

	if token.SessionToken != "" {
		values["Token"] = token.SessionToken
	}

	return values, nil
}
