package controller

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
	"strings"
)

type ApiControl struct {
	DataControl
	DoApi		ApiCmd
	DoCmd		string
	DoTable		string
	DoId		string
	DoPhone		string
	DoCode		string
	DoPrm		string
	Answer		lik.Seter
}

type Apier interface {
	Controller
}

type ApiCmd struct {
	Cmd		string
	Title	string
	Params	string
}

var (
	ApiList		[]ApiCmd = []ApiCmd {
		{ "", "выберите команду", ""},
		{ "get", "получить объект", "TI"},
		{ "list", "получить список", "T"},
		{ "update", "изменить объект", "TIJ"},
		{ "insert", "вставить объект", "TJ"},
		{ "delete", "удалить объект", "TI"},
		{ "search", "поиск", "TS"},
		{ "searchforum", "поиск на форуме", "TIS"},
		{ "sendsmscode", "послать код в SMS", "P"},
		{ "probesmscode", "проверить код из SMS", "PS"},
		{ "grcode", "штрихкод", "S"},
		{ "ean13", "штрихкод", "S"},
	}
)

func BuildApi(rule ruler.DataRuler, level int) Apier {
	it := &ApiControl{ }
	rule.BindControl(level, it)
	return it
}

func (it *ApiControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.menuPrepare(rule, false)
	row := tbl.BuildTr()
	it.menuItemCmd(rule, row,"", "API", "seek")
	it.menuItemText(rule, row,"|")
	it.menuItemCmd(rule, row, "","Выполнить", "execute")
	it.menuTools(rule, row)
	return tbl
}

func (it *ApiControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	tbl := div.BuildTableClass("api")
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("api_title").BuildString("Команда")
		row.BuildTdClass("api_info").AppendItem(it.buildLineCmd(rule))
	}
	if strings.Contains(it.DoApi.Params, "T") {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Таблица")
			row.BuildTdClass("api_info").AppendItem(it.buildLineTable(rule))
		}
	}
	if strings.Contains(it.DoApi.Params, "I") {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("ID")
			row.BuildTdClass("api_info").AppendItem(it.buildLineId(rule))
		}
	}
	if strings.Contains(it.DoApi.Params, "P") {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Телефон")
			row.BuildTdClass("api_info").AppendItem(it.buildLinePhone(rule))
		}
	}
	if strings.Contains(it.DoApi.Params, "S") {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Строка")
			row.BuildTdClass("api_info").AppendItem(it.buildLineString(rule))
		}
	}
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("api_title").BuildString("Вызов")
		row.BuildTdClass("api_info").AppendItem(it.buildLineCall(rule))
	}
	if strings.Contains(it.DoApi.Params, "J") {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Параметры")
			row.BuildTdClass("api_info").AppendItem(it.buildLineParam(rule))
		}
	}
	if it.DoApi.Cmd != "" {
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Выполнение")
			row.BuildTdClass("api_info").AppendItem(it.buildLineExec(rule))
		}
	}
	if it.Answer != nil {
		tbl.BuildTrTd("colspan=2").BuildString("<hr>")
		if row := tbl.BuildTr(); row != nil {
			row.BuildTdClass("api_title").BuildString("Ответ")
			row.BuildTdClass("api_info").BuildItem("pre").BuildString(it.Answer.Format(""))
		}
		if url := it.Answer.GetString("url"); url != "" {
			if row := tbl.BuildTr(); row != nil {
				row.BuildTdClass("api_title").BuildString("Изображение")
				row.BuildTdClass("api_info").BuildUnpairItem("img", "src", url)
			}
		}
	}
	return div
}

func (it *ApiControl) buildLineCmd(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildItem("select", "class=api", "id=apicmd")
	line.SetAttr("onchange", it.buildProc("select_cmd"))
	for _,api := range ApiList {
		opt := line.BuildItem("option", "value", api.Cmd)
		opt.BuildString(api.Cmd + " (" + api.Title + ")")
		if api.Cmd == it.DoCmd {
			it.DoApi = api
			opt.SetAttr("selected")
		}
	}
	return line
}

func (it *ApiControl) buildLineTable(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildItem("select", "class=api", "id=apitable")
	line.SetAttr("onchange", it.buildProc("select_table"))
	line.BuildItem("option").BuildString("")
	for _, table := range base.ListTables {
		opt := line.BuildItem("option", "value", table.Part)
		opt.BuildString(fmt.Sprintf("%s (%s)", table.Part, table.Title))
		if table.Part == it.DoTable {
			opt.SetAttr("selected")
		}
	}
	return line
}

func (it *ApiControl) buildLineId(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildUnpairItem("input", "type=text", "class=api", "id=apiid")
	line.SetAttr("onkeypress", it.buildProc("select_id"))
	if it.DoId != "" {
		line.SetAttr("value", it.DoId)
	}
	return line
}

func (it *ApiControl) buildLinePhone(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildUnpairItem("input", "type=text", "class=api", "id=apiphone")
	line.SetAttr("onkeypress", it.buildProc("select_phone"))
	if it.DoPhone != "" {
		line.SetAttr("value", it.DoPhone)
	}
	return line
}

func (it *ApiControl) buildLineString(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildUnpairItem("input", "type=text", "class=api", "id=apistring")
	line.SetAttr("onkeypress", it.buildProc("select_string"))
	if it.DoPrm != "" {
		line.SetAttr("value", it.DoPrm)
	}
	return line
}

func (it *ApiControl) buildLineCall(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildItem("b", "id=apicall", "redraw=request_build")
	return line
}

func (it *ApiControl) buildLineParam(rule ruler.DataRuler) likdom.Domer {
	line := likdom.BuildItem("textarea", "class=api", "id=apiprm")
	line.SetAttr("onchange", it.buildProc("select_prm"))
	if it.DoPrm != "" {
		line.BuildString(it.DoPrm)
	}
	return line
}

func (it *ApiControl) buildLineExec(rule ruler.DataRuler) likdom.Domer {
	cmd := likdom.BuildItemClass("a", "api button", "href=#")
	cmd.SetAttr("onclick", it.buildProc("call"))
	cmd.BuildString("Старт")
	return cmd
}

func (it *ApiControl) buildProc(part string) string {
	path := it.buildPart(part)
	return fmt.Sprintf("%s('%s')", "request_" + part, path)
}

func (it *ApiControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("select_cmd") {
		it.DoCmd = lik.StringFromXS(rule.GetContext("data"))
		rule.SetNeedRedraw()
	} else if rule.IsShift("select_table") {
		it.DoTable = lik.StringFromXS(rule.GetContext("data"))
	} else if rule.IsShift("select_id") {
		it.DoId = lik.StringFromXS(rule.GetContext("data"))
	} else if rule.IsShift("select_phone") {
		it.DoPhone = lik.StringFromXS(rule.GetContext("data"))
	} else if rule.IsShift("select_string") {
		it.DoPrm = lik.StringFromXS(rule.GetContext("data"))
	} else if rule.IsShift("select_prm") {
		it.DoPrm = lik.StringFromXS(rule.GetContext("data"))
	} else if rule.IsShift("call") {
		if ans := rule.GetContext("answer"); ans != "" {
			it.Answer = lik.SetFromRequest(lik.StringFromXS(ans))
		}
		rule.SetNeedRedraw()
	} else {
		it.execute(rule)
	}
}

func (it *ApiControl) Marshal(rule ruler.DataRuler) {
}

