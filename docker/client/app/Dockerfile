FROM node:20.15-slim as base

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV PORT 3000
ENV ROOTPATH /app
ENV NPM_CONFIG_PREFIX /home/node/.npm-global
ENV PATH $PATH:/home/node/.npm-global/bin

WORKDIR $ROOTPATH

COPY packages/client $ROOTPATH

RUN npm ci \
    && npm cache clean --force

EXPOSE 3000

# 開発用イメージ
FROM base as dev
RUN npm install prettier -g
ENTRYPOINT ["npm", "run", "dev"]
# CMD ["/bin/bash"]

# 本番用イメージ
FROM base as prd
RUN npm install pm2 -g
COPY --from=base $ROOTPATH/node_modules ./node_modules
COPY --from=base $ROOTPATH/.output ./.output
COPY --from=base $ROOTPATH/ecosystem.config.json ./ecosystem.config.json
COPY --from=base $ROOTPATH/package.json ./package.json
ENTRYPOINT ["npm", "run", "prd"]
