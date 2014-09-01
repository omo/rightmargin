#
#
#

# For coding workflow
run: web
	docker run --rm -p 8080:8080 -v "${PWD}/web:/usr/local/web:ro" rightmargin

web: web/rightmargin

web/rightmargin:
	cd web && go build

# For operating image
run-image:
	docker run --rm -p 8080:8080 rightmargin

build-image:
	-docker rmi rightmargin
	docker build --rm=true -no-cache -t rightmargin web

.PHONY: run build-image run-image
