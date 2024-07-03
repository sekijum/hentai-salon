FROM ubuntu:22.04

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /infra

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y less vim curl unzip git jq dpkg lsof openssh-server mysql-client iputils-ping net-tools

# aws cli v2
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
RUN ./aws/install
RUN rm awscliv2.zip

# ssm
RUN curl "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_64bit/session-manager-plugin.deb" -o "session-manager-plugin.deb"
RUN dpkg -i session-manager-plugin.deb
RUN rm session-manager-plugin.deb

# terraform
ENV TF_VERSION=1.5.3
RUN curl "https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip" -o "terraform.zip"
RUN unzip -d /usr/local/bin/ terraform.zip
RUN rm terraform.zip

# ssh
RUN mkdir /var/run/sshd
RUN echo 'root:root' | chpasswd
RUN sed -ri 's/^#?PermitEmptyPasswords\s+.*/PermitEmptyPasswords yes/' /etc/ssh/sshd_config
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -ri 's/^#?UsePAM\s+.*/UsePAM no/' /etc/ssh/sshd_config
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

EXPOSE 22
CMD service ssh start && /bin/bash
