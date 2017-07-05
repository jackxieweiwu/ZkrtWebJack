/**
 * Created by root on 17-7-3.
 */
$(function () {
    var ws = new WebSocket('ws://'+window.location.host+'/ws/dronestatus')
    ws.onmessage = function (e) {
        var obj_status = new Function("return" + event.data)();
        var objstatus = JSON.parse(obj_status.MessageSta)
        document.write(objstatus.DroneTmp+"<br>")
    }
})
