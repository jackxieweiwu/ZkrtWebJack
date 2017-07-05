/**
 * Created by root on 17-6-30.
 */
var gaoDeMap, ruler1,mouseTool;
gaoDeMap = new AMap.Map('map_div', {
    resizeEnable: true,
    zoom: 12
});

if (document.createElement('canvas')
    && document.createElement('canvas').getContext
    && document.createElement('canvas').getContext('2d')) {
    // 实例化3D楼块图层
    var buildings = new AMap.Buildings();
    // 在map中添加3D楼块图层
    buildings.setMap(gaoDeMap);
} else {
    document.getElementById('tip').innerHTML = "对不起，运行该示例需要浏览器支持HTML5！";
}

gaoDeMap.plugin(["AMap.RangingTool"], function() {
    ruler1 = new AMap.RangingTool(gaoDeMap);
    AMap.event.addListener(ruler1, "end", function(e) {
        ruler1.turnOff();
    });
    var sMarker = {
        icon: new AMap.Icon({
            size: new AMap.Size(19, 31),//图标大小
            image: "http://webapi.amap.com/theme/v1.3/markers/n/mark_b1.png"
        })
    };
    var eMarker = {
        icon: new AMap.Icon({
            size: new AMap.Size(19, 31),//图标大小
            image: "http://webapi.amap.com/theme/v1.3/markers/n/mark_b2.png"
        }),
        offset: new AMap.Pixel(-9, -31)
    };
    var lOptions = {
        strokeStyle: "solid",
        strokeColor: "#FF33FF",
        strokeOpacity: 1,
        strokeWeight: 2
    };
    var rulerOptions = {startMarkerOptions: sMarker, endMarkerOptions: eMarker, lineOptions: lOptions};
});

//开始启用面积测量
gaoDeMap.plugin(["AMap.MouseTool"], function() {
    mouseTool = new AMap.MouseTool(gaoDeMap);
    AMap.event.addListener(mouseTool, "draw", function callback(e) {
        var eObject = e.obj;
    });
    mouseTool.close(false);
});

//启用默认样式测距
function startRuler1() {
    ruler1.turnOn();
}

function startRuler2() {
    gaoDeMap.clearMap();
}

function startRuler3() {
    var iid =document.getElementById("acreage")
    if(iid.value == "面积测量"){
        iid.value = "取消测量"
        mouseTool.measureArea();
    }else if(iid.value == "取消测量"){
        iid.value = "面积测量"
        mouseTool.close(false);
    }
}
