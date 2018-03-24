# API 文档

### 存储结构

* 第一层 命名空间(或者叫组织机构)
* 第二层 项目名称
* 第三层 版本+二进制文件

eq.
```
google
  |- kubernetes
      |- 文件（文件名，版本，大小，hash）
      |- 1.6.x
          |- kubecli.tar.gz
             # https://example.com/google/kubernetes/kubecli.tar.gz?version=x.y.z
          |- kubelet.tar.gz
          |- kube-proxy.tar.gz
      |- 1.7.x
          |- '/'.tar.gz
          |- kubelet.tar.gz
          |- kube-proxy.tar.gz
      |- 1.9.x
          |- kubecli.tar.gz
          |- kubelet.tar.gz
          |- kube-proxy.tar.gz
  |- angularjs
      |- 1.4
          |- ng-cli.tar.gz
      |- 2.0

```


### 项目组
