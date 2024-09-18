SQLBOILER_SED_EXPRESSION := "s/{{GOPATH}}/$(subst /,\/,$(GOPATH))/g"


.PHONY: build
build:
	go build -o ./.bin/ts ./
