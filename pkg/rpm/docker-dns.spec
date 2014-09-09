%global debug_package %{nil}

Name:		docker-dns
Version:	0.1
Release:	1%{?dist}
Summary:	Poor mans service discovery for docker.

Group:		Development/Tools
License:	MIT
URL:		https://github.com/svenwltr/docker-dns
Source0:	https://github.com/svenwltr/docker-dns/archive/%{name}-%{version}.tar.gz

BuildRequires:	golang

%description


%prep
%setup


%build


%install
rm -rf %{buildroot}

#install -dm 755 %{buildroot}%{_datadir}/%{name}
#install -pm 755 docker-dns %{buildroot}%{_datadir}/%{name}
#install -pm 755 README.md %{buildroot}%{_datadir}/%{name}
#install -pm 644 docker-dns.service %{buildroot}%{_datadir}/%{name}

install -dm 755 %{buildroot}/usr/bin
install -pm 755 docker-dns %{buildroot}/usr/bin/%{name}

install -dm 755 %{buildroot}/usr/lib/systemd/system
install -pm 644 docker-dns.service %{buildroot}/usr/lib/systemd/system/%{name}.service


%files
#%defattr(-,root,root,-)
#%{_datadir}/%{name}
#%doc README.md
/usr/bin/docker-dns
/usr/lib/systemd/system/docker-dns.service



%changelog

