# What is this?

This is supposed to be a small project to export some metrics from argocd.
It will be a prometheus exporter for the argocd branch information.
it's mainly a learning project for me and not suited for production use.

# Why?

I want to be able to monitor the argocd branch information in prometheus.
This information can then be used for alerting. Fo example, if the branch is `hotfix/foo` and the `hotfix` branch is not merged into `main` for a long time, then it might be a good idea to alert.

# Plans / Features

- [ ] Figure out how to export any merics at all
- [ ] Figure out how to interact with the Kubernetes API
    - https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
- [ ] Figure out how proper testing works in golang
- [ ] Figure out how to work with `make`
- [ ] Figure out how to package this in a container
    - [ ] Make it small
    - [ ] Make it secure
    - [ ] Make it distroless (?)
- [ ] Figure out github actions for automations
    - [ ] Testing
    - [ ] Packaging
- [ ] Figure out github container registry for publishing
- [ ] Make some nice docs
- [ ] Maybe have other people contribute

