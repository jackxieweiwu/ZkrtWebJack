
## Use
2. run  `livego` to start livego server
3. push `RTMP` stream to `rtmp://localhost:1935/live/movie`, eg use `ffmpeg -re -i demo.flv -c copy -f flv rtmp://localhost:1935/live/movie`
4. play live stream form:
    - `RTMP`:`rtmp://localhost:1935/live/movie`
    - `FLV`:`http://127.0.0.1:7001/live/movie.flv`
    - `HLS`:`http://127.0.0.1:7002/live/movie.m3u8`
