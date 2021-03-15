up:
up:
	docker-compose up -d
up-d:
	docker-compose -f docker-compose.dev.yml up -d
down-d:
	docker-compose -f docker-compose.dev.yml down
build:
	docker-compose build --no-cache --force-rm
build-d:
	docker-compose -f docker-compose.dev.yml build --no-cache --force-rm

# module list
m-list:
	go list -f '{{join .Deps "\n"}}' | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'