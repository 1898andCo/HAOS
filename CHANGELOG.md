<a name="unreleased"></a>
## [Unreleased]

### Documentation Improvements
- **changelog:** auto changelog update v1.5.4-v1.21.1-k3s1 [CI SKIP]


<a name="v1.5.4-v1.21.1-k3s1"></a>
## [v1.5.4-v1.21.1-k3s1] - 2021-06-08
### Bug Fixes
- **drone:** add bash to haos prep

### Dependency Updates
- **00-base:** bump openrc version

### Documentation Improvements
- **changelog:** auto changelog update v1.5.4-v1.21.1-k3s1 [CI SKIP]
- **changelog:** auto changelog update v1.5.3-v1.21.1-k3s1 [CI SKIP]
- **changelog:** auto changelog update v1.5.3-v1.21.1-k3s1 [CI SKIP]


<a name="v1.5.3-v1.21.1-k3s1"></a>
## [v1.5.3-v1.21.1-k3s1] - 2021-06-08
### Bug Fixes
- **haos:** corrected haos build

### Documentation Improvements
- **changelog:** auto changelog update v1.5.2-v1.21.0-k3s1 [CI SKIP]

### Features
- **HAOS:** add haos initial support
- **upgrade plan:** switched channel to BLAOS

### Maintenance
- **drone:** add haos build to drone


<a name="v1.5.2-v1.21.0-k3s1"></a>
## [v1.5.2-v1.21.0-k3s1] - 2021-06-08
### Bug Fixes
- **dapper:** musl-dev version

### Dependency Updates
- **10-gobuild:** bump musl-dev

### Documentation Improvements
- **changelog:** auto changelog update v1.5.1-v1.21.0-k3s1 [CI SKIP]

### Features
- **images 01-vm:** output virtual version of kernel for update
- **kernel:** bump kernel version
- **package-update:** add package-update step

### Maintenance
- **blaos-vm:** add kernel to vm output
- **ci:** add package-update to ci
- **dapper:** version lock deps
- **dapper:** add makeself to dapper env
- **deps:** remove redundant U arg
- **drone:** Add echo to track each section
- **drone:** reorganized drone yaml
- **drone:** add master ref to mirror
- **drone:** adjusted file path
- **drone:** add support for package-update
- **drone:** removed multi steps
- **drone:** corrected to capture amd non vm update
- **updates:** adjust path for update package script

### Reverts
- fix(dapper): musl-dev version


<a name="v1.5.1-v1.21.0-k3s1"></a>
## [v1.5.1-v1.21.0-k3s1] - 2021-06-03
### Dependency Updates
- **alpine:** bump to version 3.13
- **package managers:** lock package versions for deterministic installation

### Documentation Improvements
- **changelog:** auto changelog update v1.4.0-v1.21.0-k3s1 [CI SKIP]

### Maintenance
- **gitignore:** updated gitignore for dapper files


<a name="v1.4.0-v1.21.0-k3s1"></a>
## [v1.4.0-v1.21.0-k3s1] - 2021-06-01
### Bug Fixes
- **Images 06-base:** correct kots install
- **images 06-base:** change name of kots to kubectl-kots
- **images 06-base:** moved kots to sbin

### Documentation Improvements
- **changelog:** auto changelog update v1.3.0-v1.21.0-k3s1 [CI SKIP]
- **readme:** updated shields

### Features
- **k3s:** bump to version v1.21.1+k3s1
- **kubectl pugins:** add kots plugin


<a name="v1.3.0-v1.21.0-k3s1"></a>
## [v1.3.0-v1.21.0-k3s1] - 2021-05-30
### Bug Fixes
- **apparmor:** disable apparmor
- **drone:** add from_secrets flag
- **drone:** add missing brackets
- **startup:** Apply hostname, DNS, NTP, Wifi, password furing startup
- **update-issue:** improved font choice

