all: builddocker rundocker mv_fetch mv_web killdocker

builddocker:
	docker build -t wow_build:latest .

ID=$(shell docker images -q wow_build)

rundocker:
	docker run -d --name wow_build -it $(ID) bash

mv_fetch:
	docker cp wow_build:/go/src/github.com/alexstoick/wow/datafetch/datafetch datafetch/datafetch
build_fetch:
	cd datafetch && docker build -t wow_datafetch:latest .
tag_and_push_fetch:
	docker tag wow_datafetch registry.management.stoica.xyz/wow_datafetch
	docker push registry.management.stoica.xyz/wow_datafetch

mv_web:
	docker cp wow_build:/go/src/github.com/alexstoick/wow/web/web web/web
build_web:
	cd web && docker build -t wow_web:latest .
tag_and_push_web:
	docker tag wow_web registry.management.stoica.xyz/wow_web
	docker push registry.management.stoica.xyz/wow_web

killdocker:
	docker stop wow_build
	docker rm wow_build

deploy_web: mv_web build_web tag_and_push_web
deploy_fetch: mv_fetch build_fetch tag_and_push_fetch

deploy: deploy_web deploy_fetch
