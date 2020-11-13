package base

import (
	"github.com/massarakhsh/lik"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const AmsApiId = "1DDB039C-175B-AA16-6240-EAA5BB6A71C8"

func SendSMS(phone string, code string) bool {
	lik.SayInfo("Send SMS " + phone + ": " + code)
	ok := false
	if len(phone) == 10 {
		ok = true
	} else if len(phone) == 11 && phone[0] == '7' {
		phone = phone[1:]
		ok = true
	} else if len(phone) == 11 && phone[0] == '8' {
		phone = phone[1:]
		ok = true
	}
	if ok {
		sphone := "7" + phone
		scode := "PolyForum+" + code
		uri := fmt.Sprintf("https://sms.ru/sms/send?api_id=%s&to=%s&msg=%s&json=0", AmsApiId, sphone, scode)
		answer := httpRequestGet(uri)
		ok = strings.HasPrefix(answer, "100")
	}
	return ok
}

func httpRequestGet(uri string) string {
	answer := ""
	if resp, err := http.Get(uri); err == nil {
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			answer = string(body)
		}
	}
	return answer
}

