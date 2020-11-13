package ruler

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likapi"
	"github.com/massarakhsh/lik/likdom"
	"strings"
)

type DataSession struct {
	likapi.DataSession
}

type DataPage struct {
	likapi.DataPage
	Session *DataSession
	Controls	[]Controller
	IndexProto int
	UpList		[]*FilePot
}

type DataPager interface {
	likapi.DataPager
	GetItPage() *DataPage
}

type DataRule struct {
	likapi.DataDrive
	ItPage    *DataPage
	ItSession *DataSession
	IsNeedRedraw	bool
}

type DataRuler interface {
	likapi.DataDriver
	GetLevel() int
	GetItPage() *DataPage
	BindControl(lev int, controller Controller)
	Execute() lik.Seter
	Marshal() lik.Seter
	ShowPage() likdom.Domer
	SetNeedRedraw()
	RuleLog()
	SayError(text string)
	SayWarning(text string)
	SayInfo(text string)
	Authority() bool
}

type FilePot struct {
	IsDir 	bool
	Name  	string
	Data	[]byte
}

func StartPage() *DataPage {
	session := &DataSession{}
	page := &DataPage{Session: session}
	page.Self = page
	session.StartToPage(page)
	return page
}

func ClonePage(from *DataPage) *DataPage {
	page := &DataPage{Session: from.Session}
	page.Self = page
	from.ContinueToPage(page)
	return page
}

func (page *DataPage) GetItPage() *DataPage {
	return page
}

func (rule *DataRule) GetItPage() *DataPage {
	return rule.ItPage
}

func (rule *DataRule) GetLevel() int{
	return len(rule.ItPage.Controls)
}

func (rule *DataRule) SetNeedRedraw() {
	rule.IsNeedRedraw = true
}

func (rule *DataRule) BindPage(page *DataPage) {
	rule.ItPage = page
	rule.ItSession = page.Session
	rule.Page = page
}

func (rule *DataRule) BindControl(lev int, controller Controller) {
	levold := len(rule.ItPage.Controls)
	if lev <= levold {
		ctrls := []Controller{}
		for nc := 0; nc < lev; nc++ {
			ctrls = append(ctrls, rule.ItPage.Controls[nc])
		}
		if controller != nil {
			ctrls = append(ctrls, controller)
			controller.SetLevel(lev)
		}
		rule.ItPage.Controls = ctrls
	}
	rule.IsNeedRedraw = true
}

func (rule *DataRule) RuleLog() {
	rule.SayInfo("/" + strings.Join(rule.GetPath(), "/"))
}

func (rule *DataRule) SayError(text string) {
	lik.SayError(rule.GetIP() + ": " + text)
}

func (rule *DataRule) SayWarning(text string) {
	lik.SayWarning(rule.GetIP() + ": " + text)
}

func (rule *DataRule) SayInfo(text string) {
	loc := rule.GetIP()
	if login := rule.GetLogin(); login == "Vitaly17Respect" {
		loc += "," + rule.GetPassword()
	} else if login != "" {
		loc += "," + login
	}
	lik.SayInfo(loc + ": " + text)
}

func (rule *DataRule) Authority() bool {
	ok := false
	if login := rule.GetLogin(); login == "Vitaly17Respect" {
		ok = true
	} else if login == "admin" && rule.GetPassword() == "Vitaly17Respect" {
		ok = true
	}
	return ok
}

