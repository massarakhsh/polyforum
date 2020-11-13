package controller

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/polyforum/base"
	"fmt"
	"math/rand"
)

func (it *DataControl) ShowTable(lev int, part string) likdom.Domer {
	tbl := it.ShowGrid(it.buildPart("gridinit"))
	if table := base.GetTable(part); table != nil {
		t1 := tbl.BuildItem("thead").BuildTr()
//		t2 := tbl.BuildItem("thead").BuildTr()
		for _, fld := range table.Fields {
			t1.BuildItem("th").BuildString(fld.Title + "<br/><small>" + fld.Key + "</small>")
//			t2.BuildItem("th").BuildString(fld.Title)
		}
	}
	return tbl
}

func (it *DataControl) ShowGrid(path string) likdom.Domer {
	id := fmt.Sprintf("id_%d", 100000 + rand.Int31n(900000))
	return likdom.BuildTableClassId("grid",id, "path", path, "redraw=grid_redraw")
}

