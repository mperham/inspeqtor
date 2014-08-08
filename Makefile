NAME=inspeqtor
VERSION=0.0.1

# when fixing packaging bugs but not changing the binary, we increment this number
ITERATION=1
BASENAME=$(NAME)_$(VERSION)-$(ITERATION)

# Include the secret API key which is needed to upload releases to bintray
-include .local.sh

all: clean package

test:
	go test ./...

build: test
	GOOS=linux GOARCH=amd64 go build -o $(NAME) cmd/main.go

clean:
	rm -f main $(NAME)
	rm -rf output
	mkdir output

package: build_deb build_rpm

build_rpm: build
	# gem install fpm
	# brew install rpm
	fpm -f -s dir -t rpm -n $(NAME) -v $(VERSION) -p output \
		--rpm-compression bzip2 --rpm-os linux -a x86_64 $(NAME)

build_deb: build
	# gem install fpm
	fpm -f -s dir -t deb -n $(NAME) -v $(VERSION) -p output \
		--deb-priority optional --category admin \
		--deb-compression bzip2 --after-install packaging/postinst.sh \
	 	--before-remove packaging/prerm.sh --after-remove packaging/postrm.sh \
		--url http://contribsys.com/inspeqtor --description "Modern service monitoring" \
		-m "Mike Perham <oss@contribsys.com>" --iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -d "runit" -a amd64 $(NAME)

upload: clean package
	curl \
		-T output/$(BASENAME)_amd64.deb \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases/$(NAME)/${VERSION}/$(BASENAME)_amd64.deb;publish=1"
	curl \
		-T output/$(NAME)_$(VERSION)-$(ITERATION).x86_64.rpm \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases/$(NAME)/${VERSION}/$(BASENAME).x86_64.rpm;publish=1"

.PHONY: all clean test build package upload
