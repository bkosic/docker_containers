#Download base image ubuntu 16.04
FROM ubuntu:16.04
MAINTAINER bkosic

#Update repository
RUN apt-get update && apt-get upgrade -y

#Install dependencies
RUN apt-get install wget checkinstall apt-utils dpkg -y

#Install PCRE
RUN cd /tmp \
	&& mkdir pcre \
	&& cd pcre \
	&& wget ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-8.41.tar.gz \
	&& tar -zxf pcre-8.41.tar.gz \
	&& cd pcre-8.41 \
	&& ./configure \
	&& make \
	&& make install

#Install ZLIB
RUN cd /tmp \
	&& mkdir zlib \
	&& cd zlib \
	&& wget http://zlib.net/zlib-1.2.11.tar.gz \
	&& tar -zxf zlib-1.2.11.tar.gz \
	&& cd zlib-1.2.11 \
	&& ./configure \
	&& make \
	&& make install

#Install OpenSSL
RUN apt-get install openssl -y
#RUN cd /tmp \
#	&& mkdir openssl \
#	&& cd openssl \
#	&& wget wget http://www.openssl.org/source/openssl-1.0.2k.tar.gz \
#	&& tar -zxf openssl-1.0.2k.tar.gz \
#	&& cd openssl-1.0.2k \
#	&& ./Configure darwin64-x86_64-cc --prefix=/usr \
#	&& make \
#	&& make install

# Install nginx
RUN cd /tmp \
	&& wget http://nginx.org/download/nginx-1.12.1.tar.gz \
	&& tar zxf nginx-1.12.1.tar.gz \
	&& cd nginx-1.12.1 \
	&& wget https://github.com/openresty/headers-more-nginx-module/archive/v0.33.tar.gz \
	&& tar zxf v0.33.tar.gz \
	&& ./configure --prefix=/opt/nginx --add-dynamic-module=/tmp/nginx-1.12.1/headers-more-nginx-module-0.33 \
	&& make \
	&& checkinstall \
	&& dpkg -i /tmp/nginx-1.12.1/nginx_1.12.1-1_amd64.deb 

# Start nginx
CMD ["/opt/nginx/sbin/nginx"]
EXPOSE 80
