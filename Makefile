
.PHONY: all
all: clean wpw

wpw:
	@GOOS=linux GOARCH=amd64 go build  -o $@
	chmod +x $@

.PHONY: clean
clean:
	@-rm -rf wpw
	