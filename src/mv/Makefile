RM=rm -rf
FMT=go fmt 
LINT=golint
GET=go get 
BUILD=go build
VET=go vet 
FIX=go fix 
INST=go install
DEBUG_FLAGS=-gcflags "-N -l"
RELEASE_FLAGS=-ldflags "-w"

.PHONY : all 
all: release
release: setup format lint build_release release-install
debug: setup format lint build_debug debug-install
dock_all: setup format lint  build_release release-install post_install

setup:
	rm -rf bin
	mkdir -p bin

lint:
	$(LINT) mv/auth
	$(LINT) mv/rr
	$(LINT) mv/item

go_install:
	$(INST) mv/auth
	$(INST) mv/rr
	$(INST) mv/item

fix: 
	$(FIX) mv/auth
	$(FIX) mv/rr
	$(FIX) mv/item

vet: 
	$(VET) mv/auth
	$(VET) mv/rr
	$(VET) mv/item

format: 
	$(FMT) mv/auth
	$(FMT) mv/rr
	$(FMT) mv/item

################ RELEASE VERSION #####################
build_release: 
	mkdir -p bin 
	$(BUILD) $(RELEASE_FLAGS) -o bin/mv/auth.bin mv/auth
	$(BUILD) $(RELEASE_FLAGS) -o bin/mv/rr.bin mv/rr
	$(BUILD) $(RELEASE_FLAGS) -o bin/mv/item.bin mv/item

################ DEBUG VERSION #####################

build_debug: 
	mkdir -p bin 
	$(BUILD) $(DEBUG_FLAGS) -o bin/mv/auth.bin mv/auth
	$(BUILD) $(DEBUG_FLAGS) -o bin/mv/rr.bin mv/rr
	$(BUILD) $(DEBUG_FLAGS) -o bin/mv/item.bin mv/item

post_install:

clean: 
	$(RM) bin *.bin

debug-install: 
	mv bin/mv bin/debug-mv

release-install: 
	mv bin/mv bin/release-mv
