BASE_FILE := batch-resizer.go
BIN := bin
OUTPUT := bin/batch-resizer
WINOUT32 := bin/batch-resizer-windows-32.exe
MACOUT32 := bin/batch-resizer-mac-32
LINOUT32 := bin/batch-resizer-lunix-32
WINOUT64 := bin/batch-resizer-windows-64.exe
MACOUT64 := bin/batch-resizer-mac-64
LINOUT64 := bin/batch-resizer-lunix-64

all: $(OUTPUT)

$(OUTPUT): $(BASE_FILE)
	go build -o $(OUTPUT) $(BASE_FILE)

$(BIN):
	mkdir bin

cross-win-32: $(BUILD_PATH)
	env GOOS="windows" GOARCH="386" go build -o $(WINOUT32) $(BASE_FILE)

cross-mac-32: $(BUILD_PATH)
	env GOOS="darwin" GOARCH="386" go build -o $(MACOUT32) $(BASE_FILE)

cross-lin-32: $(BUILD_PATH)
	env GOOS="linux" GOARCH="386" go build -o $(LINOUT32) $(BASE_FILE)

cross-win-64: $(BUILD_PATH)
	env GOOS="windows" GOARCH="amd64" go build -o $(WINOUT64) $(BASE_FILE)

cross-mac-64: $(BUILD_PATH)
	env GOOS="darwin" GOARCH="amd64" go build -o $(MACOUT64) $(BASE_FILE)

cross-lin-64: $(BUILD_PATH)
	env GOOS="linux" GOARCH="amd64" go build -o $(LINOUT64) $(BASE_FILE)

cross-all: cross-win-32 cross-mac-32 cross-lin-32 cross-win-64 cross-mac-64 cross-lin-64

clean:
	rm $(OUTPUT)

cross-clean:
	rm $(WINOUT)
	rm $(MACOUT)
	rm $(LINOUT)