package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTP GET
// url 请求地址
// data 请求成功返回的数据, 为指针
// errData 请求错误返回的数据, 为指针
func HttpGetJson(url string, data, errData interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("json http get url: %s err: %v", url, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("json http get url: %s body_err: %v", url, err)
	}

	// 只要响应不是 200
	if http.StatusOK != resp.StatusCode {

		if errData != nil {
			err = json.Unmarshal(body, errData)
			if err != nil {
				return fmt.Errorf("json http get data_err: %s json_err: %v", string(body), err)
			}
		}

		return fmt.Errorf("json http get err: %s http_code: %d", string(body), resp.StatusCode)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("json http get data: %s json_err: %v", string(body), err)
	}

	return nil
}

// HTTP GET
// url 请求地址
// data 请求成功返回的数据, 为指针
func HttpGetJsons(url string, data interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("json http get url: %s err: %v", url, err)
	}

	defer resp.Body.Close()

	// 只要响应不是 200
	if http.StatusOK != resp.StatusCode {
		return fmt.Errorf("json http get err: http_code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("json http get url: %s body_err: %v", url, err)
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return fmt.Errorf("json http get data: %s json_err: %v", string(body), err)
	}

	return nil
}
