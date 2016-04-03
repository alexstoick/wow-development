all: builddocker rundocker mv_fetch mv_web killdocker

create_ca:
	sudo mkdir -p /etc/docker/certs.d/wow.stoica.xyz:5000
	sudo mv ca.crt /etc/docker/certs.d/wow.stoica.xyz:5000/ca.crt
	sudo cat /etc/docker/certs.d/wow.stoica.xyz:5000/ca.crt
	sudo service docker restart

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
	docker tag wow_datafetch wow.stoica.xyz:5000/wow_datafetch
	docker push wow.stoica.xyz:5000/wow_datafetch

mv_web:
	docker cp wow_build:/go/src/github.com/alexstoick/wow/web/web web/web
build_web:
	cd web && docker build -t wow_web:latest .
tag_and_push_web:
	docker tag wow_web wow.stoica.xyz:5000/wow_web
	docker push wow.stoica.xyz:5000/wow_web

killdocker:
	docker stop wow_build
	docker rm wow_build

deploy_web: mv_web build_web tag_and_push_web
deploy_fetch: mv_fetch build_fetch tag_and_push_fetch

deploy: deploy_web deploy_fetch
