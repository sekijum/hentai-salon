FROM node:20.15-slim as base

ARG APP_ENV=development
ARG NUXT_PUBLIC_API_BASE_URL=http://localhost:8080
ARG STATIC_URL
ARG GA_MEASUREMENT_ID

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
ENV PORT 3000
ENV ROOTPATH /app
ENV NPM_CONFIG_PREFIX /home/node/.npm-global
ENV PATH $PATH:/home/node/.npm-global/bin
ENV APP_ENV=${APP_ENV}
ENV NUXT_PUBLIC_API_BASE_URL=${NUXT_PUBLIC_API_BASE_URL}
ENV STATIC_URL=${STATIC_URL}
ENV GA_MEASUREMENT_ID=${GA_MEASUREMENT_ID}

RUN echo "APP_ENV: ${APP_ENV}" && \
    echo "NUXT_PUBLIC_API_BASE_URL: ${NUXT_PUBLIC_API_BASE_URL}" && \
    echo "STATIC_URL: ${STATIC_URL}" && \
    echo "GA_MEASUREMENT_ID: ${GA_MEASUREMENT_ID}"

WORKDIR $ROOTPATH

COPY packages/client $ROOTPATH

RUN npm ci \
    && npm cache clean --force \
    && npm run build

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
