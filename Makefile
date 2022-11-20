PKG=github.com/jiyee/cubox-cli
NAME=cubox-cli
BINDIR=Downloads
VERSION=$(shell git describe --tags || echo "unknown")
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "main.version=$(VERSION)" \
		-w -s -buildid='

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \

test:
	go test ./...

all: darwin-amd64 darwin-arm64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@
	
darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$(NAME)-$@

gz_releases=$(addsuffix .gz, $(PLATFORM_LIST))

$(gz_releases): %.gz : %
	chmod +x $(BINDIR)/$(NAME)-$(basename $@)
	gzip -f -S .gz $(BINDIR)/$(NAME)-$(basename $@)

all-arch: $(PLATFORM_LIST) 

releases: $(gz_releases) $(zip_releases)

clean:
	rm $(BINDIR)/*