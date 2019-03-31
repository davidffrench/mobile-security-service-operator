APP_NAME = mobile-security-service-operator
ORG_NAME = aerogear
PKG = github.com/$(ORG_NAME)/$(APP_NAME)
APP_FILE=./cmd/manager/main.go
BIN_DIR := $(GOPATH)/bin
BINARY ?= mobile-security-service-operator
TAG= 0.1.0
DOCKER-ORG=aerogear
DOCKER-REPO=mobile-security-service-operator

# This follows the output format for goreleaser
BINARY_LINUX_64 = ./dist/linux_amd64/$(BINARY)

LDFLAGS=-ldflags "-w -s -X main.Version=${TAG}"

.PHONY: deploy
deploy:
	@echo Deploying Mobile Security Service Operator:
	kubectl create namespace mobile-security-service-operator
	kubectl create -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_crd.yaml
	kubectl create -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_crd.yaml
	kubectl create -f deploy/cluster_role.yaml
	kubectl create -f deploy/cluster_role_binding.yaml
	kubectl create -f deploy/role.yaml
	kubectl create -f deploy/role_binding.yaml
	kubectl create -f deploy/service_account.yaml
	kubectl create -f deploy/operator.yaml

.PHONY: undeploy
undeploy:
	@echo Undeploy Mobile Security Service Operator:
	- kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_crd.yaml
	- kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_crd.yaml
	- kubectl delete -f deploy/cluster_role.yaml
	- kubectl delete -f deploy/cluster_role_binding.yaml
	- kubectl delete -f deploy/service_account.yaml
	- kubectl delete -f deploy/operator.yaml
	- kubectl delete -f deploy/role.yaml
	- kubectl delete -f deploy/role_binding.yaml
	- kubectl delete namespace mobile-security-service-operator

.PHONY: deploy-app
deploy-app:
	@echo Deploying Mobile Security Service and Database into project:
	kubectl apply -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_cr.yaml
	kubectl apply -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_cr.yaml

.PHONY: deploy-app-only
deploy-app-only:
	@echo Deploying Mobile Security Service and Database into project:
	kubectl apply -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_cr.yaml

.PHONY: deploy-db-only
deploy-db-only:
	@echo Deploying Mobile Security Service and Database into project:
	kubectl apply -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_cr.yaml

.PHONY: undeploy-app
undeploy-app:
	@echo Undeploying Mobile Security Service and Database from the project:
	- kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_cr.yaml
	- kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_cr.yaml

.PHONY: undeploy-app-only
undeploy-app-only:
	@echo Undeploying Mobile Security Service and Database from the project:
	kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservice_cr.yaml

.PHONY: undeploy-db-only
undeploy-db-only:
	@echo Undeploying Mobile Security Service and Database from the project:
	kubectl delete -f deploy/crds/mobile-security-service_v1alpha1_mobilesecurityservicedb_cr.yaml

.PHONY: build
build:
	@echo Buinding operator with the tag $(TAG):
	operator-sdk build $(DOCKER-ORG)/$(DOCKER-REPO):$(TAG)

.PHONY: publish
publish:
	@echo Publishing operator in $(DOCKER-ORG)/$(DOCKER-REPO) with the tag $(TAG):
	docker push $(DOCKER-ORG)/$(DOCKER-REPO):$(TAG)

.PHONY: vet
vet:
	@echo go vet
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: fmt
fmt:
	@echo go fmt
	go fmt $$(go list ./... | grep -v /vendor/)