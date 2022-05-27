# kubectl-img
kubectl-img is a kubectl plugin that allows you to show kubernetes resource image.

## Install
Linux
```shell
export release=1.0.0
curl -L -o kubectl-img.tar.gz https://github.com/mark8s/kubectl-img/releases/download/${release}/kubectl-img-${release}.tar.gz
tar -xvf kubectl-img.tar.gz
cp kubectl-img /usr/local/bin/kubectl-img
```

## Usage
You don't need to setup anything on your cluster before using it, please don't use it already on a production system, just because this isn't yet 100% ready.

```shell
$ kubectl-img 
kubectl-img show the Kubernetes workload image.

Usage:
  kubectl-img [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  image       kubernetes workload image

Flags:
  -h, --help     help for kubectl-img
  -t, --toggle   Help message for toggle

Use "kubectl-img [command] --help" for more information about a command.
```

### View Deployments Images
```shell
# 查看所有namespace的deployment的image
kubectl-img image -t deploy -A
# 查看指定namespace的deployment的image
kubectl-img image -t deploy -n default
```
### View format
Table display is used by default,you can use `-f json` set format=json
```shell
# Table display is used by default
$ kubectl-img image -f json
[
   {
      "IMAGE": "nginx",
      "NAME": "nginx",
      "NAMESPACE": "default",
      "TYPE": "pod"
   }
]
```

### Example
1. 显示所有namespace的pod images
```shell
$ kubectl-img image -A
+----------------+------+---------------------------------------+---------------------------------------------------------------------+
|   NAMESPACE    | TYPE |                 NAME                  |                                IMAGE                                |
+----------------+------+---------------------------------------+---------------------------------------------------------------------+
|    default     | pod  |                 nginx                 |                                nginx                                |
| istio-operator | pod  |    istio-operator-64b44774cf-sl75w    |                   docker.io/istio/operator:1.13.0                   |
|  istio-system  | pod  | istio-egressgateway-6cf5fb4756-52scr  |                   docker.io/istio/proxyv2:1.13.0                    |
|  istio-system  | pod  | istio-ingressgateway-dc9c8f588-jc9kw  |                   docker.io/istio/proxyv2:1.13.0                    |
|  istio-system  | pod  |        istiod-6874697f7d-fpzrk        |                    docker.io/istio/pilot:1.13.0                     |
|  kube-system   | pod  |       coredns-558bd4d5db-4qgcs        |                  k8s.gcr.io/coredns/coredns:v1.8.0                  |
|  kube-system   | pod  |       coredns-558bd4d5db-z89dd        |                  k8s.gcr.io/coredns/coredns:v1.8.0                  |
|  kube-system   | pod  |          etcd-biz-master-48           |                      k8s.gcr.io/etcd:3.4.13-0                       |
|  kube-system   | pod  |     kube-apiserver-biz-master-48      |                  k8s.gcr.io/kube-apiserver:v1.21.0                  |
|  kube-system   | pod  | kube-controller-manager-biz-master-48 |             k8s.gcr.io/kube-controller-manager:v1.21.0              |
|  kube-system   | pod  |           kube-proxy-cgr4w            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |           kube-proxy-dgmm7            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |           kube-proxy-jmkb6            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |     kube-scheduler-biz-master-48      |                  k8s.gcr.io/kube-scheduler:v1.21.0                  |
|  kube-system   | pod  |            weave-net-hxvkn            |                docker.io/weaveworks/weave-kube:2.8.1                |
|  kube-system   | pod  |            weave-net-ts2tg            |                docker.io/weaveworks/weave-kube:2.8.1                |
|  kube-system   | pod  |            weave-net-xz2qx            |                docker.io/weaveworks/weave-kube:2.8.1                |
|      mall      | pod  |    elasticsearch-6775b87995-ffr4s     |                         elasticsearch:7.9.2                         |
|      mall      | pod  |        eureka-7896f8857f-ktjgv        |                      springcloud/eureka:latest                      |
|      mall      | pod  |      mall-admin-cb648c896-4lrqp       |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-admin:1.0-SNAPSHOT  |
|      mall      | pod  |      mall-auth-5fb9f5f7f9-gwzpm       |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-auth:1.0-SNAPSHOT   |
|      mall      | pod  |     mall-gateway-597f958559-mcg2f     | registry.cn-shenzhen.aliyuncs.com/k8small/mall-gateway:1.0-SNAPSHOT |
|      mall      | pod  |     mall-monitor-77c74fbf4f-gvvcd     | registry.cn-shenzhen.aliyuncs.com/k8small/mall-monitor:1.0-SNAPSHOT |
|      mall      | pod  |      mall-mysql-66756d9b7b-cmwj4      |                              mysql:5.7                              |
|      mall      | pod  |     mall-portal-66b8446cdd-h5m76      | registry.cn-shenzhen.aliyuncs.com/k8small/mall-portal:1.0-SNAPSHOT  |
|      mall      | pod  |     mall-search-5df6b697d5-9jh2d      | registry.cn-shenzhen.aliyuncs.com/k8small/mall-search:1.0-SNAPSHOT  |
|      mall      | pod  |        mongodb-7ddd4c44f-fz9th        |                             mongo:4.2.1                             |
|      mall      | pod  |       rabbitmq-7bcc69d695-8jsx4       |                     rabbitmq:3.7.15-management                      |
|      mall      | pod  |        redis-5dc8bc7c9f-mrqb7         |                                redis                                |
+----------------+------+---------------------------------------+---------------------------------------------------------------------+

```

