#Download base image ubuntu 16.04
FROM ubuntu:16.04
MAINTAINER bkosic

#Update repository
RUN apt-get update && apt-get upgrade -y

#Install curl,python2.7 and OpenJDK 8
RUN apt-get install python2.7 openjdk-8-jre curl -y

# Install cassandra
RUN echo "deb http://www.apache.org/dist/cassandra/debian 311x main" | tee -a /etc/apt/sources.list.d/cassandra.sources.list \
    && curl https://www.apache.org/dist/cassandra/KEYS | apt-key add - \
    && apt-get update \
    && apt-get install cassandra -y \
    && sed  -i "/^[^#]*ulimit/ s/.*/#&/"  /etc/init.d/cassandra \
    && sed -i "s\Test Cluster\db-cassandra\g" /etc/cassandra/cassandra.yaml \
    && sed -i "s\rpc_address: localhost\rpc_address: eth0@if151\g" /etc/cassandra/cassandra.yaml \
    && sed -i "s\listen_address: localhost\rpc_address: eth0@if151\g" /etc/cassandra/cassandra.yaml

COPY start.sh start.sh
RUN chmod +x /start.sh
# Start cassandra
CMD ["/start.sh"]

# 7000: intra-node communication
# 7001: TLS intra-node communication
# 7199: JMX
# 9042: CQL
# 9160: thrift service
EXPOSE 7000 7001 7199 9042 9160
