# Kubectl Search

[![CircleCI](https://circleci.com/gh/guessi/kubectl-grep.svg?style=svg)](https://circleci.com/gh/guessi/kubectl-grep)
[![GoDoc](https://godoc.org/github.com/guessi/kubectl-grep?status.svg)](https://godoc.org/github.com/guessi/kubectl-grep)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/kubectl-grep)](https://goreportcard.com/report/github.com/guessi/kubectl-grep)
[![GitHub release](https://img.shields.io/github/release/guessi/kubectl-grep.svg)](https://github.com/guessi/kubectl-grep/releases/latest)

simple kubernetes plugins for searching resources with keyword

# Requirements

- Kubernetes 1.10.0+
- Kubectl 1.13.0+

# Why we need it?

playing with Kubernetes is my daily job, and I normally search pods by `pipe`,
`grep`, `--label`, `--field-selector`, etc. while hunting abnormal pods, but
typing such long commands are quite annoyed.

Before change, we usually search pods by the following commands,

    $ kubectl get pods                                 | grep "keyword"
    $ kubectl get pods -o wide                         | grep "keyword"
    $ kubectl get pods -n "my-ns"                      | grep "keyword"
    $ kubectl get pods --all-namespaces                | grep "keyword"
    $ kubectl get pods -l "key=value"                  | grep "keyword"
    $ kubectl get pods -l "key=value" -n "my-ns"       | grep "keyword"
    $ kubectl get pods -l "key=value" --all-namespaces | grep "keyword"
    $ kubectl get pods --field-selector "key=value"    | grep "keyword"

With this plugin installed, you can search pod with `kubectl grep` easily

    $ kubectl grep pods "keyword"
    $ kubectl grep pods "keyword" -o wide
    $ kubectl grep pods "keyword" -n "my-ns"
    $ kubectl grep pods "keyword" --all-namespaces
    $ kubectl grep pods "keyword" -l "key=value"
    $ kubectl grep pods "keyword" -l "key=value" -n "my-ns"
    $ kubectl grep pods "keyword" -l "key=value" --all-namespaces
    $ kubectl grep pods "keyword" --field-selector "key=value"

# Installation

    $ export KUBECTL_VERSION="v1.0.5"
    $ wget https://github.com/guessi/kubectl-grep/releases/download/${KUBECTL_VERSION}/kubectl-grep-`uname -s`-`uname -m`.tar.gz
    $ tar zxvf kubectl-grep-`uname -s`-`uname -m`.tar.gz
    $ mv kubectl-grep /usr/local/bin
    $ chmod +x /usr/local/bin/kubectl-grep

# Basic Usage

    $ kubectl grep pods "keyword"

# How to get developer build?

    $ go get -u github.com/guessi/kubectl-grep

    $ cd ${GOPATH}/src/github.com/guessi/kubectl-grep

    $ make all

# FAQ

What kind of resource does current release supported?

    now it support Nodes, Pods, Deployments, DaemonSets, and HPAs

Is it possible to search StatefulSets, etc?

    sure, but not yet supported, I need more time to write code, and tests

How do I check the tool's version?

    $ kubectl grep version

Any plan to support Kubernetes version before 1.10.0?

    sorry, it only support Kubernetes 1.10.0+

I'm now running Kubernetes 1.10.0, do I need to upgrade my cluster?

    nope, the only requirement is to upgrade your `kubectl` to 1.13.0+

Can I run Kubernetes 1.12.0 with kubectl 1.13.0?

    sure, no problem

# Reference

- [Kubectl Plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)

# License

[Apache-2.0](LICENSE)
