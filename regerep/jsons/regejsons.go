package jsons

import (
	"fmt"
	"github.com/typenameman/gotools/unicodes"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func ShowJsonContentsWithUnicode(url string, method string, header map[string]string, payload io.Reader, tag string) []string {
	var resu []string
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		panic(err)
	}
	for a, b := range header {
		req.Header.Add(a, b)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//tags1
	exp := "\"" + tag + "\":(\u0020){0,1}\"(.*?)\"(,|\n}|})"
	rub1 := fmt.Sprintf("\"%s\": ", tag)
	cont := regesubmatch(string(body), exp, rub1, "\"", " ", ",", "\u007d", "\n")
	for _, s := range cont {
		s = unicodeToString(s)
		resu = append(resu, s)
	}
	//tags2

	return resu
}
func ShowJsonContents(url string, method string, header map[string]string, payload io.Reader, tag string) []string {
	var resu []string
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		panic(err)
	}
	for a, b := range header {
		req.Header.Add(a, b)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//tags1
	exp := "\"" + tag + "\":(\u0020){0,1}\"(.*?)\"(,|\n}|})"
	rub1 := fmt.Sprintf("\"%s\":", tag)
	cont := regesubmatch(string(body), exp, rub1, "\"", " ", ",", "\u007d", "\n")
	resu = append(resu, cont...)
	//tags2

	return resu
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
func unicodeToString(unictext string) string {
	res := regesubmatch(unictext, "\\\\u[0-9||a-z]{4}")
	for _, i := range res {
		unictext = strings.ReplaceAll(unictext, i, unicodes.UnicodeDecodeToStrings(i))
	}
	return unictext
}
