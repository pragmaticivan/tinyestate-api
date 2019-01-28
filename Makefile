build:
	go build

dev-dependencies:
	./scripts/dev-dependencies.sh

list-packages:
	go list ./...

test:
	GO_ENV=test ./scripts/test.sh

coverage:
	./scripts/coverage.sh

coverage.html:
	./scripts/coverage.sh --html

coverage.coveralls:
	./scripts/coverage.sh --coveralls

watch:
	gin run main.go

db.create:
	soda create -a

db.migrate:
	soda migrate up

.PHONY: install watch test list-packages dev-dependencies db.create db.migrate coverage
