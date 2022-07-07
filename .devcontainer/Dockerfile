FROM golang:1.18.3

ARG VSCODE_SCRIPTS_VERSION="v0.193.0"

RUN apt-get update && \
    wget "https://raw.githubusercontent.com/microsoft/vscode-dev-containers/${VSCODE_SCRIPTS_VERSION}/script-library/common-debian.sh" && \
    chmod +x ./common-debian.sh && \
    ./common-debian.sh false vscode automatic automatic true false && \
    wget "https://raw.githubusercontent.com/microsoft/vscode-dev-containers/${VSCODE_SCRIPTS_VERSION}/script-library/go-debian.sh" && \
    chmod +x ./go-debian.sh && \
    ./go-debian.sh none /usr/local/go /go vscode false true


# Install Bazelisk as bazel to manage Bazel
RUN go install github.com/bazelbuild/bazelisk@latest && \
    mv $(which bazelisk) /usr/local/bin/bazel

# Install buildifier for bazel formatting
RUN go install github.com/bazelbuild/buildtools/buildifier@latest

RUN go install github.com/bufbuild/buf/cmd/buf@v1.3.1

# Give vscode ownership of GOPATH
RUN chown -R vscode: /go

USER vscode