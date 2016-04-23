OS := linux darwin windows
ARCH := 386 amd64

.PHONY: build
build: pkg/ gox
	gox -os="$(OS)" -arch="$(ARCH)" -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

pkg/:
	@mkdir pkg

.PHONY: gox
gox:
ifeq (,$(shell, which gox))
else
	go get -v github.com/mitchellh/gox
endif
