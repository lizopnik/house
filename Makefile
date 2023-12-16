test:
	@echo 'Мы сделали Makefile'

up:
	sudo docker-compose up --build awesome-project

stop:
	sudo docker-compose stop