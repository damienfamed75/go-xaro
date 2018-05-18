output:
	go build -o bin/main.exe

r:
	go build -o bin/main.exe
	./bin/main.exe

run:
	go build -o bin/main.exe
	./bin/main.exe

clean:
	-rm -f *.o bin/* output