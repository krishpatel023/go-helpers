#!/bin/sh
# If you want to run the benchmark with a config, use the following command
# wrk -t4 -c100 -d60s -s /add-config.lua "$@"

wrk -t4 -c100 -d1m "$@"