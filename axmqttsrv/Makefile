CC=gcc
CFLAGS=-I.
LDFLAGS=-L.
SRC=main.o
EXE=axmqttsrv
HEAD=libaxweb.h

ANDROID_SDK=$(HOME)/Android/Sdk
NDK_BIN=$(ANDROID_SDK)/ndk/25.1.8937393/toolchains/llvm/prebuilt/linux-x86_64/bin

# all: 	
# 	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=$(CC) go build -o $(EXE) -buildmode=c-shared 

# $(EXE): $(SRC)
# 	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ $^

# %.o: %.c
# 	$(CC) $(CFLAGS) $(LDFLAGS) -c $<

android:
	CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 CC=$(NDK_BIN)/armv7a-linux-androideabi21-clang go build -o $(EXE) -buildmode=c-shared 
	$(NDK_BIN)/llvm-strip $(EXE)

android64:
	CGO_ENABLED=1 GOOS=android GOARCH=arm64 CC=$(NDK_BIN)/aarch64-linux-android21-clang go build -o $(EXE)		
	$(NDK_BIN)/llvm-strip $(EXE)

linux64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc go build -o $(EXE) -buildmode=c-shared 	

pydist:
	@mkdir -p $@
	@cp -a $(EXE) $@

clean:
	rm  -rf *.o $(EXE)
	rm -rf $(HEAD)
	rm -rf node/node_modules
