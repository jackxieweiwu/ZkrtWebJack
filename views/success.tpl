<html>
<head>
    <meta content="text/html; charset=utf-8" />
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no" />
    <style type="text/css">
        body, html,#map {width: 100%;height: 100%;overflow: hidden;margin:0;font-family:"微软雅黑";}
    </style>
    <title>展示</title>
    <link rel="stylesheet" href="http://cache.amap.com/lbs/static/main1119.css"/>
    <script type="text/javascript" src="http://webapi.amap.com/maps?v=1.3&key=471fdfd1038ae0f3ea42d9440884fc58"></script>
    <script type="text/javascript" src="http://cache.amap.com/lbs/static/addToolbar.js"></script>
    <script src="//webapi.amap.com/ui/1.0/main.js"></script>

</head>
    <body>
        <div class="author">
            Official website:
            <a class="lat" href="mailto:{{.Lat}}">{{.Lat}}</a>
            <a class="lng" href="mailto:{{.Lng}}">{{.Lng}}</a>
        </div>
        <div id='container'></div>
        <div class="button-group">
            <input type="button" class="button" value="距离量测" onClick="javascript:startRuler1()"/>
            <input type="button" class="button" value="面积测量" id="acreage" onClick="javascript:startRuler3()"/>
            <input type="button" class="button" value="清除痕迹"  onClick="javascript:startRuler2()"/>
        </div>
        <div id="tip"></div>
        <script type="text/javascript">
            var map, ruler1,mouseTool;
            map = new AMap.Map('container', {
                resizeEnable: true,
                zoom: 12
            });

            if (document.createElement('canvas')
                && document.createElement('canvas').getContext
                && document.createElement('canvas').getContext('2d')) {
                // 实例化3D楼块图层
                var buildings = new AMap.Buildings();
                // 在map中添加3D楼块图层
                buildings.setMap(map);
            } else {
                document.getElementById('tip').innerHTML = "对不起，运行该示例需要浏览器支持HTML5！";
            }


            map.plugin(["AMap.RangingTool"], function() {
                ruler1 = new AMap.RangingTool(map);
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
            map.plugin(["AMap.MouseTool"], function() {
                mouseTool = new AMap.MouseTool(map);
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
                map.clearMap();
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

            //添加基本的UI控件
            AMapUI.loadUI(['control/BasicControl'], function(BasicControl) {
                //图层切换控件
                map.addControl(new BasicControl.LayerSwitcher({
                    position: 'rt'
                }));
            });

            var marker1 = new AMap.Marker({
                icon: "../static/img/icon_uav.png",
                position: [123.400973, 41.784646]
            });

            marker1.setMap(map);//将标记小车图片加到地图
            marker1.moveAlong(lineArr);
        </script>
        <script src="/static/js/reload.min.js"></script>
    </body>
</html>
