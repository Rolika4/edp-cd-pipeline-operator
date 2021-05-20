module github.com/epam/edp-cd-pipeline-operator/v2

go 1.14

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/kubernetes-incubator/reference-docs => github.com/kubernetes-sigs/reference-docs v0.0.0-20170929004150-fcf65347b256
	github.com/markbates/inflect => github.com/markbates/inflect v1.0.4
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20210226220824-aa7126864d82
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => k8s.io/api v0.20.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.2
	github.com/epam/edp-cd-pipeline-operator/v2 => github.com/epam/edp-cd-pipeline-operator/v2 v2.3.0-58.0.20210519064332-021c2b1539f2
	github.com/epam/edp-codebase-operator/v2 => github.com/epam/edp-codebase-operator/v2 v2.3.0-95.0.20210520182756-7a216ea8871c
	github.com/epam/edp-gerrit-operator/v2 => github.com/epam/edp-gerrit-operator/v2 v2.3.0-73.0.20210505081511-94421d72acf6
	github.com/epam/edp-keycloak-operator => github.com/epam/edp-keycloak-operator v1.3.0-alpha-81.0.20210427070516-9b6232f72684
)

require (
	github.com/bndr/gojenkins v0.2.1-0.20181125150310-de43c03cf849
	github.com/epam/edp-codebase-operator/v2 v2.3.0-95.0.20210520182756-7a216ea8871c
	github.com/epam/edp-component-operator v0.1.1-0.20210427065236-c7dce7f4ea2b
	github.com/epam/edp-jenkins-operator/v2 v2.3.0-130.0.20210520181332-d15199eed95d
	github.com/go-logr/logr v0.4.0
	github.com/go-openapi/spec v0.19.5
	github.com/loft-sh/kiosk v0.2.4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	k8s.io/api v0.21.0-rc.0
	k8s.io/apimachinery v0.21.0-rc.0
	k8s.io/client-go v0.20.2
	k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
	sigs.k8s.io/controller-runtime v0.8.3
)
