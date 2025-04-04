# hive-cli

Command line tool for building / deploying AI agents to Kubernetes. Work in progress.

## Usage

### Install CLI

Download the [latest release](https://github.com/psschwei/hive-cli/releases) and place the binary on your PATH.

### Create an Agent

At a minimum, you need to create two files, `requirements.txt` and `run_agents.py`.

`requirements.txt` is the standard Python [requirements file](https://pip.pypa.io/en/stable/reference/requirements-file-format/). For example

```
beeai-framework==0.1.8
```

`run_agents.py` is the source code for your agent. It needs to implement this function:

```python
async def run_agent(prompt: str) -> str:
    # Do your stuff here, then return a string
```

A full working example, using the BeeAI Framework, can be found [here](https://github.com/psschwei/bees-in-a-pod/tree/main/agents/custom-agent).

### Build your Agent

Use the `hive build` command to build your agent using Docker. (If you don't have `docker`, this won't work)

```bash
hive build --dir /home/user/my-agent/src --tag icr.io/username/agent:latest
```

### Deplpy your Agent to Kubernetes

We assume you have `kubectl` configured to access some Kubernetes cluster.

We also assume that you have deployed an `agent-configmap` and an `agent-secrets` as shown [here](https://github.com/psschwei/bees-in-a-pod/blob/main/config/deploy.yaml). This is required, even if your agent isn't using values from either (just deploy them with dummy values).

Use the `hive deploy` command to deploy your containerized agent to Kubernetes.

```bash
hive deploy --name my-agent --tag icr.io/username/agent:latest
```
