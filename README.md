下载kubebuilder curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
mkdir network-operator
cd network-operator
go mod init network-operator
kubebuilder init --domain linshanyu.com --owner "linshanyu" --skip-go-version-check
kubebuilder create api --group net --version v1 --kind MutilCniNetwork


打算做一个支持多网络的crd