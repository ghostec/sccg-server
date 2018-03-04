MY_IP=`ifconfig | grep --color=none -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep --color=none -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1`
project=$(shell basename $(PWD))
project_sanitized=$(shell echo $(project) | sed -e "s/\-//")
pg_dep=$(project_sanitized)_postgres_1
test_packages=`find . -type f -name "*.go" ! \( -path "*vendor*" \) | sed -En "s/([^\.])\/.*/\1/p" | uniq`
database=postgres://postgres:$(project)@localhost:5432/$(project)?sslmode=disable

setup: setup-tests setup-project

setup-project:
	@go get -u github.com/golang/dep/...
	@dep ensure

setup-tests:
	@go get github.com/onsi/ginkgo/ginkgo
	@go get github.com/onsi/gomega/...

deps:
	@mkdir -p docker_data && docker-compose up -d postgres
	@until docker exec $(pg_dep) pg_isready; do echo 'Waiting Postgres...' && sleep 1; done
	@docker exec $(pg_dep) createuser -s -U postgres $(project) 2>/dev/null || true
	@docker exec $(pg_dep) createdb -U $(project) $(project) 2>/dev/null || true

deps-test: deps

stop-deps:
	@docker-compose down

stop-deps-test: stop-deps

build:
	@mkdir -p bin && go build -o ./bin/$(project) .

build-docker:
	@docker build -t $(project) .

run:
	@go run main.go start-api

run-docker:
	@docker-compose up -d

stop-docker:
	@docker-compose down

setup-migrate:
	@go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
	@go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

migrate:
	@migrate -path migrations -database $(database) up

migrate-test:
	@migrate -database $(database)-test up

drop:
	@migrate -path migrations -database $(database) drop

drop-test:
	@migrate -database $(database)-test drop

reset-db: drop migrate

test: deps-test unit stop-deps-test

test-fast: unit

unit: unit-board unit-run

unit-run:
	@ginkgo -tags unit -cover -r -randomizeAllSpecs -randomizeSuites -skipMeasurements $(test_packages)

unit-board:
	@echo
	@echo "\033[1;34m=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\033[0m"
	@echo "\033[1;34m=         Unit Tests         -\033[0m"
	@echo "\033[1;34m=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-\033[0m"
