#!/bin/bash

oc login --insecure-skip-tls-verify ${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT_HTTPS} -u "${USER}" -p "${PASSWORD}"

bundle-controller
