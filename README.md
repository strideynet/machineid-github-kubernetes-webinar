# machineid-github-kubernetes-webinar

This README will be expanded.

Pre-reqs:
- Kubernetes cluster connected to Teleport cluster

Steps:
- Fork this repo, and replace:
  - `strideynet/machineid-github-kubernetes-webinar` with your own GitHub repo
  - `docker-desktop` with the name of your Kubernetes cluster
  - `root.tele.ottr.sh:443` with the address of your proxy
- `tctl create -f teleport/role.yaml`
- `tctl create -f teleport/github-bot-token.yaml`
- `tctl bots add bots add colormatic-deployer --roles=colormatic-deployer --token=colormatic-deployer`
- `kubectl apply` `kubernetes/00-namespace.yaml`, `kubernetes/01-role.yaml`, `kubernetes/02-rolebinding.yaml`
- Your Teleport and Kubernetes cluster now have the correct RBAC to allow your GitHub repo to deploy to the Kubernetes cluster through Teleport.
- Trigger CI with color change.
- Check to see service deployed.
- Take a peek at your Teleport audit log.