2. 显示mall命名空间的下面的deployment images
```shell
$ kubectl-img image -t deploy -n mall
+-----------+--------+---------------+---------------------------------------------------------------------+
| NAMESPACE |  TYPE  |     NAME      |                                IMAGE                                |
+-----------+--------+---------------+---------------------------------------------------------------------+
|   mall    | deploy | elasticsearch |                         elasticsearch:7.9.2                         |
|   mall    | deploy |    eureka     |                      springcloud/eureka:latest                      |
|   mall    | deploy |  mall-admin   |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-admin:1.0-SNAPSHOT  |
|   mall    | deploy |   mall-auth   |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-auth:1.0-SNAPSHOT   |
|   mall    | deploy | mall-gateway  | registry.cn-shenzhen.aliyuncs.com/k8small/mall-gateway:1.0-SNAPSHOT |
|   mall    | deploy | mall-monitor  | registry.cn-shenzhen.aliyuncs.com/k8small/mall-monitor:1.0-SNAPSHOT |
|   mall    | deploy |  mall-mysql   |                              mysql:5.7                              |
|   mall    | deploy |  mall-portal  | registry.cn-shenzhen.aliyuncs.com/k8small/mall-portal:1.0-SNAPSHOT  |
|   mall    | deploy |  mall-search  | registry.cn-shenzhen.aliyuncs.com/k8small/mall-search:1.0-SNAPSHOT  |
|   mall    | deploy |    mongodb    |                             mongo:4.2.1                             |
|   mall    | deploy |   rabbitmq    |                     rabbitmq:3.7.15-management                      |
|   mall    | deploy |     redis     |                                redis                                |
+-----------+--------+---------------+---------------------------------------------------------------------+
```

3.显示mall命名空间的下面的名称为eureka的deployment的image
```shell
$ kubectl-img image -t deploy -c eureka -n mall
+-----------+--------+--------+---------------------------+
| NAMESPACE |  TYPE  |  NAME  |           IMAGE           |
+-----------+--------+--------+---------------------------+
|   mall    | deploy | eureka | springcloud/eureka:latest |
+-----------+--------+--------+---------------------------+
```


## Extensions
使用`Krew`,将`kubectl-img`交由kubectl管理。

`Krew` 是 类似于系统的apt、dnf或者brew的 kubectl插件包管理工具，利用其可以轻松的完成kubectl 插件的全上面周期管理，包括搜索、下载、卸载等。

kubectl 其工具已经比较完善，但是对于一些个性化的命令，其宗旨是希望开发者能以独立而紧张形式发布自定义的kubectl子命令，插件的开发语言不限，需要将最终的脚步或二进制可执行程序以kubectl- 的前缀命名，然后放到PATH中即可，可以使用kubectl plugin list查看目前已经安装的插件。

### Install 

```shell
(
  set -x; cd "$(mktemp -d)" &&
  OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
  ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
  curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" &&
  tar zxvf krew.tar.gz &&
  KREW=./krew-"${OS}_${ARCH}" &&
  "$KREW" install krew
)
```
执行这个 shell，可能 `` curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/krew.tar.gz" `` 会报错，那么自己改改脚本就好。

添加环境变量：
```shell
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
source ~/.bashrc
```
确认插件安装:
```shell
$ kubectl plugin list
The following compatible plugins are available:
/usr/local/bin/kubectl-img
/usr/local/bin/kubectl-krew
```

### Usage
```shell
$ kubectl img
kubectl-img show the Kubernetes workload image.

Usage:
  kubectl-img [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  image       kubernetes workload image

Flags:
  -h, --help     help for kubectl-img
  -t, --toggle   Help message for toggle

Use "kubectl-img [command] --help" for more information about a command.
```

