#!/usr/bin/env bash
#Start cassandra  service and infinite loop
service cassandra start && while :; do echo 'Hit CTRL+C'; sleep 1; done
