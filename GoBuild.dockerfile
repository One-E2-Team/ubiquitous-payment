FROM golang AS build
COPY . src/ubiquitous-payment/
ARG ARG_TARGET_GO
RUN cd src/ubiquitous-payment && go mod download && go mod verify && go build -o exec ${ARG_TARGET_GO}/main.go


FROM alpine AS image
COPY --from=build /go/src/ubiquitous-payment/exec /ubiquitous-payment/exec
RUN apk update && apk add gcompat && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY ./conf/certs/pem/* /usr/local/share/ca-certificates
RUN update-ca-certificates
EXPOSE 443
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/sh", "-c" , "cd ubiquitous-payment && ./exec"]