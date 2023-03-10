# machineid-github-kubernetes-webinar

The accompanying demo repository for 
https://teleport.registration.goldcast.io/events/cc92262f-da46-44c3-9c3a-c6278f40a043

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

## Highlights

### Rich auditing

Teleport tracks audit events for the bot join, as well for each Kubernetes
request made by the bot through Teleport.

Bot join audit event in Teleport includes rich information about the CI workflow
that join originated from:

```json
{
  "attributes": {
    "actor": "strideynet",
    "actor_id": "16336790",
    "base_ref": "",
    "environment": "production",
    "event_name": "push",
    "head_ref": "",
    "job_workflow_ref": "strideynet/machineid-github-kubernetes-webinar/.github/workflows/deploy.yaml@refs/heads/main",
    "ref": "refs/heads/main",
    "ref_type": "branch",
    "repository": "strideynet/machineid-github-kubernetes-webinar",
    "repository_id": "611387043",
    "repository_owner": "strideynet",
    "repository_owner_id": "16336790",
    "repository_visibility": "public",
    "run_attempt": "1",
    "run_id": "4382935357",
    "run_number": "15",
    "sha": "91ea9ae56b5686a85e5e412570d1391294606089",
    "sub": "repo:strideynet/machineid-github-kubernetes-webinar:environment:production",
    "workflow": "Deploy Colormatic!"
  },
  "bot_name": "colormatic-deployer",
  "cluster_name": "root.tele.ottr.sh",
  "code": "TJ001I",
  "ei": 0,
  "event": "bot.join",
  "method": "github",
  "success": true,
  "time": "2023-03-10T09:04:34.621Z",
  "token_name": "colormatic-deployer",
  "uid": "e623f01a-94dd-4bd2-9914-19c0fbc672d6"
}
```