### Example
显示所有namespace的pod images
```shell
$ kubectl img image -A
+----------------+------+---------------------------------------+---------------------------------------------------------------------+
|   NAMESPACE    | TYPE |                 NAME                  |                                IMAGE                                |
+----------------+------+---------------------------------------+---------------------------------------------------------------------+
|    default     | pod  |                 nginx                 |                                nginx                                |
| istio-operator | pod  |    istio-operator-64b44774cf-sl75w    |                   docker.io/istio/operator:1.13.0                   |
|  istio-system  | pod  | istio-egressgateway-6cf5fb4756-52scr  |                   docker.io/istio/proxyv2:1.13.0                    |
|  istio-system  | pod  | istio-ingressgateway-dc9c8f588-jc9kw  |                   docker.io/istio/proxyv2:1.13.0                    |
|  istio-system  | pod  |        istiod-6874697f7d-fpzrk        |                    docker.io/istio/pilot:1.13.0                     |
|  kube-system   | pod  |       coredns-558bd4d5db-4qgcs        |                  k8s.gcr.io/coredns/coredns:v1.8.0                  |
|  kube-system   | pod  |       coredns-558bd4d5db-z89dd        |                  k8s.gcr.io/coredns/coredns:v1.8.0                  |
|  kube-system   | pod  |          etcd-biz-master-48           |                      k8s.gcr.io/etcd:3.4.13-0                       |
|  kube-system   | pod  |     kube-apiserver-biz-master-48      |                  k8s.gcr.io/kube-apiserver:v1.21.0                  |
|  kube-system   | pod  | kube-controller-manager-biz-master-48 |             k8s.gcr.io/kube-controller-manager:v1.21.0              |
|  kube-system   | pod  |           kube-proxy-cgr4w            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |           kube-proxy-dgmm7            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |           kube-proxy-jmkb6            |                    k8s.gcr.io/kube-proxy:v1.21.0                    |
|  kube-system   | pod  |     kube-scheduler-biz-master-48      |                  k8s.gcr.io/kube-scheduler:v1.21.0                  |
|  kube-system   | pod  |            weave-net-hxvkn            |                docker.io/weaveworks/weave-kube:2.8.1                |
|  kube-system   | pod  |            weave-net-ts2tg            |                docker.io/weaveworks/weave-kube:2.8.1                |
|  kube-system   | pod  |            weave-net-xz2qx            |                docker.io/weaveworks/weave-kube:2.8.1                |
|      mall      | pod  |    elasticsearch-6775b87995-ffr4s     |                         elasticsearch:7.9.2                         |
|      mall      | pod  |        eureka-7896f8857f-ktjgv        |                      springcloud/eureka:latest                      |
|      mall      | pod  |      mall-admin-cb648c896-4lrqp       |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-admin:1.0-SNAPSHOT  |
|      mall      | pod  |      mall-auth-5fb9f5f7f9-gwzpm       |  registry.cn-shenzhen.aliyuncs.com/k8small/mall-auth:1.0-SNAPSHOT   |
|      mall      | pod  |     mall-gateway-597f958559-mcg2f     | registry.cn-shenzhen.aliyuncs.com/k8small/mall-gateway:1.0-SNAPSHOT |
|      mall      | pod  |     mall-monitor-77c74fbf4f-gvvcd     | registry.cn-shenzhen.aliyuncs.com/k8small/mall-monitor:1.0-SNAPSHOT |
|      mall      | pod  |      mall-mysql-66756d9b7b-cmwj4      |                              mysql:5.7                              |
|      mall      | pod  |     mall-portal-66b8446cdd-h5m76      | registry.cn-shenzhen.aliyuncs.com/k8small/mall-portal:1.0-SNAPSHOT  |
|      mall      | pod  |     mall-search-5df6b697d5-9jh2d      | registry.cn-shenzhen.aliyuncs.com/k8small/mall-search:1.0-SNAPSHOT  |
|      mall      | pod  |        mongodb-7ddd4c44f-fz9th        |                             mongo:4.2.1                             |
|      mall      | pod  |       rabbitmq-7bcc69d695-8jsx4       |                     rabbitmq:3.7.15-management                      |
|      mall      | pod  |        redis-5dc8bc7c9f-mrqb7         |                                redis                                |
+----------------+------+---------------------------------------+---------------------------------------------------------------------+
```

## Reference
[redhatxl/kubectl-img](https://github.com/redhatxl/kubectl-img)



