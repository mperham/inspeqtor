NAME=inspeqtor
VERSION=0.8.1

# when fixing packaging bugs but not changing the binary, we increment ITERATION
ITERATION=1
BASENAME=$(NAME)_$(VERSION)-$(ITERATION)

# contains various secret or machine-specific variables.
# DEB_PRODUCTION: hostname of a debian-based upstart machine (e.g. Ubuntu {12,14}.04 LTS)
# RPM_PRODUCTION: hostname of a redhat-based systemd machine (e.g. CentOS 7)
-include .local.sh

# TODO I'd love some help making this a proper Makefile
# with real file dependencies.

all: test

prepare:
	#wget https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
	#sudo tar -C /usr/local -xzf go1.3.3.linux-amd64.tar.gz
	go get github.com/stretchr/testify/...
	go get github.com/jteeuwen/go-bindata/...
	go get code.google.com/p/gocc/...
	# needed for `make fmt`
	go get golang.org/x/tools/cmd/goimports
	@echo Now you should be ready to run "make"
	# To cross-compile from OSX to Linux, you need to run this:
	# cd $GOROOT/src && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

test:
	@go generate
	@go test -parallel 4 ./... | grep -v "no test files"

# gocc produces ill-formated code, clean it up with fmt
parsers: gocc fmt

# Generate parser and token package for conf/global/format.bnf and
# conf/inq/format.bnf
gocc:
	cd $(shell pwd)/conf/global && $(GOPATH)/bin/gocc format.bnf
	cd $(shell pwd)/conf/inq && $(GOPATH)/bin/gocc format.bnf

build: test
	@GOOS=linux GOARCH=amd64 go build -o inspeqtor cmd/main.go

# goimports produces slightly different formatted code from go fmt
fmt:
	find . -name "*.go" -exec goimports -w {} \;

lint:
	go vet ./...
	golint ./... | grep -v exported

clean:
	rm -f main inspeqtor templates.go
	rm -rf packaging/output
	mkdir -p packaging/output/upstart
	mkdir -p packaging/output/systemd

real:
	# Place real configuration with passwords, etc in "realtest".
	# git will ignore that directory and you can integration test
	# Inspeqtor on your local machine just by running "make real"
	GOMAXPROCS=4 go run -race cmd/main.go -l debug -s i.sock -c realtest

package: clean version_check build_deb build_rpm

version_check:
	@grep -q $(VERSION) inspeqtor.go || (echo VERSIONS OUT OF SYNC && false)

purge_deb:
	ssh -t $(DEB_PRODUCTION) 'sudo apt-get purge -y $(NAME) && sudo rm -f /etc/inspeqtor' || true

purge_rpm:
	ssh -t $(RPM_PRODUCTION) 'sudo rpm -e $(NAME) && sudo rm -f /etc/inspeqtor' || true

deploy_deb: clean build_deb purge_deb
	scp packaging/output/upstart/*.deb $(DEB_PRODUCTION):~
	ssh $(DEB_PRODUCTION) 'sudo rm -f /etc/inspeqtor && sudo dpkg -i $(NAME)_$(VERSION)-$(ITERATION)_amd64.deb && sudo ./fix && sudo restart inspeqtor || true'

deploy_rpm: clean build_rpm purge_rpm
	scp packaging/output/systemd/*.rpm $(RPM_PRODUCTION):~
	ssh -t $(RPM_PRODUCTION) 'sudo rm -f /etc/inspeqtor && sudo yum install -q -y $(NAME)-$(VERSION)-$(ITERATION).x86_64.rpm && sudo ./fix && sudo systemctl restart inspeqtor'

update_deb: clean build_deb
	scp packaging/output/upstart/*.deb $(DEB_PRODUCTION):~
	ssh $(DEB_PRODUCTION) 'sudo dpkg -i $(NAME)_$(VERSION)-$(ITERATION)_amd64.deb'

update_rpm: clean build_rpm
	scp packaging/output/systemd/*.rpm $(RPM_PRODUCTION):~
	ssh -t $(RPM_PRODUCTION) 'sudo yum install -q -y $(NAME)-$(VERSION)-$(ITERATION).x86_64.rpm'

deploy: deploy_deb deploy_rpm
purge: purge_deb purge_rpm

cover:
	go test -cover -coverprofile cover.out
	go tool cover -html=cover.out

build_rpm: build_rpm_upstart build_rpm_systemd
build_deb: build_deb_upstart

build_rpm_upstart: build
	# gem install fpm
	# brew install rpm
	fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -p packaging/output/upstart \
		--rpm-compression bzip2 --rpm-os linux \
	 	--after-install packaging/scripts/postinst.rpm.upstart \
	 	--before-remove packaging/scripts/prerm.rpm.upstart \
		--after-remove packaging/scripts/postrm.rpm.upstart \
		--url http://contribsys.com/inspeqtor \
		--description "Application infrastructure monitoring" \
		-m "Contributed Systems LLC <oss@contribsys.com>" \
		--iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -a amd64 \
		inspeqtor=/usr/bin/inspeqtor \
		packaging/root/=/

build_rpm_systemd: build
	# gem install fpm
	# brew install rpm
	fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -p packaging/output/systemd \
		--rpm-compression bzip2 --rpm-os linux \
	 	--after-install packaging/scripts/postinst.rpm.systemd \
	 	--before-remove packaging/scripts/prerm.rpm.systemd \
		--after-remove packaging/scripts/postrm.rpm.systemd \
		--url http://contribsys.com/inspeqtor \
		--description "Application infrastructure monitoring" \
		-m "Contributed Systems LLC <oss@contribsys.com>" \
		--iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -a amd64 \
		inspeqtor=/usr/bin/inspeqtor \
		packaging/root/=/

# TODO build_deb_systemd
build_deb_upstart: build
	# gem install fpm
	fpm -s dir -t deb -n $(NAME) -v $(VERSION) -p packaging/output/upstart \
		--deb-priority optional --category admin \
		--deb-compression bzip2 \
	 	--after-install packaging/scripts/postinst.deb.upstart \
	 	--before-remove packaging/scripts/prerm.deb.upstart \
		--after-remove packaging/scripts/postrm.deb.upstart \
		--url http://contribsys.com/inspeqtor \
		--description "Application infrastructure monitoring" \
		-m "Contributed Systems LLC <oss@contribsys.com>" \
		--iteration $(ITERATION) --license "GPL 3.0" \
		--vendor "Contributed Systems" -a amd64 \
		inspeqtor=/usr/bin/inspeqtor \
		packaging/root/=/

tag:
	git tag v$(VERSION)-$(ITERATION)
	git push --tags

upload:	package tag
	package_cloud push contribsys/inspeqtor/ubuntu/precise packaging/output/upstart/$(NAME)_$(VERSION)-$(ITERATION)_amd64.deb
	package_cloud push contribsys/inspeqtor/ubuntu/trusty packaging/output/upstart/$(NAME)_$(VERSION)-$(ITERATION)_amd64.deb
	package_cloud push contribsys/inspeqtor/el/7 packaging/output/systemd/$(NAME)-$(VERSION)-$(ITERATION).x86_64.rpm
	package_cloud push contribsys/inspeqtor/el/6 packaging/output/upstart/$(NAME)-$(VERSION)-$(ITERATION).x86_64.rpm

.PHONY: all clean test build package upload
