.PHONY: wrk-test wrk-build wrk-run
wrk-test:
	@echo "Running wrk benchmark"
	make wkr-build
	make wkr-run

wrk-build:
	docker build -t wrk-benchmark .

wrk-run:
	docker run --rm --add-host=host.docker.internal:host-gateway wrk-benchmark http://host.docker.internal:8080/
	# You need to keep host.docker.internal as it is, it is a special DNS name that resolves to the host machine from a container