FROM node:20.15-slim

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV PORT 3000
ENV ROOTPATH /client
ENV NPM_CONFIG_PREFIX /home/node/.npm-global
ENV PATH $PATH:/home/node/.npm-global/bin

WORKDIR $ROOTPATH

RUN npm install prettier -g

EXPOSE 3000
