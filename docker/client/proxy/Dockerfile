FROM nginx:1.24-bullseye

COPY ./docker/client/proxy/nginx.conf /etc/nginx/nginx.conf
COPY ./docker/client/proxy/default.conf.template /etc/nginx/templates/default.conf.template

ENV TZ=Asia/Tokyo
