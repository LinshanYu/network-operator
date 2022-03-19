下载kubebuilder curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
mkdir network-operator
cd network-operator
go mod init network-operator
kubebuilder init --domain linshanyu.com --owner "linshanyu" --skip-go-version-check
//创建api/net version v1 kind MutilCniNetwork
kubebuilder create api --group net --version v1 --kind MutilCniNetwork


在/api/v1/MutilCniNetwork 把MutilCniNetworkSpec 新增逻辑代码 例如我们要支持多个网络 那么可以定义cnis，然后，重新生成代码和二进制
make manifests generate

打算做一个支持多网络的crd
1.用户cr触发当前controller
2.该controller启动daemonset,在/etc/cni/net.d下删除配置，并去cr中读取相关的数据，将数据组装成一个cni的配置
3.