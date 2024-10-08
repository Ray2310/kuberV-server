package k8s

import (
	"kubeimooc.com/service"
	"kubeimooc.com/validate"
)

// k8s 包下所有接口的集合
type ApiGroup struct {
	PodApi
	NamespaceApi
	NodeApi
	ConfigMapApi
	SecretApi
}

var podValidate = validate.ValidateGroupApp.PodValidate
var podService = service.ServiceGroupApp.PodServiceGroup.PodService
var nodeService = service.ServiceGroupApp.NodeServiceGroup.NodeService
var configMapService = service.ServiceGroupApp.ConfigMapServiceGroup.ConfigMapService
var secretService = service.ServiceGroupApp.SecretServiceGroup.SecretService
