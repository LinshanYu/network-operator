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


bin/controller-gen的作用可以手动生成crd rbac等文件


支持pod多网络
1.Initializer，给pod添加一个网络代理容器
2.用户的cr指定labelselector的pod
3.当用户新创建出来的pod 符合labelselector的做一下处理，容器中的网络代理负责处理容器中的路由等
4.容器外的daemonset负责host的路由等操作
