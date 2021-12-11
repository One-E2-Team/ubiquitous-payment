# ubiquitous-payment

to build plugins, run ```docker build  --output temp --progress=plain -f GoPluginBuild.dockerfile .```
build psp ```docker build  -f GoBuild.dockerfile -t gotest . --build-arg ARG_TARGET_GO=psp --no-cache```
run it ```docker run -v C:\Users\igors\source\repos\ubiquitous-payment\temp:/ubiquitous-payment/temp gotest```


```curl
curl -v --digest -X POST host.docker.internal:18332/wallet/secondarytest -d "{\"jsonrpc\":\"2.0\",\"id\":\"0\",\"method\":\"sendtoaddress\",\"params\":[\"tb1qy7sxj94edsd4fcn2ckzk0taafqlqed80vawppu\"
,0.00005176,\"\",\"\",true,true,6,\"unset\",false]}" -H 'Content-Type:application/json' -H 'Authorization: Basic cm9vdDpyb290'
```

```ps
.\bitcoind.exe -bind="0.0.0.0:18333" -rest -rpcbind="0.0.0.0:18332" -chain=test -debug=rpc -server -listen -rpcuser=root -rpcpassword=root -rpcallowip="0.0.0.0/0"
```