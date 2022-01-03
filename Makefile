build-example:
	docker build -t go-cgroups-reader-example -f ./example/Dockerfile .
run-example:
	docker run --rm --memory 10485760 --cpu-quota 800000 --cpu-period 100000 go-cgroups-reader-example
clean:
	docker rmi -f go-cgroups-reader-example