### Dependency Updates
- **k3os:** update to golang 1.16.x
- **vendor:** tidy vendors

### Documentation Improvements
- **changelog:** auto changelog update v1.2.1-v1.20.4-k3s1 [CI SKIP]
- **changelog:** auto changelog update v1.2.1-v1.20.4-k3s1 [CI SKIP]

### Features
- **k3s:** bump to v1.21.0+k3s1
- **kernel:** bump kernel to v5.4.0-72.80-rancher1
- **package:** add smartmontools
- **system-upgrade-controller:** bump to v0.7.0

### Maintenance
- **docker:** update docker ignore
- **drone:** revert package names to k3os
- **drone:** add cache step to build
- **drone:** swapped git sync method
- **update-issue:** update to BLAOS

### Style Improvements
- **motd:** improve the motd
- **update-issue:** improved issue text

### Reverts
- fix(drone): add from_secrets flag
- fix(drone): add missing brackets


<a name="v1.2.1-v1.20.4-k3s1"></a>
## [v1.2.1-v1.20.4-k3s1] - 2021-05-25
### Documentation Improvements
- **changelog:** auto changelog update v1.2.0-v1.20.4-k3s1 [CI SKIP]

### Maintenance
- **drone:** change sync plugin


<a name="v1.2.0-v1.20.4-k3s1"></a>
## [v1.2.0-v1.20.4-k3s1] - 2021-05-24
### Bug Fixes
- **drone:** apply change to drone
- **drone:** add target and trigger

### Documentation Improvements
- **changelog:** auto changelog update v1.0.0-v1.20.4-k3s1 [CI SKIP]

### Features
- **drone:** add repo sync
- **images:** add vm version iso


<a name="v1.0.0-v1.20.4-k3s1"></a>
## v1.0.0-v1.20.4-k3s1 - 2021-05-24
### Bug Fixes
- **05-base:** add missing open-vm-tools packages to dockerfile
- **cc:** add missing struct
- **drone:** add additional conditions
- **drone:** change to refs/tags
- **drone:** add privileged
- **scripts:** strip + from tag name

### Documentation Improvements
- **CONTRIBUTING:** Add contributing guide
- **issue template:** Create issue templates
- **pull requests:** create general pull request
- **templates:** add research proposal

### Features
- **CC:** Add manifest to cloud config

### Maintenance
- **.gitignore:** add gitignore file
- **changelog:** add changelog config & Template
- **dockerignore:** add docker ignore file
- **drone:** updated drone config
- **drone:** remove privileged
- **drone:** add drone config
- **drone:** change to ref
- **drone:** updated build tasks

### Style Improvements
- **cc:** correct gofmt style errors
- **drone:** specify host


[Unreleased]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.5.4-v1.21.1-k3s1...HEAD
[v1.5.4-v1.21.1-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.5.3-v1.21.1-k3s1...v1.5.4-v1.21.1-k3s1
[v1.5.3-v1.21.1-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.5.2-v1.21.0-k3s1...v1.5.3-v1.21.1-k3s1
[v1.5.2-v1.21.0-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.5.1-v1.21.0-k3s1...v1.5.2-v1.21.0-k3s1
[v1.5.1-v1.21.0-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.4.0-v1.21.0-k3s1...v1.5.1-v1.21.0-k3s1
[v1.4.0-v1.21.0-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.3.0-v1.21.0-k3s1...v1.4.0-v1.21.0-k3s1
[v1.3.0-v1.21.0-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.2.1-v1.20.4-k3s1...v1.3.0-v1.21.0-k3s1
[v1.2.1-v1.20.4-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.2.0-v1.20.4-k3s1...v1.2.1-v1.20.4-k3s1
[v1.2.0-v1.20.4-k3s1]: https://github.com/BOHICA-LABS/BLAOS/compare/v1.0.0-v1.20.4-k3s1...v1.2.0-v1.20.4-k3s1
