package controller

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
)

type EditControl struct {
	DataControl
	Part	string
	Id		lik.IDB
	IsShow	bool
}

type Editor interface {
	Controller
}

func BuildEditor(rule ruler.DataRuler, level int, part string, id lik.IDB, mode string) Editor {
	it := &EditControl{ Part: part, Id: id }
	it.Mode = mode
	if mode == "show" {
		it.IsShow = true
	}
	rule.BindControl(level, it)
	return it
}

func (it *EditControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.menuPrepare(rule, true)
	row := tbl.BuildTr()
	if table := base.GetTable(it.Part); table != nil {
		if it.Mode == "edit" {
			it.menuItemText(rule, row, fmt.Sprintf("Редактирование записи ID=%d", int(it.Id)))
			it.menuItemText(rule, row, "|")
			path := it.buildPart("write")
			proc := fmt.Sprintf("edit_write('%s')", path)
			it.menuItemProc(rule, row, "","Записать изменения", proc)
			it.menuItemCmd(rule, row, "","Отменить", "cancel")
		} else if it.Mode == "create" {
			it.menuItemText(rule, row, "Новая запись")
			it.menuItemText(rule, row, "|")
			path := it.buildPart("write")
			proc := fmt.Sprintf("edit_write('%s')", path)
			it.menuItemProc(rule, row, "","Записать новую", proc)
			it.menuItemCmd(rule, row, "","Отменить", "cancel")
		} else if it.Mode == "delete" {
			it.menuItemText(rule, row, fmt.Sprintf("Удаление записи ID=%d", int(it.Id)))
			it.menuItemText(rule, row, "|")
			it.menuItemCmd(rule, row, "","ДЕЙСТВИТЕЛЬНО УДАЛИТЬ", "realdelete")
			it.menuItemCmd(rule, row, "","Отменить", "cancel")
		} else {
			it.menuItemText(rule, row, fmt.Sprintf("Запись ID=%d", int(it.Id)))
			it.menuItemText(rule, row, "|")
			it.menuItemCmd(rule, row, "","Изменить", "edit")
			it.menuItemCmd(rule, row, "","Удалить", "delete")
			it.menuItemCmd(rule, row, "","Закрыть", "exit")
		}
	}
	it.menuTools(rule, row)
	return tbl
}

func (it *EditControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	if table := base.GetTable(it.Part); table != nil {
		var elm lik.Seter
		if it.Mode != "create" && it.Id > 0 {
			elm = base.GetElm(it.Part, it.Id)
		}
		tbl := div.BuildTableClass("edit")
		for _,field := range table.Fields {
			key := field.Key
			row := tbl.BuildTr()
			if td := row.BuildTdClass("edit edit_title"); td != nil {
				td.BuildString(field.Title)
			}
			if td := row.BuildTdClass("edit edit_info"); td != nil {
				isdate := lik.RegExCompare(key, "Dt$")
				istime := lik.RegExCompare(key, "At$")
				input := td.BuildUnpairItem("input", "type=text")
				if isdate {
					input.SetAttr("class", "tcal")
				} else if istime {
					input.SetAttr("class", "tmcl")
				}
				if key == "Id" {
					input.SetAttr("readonly", "")
				} else if it.Mode != "edit" && it.Mode != "create" {
					input.SetAttr("readonly", "")
				} else {
					input.SetAttr("id", "up_"+key)
				}
				if elm != nil {
					value := elm.GetString(key)
					input.SetAttr("value", value)
				}
				input.SetAttr("redraw", "form_field_init")
			}
		}
	}
	return div
}

func (it *EditControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("cancel") {
		it.Mode = "show"
	} else if rule.IsShift("edit") {
		it.Mode = "edit"
	} else if rule.IsShift("delete") {
		it.Mode = "delete"
	} else if rule.IsShift("realdelete") {
		base.DeleteElm(it.Part, it.Id)
		rule.BindControl(it.GetLevel(), nil)
	} else if rule.IsShift("write") {
		it.write(rule)
	} else {
		it.execute(rule)
	}
	rule.SetNeedRedraw()
}

func (it *EditControl) write(rule ruler.DataRuler) {
	if parms := it.collectParms(rule, "up_"); parms != nil {
		sets := lik.BuildSet()
		if table := base.GetTable(it.Part); table != nil {
			for _, field := range table.Fields {
				key := field.Key
				if val := parms.GetItem(key); val != nil {
					if lik.RegExCompare(field.Proto, "L") {
						sets.SetItem(val.ToInt(), key)
					} else {
						sets.SetItem(val.ToString(), key)
					}
				}
			}
			if it.Mode == "create" {
				it.Id = base.InsertElm(it.Part, sets)
			} else if it.Mode == "edit" && it.Id > 0 {
				base.UpdateElm(it.Part, it.Id, sets)
			}
		}
	}
	if it.IsShow {
		it.Mode = "show"
	} else {
		rule.BindControl(it.GetLevel(), nil)
	}
}

func (it *EditControl) Marshal(rule ruler.DataRuler) {
}

