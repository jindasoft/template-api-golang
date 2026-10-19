# template-api-golang

## rename file

rename filename xxxxx with module name

## change config

rename exm_example ==> xxx_service
change kv_name ==> xxx-api-dev-vault
change port
setup secret.json

## replace module

template-api-golang ==> xxx-api-service
xxxxxs ==> modules
Xxxxx ==> Module
xxxxx ==> module

## change lib

```sh
go mod edit -droprequire=github.com/jindasoft/template-platform-go
go get github.com/xxxxx/template-platform-go@develop
```

rename ==> github.com/jindasoft/template-platform-go

## setup

```sh
make check
make init
```
