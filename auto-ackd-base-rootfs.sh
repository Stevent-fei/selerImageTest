#!/bin/bash

set -e

for i in "$@"; do
  case $i in
  -c=* | --cri=*)
    cri="${i#*=}"
    if [ "$cri" != "docker" ] && [ "$cri" != "containerd" ]; then
      echo "Unsupported container runtime: ${cri}"
      exit 1
    fi
    shift # past argument=value
    ;;
  --push)
    push="true"
    shift # past argument=value
    ;;
  -p=* | --password=*)
    password="${i#*=}"
    shift # past argument=value
    ;;
    --platform=*)
    platform="${i#*=}"
    shift # past argument=value
    ;;
  -u=* | --username=*)
    username="${i#*=}"
    shift # past argument=value
    ;;
  --tag=*)
    tag="${i#*=}"
    shift # past argument=value
    ;;
  -h | --help)
    echo "
### Options
  --push                push clusterimage after building the clusterimage. The image name must contain the full name of the repository, and use -u and -p to specify the username and password.
  -u, --username        specify the user's username for pushing the Clusterimage
  -p, --password        specify the user's password for pushing the Clusterimage
  -d, --debug           show all script logs
  -h, --help            help for auto build shell scripts"
    exit 0
    ;;
  -d | --debug)
    set -x
    shift
    ;;
  -*)
    echo "Unknown option $i"
    exit 1
    ;;
  *) ;;
  esac
done

cri=docker

workdir="$(mktemp -d auto-build-XXXXX)" && sudo cp -r context-ackd "${workdir}" && cd "${workdir}/context-ackd" && sudo cp -rf "${cri}"/* .
platform=$(if [[ -z "$platform" ]]; then echo "linux/arm64,linux/amd64"; else echo "$platform"; fi)
# shellcheck disable=SC1091
sudo chmod +x version.sh download.sh  && source version.sh
./download.sh "${cri}"

sudo mkdir manifests
sudo sealer build -t "registry.cn-qingdao.aliyuncs.com/sealer-io/ackdistro-multi:${tag}" -f Kubefile --platform "${platform}"
if [[ "$push" == "true" ]]; then
  if [[ -n "$username" ]] && [[ -n "$password" ]]; then
    sudo sealer login "$(echo "docker.io" | cut -d "/" -f1)" -u "${username}" -p "${password}"
  fi
  sudo sealer push "registry.cn-qingdao.aliyuncs.com/sealer-io/ackdistro-multi:${tag}"
fi
