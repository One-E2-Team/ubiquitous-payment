FROM node:16-bullseye as build
ARG WEBSHOP_SERVER_HOST_PORT
ARG PSP_SERVER_HOST_PORT
ARG BANK_SERVER_HOST_PORT
ARG WEBSHOP_PROTOCOL
ARG BANK_PROTOCOL
ARG PROTOCOL
ARG ARG_TARGET_STATIC
ENV NODE_ENV=development
COPY ${ARG_TARGET_STATIC} /${ARG_TARGET_STATIC}
RUN npm install -g cross-var preprocessor
WORKDIR /${ARG_TARGET_STATIC}
RUN npm install -d && npm run deploy

FROM nginx as static
ARG ARG_TARGET_STATIC
COPY --from=build /${ARG_TARGET_STATIC}/dist /usr/share/nginx/html/dist/web
#RUN touch /etc/nginx/passthrough.conf
#RUN echo "include /etc/nginx/passthrough.conf;" >> /etc/nginx/nginx.conf
EXPOSE 80
EXPOSE 443
EXPOSE 1081
STOPSIGNAL SIGTERM
CMD ["nginx", "-g", "daemon off;"]