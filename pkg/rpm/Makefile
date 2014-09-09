build:
	cd ../.. && make
	tar -cf docker-dns-0.1.tar docker-dns.spec
	tar \
		--append \
		--directory ../.. \
		--transform "s#^#docker-dns-0.1/#" \
		--file docker-dns-0.1.tar \
		README.md
	tar \
		--append \
		--directory ../.. \
		--transform "s#^#docker-dns-0.1/#" \
		--file docker-dns-0.1.tar \
		docker-dns
	tar \
		--append \
		--directory ../../init/systemd \
		--transform "s#^#docker-dns-0.1/#" \
		--file docker-dns-0.1.tar \
		docker-dns.service
	gzip -c docker-dns-0.1.tar  > docker-dns-0.1.tar.gz
	rpmbuild -ta docker-dns-0.1.tar.gz



