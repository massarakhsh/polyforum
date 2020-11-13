package controller

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
	"strings"
)

type DataControl struct {
	ruler.DataControl
}

type Controller interface {
	ruler.Controller
}

func (it *DataControl) buildPart(part string) string {
	return "/" + ruler.GetIdLevel(it.Level) + "/" + part
}

func (it *DataControl) buildUrl(rule ruler.DataRuler, part string) string {
	return rule.BuildUrl(it.buildPart(part))
}

func (it *DataControl) buildProc(proc string, part string, parm string) string {
	parms := "'" + it.buildPart(part) + "'"
	if parm != "" {
		parms += "," + parm
	}
	return proc + "(" + parms + ")"
}

func (it *DataControl) menuPrepare(rule ruler.DataRuler, state bool) likdom.Domer {
	id := fmt.Sprintf("menu_%d", it.Level)
	tbl := likdom.BuildTableClass("menu", "id", id)
	if !state && it.Level + 1 >= rule.GetLevel() {
		it.Mode = ""
	}
	return tbl
}

func (it *DataControl) menuItemText(rule ruler.DataRuler, row likdom.Domer, text string) {
	it.menuItemCmd(rule, row, "", text, "")
}

func (it *DataControl) menuItemImg(rule ruler.DataRuler, row likdom.Domer, mode string, txt string, img string, cmd string) {
	text := ""
	if img != "" {
		item := likdom.BuildUnpairItem("img", "src", img)
		if txt != "" {
			item.SetAttr("title", txt)
		}
		text = item.ToString()
	}
	it.menuItemCmd(rule, row, mode, text, cmd)
}

func (it *DataControl) menuItemCmd(rule ruler.DataRuler, row likdom.Domer, mode string, txt string, cmd string) {
	proc := ""
	if cmd != "" {
		path := it.buildPart(cmd)
		proc = fmt.Sprintf("front_get('%s')", path)
	}
	it.menuItemProc(rule, row, mode, txt, proc)
}

func (it *DataControl) menuItemProc(rule ruler.DataRuler, row likdom.Domer, mode string, txt string, proc string) {
	cls := "menu"
	if mode != "" && mode == it.GetMode() {
		cls += " menu_select"
	}
	td := row.BuildTdClass(cls)
	if proc != "" {
		a := td.BuildItem("a", "href=#", "onclick", proc)
		a.BuildString(txt)
	} else {
		td.BuildString(txt)
	}
}

func (it *DataControl) menuItemSep(rule ruler.DataRuler, row likdom.Domer) {
	td := row.BuildTdClass("menu fill")
	td.BuildString("&nbsp;")
}

func (it *DataControl) menuTools(rule ruler.DataRuler, row likdom.Domer) {
	it.menuItemSep(rule, row)
	if it.Level == 0 {
		row.BuildTdClass("menu").BuildString("<span id=srvtime class=srvtime></span>")
	} else {
		it.menuItemImg(rule, row,"", "Закрыть", "/images/menuexit.png", "exit")
	}
}

func (it *DataControl) linkItemImg(pic string, txt string, cmd string, cls string) likdom.Domer {
	img := likdom.BuildUnpairItem("img", "src", pic)
	if txt != "" {
		img.SetAttr("title", txt)
	}
	return it.linkItemCmd(img.ToString(), cmd, cls)
}

func (it *DataControl) linkItemCmd(txt string, cmd string, cls string) likdom.Domer {
	proc := ""
	if cmd != "" {
		path := it.buildPart(cmd)
		proc = fmt.Sprintf("front_get('%s')", path)
	}
	return it.linkItemProc(txt, proc, cls)
}

func (it *DataControl) linkItemProc(txt string, proc string, cls string) likdom.Domer {
	a := likdom.BuildItemClass("a", cls, "href=#", "onclick", proc)
	a.BuildString(txt)
	return a
}

func (it *DataControl) collectParms(rule ruler.DataRuler, prefix string) lik.Seter {
	parms := lik.BuildSet()
	if context := rule.GetAllContext(); context != nil {
		for _,set := range(context.Values()) {
			if strings.HasPrefix(set.Key, prefix) && set.Val != nil {
				str := lik.StringFromXS(set.Val.ToString())
				parms.SetItem(str, set.Key[len(prefix):])
			}
		}
	}
	return parms
}

func (it *DataControl) execute(rule ruler.DataRuler) {
	if rule.IsShift("exit") {
		rule.BindControl(it.GetLevel(), nil)
	} else if rule.IsShift("seek") {
		rule.BindControl(it.GetLevel() + 1, nil)
	}
}

