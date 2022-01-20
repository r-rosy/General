package html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var unicodelist string = "\u4e00-\u9fa5||0-9||a-z||A-Z||\u0020\u3000\u00A0。；，：“”（）||\u002f||、？《》|!:,\"?~℃%+.()\ufe0f\u002d"

// decodelist-explain 中文汉字  三种空格 页面编写人员未知符号 \
func ShowTags(content string) []string {
	var res []string
	exp1 := fmt.Sprintf("<.{1,}?>[%s]{1,}?<.{1,}?>", unicodelist)
	exp2 := fmt.Sprintf(">[%s]{1,}?<", unicodelist)
	reg1, _ := regexp.Compile(exp1)
	reg2, _ := regexp.Compile(exp2)
	y := reg1.FindAllStringSubmatch(content, -1)
	for _, b := range y {
		for _, a := range b {
			str := reg2.FindAllStringSubmatch(a, -1)
			for _, p := range str {
				for i, _ := range p {
					p[i] = clearrub(p[i], ">")
					p[i] = clearrub(p[i], "<")
					res = append(res, p[i])
				}
			}
		}
	}
	return res
}
func RequestAndShowTags(url string, cookie string, Header map[string]string) []string {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	for a, b := range Header {
		req.Header.Add(a, b)
	}
	if cookie != "" {
		req.Header.Add("Cookie", cookie)
	}
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	return ShowTags(string(body))

}
func clearrub(str string, rub string) string {
	return strings.ReplaceAll(str, rub, "")
}
func regesubmatch(content string, exp string, rub ...string) []string {
	reg, err := regexp.Compile(exp)
	if err != nil {
		panic(err)
	}
	var res []string
	y := reg.FindAllString(content, -1)
	for i, _ := range y {
		if len(rub) != 0 {
			for a, _ := range rub {
				y[i] = clearrub(y[i], rub[a])
			}
		}
		res = append(res, y[i])
	}
	return res
}
