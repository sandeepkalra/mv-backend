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
MODULES= auth rr item
.PHONY : all 
all: release
release: setup format lint build_release release-install 
debug: setup format lint build_debug debug-install
dock_all: setup format lint  build_release release-install post_install

setup:
	rm -rf bin
	mkdir -p bin

lint:
	for mod in $(MODULES); do \
		$(LINT) ${WORKSPACE_PATH}/$$mod; \
	done

go_install:
	for mod in $(MODULES); do \
		$(INST) ${WORKSPACE_PATH}/$$mod;\
	done

fix:
	for mod in $(MODULES); do \
		$(FIX) ${WORKSPACE_PATH}/$$mod;\
	done

vet: 
	for mod in $(MODULES); do \
		$(VET) ${WORKSPACE_PATH}/$$mod;\
	done

format: 
	for mod in $(MODULES); do \
		$(FMT) ${WORKSPACE_PATH}/$$mod;\
	done

################ RELEASE VERSION #####################
build_release: 
	@echo "=== building release ==="
	mkdir -p bin 
	for mod in $(MODULES); do \
			$(BUILD) $(RELEASE_FLAGS) -o bin/$$mod.bin  ${WORKSPACE_PATH}/$$mod;\
	done

################ DEBUG VERSION #####################

build_debug: 
	@echo "=== building debug ==="
	mkdir -p bin 
	for mod in $(MODULES); do \
			$(BUILD) $(DEBUG_FLAGS) -o bin/$$mod.bin  ${WORKSPACE_PATH}/$$mod;\
	done

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
