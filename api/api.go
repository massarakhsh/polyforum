package api

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/generate"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
	"time"
	"crypto/md5"
	"encoding/base64"
)

const MAX_SEARCH = 999

type DataRule struct {
	ruler.DataRule
}

type DataRuler interface {
	ruler.DataRuler
}

type SendSMSData struct {
	At		time.Time
	Phone	string
	Code	string
}

var (
	smsSync		sync.Mutex
	smsMap		map[string]SendSMSData
)

func BuildRule(page *ruler.DataPage) *DataRule {
	rule := &DataRule{ }
	rule.BindPage(page)
	return rule
}

func (rule *DataRule) Execute() lik.Seter {
	rule.SeekPageSize()
	rule.execute()
	responce := rule.GetAllResponse()
	if responce == nil || responce.Count() == 0 {
		responce = lik.BuildSet("disagnosis=error")
	}
	return responce
}

func (rule *DataRule) Marshal() lik.Seter {
	return nil
}

func (rule *DataRule) ShowPage() likdom.Domer {
	return nil
}

func (rule *DataRule) execute() {
	if rule.IsShift("api") {
		rule.execute()
	} else if rule.IsShift("get") {
		rule.executeGet()
	} else if rule.IsShift("list") {
		rule.executeList()
	} else if rule.IsShift("update") {
		rule.executeUpdate()
	} else if rule.IsShift("insert") {
		rule.executeInsert()
	} else if rule.IsShift("delete") {
		rule.executeDelete()
	} else if rule.IsShift("lastid") {
		rule.executeLastId()
	} else if rule.IsShift("search") {
		rule.executeSearch()
	} else if rule.IsShift("searchforum") {
		rule.executeSearchForum()
	} else if rule.IsShift("ean13") {
		rule.executeEan13()
	} else if rule.IsShift("qrcode") {
		rule.executeQRCode()
	} else if rule.IsShift("sendsmscode") {
		rule.executeSendSMS()
	} else if rule.IsShift("createperson") {
		rule.executeCreatePerson()
	} else if rule.IsShift("seekperson") {
		rule.executeSeekPerson()
	} else if rule.IsShift("probesmscode") {
		rule.executeProbeSMS()
	}
}

