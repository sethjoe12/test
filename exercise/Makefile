build:
	
	go build -o bin/exercise

run: build
	./bin/exercise
	
mysql-create-db:
	@docker-compose exec -T mysqldb mysql -u"root" -p"Sethjoe12345" -e "'test_demo'"

test:
	go test -v ./...	 
		   
