# Sami-cli specs

The **sami-cli** is a tool that allows us to deploy our applications on any infrastructure,
whether it's on [swarm](https://docs.docker.com/engine/swarm/) or [k8s](https://kubernetes.io/).

## How it works ?
Each application deployed by the **sami-cli** should follow a certain structure.
The application deployments live in a folder that contains its deployment manifests
and a `sami.yaml` that gathers all the information and metadata required to deploy
the application.

### The folder structure

An application folder will live in a [deployments](https://github.com/osscameroon/deployments) repository.
The folder should be named after the application you want to deploy and stored in the `services` folder.
as it is done for the [camerdevs](https://github.com/osscameroon/deployments/tree/main/services) application.

Under this `deployments/services/<application_name>` folder should be added a folder named
after the type of deployment you want to perform. The deployment can be either `compose` or `k8s` deployment.

In the `deployments/services/<application_name>/<deployment_type>` should be added your
application deployment manifests along with a `sami.yaml` file.

An example of deployment manifests can be found [here](https://github.com/osscameroon/deployments/tree/main/services/camerdevs/compose/stage).


### The `sami.yaml` file

The `sami.yaml` file contains your application metadata along with informations required
to deploy your applications.

| Fields       | Description           |
|--------------|-----------------------|
|`version`       | The version of the `sami.yaml` file it is set using semantic versionning `v0`|
|`service_name`  | The name of the service you want to deploy. **Example: camerdevs**
|`type`          | The type of deployment you want to use can be `swarm|k8s|cron`
|`pr_deploy`     | Set wheter it is a pull request deployment or not. set to `false` by default
|`files`         | Is a list of deployment files. for camerdevs it should be set to `- camerdevs-stack.yml`
|`oob`           | The out of band field indicate which field and file should be edited when a deployment is performed, we could change the image sha of `service.api.image` field in the `camerdevs-stack.yml` file
|`oob.[].key.path`   | Path to the field we want to edit in the deployment file
|`oob.[].key.file`   | name of the file to the field we want to edit in the deployment file

### `sami-cli` commands

The `sami-cli` has several commands that help us deploy, monitor and apply changes to our deployment.


#### `sami deploy`

This command takes a `sami.yaml` file and apply your deployment manifest to deploy the application in your infrastructure

```
sami deploy -f deployments/services/camerdevs/compose/sami.yaml
```

*NOTE: If the file is not provided to the command, sami will look for a file in the current directory.*

#### `sami logs`

Displays logs of a given service

```
sami logs
```

OR

To watch the service logs as they appear

```
sami logs --watch
```


#### `sami oob`

The `oob` (out of bands) command is used to change the value of a specific field in our deployment files.

For a given `sami.yaml`, running `sami oob update-image "new-image"` will update the `stack.yaml` file
on this `.services.frontend.image` to `new-image`.

```yaml
version: v1beta

service_name: "camerdevs"
type: "swarm"
files:
 - ./stack.yaml
oob:
 - update-image:
   file: "./stack.yaml"
   path: ".services.frontend.image"
```


#### `sami status`

Shows the status of a given service

`sami status -f ./sami.yaml`
