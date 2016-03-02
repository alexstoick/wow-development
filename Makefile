
all: builddocker copyfile

builddocker:
	docker build -t wow_build:latest .

ID=$(shell docker images -q wow_build)
copyfile:
	docker cp $(ID):/go/src/github.com/alexstoick/wow/datafetch/datafetch datafetch/datafetch
