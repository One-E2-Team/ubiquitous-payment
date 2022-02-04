FROM node as build
ARG WEBSHOP_SERVER_HOST_PORT
ARG PSP_SERVER_HOST_PORT
ARG BANK_SERVER_HOST_PORT
ARG WEBSHOP_PROTOCOL
ARG BANK_PROTOCOL
ARG PROTOCOL
ENV NODE_ENV=development
COPY webshop-front /webshop-front
COPY psp-front /psp-front
COPY bank-front /bank1-front
COPY bank-front /bank2-front
RUN npm install --save cross-var preprocessor
WORKDIR /webshop-front
RUN npm install -d && npm run deploy
WORKDIR /psp-front
RUN npm install -d && npm run deploy
WORKDIR /bank1-front
RUN npm install -d && npm run deploy
WORKDIR /bank2-front
RUN npm install -d && npm run deploy

FROM nginx as static
ARG ARG_TARGET_STATIC
COPY --from=build /${ARG_TARGET_GO}/dist /usr/share/nginx/html/dist/web
EXPOSE 80
EXPOSE 443
EXPOSE 1081
STOPSIGNAL SIGTERM
CMD ["nginx", "-g", "daemon off;"]