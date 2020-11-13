package controller

import (
	"github.com/massarakhsh/lik"
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"github.com/massarakhsh/polyforum/ruler"
)

type GridControl struct {
	DataControl
	Part	string
	Start	int
	Length	int
	Total	int
	IdSel	lik.IDB
}

type Grider interface {
	Controller
}

func BuildGrider(rule ruler.DataRuler, level int, part string) Grider {
	it := &GridControl{ }
	it.Part = part
	rule.BindControl(level, it)
	return it
}

func (it *GridControl) ShowMenu(rule ruler.DataRuler) likdom.Domer {
	tbl := it.menuPrepare(rule, false)
	row := tbl.BuildTr()
	if table := base.GetTable(it.Part); table != nil {
		it.menuItemCmd(rule, row, "", table.Title+" ("+table.Part+")", "seek")
		it.menuItemText(rule, row,"|")
		it.menuItemCmd(rule, row,"append", "Добавить", "append")
	} else {
		it.menuItemText(rule, row, it.Part)
	}
	it.menuTools(rule, row)
	return tbl
}

func (it *GridControl) ShowInfo(rule ruler.DataRuler) likdom.Domer {
	div := likdom.BuildDivClass("grid")
	div.AppendItem(it.ShowTable(it.Level, it.Part))
	return div
}

func (it *GridControl) Execute(rule ruler.DataRuler) {
	if rule.IsShift("gridinit") {
		it.execGridInit(rule)
	} else if rule.IsShift("griddata") {
		it.execGridData(rule)
	} else if rule.IsShift("select") {
		it.execSelect(rule)
	} else if rule.IsShift("open") {
		it.Mode = "open"
		it.execOpen(rule)
	} else if rule.IsShift("append") {
		it.Mode = "append"
		it.execAppend(rule)
	} else {
		it.execute(rule)
	}
}

func (it *GridControl) execGridInit(rule ruler.DataRuler) {
	grid := lik.BuildSet()
	grid.SetItem(true, "serverSide")
	grid.SetItem(true, "processing")
	grid.SetItem(it.execInitLanguage(rule), "language")
	grid.SetItem(false, "searching")
	grid.SetItem(false, "lengthChange")
	grid.SetItem("single", "select/style")
	if it.IdSel > 0 {
		grid.SetItem(it.IdSel, "likSelect")
	}
	if table := base.GetTable(it.Part); table != nil {
		columns := lik.BuildList()
		for _,field := range table.Fields {
			columns.AddItemSet("data", field.Key, "width=100px")
		}
		grid.SetItem(columns, "columns")
	}
	grid.SetItem(rule.BuildUrl("/front/" + ruler.GetIdLevel(it.Level) + "/griddata"), "ajax")
	rule.SetResponse(grid, "grid")
}

func (it *GridControl) execInitLanguage(rule ruler.DataRuler) lik.Seter {
	data := lik.BuildSet()
	data.SetItem("Поиск", "search")
	data.SetItem("Таблица пуста", "emptyTable")
	data.SetItem("Строки от _START_ до _END_, всего _TOTAL_", "info")
	data.SetItem("Загрузка ...", "loadingRecords")
	data.SetItem("Обработка ...", "processing")
	data.SetItem("Нет строк в таблице", "infoEmpty")
	data.SetItem("В начало", "paginate/first")
	data.SetItem("Назад", "paginate/previos")
	data.SetItem("Вперёд", "paginate/next")
	data.SetItem("В конец", "paginate/last")
	return data
}

func (it *GridControl) execGridData(rule ruler.DataRuler) {
	if parm := rule.GetContext("draw"); parm != "" {
		rule.SetResponse(lik.StrToInt(parm), "draw")
	}
	if parm := rule.GetContext("start"); parm != "" {
		it.Start = lik.StrToInt(parm)
	}
	if parm := rule.GetContext("length"); parm != "" {
		it.Length = lik.StrToInt(parm)
	}
	if it.Length == 0 {
		it.Length = 10
	}
	data := lik.BuildList()
	if table := base.GetTable(it.Part); table != nil {
		if list := base.DB.GetListAll(it.Part); list != nil {
			it.Total = list.Count()
			rule.SetResponse(it.Total, "recordsTotal")
			rule.SetResponse(it.Total, "recordsFiltered")
			for n := 0; n < it.Length && it.Start + n < it.Total; n++ {
				nr := (it.Start + n)
				if elm := list.GetSet(nr); elm != nil {
					id := elm.GetIDB("Id")
					row := lik.BuildSet("DT_RowId", id)
					for _, fld := range table.Fields {
						if fld.Key == "Id" {
							ids := lik.IDBToStr(id)
							ent := it.linkItemCmd("[" + ids + "]", "open/" + ids, "cmd")
							row.SetItem(ent.ToString(), fld.Key)
						} else {
							data := lik.LimitString(elm.GetString(fld.Key), 30)
							row.SetItem(data, fld.Key)
						}
					}
					data.AddItems(row)
				}
			}
		}
	}
	rule.SetResponse(data, "data")
}

func (it *GridControl) execSelect(rule ruler.DataRuler) {
	it.IdSel = lik.IDB(lik.StrToInt(rule.Shift()))
	//rule.StoreItem(it.ShowMenu(rule))
}

func (it *GridControl) execOpen(rule ruler.DataRuler) {
	key := lik.StrToIDB(rule.Shift())
	if key > 0 {
		it.IdSel = key
	} else {
		key = it.IdSel
	}
	if key > 0 {
		BuildEditor(rule, it.Level+1, it.Part, key, "show")
		rule.SetNeedRedraw()
	}
}

func (it *GridControl) execAppend(rule ruler.DataRuler) {
	BuildEditor(rule, it.Level + 1, it.Part, 0, "create")
	rule.SetNeedRedraw()
}

func (it *GridControl) Marshal(rule ruler.DataRuler) {
}

