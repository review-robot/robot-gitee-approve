module github.com/opensourceways/robot-gitee-approve

go 1.15

replace (
	cloud.google.com/go => cloud.google.com/go v0.44.3
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
	k8s.io/api => k8s.io/api v0.17.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.3
	k8s.io/client-go => k8s.io/client-go v0.17.3
	k8s.io/code-generator => k8s.io/code-generator v0.17.3
)

require (
	github.com/opensourceways/community-robot-lib v0.0.0-20211220063904-5d625d7719ea
	github.com/opensourceways/go-gitee v0.0.0-20211230032551-d653a809e178
	github.com/opensourceways/repo-owners-cache v0.0.0-20211227074144-9ad8301da957
	github.com/sirupsen/logrus v1.8.1
	k8s.io/apimachinery v0.23.1
	k8s.io/test-infra v0.0.0-20200522021239-7ab687ff3213
)
