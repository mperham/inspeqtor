NAME=inspeqtor
VERSION=0.0.2

# when fixing packaging bugs but not changing the binary, we increment this number
ITERATION=1
BASENAME=$(NAME)_$(VERSION)-$(ITERATION)

# Include the secret API key which is needed to upload releases to bintray
# Also you can set PRODUCTION to a Debian hostname you want Inspeqtor deployed to.
-include .local.sh

# TODO I'd love some help making this a proper Makefile
# with real file dependencies.

all: test

prepare:
	#wget https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz
	#tar -C /usr/local -xzf go1.3.1.linux-amd64.tar.gz
	go get github.com/stretchr/testify/...
	go get github.com/jteeuwen/go-bindata/...
	@echo Now you should be ready to run "make"

test:
	@go-bindata -pkg inspeqtor -o templates.go templates/...
	@go test -parallel 4 ./... | grep -v "no test files"

build: test
	@GOOS=linux GOARCH=amd64 go build -o $(NAME) cmd/main.go

clean:
	rm -f main $(NAME)
	rm -rf packaging/output
	mkdir packaging/output

real:
	go run cmd/main.go -l debug -s i.sock -c realtest

# TODO add build_rpm when working
package: clean build_deb

deploy: clean build_deb
	scp packaging/output/$(BASENAME)_amd64.deb $(PRODUCTION):~
	ssh $(PRODUCTION) 'sudo dpkg -i $(BASENAME)_amd64.deb && sudo ./fix && sudo sv restart inspeqtor'

build_rpm: build
	# gem install fpm
	# brew install rpm
	fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -p packaging/output \
		--config-files /etc/$(NAME) --config-files /var/log/$(NAME) \
		--rpm-compression bzip2 --rpm-os linux -a x86_64 \
		$(NAME)=/usr/bin/$(NAME) \
		packaging/root/=/

build_deb: build
	# gem install fpm
	fpm -s dir -t deb -n $(NAME) -v $(VERSION) -p packaging/output \
		--deb-priority optional --category admin \
		--deb-compression bzip2 \
	 	--after-install packaging/debian/postinst \
	 	--before-remove packaging/debian/prerm \
		--after-remove packaging/debian/postrm \
		--url http://contribsys.com/$(NAME) \
		--description "Modern service and host monitoring" \
		-m "Contributed Systems LLC <oss@contribsys.com>" \
		--iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -d "runit" -a amd64 \
	 	$(NAME)=/usr/bin/$(NAME) \
		packaging/root/=/

upload:	clean package
	curl \
		-T packaging/output/$(BASENAME)_amd64.deb \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases-deb/$(NAME)/${VERSION}/$(BASENAME)_amd64.deb;publish=1"
	curl \
		-X POST -H "X-GPG-PASSPHRASE: $(PASSPHRASE)" \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/gpg/contribsys/releases-deb/$(NAME)/versions/${VERSION}"

.PHONY: all clean test build package upload
