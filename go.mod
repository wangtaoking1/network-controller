module github.com/wangtaoking1/network-controller

go 1.15

require (
	k8s.io/api v0.0.0-20200821172135-21b59c1ded36
	k8s.io/apimachinery v0.0.0-20200821171749-b63a0c883fbf
	k8s.io/client-go v0.0.0-20200821172742-3e55cca68bbc
	k8s.io/klog/v2 v2.2.0
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20200821172135-21b59c1ded36
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200821171749-b63a0c883fbf
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200821172742-3e55cca68bbc
)