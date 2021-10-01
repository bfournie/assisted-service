#!/usr/bin/env bash
export ASSISTED_OPENSHIFT_VERSION="${ASSISTED_OPENSHIFT_VERSION_OVERRIDE:-openshift-v4.9.0}"
export ASSISTED_UPGRADE_OPERATOR="${ASSISTED_UPGRADE_OPERATOR_OVERRIDE:-false}"
export ASSISTED_STOP_AFTER_AGENT_DISCOVERY="${ASSISTED_STOP_AFTER_AGENT_DISCOVERY_OVERRIDE:-true}"
export ASSISTED_CLUSTER_NAME="${ASSISTED_CLUSTER_NAME:-assisted-test-cluster-upgraded}"
export ASSISTED_CLUSTER_DEPLOYMENT_NAME="${ASSISTED_CLUSTER_DEPLOYMENT_NAME:-assisted-test-cluster-upgraded}"
export ASSISTED_AGENT_CLUSTER_INSTALL_NAME="${ASSISTED_AGENT_CLUSTER_INSTALL_NAME:-assisted-agent-cluster-install-upgraded}"
export ASSISTED_INFRAENV_NAME="${ASSISTED_INFRAENV_NAME:-assisted-infra-env-upgraded}"