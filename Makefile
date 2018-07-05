.PHONY: clean all

BUILD_DIR=build

clean:
	rm -rf $(BUILD_DIR)/*

all: \
  $(BUILD_DIR)/hello \
  $(BUILD_DIR)/foo \
  $(BUILD_DIR)/bar

$(BUILD_DIR)/%: cmd/%.go
	go build -i -o $@ $^
