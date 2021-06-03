mode =?
LINUX_BUILD_CMD=CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo

.PHONY: all linux

darwin:
	go build .

clean:
	rm -f gocomposer*

linux:
	$(LINUX_BUILD_CMD) -o gocomposer-amd64 .