##jack xie
## Use

  ffmpeg -re -i demo.flv -c copy -f flv rtmp://ip/live/phone
  ffmpeg -report -f dshow -i audio="virtual-audio-capturer" -f dshow -i video="screen-capture-recorder" -vcodec 
    libx264 -acodec aac -s 1920*1080 -r 25 -g 25 -pix_fmt yuv420p -preset veryfast -tune zerolatency -f flv rtmp://ip/live/phone

## 直播端
    ffplay.exe -i rtmp://ip/live/phone -fflags nobuffer
    
    
    