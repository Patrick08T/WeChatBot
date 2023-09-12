package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func StructToQueryString(data interface{}) (string, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("Input data is not a struct")
	}

	query := url.Values{}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" {
			continue // 跳过没有 "url" 标签的字段
		}

		fieldValue := v.Field(i)

		// 将字段值格式化为字符串并添加到查询参数
		query.Add(tag, fmt.Sprintf("%v", fieldValue.Interface()))
	}

	return query.Encode(), nil
}

func GetWithHeaderAndBody(url string, msg []byte, headers map[string]string, bodys map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(string(msg)))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	q := req.URL.Query()
	for k, v := range bodys {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
