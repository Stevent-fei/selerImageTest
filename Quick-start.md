#Quick start

#Install a kubernetes cluster

```shell
# install Sealer binaries
wget https://github.com/sealerio/sealer/releases/download/v0.9.0/sealer-v0.9.0-linux-amd64.tar.gz && \
tar zxvf sealer-v0.9.0-linux-amd64.tar.gz && mv sealer /usr/bin
# run a kubernetes cluster
sealer run registry.cn-qingdao.aliyuncs.com/sealer-io/kubernetes:v1.22.15 \
  --masters 192.168.0.2,192.168.0.3,192.168.0.4 \
  --nodes 192.168.0.5,192.168.0.6,192.168.0.7 --passwd xxx
```

```shell
[root@iZm5e42unzb79kod55hehvZ ~]# kubectl get node
NAME                    STATUS ROLES AGE VERSION
izm5e42unzb79kod55hehvz Ready master 18h v1.22.15
izm5ehdjw3kru84f0kq7r7z Ready master 18h v1.22.15
izm5ehdjw3kru84f0kq7r8z Ready master 18h v1.22.15
izm5ehdjw3kru84f0kq7r9z Ready <none> 18h v1.22.15
izm5ehdjw3kru84f0kq7raz Ready <none> 18h v1.22.15
izm5ehdjw3kru84f0kq7rbz Ready <none> 18h v1.22.15
```

#Clean the cluster

Some information of the basic settings will be written to the Clusterfile and stored in /root/.sealer/Clusterfile.

```shell
sealer delete -af /root/.sealer/Clusterfile
```

# build app image

nginx.yaml:

```shell
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-nginx
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      run: my-nginx
  template:
    metadata:
      labels:
        run: my-nginx
    spec:
      containers:
        - name: my-nginx
          image: nginx
          ports:
            - containerPort: 80
```

Kubefile:

```shell
FROM scratch
APP nginx local://nginx.yaml
LAUNCH ["nginx"]
```

```shell
sealer build -f Kubefile -t registry.cn-qingdao.aliyuncs.com/sealer-io/nginx:latest --type app-installer
```
# run app image

```shell
sealer run registry.cn-qingdao.aliyuncs.com/sealer-io/nginx:latest
# check the pod
kubectl get pod -A
```

# push the app image to the registry

```shell
# you can push the app image to docker hub, Ali ACR, or Harbor
sealer push registry.cn-qingdao.aliyuncs.com/sealer-io/nginx:latest
```



# Build your own ClusterImage

For example, build a dashboard ClusterImage:

```shell
wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.2.0/aio/deploy/recommended.yaml
```
Kubefile:

```shell
# base ClusterImage contains all the files that run a kubernetes cluster needed.
#    1. kubernetes components like kubectl kubeadm kubelet and apiserver images ...
#    2. docker engine, and a private registry
#    3. config files, yaml, static files, scripts ...
FROM registry.cn-qingdao.aliyuncs.com/sealer-io/kubernetes:v1.22.15
# download kubernetes dashboard yaml file
APP recommended local://recommended.yaml
# when run this ClusterImage, will apply a dashboard manifests
LAUNCH ["recommended"]
```

Build dashboard ClusterImage:

```shell
sealer build -f Kubefile -t registry.cn-qingdao.aliyuncs.com/sealer-io/dashboard:latest --type kube-installer
```

Run your kubernetes cluster with dashboard:

```shell
# sealer will install a kubernetes on host 192.168.0.2 then apply the dashboard manifests
sealer run registry.cn-qingdao.aliyuncs.com/sealer-io/dashboard:latest --masters 192.168.0.2 --passwd xxx
# check the pod
kubectl get pod -A|grep dashboard
```
