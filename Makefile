BINARY=engine

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

engine:
	go build -o ${BINARY} ./*.go

run:
	docker-compose up --build -d

stop:
	docker-compose down

remove:
	docker-compose down -v

ps:
	docker-compose ps

dbclean:
	rm -rf ./docker/db/mysql/data/*

gswag:
	swag init -g src/infrastructure/router.go

test:
	go test -v ./tests/*/*_test.go

.PHONY: clean run stop remove gswag dbclean ps engine test
