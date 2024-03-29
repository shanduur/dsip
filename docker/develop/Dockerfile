FROM golang:latest

# update packages
RUN apt-get update && apt-get upgrade -y

# add protoc
RUN apt-get install -y protobuf-compiler

# install SSH server
RUN apt-get install -y openssh-server

# install SUDO
RUN apt-get install -y sudo

# install Clang Format
RUN apt-get install -y clang-format

# add GOLANG to path for SSH user
RUN echo 'export PATH=$PATH:/usr/local/go/bin:/root/go/bin' >> /root/.bashrc

# this should be changed
ARG GIT_NAME=name
ARG GIT_EMAIL=email
ARG SSH_AUTH_KEY

RUN echo '[user]' > /root/.gitconfig
RUN echo "\tname = $GIT_NAME" >> /root/.gitconfig
RUN echo "\temail = $GIT_EMAIL" >> /root/.gitconfig
RUN echo "[credential]" >> /root/.gitconfig
RUN echo "\thelper = store --file ~/.git_credentials" >> /root/.gitconfig

RUN mkdir /root/.ssh && echo $SSH_AUTH_KEY > /root/.ssh/authorized_keys
RUN mkdir /root/code

# configure SSH server
RUN mkdir /var/run/sshd
RUN sed -i 's/#*PermitRootLogin prohibit-password/PermitRootLogin prohibit-password/g' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
RUN sed -i 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' /etc/pam.d/sshd

ENV NOTVISIBLE="in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

ARG PG_DATABASE
RUN echo "export PG_DATABASE=$PG_DATABASE" >> /etc/profile

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]