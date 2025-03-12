FROM williamyeh/wrk
# COPY add-config.lua /add-config.lua
COPY run-benchmark.sh /run-benchmark.sh
RUN chmod +x /run-benchmark.sh
ENTRYPOINT ["/run-benchmark.sh"]