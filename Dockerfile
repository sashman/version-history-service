FROM centos:6
MAINTAINER Oleksandr Stasyk <oleksandr.stasyk@logicnow.com>

RUN	yum -y update && yum clean all \
	&& yum -y install http://www.percona.com/downloads/percona-release/redhat/0.1-3/percona-release-0.1-3.noarch.rpm \
	&& yum -y install Percona-Server-server-56.x86_64 \
	&& yum -y install Percona-Server-client-56.x86_64 \
	&& yum -y install Percona-Server-devel-56.x86_64 \
	&& yum -y install Percona-Server-56-debuginfo.x86_64 \
	&& yum -y install Percona-Server-shared-56.x86_64 \
	&& yum -y install Percona-Server-shared-compat.x86_64

RUN	yum -y update && yum clean all \
	&& yum -y install wget curl tar rpm bzip2 \
	&& rpm -ivh http://dl.iuscommunity.org/pub/ius/stable/CentOS/6/x86_64/epel-release-6-5.noarch.rpm \
	&& rpm -ivh http://dl.iuscommunity.org/pub/ius/stable/CentOS/6/x86_64/ius-release-1.0-11.ius.centos6.noarch.rpm \
	&& yum -y groupinstall 'Development Tools' \
	&& yum -y install bzip2-devel libcurl-devel t1lib-devel mcrypt libmcrypt libmcrypt-devel \
	&& yum -y install openssl openssl-devel \
	&& yum -y install libxml2 libxml2-devel libtool-ltdl-devel \
	&& yum -y install libjpeg-turbo-devel libpng-devel libXpm-devel freetype-devel t1lib-devel \
	&& yum -y install gmp-devel mcrypt libmcrypt libmcrypt-devel libtidy-devel tidy bison libtool-ltdl-devel \
	&& yum -y install autoconf213 \
	&& yum -y install unixODBC unixODBC-devel libsodium libsodium-devel \
	&& yum -y install sqlite-devel xz-libs 

RUN	yum -y install http://ftp.riken.jp/Linux/fedora/epel/6/i386/epel-release-6-8.noarch.rpm

RUN rpm -i 'http://pkgs.repoforge.org/rpmforge-release/rpmforge-release-0.5.3-1.el6.rf.x86_64.rpm'
RUN rpm --import http://apt.sw.be/RPM-GPG-KEY.dag.txt
RUN yum -y --enablerepo=rpmforge-extras update git

ENV GO_VERSION 1.5
ENV GO_WRAPPER_COMMIT 6ea1f29b1fe7e6b0b8eb89493ed5e06bac454654

RUN curl -sSL https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz \
    | tar -v -C /usr/local -xz

RUN mkdir -p /Application
ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /Application
ENV GOBIN $GOPATH/bin

RUN mkdir -p /go/src/app /go/bin && chmod -R 777 /go

RUN curl https://raw.githubusercontent.com/docker-library/golang/${GO_WRAPPER_COMMIT}/1.5/go-wrapper \
    -o /usr/local/bin/go-wrapper \
    && chmod 755 /usr/local/bin/go-wrapper

RUN yum -y install epel-release; yum clean all
RUN yum -y install mongodb-server; yum clean all
RUN mkdir -p /data/db

COPY	Application /Application
RUN	cd /Application && go get . 
RUN	cd /Application && go build 

EXPOSE 8080
