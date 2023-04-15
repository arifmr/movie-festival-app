## run: run app on local machine
.PHONY: run
run:
	@echo "Running app on local machine"
	@APP_ENV="local" go run cmd/main.go