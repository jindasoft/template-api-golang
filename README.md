# template-api-golang

## rename file

## change config

rename exm_example ==> xxx_service
change kv_name - template-api-dev-vault
change port

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
make init
```
