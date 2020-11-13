package controller

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/ruler"
	"fmt"
)

type AdminControl struct {
	DataControl
}

type Adminer interface {
	Controller
}

func BuildAdmin(rule ruler.DataRuler, level int) Adminer {
	it := &AdminControl{ }
	rule.BindControl(level, it)
	return it
}

func (it *AdminControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := likdom.BuildTableClass("menu")
	row := tbl.BuildTr()
	it.menuItemText(rule, row,"Администрирование")
	it.menuItemText(rule, row, "|")
	it.menuTools(rule, row)
	return tbl
}

func (it *AdminControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	tbl := div.BuildTableClass("api")
	if row := tbl.BuildTr(); row != nil {
		row.BuildTdClass("api_title").BuildString("Очистка базы данных")
		proc := fmt.Sprintf("purge_all_base('%s')", it.buildPart("purgebase"))
		row.BuildTdClass("api_info").AppendItem(it.linkItemProc("Очистить", proc, "cmd"))
	}
	return div
}

func (it *AdminControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("purgebase") {
		fmt.Println("purgebase")
	} else {
		it.execute(rule)
	}
}

func (it *AdminControl) Marshal(rule ruler.DataRuler) {
}

//func (it *AdminControl) buildLinePurge(rule ruler.DataRuler) likdom.Domer {
//	proc := it.buildProc("purge")
//	line := LinkTe("api", "Очистить", proc)
//	return line
//}

func (it *AdminControl) buildProc(part string) string {
	path := it.buildPart(part)
	return fmt.Sprintf("%s('%s')", "admin_" + part, path)
}

