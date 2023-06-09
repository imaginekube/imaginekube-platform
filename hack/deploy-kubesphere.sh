#!/usr/bin/env bash

# Copyright 2023 The ImagineKube Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

function wait_for_installation_finish() {
    echo "waiting for ks-installer pod ready"
    kubectl -n imaginekube-system wait --timeout=180s --for=condition=Ready "$(kubectl -n imaginekube-system get pod -l app=ks-install -oname)"
    echo "waiting for ImagineKube ready"
    while IFS= read -r line; do
        if [[ $line =~ "Welcome to ImagineKube" ]]
            then
                break
        fi
    done < <(timeout 900 kubectl logs -n imaginekube-system deploy/ks-installer -f)
}

# Use imaginekubedev and latest tag as default image
TAG="${TAG:-latest}"
REPO="${REPO:-imaginekubedev}"

# Use KIND_LOAD_IMAGE=y .hack/deploy-imaginekube.sh to load
# the built docker image into kind before deploying.
if [[ "${KIND_LOAD_IMAGE:-}" == "y" ]]; then
    kind load docker-image "$REPO/ks-apiserver:$TAG" --name="${KIND_CLUSTER_NAME:-kind}"
    kind load docker-image "$REPO/ks-controller-manager:$TAG" --name="${KIND_CLUSTER_NAME:-kind}"
fi

#TODO: override ks-apiserver and ks-controller-manager images with specific tag
kubectl apply -f https://raw.githubusercontent.com/imaginekube/ks-installer/master/deploy/imaginekube-installer.yaml
kubectl apply -f https://raw.githubusercontent.com/imaginekube/ks-installer/master/deploy/cluster-configuration.yaml


wait_for_installation_finish

# Expose ImagineKube API Server
kubectl -n imaginekube-system patch svc ks-apiserver -p '{"spec":{"type":"NodePort","ports":[{"name":"ks-apiserver","port":80,"protocol":"TCP","targetPort":9090,"nodePort":30881}]}}'
