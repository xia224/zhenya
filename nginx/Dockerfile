FROM hub.agoralab.co/uap/rtmp_pusher/nginx-rtmp:latest

COPY ./nginx_latest /usr/local/sbin/nginx
RUN mkdir -p /usr/local/nginx/logs

ENTRYPOINT [ "/usr/local/sbin/nginx", "-c",  "/etc/nginx/nginx.conf", "-g", "daemon off;" ]

