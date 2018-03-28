FROM centos:7
MAINTAINER Ansible Playbook Bundle Community

RUN yum -y update \
 && yum -y install epel-release centos-release-openshift-origin \
 && yum -y install origin-clients net-tools bind-utils \
 && yum clean all

ENV BASE_DIR=/opt/bundle-controller
ENV HOME=${BASE_DIR}

RUN mkdir -p /opt/bundle-controller/.kube
COPY config /opt/bundle-controller/.kube/config
RUN chmod -R g=u /opt/bundle-controller

COPY entrypoint.sh /usr/bin/
COPY main /usr/bin/bundle-controller

USER bundle-controller
ENTRYPOINT ["entrypoint.sh"]
