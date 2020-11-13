package controller

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/ruler"
)

type RootControl struct {
	DataControl
}

type Rooter interface {
	Controller
}

func BuildRoot(rule ruler.DataRuler, level int) Rooter {
	it := &RootControl{ }
	rule.BindControl(level, it)
	return it
}

func (it *RootControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.menuPrepare(rule, false)
	row := tbl.BuildTr()
	it.menuItemCmd(rule, row, "", "POLYFORUM", "seek")
	it.menuItemText(rule, row, base.Version)
	it.menuItemText(rule, row,"|")
	it.menuItemCmd(rule, row,"file", "Файлы", "file")
	it.menuItemCmd(rule, row,"api", "API", "api")
	it.menuItemCmd(rule, row,"diction", "Справочники", "diction")
	it.menuItemCmd(rule, row,"table", "Объекты", "table")
	it.menuTools(rule, row)
	return tbl
}

func (it *RootControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	return div
}

func (it *RootControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("admin") {
		it.Mode = "admin"
		BuildAdmin(rule, it.Level + 1)
	} else if rule.IsShift("api") {
		it.Mode = "api"
		BuildApi(rule, it.Level + 1)
	} else if rule.IsShift("diction") {
		it.Mode = "diction"
		BuildTabler(rule, it.Level + 1, false)
	} else if rule.IsShift("table") {
		it.Mode = "table"
		BuildTabler(rule, it.Level + 1, true)
	} else if rule.IsShift("file") {
		it.Mode = "file"
		BuildFile(rule, it.Level + 1)
	} else if rule.IsShift("db") {
		it.Mode = "db"
		BuildDb(rule, it.Level + 1)
	} else {
		it.execute(rule)
	}
}

func (it *RootControl) Marshal(rule ruler.DataRuler) {
}

