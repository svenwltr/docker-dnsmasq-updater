
version=$(shell git describe --tags | tr '-' '.')
name=docker-dnsmasq-updater
base=$(shell readlink -f ../..)
build=$(shell readlink -f build)
srcpkg=$(name)-$(version).tar

build:
	mkdir -p $(build)
	cd $(base) && \
	git archive \
		--format tar \
		--prefix=$(name)-$(version)/ \
		HEAD \
		> $(build)/$(srcpkg)
	tar \
		--append \
		--file $(build)/$(srcpkg) \
		docker-dnsmasq-updater.spec
	# a other spec file causes problems with rpmbuild
	tar \
		--delete \
		--file $(build)/$(srcpkg) \
		$(name)-$(version)/pkg/rpm/docker-dnsmasq-updater.spec
	gzip -c $(build)/$(srcpkg) > $(build)/$(srcpkg).gz
	cd $(build) && rpmbuild \
		-vv \
		--clean \
		--define 'gitversion $(version)' \
		-ta "$(srcpkg).gz"

clean:
	rm -rf build


