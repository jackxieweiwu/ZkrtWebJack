/**
 * Created by root on 17-6-30.
 */
var message, player
(function(){
    function getParams(name) {
        var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
        var r = window.location.search.substr(1).match(reg);
        if (r != null) {
            return decodeURIComponent(r[2]);
        }
        return null;
    }
    var rtmp = getParams('rtmp'),
        flv  = getParams('flv'),
        m3u8 = getParams('m3u8'),
        mp4  = getParams('mp4'),
        live = (getParams('live') == 'true' ? true : false),
        coverpic = getParams('coverpic'),
        width = getParams('width'),
        height = getParams('height'),
        autoplay = (getParams('autoplay') == 'false' ? true : false);

    var options = {
        rtmp: rtmp || 'rtmp://localhost:1935/live/movie',
        flv: flv || 'http://localhost:7001/live/movie.flv',
        m3u8: m3u8 || 'http://localhost:7002/live/movie.m3u8',
        mp4 : mp4,
        coverpic: coverpic || {style:'cover', src:'//vodplayerinfo-10005041.file.myqcloud.com/3035579109/vod_paster_pause/paster_pause1469013308.jpg'},
        autoplay: autoplay ? true : false,
        live: live,
        width : width,
        height : height,
        wording: {
            1002: '即将直播，请稍等',
            2032: '请求视频失败，请检查网络',
            2048: '请求m3u8文件失败，可能是网络错误或者跨域问题'
        },
        /*listener: function (msg) {
            // console.log('listener',msg);
            message = msg.type
            if(message == 'error'){
                window.setTimeout(function(){
                    player.load();//进行重连
                },3000);
            }
        }*/

    };
    player = new TcPlayer('video-container', options);
    window.qcplayer  = player;
})();