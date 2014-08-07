NAME=inspeqtor
VERSION=0.0.1
ARCH=amd64

# Include the secret API key which is needed to upload releases to bintray
-include .local.sh

all: clean package

test:
	go test ./...

build: test
	GOOS=linux GOARCH=$(ARCH) go build -o $(NAME) cmd/main.go

clean:
	rm -f main $(NAME)
	rm -rf output
	mkdir output

package: build
	fpm -f -s dir -t deb -n $(NAME) -v $(VERSION) -p output $(NAME)

upload: clean package
	curl \
		-T output/$(NAME)_$(VERSION)_$(ARCH).deb \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases/inspeqtor/${VERSION}/$(NAME)_$(VERSION)_$(ARCH).deb"

.PHONY: all clean test build package upload
