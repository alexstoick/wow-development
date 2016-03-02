
all: builddocker copyfile

builddocker:
	docker build -t wow_build:latest .

ID=$(shell docker ps -q)
copyfile:
	docker cp $(ID):/go/src/github.com/alexstoick/wow/datafetch/datafetch datafetch/datafetch
