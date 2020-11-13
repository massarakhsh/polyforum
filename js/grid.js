var grid_Instant;
var row_likSelect;

function grid_redraw(elm) {
    let path = elm.attr('path');
    front_proc(path, grid_redraw_init, elm);
}

function grid_redraw_init(elm, lika) {
    let instant = ('grid' in lika) ? lika.grid : [];
    grid_prepare(instant);
    let grid = elm.DataTable(instant);
    grid_Instant = instant;
    row_likSelect = ('likSelect' in instant) ? instant.likSelect : 0;
    grid.on( 'select', grid_select);
    grid.on( 'draw', grid_draw_done);
}

function grid_prepare(data) {
    if (data !== null && typeof(data) == 'object') {
        for (var key in data) {
            let value = data[key];
            if (typeof(value) == 'string') {
                var match;
                if (match = /^function_(.+)\((.*)\)/.exec(value)) {
                    let func = match[1];
                    let parm = match[2];
                    if (func in window) {
                        data[key] = function () {
                            window[func](this, parm);
                        };
                    } else {
                        data[key] = grid_nothing;
                    }
                } else if (match = /^function_(.+)/.exec(value)) {
                    let func = match[1];
                    if (func in window) {
                        data[key] = window[func];
                    } else {
                        data[key] = grid_nothing;
                    }
                }
            } else if (value !== null && typeof(value) == 'object') {
                grid_prepare(data[key]);
            }
        }
    }
}

function grid_nothing() {
}

function grid_select( e, dt, type, indexes ) {
    if ( type === 'row' ) {
        if (indexes && indexes.length > 0) {
            var datas = dt.rows(indexes).data();
            if (datas && datas.length>0) {
                row_likSelect = datas[0].DT_RowId;
                let path = grid_Instant.ajax;
                if (match = /^(.+)griddata(.*)$/.exec(path)) {
                    path = match[1] + "select/" + row_likSelect + match[2];
                    get_data_part(path);
                }
            }
        }
    }
}

function grid_draw_done( e, settings ) {
    //alert('grid_draw_done');
    var api = new $.fn.dataTable.Api( settings );
    api.rows().eq(0).each( function ( index ) {
        var row = api.row( index );
        var data = row.data();
        if (data.DT_RowId == row_likSelect) {
            row.select();
        }
        // ... do something with data(), or row.node(), etc
    } );    //api.row(':eq(7)').select();
    //alert(api.rows().length);
    //console.log( api.rows( {page:'current'} ).data() );
}
