# machineid-github-kubernetes-webinar

The accompanying demo repository for the [Machine ID GitHub and Kubernetes
Webinar](https://teleport.registration.goldcast.io/events/cc92262f-da46-44c3-9c3a-c6278f40a043).

The demo shows a simple Go application that uses GitHub Actions and Teleport to
build a docker image and deploy this to a Kubernetes cluster securely and
without the use of secrets.

## Highlights

### No long-lived secrets

Unlike some traditional methods of configuring GitHub Actions to deploy to a
Kubernetes cluster, there are no secrets being stored in GitHub Action's
secret store. This prevents these powerful secrets from being exfiltrated from
your CI/CD system and then used for malicious purposes.

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

### Fine-grained access control

Access is limited to only jobs within this repository attached to the
environment `production`.

## Trying it yourself

Pre-requisites:
- Kubernetes cluster connected to Teleport cluster.
- Your cluster should support LoadBalancer type services (this is for the demo
  to be internet accessible - you can omit this).

Steps:
- Fork this repo and replace:
  - `strideynet/machineid-github-kubernetes-webinar` with your own GitHub repo's path.
  - `gcp` with the name of your Kubernetes cluster as configured in Teleport.
  - `noah-demo.teleport.sh:443` with the publicly accessible address of your proxy.
- `tctl create -f teleport/role.yaml` - creates the role your bot will use.
- `tctl create -f teleport/github-bot-token.yaml` - defines the rules for which GitHub Action will be able to access the bot you create in the next step.
- `tctl bots add colormatic-deployer --roles=colormatic-deployer --token=colormatic-deployer` - creates the Bot user that your GitHub Actions will authenticate as when connecting to Teleport.
- `kubectl apply` `kubernetes/00-namespace.yaml`, `kubernetes/01-role.yaml`, `kubernetes/02-rolebinding.yaml` - create roles and role bindings in Kubernetes to assign that bot permissions to deploy to your cluster.
- Your Teleport and Kubernetes cluster now have the correct RBAC to allow your GitHub repo to deploy to the Kubernetes cluster through Teleport.
- Commit and push a change to `main.go` to trigger the CI.
- Check to see service deployed.
- Take a peek at your Teleport audit log.