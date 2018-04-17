package oschina

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

//HTTPPost post方式返回数据
func HTTPPost(url string, params interface{}, response interface{}) error {

	nv := Struct2Map(params)
	resp, err := http.PostForm(url, nv)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}

func Struct2Map(obj interface{}) url.Values {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = url.Values{}
	for i := 0; i < t.NumField(); i++ {

		name := t.Field(i).Tag.Get("org")
		if name == "" {
			name = t.Field(i).Name
		}

		var value string
		sv := v.Field(i)
		switch sv.Kind() {
		case reflect.String:
			value = sv.Interface().(string)
		case reflect.Int:
			value = strconv.FormatInt(int64(sv.Interface().(int)), 10)
		}
		if value == "" {
			continue
		}
		data.Add(name, value)
	}
	return data
}