func (rule *DataRule) executeGet() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		id := lik.StrToInt(rule.Shift())
		if obj := base.GetElm(part, lik.IDB(id)); obj != nil {
			rule.SetResponse(obj, part)
		} else {
			rule.SetResponse(fmt.Sprintf("Object %s/%d not found", part, id), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeList() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		if list := base.GetList(part, "Id"); list != nil {
			rule.SetResponse(list, part)
			if list != nil {
				rule.SayInfo(fmt.Sprintf("List count: %d", list.Count()))
			}
		} else {
			rule.SetResponse(fmt.Sprintf("List %s not found", part), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeUpdate() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		id := lik.StrToIDB(rule.Shift())
		if sets := rule.getDataSet(part); sets != nil {
			if base.UpdateElm(part, id, sets) {
				if obj := base.GetElm(part, id); obj != nil {
					rule.SetResponse(obj, part)
				} else {
					rule.SetResponse(fmt.Sprintf("Object %s/%d not found", part, id), "diagnosis")
				}
			} else {
				rule.SetResponse(fmt.Sprintf("Object %s.%d not updated", part, int(id)), "diagnosis")
			}
		} else {
			rule.SetResponse(fmt.Sprintf("Data %s not presented", part), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeInsert() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		if sets := rule.getDataSet(part); sets != nil {
			if obj := base.GetElm(part, base.InsertElm(part, sets)); obj != nil {
				rule.SetResponse(obj, part)
			} else {
				rule.SetResponse(fmt.Sprintf("Object %s not created", part), "diagnosis")
			}
		} else {
			rule.SetResponse(fmt.Sprintf("Data %s not presented", part), "diagnosis")
		}
	}
}

func (rule *DataRule) executeDelete() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		id := lik.StrToIDB(rule.Shift())
		if obj := base.GetElm(part, id); obj != nil {
			base.DeleteElm(part, id)
			rule.SetResponse("OK", part)
		} else {
			rule.SetResponse(fmt.Sprintf("Object %s/%d not found", part, id), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeLastId() {
	if part := rule.Shift(); base.GetTable(part) != nil {
		id := base.GetLastId(part)
		rule.SetResponse(id, "id")
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeSearch() {
	part := rule.Shift()
	if table := base.GetTable(part); table != nil {
		field := rule.GetContext("field")
		field = unBase64(field)
		sort := rule.GetContext("sort")
		sort = unBase64(sort)
		if list := base.GetList(part, sort); list != nil && list.Count() > 0 {
			code := rule.getDataText()
			code = unBase64(code)
			found := rule.searchList(table, list, 0, field, code)
			rule.SetResponse(found, part)
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeSearchForum() {
	part := rule.Shift()
	if table := base.GetTable(part); table != nil {
		id := lik.StrToIDB(rule.Shift())
		field := rule.GetContext("field")
		field = unBase64(field)
		sort := rule.GetContext("sort")
		sort = unBase64(sort)
		if list := base.GetList(part, sort); list != nil && list.Count() > 0 {
			code := rule.getDataText()
			lik.SayInfo(code);
			code = unBase64(code)
			lik.SayInfo(code);
			found := rule.searchList(table, list, id, field, code)
			rule.SetResponse(found, part)
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Table %s not exists", part), "diagnosis")
	}
}

func (rule *DataRule) executeEan13() {
	code := rule.Shift();
	if code == "" {
		code = rule.GetContext("data")
	}
	if code != "" {
		path := generate.GenerateEan13(code)
		if path != "" {
			rule.SetResponse(path, "url")
		} else {
			rule.SetResponse(fmt.Sprintf("Code not generated"), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Code not requested"), "diagnosis")
	}
}

func (rule *DataRule) executeQRCode() {
	code := rule.Shift();
	if code == "" {
		code = rule.GetContext("data")
	}
	if code != "" {
		path := generate.GenerateQR(code)
		if path != "" {
			rule.SetResponse(path, "url")
		} else {
			rule.SetResponse(fmt.Sprintf("Code not generated"), "diagnosis")
		}
	} else {
		rule.SetResponse(fmt.Sprintf("Code not requested"), "diagnosis")
	}
}

func (rule *DataRule) executeSendSMS() {
	if phone := rule.normalizePhone(rule.Shift()); phone =="" {
		rule.SetResponse(fmt.Sprintf("PhoneNumber not presented"), "diagnosis")
	} else {
		code := ""
		at := time.Now()
		smsSync.Lock()
		if smsMap == nil {
			smsMap = make(map[string]SendSMSData)
		}
		if sm,ok := smsMap[phone]; ok && sm.At.Sub(at) < 5 * time.Minute {
			code = sm.Code
		} else {
			code = fmt.Sprintf("%04d", rand.Intn(10000))
			smsMap[phone] = SendSMSData{ At: at, Phone: phone, Code: code }
		}
		smsSync.Unlock()
		if base.SendSMS(phone, code) {
			rule.SetResponse(true, "success")
		} else {
			rule.SetResponse("SMS not sent", "diagnosis")
		}
	}
}

func (rule *DataRule) executeProbeSMS() {
	if phone := rule.normalizePhone(rule.Shift()); phone =="" {
		rule.SetResponse(fmt.Sprintf("PhoneNumber not presented"), "diagnosis")
	} else if code := rule.GetContext("data"); len(code) != 4 {
		rule.SetResponse(fmt.Sprintf("Code not presented"), "diagnosis")
	} else {
		right := false
		at := time.Now()
		smsSync.Lock()
		if smsMap != nil {
			if sm, ok := smsMap[phone]; ok && sm.At.Sub(at) < 10*time.Minute {
				if code == sm.Code {
					right = true
				}
			}
		}
		smsSync.Unlock()
		if right {
			rule.SetResponse(true, "success")
		} else {
			rule.SetResponse("Codes difference", "diagnosis")
		}
	}
}

func (rule *DataRule) executeCreatePerson() {
	if phone := rule.normalizePhone(rule.Shift()); phone =="" {
		rule.SetResponse(fmt.Sprintf("PhoneNumber not presented"), "diagnosis")
	} else {
		var id lik.IDB
		if list := base.GetList("Person", "Id"); list != nil {
			for np := 0; np < list.Count(); np++ {
				if pers := list.GetSet(np); pers != nil {
					if pers.GetString("Phone") == phone {
						id = pers.GetIDB("Id")
						break
					}
				}
			}
		}
		sets := rule.getDataSet("Person")
		if id == 0 {
			id = base.InsertElm("Person", sets)
		} else if sets != nil {
			base.UpdateElm("Person", id, sets)
		}
		if person := base.GetElm("Person", id); person != nil {
			regs := lik.BuildSet()
			apps := person.GetString("AppId")
			appid := rule.GetPassword()
			if appid != "" && !strings.Contains(apps, appid) {
				if apps != "" { apps += "," }
				apps += appid
				regs.SetItem(apps, "AppId")
			}
			regs.SetItem(phone, "Phone")
			login :=""
			if sets != nil {
				login = sets.GetString("Login")
			}
			if login == "" {
				login = lik.IntToStr(10000 + rand.Intn(90000))
			}
			regs.SetItem(login, "Login")
			password := lik.IntToStr(10000 + rand.Intn(90000))
			h := md5.New()
			io.WriteString(h, password)
			bytes := h.Sum(nil)
			pass5 := base64.StdEncoding.EncodeToString(bytes)
			regs.SetItem(pass5, "PW5")
			base.UpdateElm("Person", person.GetIDB("Id"), regs)
			rule.SetResponse(person.GetIDB("Id"), "IdPerson")
			rule.SetResponse(login, "Login")
			rule.SetResponse(password, "Password")
		} else {
			rule.SetResponse(fmt.Sprintf("Object %s not created", "Person"), "diagnosis")
		}
	}
}

func (rule *DataRule) executeSeekPerson() {
	if login := rule.GetContext("login"); login =="" {
		rule.SetResponse(fmt.Sprintf("Login not presented"), "diagnosis")
	} else {
		login = unBase64(login)
		password := rule.GetContext("password")
		password = unBase64(password)
		h := md5.New()
		io.WriteString(h, password)
		bytes := h.Sum(nil)
		pass5 := base64.StdEncoding.EncodeToString(bytes)
		var id lik.IDB
		if list := base.GetList("Person", "Id"); list != nil {
			for np := 0; np < list.Count(); np++ {
				if pers := list.GetSet(np); pers != nil {
					if pers.GetString("Login") == login &&
						pers.GetString("PW5") == pass5 {
						id = pers.GetIDB("Id")
						break
					}
				}
			}
		}
		if id == 0 {
			rule.SetResponse("Person not found", "diagnosis")
		} else if person := base.GetElm("Person", id); person != nil {
			rule.SetResponse(person, "Person")
		} else {
			rule.SetResponse("Person not found", "diagnosis")
		}
	}
}

func (rule *DataRule) getDataText() string {
	code := rule.Shift();
	if code == "" {
		code = rule.GetContext("data")
	}
	//code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, "\n", "", -1)
	code = strings.Replace(code, "\r", "", -1)
	return code
}

func (rule *DataRule) getDataSet(part string) lik.Seter {
	var sets lik.Seter
	if table := base.GetTable(part); table != nil {
		if data := rule.GetAllContext(); data != nil {
			if set := data.GetSet("_json"); set != nil {
				if tet := set.GetSet(part); tet != nil {
					set = tet
				}
				sets = lik.BuildSet()
				for _, pot := range set.Values() {
					if !strings.HasPrefix(pot.Key, "_") {
						var field *base.FieldBase
						for _, fld := range table.Fields {
							if pot.Key == fld.Key {
								field = &fld
								break
							}
						}
						if field != nil {
							sets.SetItem(pot.Val, pot.Key)
						} else {
							rule.SetResponse(fmt.Sprintf("Key %s.%s not found", part, pot.Key), "diagnosis")
							return nil
						}
					}
				}
			}
		}
	}
	return sets
}

func (rule *DataRule) normalizePhone(phone string) string {
	pho := ""
	for s := 0; s < len(phone); s++ {
		if ch := phone[s]; ch >= '0' && ch <= '9' {
			pho += string(ch)
		}
	}
	if len(pho) == 11 && (pho[0] == '7' || pho[0] == '8') {
		pho = pho[1:]
	}
	return pho
}

func (rule *DataRule) searchList(table *base.TableBase, list lik.Lister, idforum lik.IDB, field string, code string) lik.Lister {
	var found lik.Lister
	if code == "" {
		found = list
	} else {
		found = lik.BuildList()
		icode := lik.StrToInt(code)
		for ne := 0; ne < list.Count(); ne++ {
			elm := list.GetSet(ne)
			if idforum > 0 {
				if table.Part == "Exhibition" && elm.GetIDB("Id") != idforum {
					continue
				} else if table.Part == "Action" && elm.GetIDB("ExhibitionId") != idforum {
					continue
				} else if table.Part == "Zone"  && elm.GetIDB("ExhibitionId") != idforum {
					continue
				} else if table.Part == "Request" && elm.GetIDB("ExhibitionId") != idforum {
					continue
				} else if table.Part == "Speaker" && elm.GetIDB("ExhibitionId") != idforum {
					continue
				} else if table.Part == "Program" && elm.GetIDB("ExhibitionId") != idforum {
					continue
				} else if table.Part == "Exhibitor" && elm.GetIDB("ExhibitionId") != idforum {
					continue
				}
			}
			for _, fld := range table.Fields {
				key := fld.Key
				proto := fld.Proto
				if key == field || field == "" {
					if strings.Contains(proto, "S") {
						if val := elm.GetString(key); val != "" {
							if strings.Contains(val, code) {
								found.AddItems(elm)
								break
							} else if strings.Contains(strings.ToLower(val), strings.ToLower(code)) {
								found.AddItems(elm)
								break
							}
						}
					} else if strings.Contains(proto, "L") {
						if val := elm.GetInt(key); val != 0 {
							if val == icode {
								found.AddItems(elm)
								break
							}
						}
					}
				}
			}
		}
	}
	return found
}

func unBase64(data string) string {
	data = strings.Replace(data, "@", "=", -1)
	if bts,err := base64.StdEncoding.DecodeString(data); err == nil {
		data = string(bts)
	}
	return data
}
