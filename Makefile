
all: builddocker rundocker mv_fetch mv_web killdocker

builddocker:
	docker build -t wow_build:latest .

ID=$(shell docker images -q wow_build)
rundocker:
	docker run -d -it $(ID) bash

RUN_ID=$(shell docker ps -q)

killdocker:
	docker stop $(RUN_ID)

mv_fetch:
	docker cp $(RUN_ID):/go/src/github.com/alexstoick/wow/datafetch/datafetch datafetch/datafetch
mv_web:
	docker cp $(RUN_ID):/go/src/github.com/alexstoick/wow/web/web web/web
