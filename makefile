BUILDPATH = $(CURDIR)
GO = $(shell which go)
GOBUILD = $(GO) build
GOCLEAN = $(GO) clean
GOGET = $(GO) get
GORUN = $(GO) run

MAINFILE = Apps

makedir:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi

# @if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

clean:
	@rm -rf $(BUILDPATH)/bin/$(MAINFILE)
# @rm -rf $(BUILDPATH)/pkg

goget:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) github.com/jmoiron/sqlx

gobuild:
	$(GOBUILD) -o bin/Apps cmd/real/main.go

gorun:
	@./bin/$(MAINFILE)

gorunapp: makedir gobuild gorun

test:
	@echo $(BUILDPATH)