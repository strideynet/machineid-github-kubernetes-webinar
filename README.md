# machineid-github-kubernetes-webinar

This README will be expanded.

Pre-reqs:
- Kubernetes cluster connected to Teleport cluster

Steps:
- `tctl create -f teleport/role.yaml`
- `tctl create -f teleport/github-bot-token.yaml`
- `kubectl apply` `kubernetes/00-namespace.yaml`, `kubernetes/01-role.yaml`, `kubernetes/02-rolebinding.yaml`
- Your Teleport and Kubernetes cluster now have the correct RBAC to allow your GitHub repo to deploy to the Kubernetes cluster through Teleport.
- Trigger CI with color change.
- Check to see service deployed.
- Take a peek at your Teleport audit log.