# bosh-concourse-tasks

Execute [BOSH](https://bosh.io/) tasks via [Concourse](https://concourse.ci/) tasks.


## Director Tasks

Requires `client`, `client_secret`, `director`, and accepts `ca_cert` parameters.

 * `clean-up.yml`
 * `clean-up-all.yml`


### Deployment Tasks

Requires `deployment` parameter.

 * `run-errand.yml` (requires `errand` parameter)


#### Instance Group / ID Tasks

Accepts `instance_group` and `id` parameters.

 * `recreate.yml`
 * `ssh.yml` (requires `command` parameter)
 * `start.yml`
 * `stop.yml`
 * `take-snapshot.yml`


## Example

```yaml
meta:
  bosh_environment: &bosh_environment
    ca_cert: ((bosh_ca_cert))
    client: ((bosh_client))
    client_secret: ((bosh_client_secret))
    environment: ((bosh_environment))
  bosh_deployment: &bosh_deployment
    <<: *bosh_environment
    deployment: legacy
jobs:
- name: start
  plan:
  - get: weekday-start
    trigger: true
  - get: bosh-concourse-tasks
  - task: start-database
    file: bosh-concourse-tasks/start.yml
    params:
      <<: *bosh_deployment
      instance_group: database
  - task: start-app-server
    file: bosh-concourse-tasks/start.yml
    params:
      <<: *bosh_deployment
      instance_group: app-server
  serial_groups:
  - restart
- name: stop
  plan:
  - get: weekday-stop
    trigger: true
  - get: bosh-concourse-tasks
  - task: stop-database
    file: bosh-concourse-tasks/stop.yml
    params:
      <<: *bosh_deployment
      instance_group: database
  - task: stop-app-server
    file: bosh-concourse-tasks/stop.yml
    params:
      <<: *bosh_deployment
      instance_group: app-server
  serial_groups:
  - restart
- name: clean-up
  plan:
  - get: daily
    trigger: true
  - get: bosh-concourse-tasks
  - task: clean-up
    file: bosh-concourse-tasks/clean-up-all.yml
    params: *bosh_environment
resources:
- name: bosh-concourse-tasks
  type: git
  source:
    uri: https://github.com/dpb587/bosh-concourse-tasks.git
- name: daily
  type: timer
  source:
    interval: "24h"
- name: weekday-start
  type: timer
  source:
    location: America/Denver
    start: "07:00"
    stop: "08:00"
    days: [ Monday, Tuesday, Wednesday, Thursday, Friday ]
- name: weekday-stop
  type: timer
  source:
    location: America/Denver
    start: "20:00"
    stop: "21:00"
    days: [ Monday, Tuesday, Wednesday, Thursday, Friday ]
```


## License

[MIT License](LICENSE)
