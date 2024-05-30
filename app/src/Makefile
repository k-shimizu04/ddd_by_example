current := $(current_env)
dockerfile_for_dev := build/docker/dev/docker-compose.yml
dockerfile_for_prd := build/docker/prd/docker-compose.yml
dockerfile := # Don't set this value. This is set automatically

all:
ifndef current
	@echo "Command argument(current_env) must be set when you execute make-command."
	@exit 1
endif

ifdef current
	@echo "#### Command is being prepared... ####"

ifeq ($(current), dev)
	@echo "Start dev mode based on your argument(current_env)"
	$(eval dockerfile = $(dockerfile_for_dev))
else
ifeq ($(current), prd)
	@echo "Start prd mode based on your argument(current_env)"
	$(eval dockerfile = $(dockerfile_for_prd))
else
	@echo "Couldn't continue. please, check your argument. Useable ethier 'dev' or 'prd'"
	@exit 1
endif
endif

	@echo "#### All complete. Ready to go!! ####"
endif

set_up:
	if [ ! -d bin ]; then mkdir bin; fi
	if [ ! -d pkg ]; then mkdir pkg; fi

build-backend: all
	docker-compose -f ${dockerfile} build backend
build-mysql: all
	docker-compose -f ${dockerfile} build mysql

up-backend: all
	docker-compose -f ${dockerfile} up -d backend
up-mysql: all
	docker-compose -f ${dockerfile} up -d mysql

down-backend: all
	docker-compose -f ${dockerfile} down backend
down-mysql: all
	docker-compose -f ${dockerfile} down mysql

wire: all
	docker-compose -f ${dockerfile} run -w /go/src/infrastructure --rm backend wire

go-tidy: all
	docker-compose -f ${dockerfile} run -w /go/src --rm backend go mod tidy