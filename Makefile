.PHONY: explain
explain:

	### Welcome to Dynamic-X
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

start: ## To start the application in same port
	@echo "Find the current proccess running in TCP PORT: 8080 and Kill it."
	lsof -ti:8080 | xargs kill -9
	go run main.go

login-db: ## To login to PSQL
	psql gocontacts