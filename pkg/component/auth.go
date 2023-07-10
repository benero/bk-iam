/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package component

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TencentBlueKing/gopkg/conv"
	"github.com/TencentBlueKing/gopkg/errorx"
	"github.com/parnurzeal/gorequest"

	"github.com/TencentBlueKing/bk-iam/pkg/logging"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock

// AuthResponse is the struct of iam backend response
type AuthResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// Error will check if the response with error
func (r *AuthResponse) Error() error {
	if r.Code == 0 {
		return nil
	}

	return fmt.Errorf("response error[code=`%d`,  message=`%s`]", r.Code, r.Message)
}

// String will return the detail text of the response
func (r *AuthResponse) String() string {
	return fmt.Sprintf("response[code=`%d`, message=`%s`, data=`%v`]", r.Code, r.Message, r.Data)
}

// AuthClient is the interface of auth client
type AuthClient interface {
	Verify(appCode, appSecret string) (bool, error)
}

type authClient struct {
	Host string

	// iam's app_code/app_secret, credentials for bkauth
	bkAppCode   string
	bkAppSecret string
}

// NewAuthClient will create a auth client
func NewAuthClient(host string, bkAppCode string, bkAppSecret string) AuthClient {
	host = strings.TrimRight(host, "/")
	return &authClient{
		Host:        host,
		bkAppCode:   bkAppCode,
		bkAppSecret: bkAppSecret,
	}
}

func (c *authClient) call(
	method Method,
	path string,
	data interface{},
	timeout int64,
) (map[string]interface{}, error) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf("component", "authClient.call")

	callTimeout := time.Duration(timeout) * time.Second
	if timeout == 0 {
		callTimeout = defaultTimeout
	}

	url := fmt.Sprintf("%s%s", c.Host, path)
	result := AuthResponse{}
	start := time.Now()
	callbackFunc := NewMetricCallback("Auth", start)

	request := gorequest.New()
	switch method {
	case POST:
		request = request.Post(url)
	case GET:
		request = request.Get(url)
	}
	request = request.Timeout(callTimeout).Type("json")

	// set headers
	request.Header.Set("X-BK-APP-CODE", c.bkAppCode)
	request.Header.Set("X-BK-APP-SECRET", c.bkAppSecret)

	// do request
	resp, respBody, errs := request.
		Send(data).
		EndStruct(&result, callbackFunc)

	// NOTE: it's a sensitive api, so, no log request detail!
	// logFailHTTPRequest(start, request, resp, respBody, errs, &result)
	logger := logging.GetComponentLogger()

	var err error
	if len(errs) != 0 {
		// 敏感信息泄漏 ip+端口号, 替换为 *.*.*.*
		errsMessage := fmt.Sprintf("gorequest errorx=`%s`", errs)
		errsMessage = ipRegex.ReplaceAllString(errsMessage, replaceToIP)
		err = errors.New(errsMessage)

		err = errorWrapf(err, "errsCount=`%d`", len(errs))
		logger.Errorf("call auth api %s fail, err=%s", path, err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("gorequest statusCode is %d not 200, respBody=%s",
			resp.StatusCode, conv.BytesToString(respBody))
		logger.Errorf("call auth api %s fail , err=%s", path, err.Error())
		return nil, errorWrapf(err, "status=%d", resp.StatusCode)
	}
	if result.Code != 0 {
		err = errors.New(result.Message)
		err = errorWrapf(err, "result.Code=%d", result.Code)
		logger.Errorf("call auth api %s ok but code in response is not 0, respBody=%s, err=%s",
			path, conv.BytesToString(respBody), err.Error())
		return nil, err
	}

	return result.Data, nil
}

// Verify will check bkAppCode, bkAppSecret is valid
func (c *authClient) Verify(appCode, appSecret string) (bool, error) {
	errorWrapf := errorx.NewLayerFunctionErrorWrapf("component", "authClient.Verify")

	path := fmt.Sprintf("/api/v1/apps/%s/access-keys/verify", appCode)

	data, err := c.call(POST, path, map[string]interface{}{
		"bk_app_secret": appSecret,
	}, 5)
	if err != nil {
		err = errorWrapf(err, "verify app_code=`%s` fail", appCode)

		return false, err
	}
	matchI, ok := data["is_match"]
	if !ok {
		return false, errors.New("no is_match in response body")
	}

	match, ok := matchI.(bool)
	if !ok {
		return false, errors.New("is_match is not a valid bool")
	}
	return match, nil
}
