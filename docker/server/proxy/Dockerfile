FROM nginx:1.24-bullseye

COPY ./docker/server/proxy/nginx.conf /etc/nginx/nginx.conf
COPY ./docker/server/proxy/default.conf.template /etc/nginx/templates/default.conf.template

ENV TZ=Asia/Tokyo
