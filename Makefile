NAME=inspeqtor
VERSION=0.0.1

# when fixing packaging bugs but not changing the binary, we increment this number
ITERATION=1
BASENAME=$(NAME)_$(VERSION)-$(ITERATION)

# Include the secret API key which is needed to upload releases to bintray
# Also you can set PRODUCTION to a Debian hostname you want Inspeqtor deployed to.
-include .local.sh

all: test

test:
	@go test -parallel 4 ./... | grep -v "no test files"

build: test
	@GOOS=linux GOARCH=amd64 go build -o $(NAME) cmd/main.go

clean:
	rm -f main $(NAME)
	rm -rf output
	mkdir output

package: clean build_deb build_rpm

deploy:
	scp output/$(BASENAME)_amd64.deb $(PRODUCTION):~
	ssh $(PRODUCTION) 'sudo dpkg -i $(BASENAME)_amd64.deb && sudo sv restart inspeqtor'

build_rpm: build
	# gem install fpm
	# brew install rpm
	fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -p output \
		--config-files /etc/$(NAME) --config-files /var/log/$(NAME) \
		--rpm-compression bzip2 --rpm-os linux -a x86_64 \
	 	$(NAME)=/usr/bin/$(NAME) \
		packaging/root/=/

build_deb: build
	# gem install fpm
	fpm -s dir -t deb -n $(NAME) -v $(VERSION) -p output \
		--deb-priority optional --category admin \
		--config-files /etc --config-files /var/log/$(NAME) \
		--deb-compression bzip2 --after-install packaging/postinst.sh \
	 	--before-remove packaging/prerm.sh --after-remove packaging/postrm.sh \
		--url http://contribsys.com/$(NAME) --description "Modern service monitoring" \
		-m "Mike Perham <oss@contribsys.com>" --iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -d "runit" -a amd64 \
	 	$(NAME)=/usr/bin/$(NAME) \
		packaging/root/=/

upload: clean package
	curl \
		-T output/$(BASENAME)_amd64.deb \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases/$(NAME)/${VERSION}/$(BASENAME)_amd64.deb;publish=1"
	curl \
		-T output/$(BASENAME).x86_64.rpm \
		-umperham:${BINTRAY_API_KEY} \
		"https://api.bintray.com/content/contribsys/releases/$(NAME)/${VERSION}/$(BASENAME).x86_64.rpm;publish=1"

.PHONY: all clean test build package upload
