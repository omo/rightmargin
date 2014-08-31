#
#
#

# For coding workflow
run: web
	docker run --rm -P -v "${PWD}/web:/usr/local/web:ro" rightmargin

web: web/rightmargin

web/rightmargin:
	cd web && go build

# For operating image
run-image:
	docker run --rm -P rightmargin

build-image:
	docker build --rm --no-cache -t rightmargin web

.PHONY: run build-image run-image
