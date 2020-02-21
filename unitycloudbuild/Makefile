.PHONY: build clean deploy apigen

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/full full/*.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: build
	sls deploy --verbose

logs:
	sls logs -f unity