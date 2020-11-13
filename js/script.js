var script_second = 0;

function script_start() {
    pool_step.push(script_step);
    if (!lik_trust) lik_set_trust("");
    if (lik_trust) lik_set_marshal(1000, "/marshal");
}

function script_step() {
    script_showtime();
    script_redraw();
}

function script_showtime() {
    if (script_second!=tick_second) {
        script_second = tick_second;
        var elm = jQuery('#srvtime');
        if (elm.size()>0) {
            var text = "";
            var ok = true;
            if (tick_total - tick_answer < 5*1000) {
                var tt = tick_server - tick_shift_minute * 60;
                text = build_showtime(tt, true);
                ok = true;
            } else if (tick_total - tick_answer < 5*60*1000) {
                var tt = Math.floor((tick_total - tick_answer) / 1000);
                text = "Нет связи " + build_showtime(tt, false);
                ok = false;
            } else {
                lik_stop();
                text = "<b>СИСТЕМА ОСТАНОВЛЕНА</b>";
                ok = false;
            }
            if (ok && elm.hasClass("srvoff")) {
                elm.removeClass("srvoff");
            } else if (!ok && !elm.hasClass("srvoff")) {
                elm.addClass("srvoff");
            }
            elm.html(text);
        }
    }
}

function build_showtime(tt, ok) {
    var ts = tt % 60;
    tt = (tt - ts) / 60;
    var tm = tt % 60;
    tt = (tt - tm) / 60;
    var th = tt % 24;
    tt = (tt - th) / 24;
    var text = "";
    if (ok) {
        text += (th >= 10) ? "" + th : "0" + th;
        text += (ts & 1) ? ":" : " ";
        text += (tm >= 10) ? tm : "0" + tm;
    } else {
        if (th > 0) text += ""+th+":";
        text += tm;
        text += (ts >= 10) ? ":" + ts : ":0" + ts;
    }
    return text;
}

function script_redraw() {
    let rdr = jQuery('[redraw]');
    if (rdr.size() > 0) {
        rdr.each(function (idx, item) {
            let elm = jQuery(item);
            let redraw = elm.attr('redraw');
            elm.removeAttr('redraw');
            setTimeout(function () {
                if (redraw in window) {
                    window[redraw](elm);
                }
            }, 50);
        });
    }
}

function combine_front(part) {
    let path = "/front";
    if (/^\//.exec(part)) {
        path += part;
    } else if (part) {
        path += "/" + part;
    }
    return path;
}

function front_get(cmd) {
    get_data_part(combine_front(cmd));
}

function front_post(cmd, data) {
    post_data_part(combine_front(cmd), data);
}

function front_proc(cmd, proc, parm) {
    get_data_proc(combine_front(cmd), proc, parm);
}

function front_post_proc(cmd, data, proc, parm) {
    post_data_proc(combine_front(cmd), data, proc, parm);
}

function edit_write(cmd) {
    var data = edit_collect();
    front_post(cmd, data);
}

function edit_collect() {
    let data = "";
    jQuery('[id^=up_]').each(function(idx,item) {
        var elm = jQuery(item);
        if (data) data += "&";
        data += elm.attr('id') + "=" + string_to_XS(elm.val());
    });
    return data;
}

function purge_all_base(cmd) {
    if (confirm("Вы действительно хотите очистить базу данных?")) {
        front_get("/" + cmd);
    }
}

function create_folder(path) {
    let name = prompt("Введите имя новой папки", "Новая папка");
    if (name) {
        front_get(path + "/" + string_to_XS(name));
    }
}

function rename_file(path, name) {
    let newname = prompt("Введите новое имя", name);
    if (newname && newname != name) {
        front_get(path + "/" + string_to_XS(newname));
    }
}

function delete_file(path, name) {
    if (confirm("Действительно удалить \"" + name + "\"?")) {
        front_get(path);
    }
}
