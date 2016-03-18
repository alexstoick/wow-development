
all: builddocker rundocker mv_fetch mv_web killdocker

builddocker:
	docker build -t wow_build:latest .

ID=$(shell docker images -q wow_build)

rundocker:
	docker run -d --name wow_build -it $(ID) bash

mv_fetch:
	docker cp wow_build:/go/src/github.com/alexstoick/wow/datafetch/datafetch datafetch/datafetch
mv_web:
	docker cp wow_build:/go/src/github.com/alexstoick/wow/web/web web/web

killdocker:
	docker stop wow_build
	docker rm wow_build
