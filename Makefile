DAY ?= 1
PART ?= 1

build:
	go build -o bin/aoc .

run:
	go run main.go -day $(DAY) -part $(PART)

run-binary: build
	bin/aoc -day $(DAY) -part $(PART)

test:
	go test ./solutions/...

gen:
	go run cmd/gen/main.go -day $(DAY)

clean:
	rm -f bin/aoc
