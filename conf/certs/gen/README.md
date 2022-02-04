```
openssl req -x509 -newkey rsa:4096 -sha256 -keyout key/pcc.key.pem -out pem/pcc.cert.pem -days 365 -nodes -config conf/pcc.conf -extensions req_ext

openssl x509 -outform der -in pem/pcc.cert.pem -out der/pcc.cert.crt
```