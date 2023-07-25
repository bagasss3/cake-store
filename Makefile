SHELL:=/bin/bash

src/model/mock/mock_cake_service.go:
	mockgen -destination=src/model/mock/mock_cake_service.go -package=mock cake-store/src/model CakeService
src/model/mock/mock_cake_repository.go:
	mockgen -destination=src/model/mock/mock_cake_repository.go -package=mock cake-store/src/model CakeRepository

mockgen: src/model/mock/mock_cake_service.go \
	src/model/mock/mock_cake_repository.go \

clean:
	rm -v src/model/mock/mock_*.go