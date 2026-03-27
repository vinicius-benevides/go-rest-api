.PHONY: test api mocks

GO := /usr/local/go/bin/go
MOCKGEN := $(shell /usr/local/go/bin/go env GOPATH)/bin/mockgen
export PATH := /usr/local/go/bin:$(PATH)

## Run all unit tests
test:
	$(GO) test ./...

## Start the API server
api:
	$(GO) run ./cmd/api

## Regenerate gomock mocks
mocks:
	$(MOCKGEN) -destination=test/mocks/ports_mock.go -package=mocks \
		github.com/vinicius-benevides/go-rest-api/src/domain/ports \
		EventRepository,RegistrationRepository,UserRepository

		$(MOCKGEN) -destination=test/mocks/services_mock.go -package=mocks \
			github.com/vinicius-benevides/go-rest-api/src/services \
			TokenService,UserService,EventService,RegistrationService
