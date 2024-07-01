run-k6:
	@docker container run --rm -i -v $$(pwd)/k6:/scripts -w /scripts --cap-add=SYS_ADMIN --network observabilidade grafana/k6:0.51.0 run --out influxdb=http://influxdb:8086/k6 index.js