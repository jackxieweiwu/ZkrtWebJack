/**
 * Created by root on 17-6-30.
 */
var elements = new Array();
var markers = []
var objj,obj
$(function () {
    var ws = new WebSocket('ws://'+window.location.host+'/ws')
    ws.onmessage = function (e) {
        var droneobj = new Function("return" + event.data)();
        obj = JSON.parse(droneobj.Message)
        /*var returnObj = {};
         returnObj.message = '还活着 ' + new Date().toLocaleString();
         console.log(returnObj);
         returnObj = JSON.stringify(returnObj);
         ws.send(returnObj);*/

        var marker
        objj = GPS.gcj_encrypt(obj.DroneGpsLng,obj.DroneGpsLat)
        if(elements.indexOf(obj.DroneId) == -1){
            elements.push(obj.DroneId)
            marker = new AMap.Marker({
                icon: "../static/img/icon_uav.png",
                map:gaoDeMap,
                position:[parseFloat(objj.lat),parseFloat(objj.lon)]
            })
            markers.push(marker)
        }else{
            marker = markers[elements.indexOf(obj.DroneId)];
            marker.setPosition(new AMap.LngLat(parseFloat(objj.lat),parseFloat(objj.lon)))
        }

        marker.on('click',markerClock)
    }
})

function markerClock() {
    setValues();
    player.load();//进行重连
}
//赋值
function setValues() {
    $("#lat_drone_msg").text(objj.lat);
    $("#lon_drone_msg").text(objj.lon);
    $("#alt_drone_msg").text(obj.Drone_alt);
    $("#r_drone_msg").text(obj.Drone_r);
    $("#p_drone_msg").text(obj.Drone_p);
    $("#y_drone_msg").text(obj.Drone_y);
}

function gas_onoe() {
    var obj = document.getElementById("tr_co");
    if(obj.style.display=="none"){
        obj.style.display = "block";
        document.getElementById("tr_nh3").style.display = "block";
        document.getElementById("tr_h2s").style.display = "block";
        document.getElementById("tr_co2").style.display = "block";
        document.getElementById("tr_voc").style.display = "block";
        document.getElementById("tr_cl2").style.display = "block";
        document.getElementById("tr_so2").style.display = "block";
        document.getElementById("tr_ch4").style.display = "block";
    }else{
        obj.style.display = "none"
        document.getElementById("tr_nh3").style.display = "none";
        document.getElementById("tr_h2s").style.display = "none";
        document.getElementById("tr_co2").style.display = "none";
        document.getElementById("tr_voc").style.display = "none";
        document.getElementById("tr_cl2").style.display = "none";
        document.getElementById("tr_so2").style.display = "none";
        document.getElementById("tr_ch4").style.display = "none";
    }
}