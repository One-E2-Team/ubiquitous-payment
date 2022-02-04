FROM golang AS build
COPY . src/ubiquitous-payment/
RUN cd src/ubiquitous-payment && \
    mkdir target &&  \
    go mod download &&  \
    go mod verify &&  \
    go build -buildmode=plugin -o target/paypal.so psp-plugins/paypal/main.go && \
    go build -buildmode=plugin -o target/bitcoin.so psp-plugins/bitcoin/main.go && \
    go build -buildmode=plugin -o target/bank.so psp-plugins/bank/main.go && \
    go build -buildmode=plugin -o target/qrcode.so psp-plugins/qrcode/main.go

FROM scratch AS export-stage
COPY --from=build /go/src/ubiquitous-payment/target /