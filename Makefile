sources := $(shell find . -name '*.go')
exe = git-timer

.PHONY: clean lint check-format vet staticcheck

$(exe): $(sources)
	go build -v -o $(exe) .

clean:
	rm -vf $(exe)

lint: check-format vet staticcheck

check-format:
	! gofmt -l . | grep --color=never .

vet:
	go vet ./...

staticcheck:
	staticcheck -checks=all ./...
