# ubiquitous-payment

to build plugins, run ```docker build  --output temp --progress=plain -f GoPluginBuild.dockerfile .```
build psp ```docker build  -f GoBuild.dockerfile -t gotest . --build-arg ARG_TARGET_GO=psp --no-cache```
run it ```docker run -v C:\Users\igors\source\repos\ubiquitous-payment\temp:/ubiquitous-payment/temp gotest```