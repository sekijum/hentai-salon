FROM nginx:1.24-bullseye

COPY ./docker/server/proxy/nginx.conf /etc/nginx/nginx.conf
COPY ./docker/server/proxy/default.conf.template /etc/nginx/templates/default.conf.template

ENV APP_HOST=127.0.0.1
ENV TZ=Asia/Tokyo
