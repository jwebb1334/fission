#!/bin/bash

set -euo pipefail

# Unbound variables cause failure, so this readable if block instead of Parameter Expansion
if [[ ${TRAVIS_EVENT_TYPE+NOVALUE} != "cronNOVALUE" ]]
then
    echo "Skipping Fission upgrade tests, not a cron job"
    exit 0
fi

if [ ! -f ${HOME}/.kube/config ]
then
    echo "Skipping end to end tests, no cluster credentials"
    exit 0
fi

ROOT_RELPATH=$(dirname $0)/../..
pushd $ROOT_RELPATH
ROOT=$(pwd)
popd

# This will change for every new release
CURRENT_VERSION=0.7.0

source $ROOT/test/test_utils.sh
source $(dirname $0)/fixture_tests.sh

id=$(generate_test_id)
ns=f-$id
fns=f-func-$id
controllerNodeport=31234
pruneInterval=1
routerServiceType=LoadBalancer

helmVars=functionNamespace=$fns,controllerPort=$controllerNodeport,pullPolicy=Always,analytics=false,pruneInterval=$pruneInterval,routerServiceType=$routerServiceType

dump_system_info


echo "Deleting old releases"
helm list -q|xargs -I@ bash -c "helm_uninstall_fission @"

# deleting ns does take a while after command is issued
while kubectl get ns| grep "fission-builder"
do
    sleep 5
done

echo "Creating namespace $ns"
kubectl create ns $ns
helm install $id \
--wait \
--timeout 540s \
--set $helmVars \
--namespace $ns \
--force https://github.com/jwebb1334/fission/releases/download/${CURRENT_VERSION}/fission-all-${CURRENT_VERSION}.tgz

mkdir temp && cd temp && curl -Lo fission https://github.com/jwebb1334/fission/releases/download/${CURRENT_VERSION}/fission-cli-linux && chmod +x fission && sudo mv fission /usr/local/bin/ && cd .. && rm -rf temp

wait_for_service $id "router"
export FISSION_ROUTER=$(kubectl -n $ns get svc router -o jsonpath='{...ip}')

## Setup - create fixtures for tests

setup_fission_objects
trap "cleanup_fission_objects $id" EXIT

## Test before upgrade

upgrade_tests

## Build images for Upgrade

REPO=gcr.io/$GKE_PROJECT_NAME
IMAGE=fission-bundle
FETCHER_IMAGE=$REPO/fetcher
BUILDER_IMAGE=$REPO/builder
TAG=upgrade-test
PRUNE_INTERVAL=1 # Unit - Minutes; Controls the interval to run archivePruner.
ROUTER_SERVICE_TYPE=ClusterIP

build_and_push_fission_bundle $IMAGE:$TAG

build_and_push_fetcher $FETCHER_IMAGE:$TAG

build_fission_cli

sudo mv $ROOT/fission/fission /usr/local/bin/

## Upgrade 

helmVars=repository=$repo,image=$IMAGE,imageTag=$TAG,fetcher.image=$FETCHER_IMAGE,fetcher.imageTag=$TAG,functionNamespace=$fns,controllerPort=$controllerNodeport,pullPolicy=Always,analytics=false,pruneInterval=$pruneInterval,routerServiceType=$routerServiceType

echo "Upgrading fission"
helm upgrade	\
 --wait			\
 --timeout 540s	        \
 --set $helmVars	\
 --namespace $ns        \
 $id $ROOT/charts/fission-all

sleep 10 # Takes a few seconds after upgrade to re-create K8S objects etc.

# No need to wait as upgrade won't create a new LB
export FISSION_ROUTER=$(kubectl -n $ns get svc router -o jsonpath='{...ip}')

## Tests
validate_post_upgrade
upgrade_tests

## Cleanup is done by trap
