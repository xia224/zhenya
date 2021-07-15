# Use nginx to receive rtmp

## Compile nginx with rtmp
./configure --add-module=../rtmp/
make install

sudo /usr/local/nginx/sbin/nginx -c /home/devops/rtmp_pusher/nginx.conf -g "daemon off;"
./configure --add-module=/path/to/nginx-rtmp-module --with-debug
NGX_RTMP_MAX_URL 256
make CFLAGS='-Wno-implicit-fallthrough'


## Run nginx 
cd objs;
./nginx -g 'daemon 0ff;'

sudo docker run -d --restart always --net=host --pid=host --log-driver json-file --log-opt max-size=10m --log-opt max-file=10  -v /home/devops/rtmp_pusher/nginx.conf:/etc/nginx/nginx.conf -v /data/uap/rtmp_pusher/log/:/usr/local/nginx/logs --name nginx-rtmp hub.agoralab.co/uap/rtmp_pusher/nginx-rtmp:custom_args_length

# Version
devops@tianjin1g-ctel-42-81-205-219:~/rtmp_pusher/compile_nginx$ ./nginx_latest -v
nginx version: nginx/1.15.0