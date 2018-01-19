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
WORKSPACE_PATH=.

.PHONY : all 
all: release
release: setup format lint build_release release-install 
debug: setup format lint build_debug debug-install
dock_all: setup format lint  build_release release-install post_install

setup:
	rm -rf bin
	mkdir -p bin

lint:
	$(LINT) ${WORKSPACE_PATH}/auth
	$(LINT) ${WORKSPACE_PATH}/rr
	$(LINT) ${WORKSPACE_PATH}/item

go_install:
	$(INST) ${WORKSPACE_PATH}/auth
	$(INST) ${WORKSPACE_PATH}/rr
	$(INST) ${WORKSPACE_PATH}/item

fix: 
	$(FIX) ${WORKSPACE_PATH}/auth
	$(FIX) ${WORKSPACE_PATH}/rr
	$(FIX) ${WORKSPACE_PATH}/item

vet: 
	$(VET) ${WORKSPACE_PATH}/auth
	$(VET) ${WORKSPACE_PATH}/rr
	$(VET) ${WORKSPACE_PATH}/item

format: 
	$(FMT) ${WORKSPACE_PATH}/auth
	$(FMT) ${WORKSPACE_PATH}/rr
	$(FMT) ${WORKSPACE_PATH}/item

################ RELEASE VERSION #####################
build_release: 
	@echo "=== building release ==="
	mkdir -p bin 
	$(BUILD) $(RELEASE_FLAGS) -o bin/auth.bin ${WORKSPACE_PATH}/auth
	$(BUILD) $(RELEASE_FLAGS) -o bin/rr.bin ${WORKSPACE_PATH}/rr
	$(BUILD) $(RELEASE_FLAGS) -o bin/item.bin ${WORKSPACE_PATH}/item

################ DEBUG VERSION #####################

build_debug: 
	@echo "=== building debug ==="
	mkdir -p bin 
	$(BUILD) $(DEBUG_FLAGS) -o bin/auth.bin ${WORKSPACE_PATH}/auth
	$(BUILD) $(DEBUG_FLAGS) -o bin/rr.bin ${WORKSPACE_PATH}/rr
	$(BUILD) $(DEBUG_FLAGS) -o bin/item.bin ${WORKSPACE_PATH}/item

post_install:


clean: 
	@echo "=== clean ==="
	$(RM) bin *.bin bin-release bin-debug

debug-install: 
	@echo "=== debug-install ==="
	mv ./bin ./bin-debug

release-install: 
	@echo "=== release install ==="
	mv ./bin ./bin-release
