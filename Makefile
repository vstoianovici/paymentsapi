BUILD=go build -ldflags="-s -w"
BUILDPATH=./cmd

all: build

build:
	@cd $(BUILDPATH); $(BUILD) -v -o ./paymentsAPI

test:
	cd $(BUILDPATH); go test ./.. -v

clean:
	@rm -f $(BUILDPATH)/paymentsAPI

.PHONY: all build clean test



