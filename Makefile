PHONY=test clean all

build: 
	go build -o ~/.local/bin/pingtray

kill:
	pgrep pingtray | sudo xargs kill 2> /dev/null

run: build
	~/.local/bin/pingtray &
