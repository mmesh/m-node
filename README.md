![mmesh](https://github.com/mmesh/assets/blob/HEAD/images/logo/mmesh_logo_v4_240x40_darkgrey.png)

[![Discord](https://img.shields.io/discord/654291649572241408?color=%236d82cb&style=flat&logo=discord&logoColor=%23ffffff&label=Chat)](https://mmesh.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/mmesh/discussions)
[![Twitter](https://img.shields.io/badge/Follow_on_Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white)](https://twitter.com/mmesh_io)

Open source projects from [mmesh](https://mmesh.io).

# mmesh-node

[![Release](https://img.shields.io/github/v/release/mmesh/m-node?display_name=tag&style=flat)](https://github.com/mmesh/m-node/releases/latest)
[![GitHub](https://img.shields.io/github/license/mmesh/m-node?style=flat)](/LICENSE)

This repository contains the `mmesh-node` agent, the component that runs on the machines you want to connect to your [mmesh](https://mmesh.io) network.

`mmesh-node` is available for a variety of Linux platforms, macOS and Windows.

## Minimum Requirements

`mmesh-node` has the same [minimum requirements](https://github.com/golang/go/wiki/MinimumRequirements#minimum-requirements) as Go:

- Linux kernel version 2.6.23 or later
- Windows 7 or later
- FreeBSD 11.2 or later
- MacOS 10.11 El Capitan or later

## Getting Started

The instructions in this repo assume you already have a mmesh account and are ready to start adding nodes.

See [Quick Start](https://mmesh.io/docs/platform/getting-started/quickstart) to learn how to start building your mmesh.

The fastest way to add Linux nodes to your mmesh is by [generating a magic link](#linux-installation-with-magic-link) in the mmesh web UI or with `mmeshctl`:

```shell
mmeshctl node add
```

See [Installation](#installation) for more details and other platforms.

## Documentation

For the complete mmesh platform documentation visit [mmesh.io/docs](https://mmesh.io/docs).

## Installation

### Binary Downloads

Linux, macOS and Windows binary downloads are available from the [Releases](https://github.com/mmesh/m-node/releases) page.

You can download the pre-compiled binaries and install them with the appropriate tools.

### Linux Installation

#### Linux installation with magic link

The easiest way to add nodes to your mmesh is by generating a magic link in the mmesh web UI or with `mmeshctl`:

```shell
mmeshctl node add
```

You will be able to use the magic link to install the `mmesh-node` agent in seconds with no additional configuration required.

Once installed you can review the configuration at `/etc/mmesh/mmesh-node.yml`.

> See the [mmesh-node configuration reference](https://mmesh.io/docs/platform/reference/mmesh-node.yml) to find all the configuration options.

#### Linux binary installation with curl

1. Download the latest release.

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/linux/amd64/mmesh-node"
    ```

2. Validate the binary (optional).

    Download the mmesh-node checksum file:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/linux/amd64/mmesh-node_checksum.sha256"
    ```

    Validate the mmesh-node binary against the checksum file:

    ```bash
    sha256sum --check < mmesh-node_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    mmesh-node: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    mmesh-node: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install mmesh-node and create its configuration file according to your needs.

    ```shell
    sudo install -o root -g root -m 0750 mmesh-node /usr/local/bin/mmesh-node
    sudo mkdir /etc/mmesh
    sudo vim /etc/mmesh/mmesh-node.yml
    ```

    See the [mmesh-node configuration reference](https://mmesh.io/docs/platform/reference/mmesh-node.yml) to find all the configuration options.

4. Create the `mmesh-node.service` for systemd.

    ```shell
    sudo cat << EOF > /etc/systemd/system/mmesh-node.service
    [Unit]
    Description=mmesh-node service
    Documentation=https://github.com/mmesh/m-node
    After=network.target

    [Service]
    Type=simple

    # Another Type: forking

    # User=
    WorkingDirectory=/var/local/mmesh
    ExecStart=/usr/local/bin/mmesh-node start
    Restart=always

    # Other restart options: always, on-abort, etc

    # The install section is needed to use

    # 'systemctl enable' to start on boot

    # For a user service that you want to enable

    # and start automatically, use 'default.target'

    # For system level services, use 'multi-user.target'

    [Install]
    WantedBy=multi-user.target
    EOF
    ```

5. Ensure the `tun` kernel module is loaded.

    ```shell
    sudo modprobe tun
    ```

6. Start the mmesh-node service.

    ```shell
    sudo systemctl daemon-reload
    sudo systemctl enable mmesh-node
    sudo systemctl restart mmesh-node
    ```

##### Uninstall Linux mmesh-node

To remove `mmesh-node` from the system, use the following commands:

```shell
sudo systemctl stop mmesh-node
sudo systemctl disable mmesh-node
sudo rm /etc/systemd/system/mmesh-node.service
sudo systemctl daemon-reload
sudo rm /usr/local/bin/mmesh-node
sudo rm /etc/mmesh/mmesh-node.yml
sudo rmdir /etc/mmesh
```

#### Package Repository

mmesh provides a package repository that contains both DEB and RPM downloads.

For DEB-based platforms (e.g. Ubuntu and Debian) run the following to setup a new APT sources.list entry and install `mmesh-node`:

```shell
echo 'deb [trusted=yes] https://repo.mmesh.io/apt/ /' | sudo tee /etc/apt/sources.list.d/mmesh.list
sudo apt update
sudo apt install mmesh-node
```

For RPM-based platforms (e.g. RHEL, CentOS) use the following to create a repo file and install `mmesh-node`:

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/mmesh.repo
[mmesh]
name=mmesh Repository - Stable
baseurl=https://repo.mmesh.io/yum
enabled=1
gpgcheck=0
EOF
sudo yum install mmesh-node
```

### macOS Installation

#### macOS binary installation with curl

1. Download the latest release.

    **Intel**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/amd64/mmesh-node"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/arm64/mmesh-node"
    ```

2. Validate the binary (optional).

    Download the mmesh-node checksum file:

    **Intel**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/amd64/mmesh-node_checksum.sha256"
    ```

    **Apple Silicon**:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/darwin/arm64/mmesh-node_checksum.sha256"
    ```

    Validate the mmesh-node binary against the checksum file:

    ```console
    shasum --algorithm 256 --check mmesh-node_checksum.sha256
    ```

    If valid, the output must be:

    ```console
    mmesh-node: OK
    ```

    If the check fails, sha256 exits with nonzero status and prints output similar to:

    ```console
    mmesh-node: FAILED
    sha256sum: WARNING: 1 computed checksum did NOT match
    ```

3. Install `mmesh-node` and create its configuration file according to your needs.

    ```console
    chmod +x mmesh-node
    sudo mkdir -p /usr/local/libexec
    sudo mv mmesh-node /usr/local/libexec/mmesh-node
    sudo chown root: /usr/local/libexec/mmesh-node
    sudo mkdir /etc/mmesh
    sudo vim /etc/mmesh/mmesh-node.yml
    sudo chmod 600 /etc/mmesh/mmesh-node.yml
    ```

    > **IMPORTANT**: In macOS, `agent.iface` must be `utun[0-9]+` in the `mmesh-node.yml`, being `utun5` usually a good choice for that setting. Use the command `ifconfig -a` before launching the `mmesh-node` service and check that the interface is not in-use.

    See the [mmesh-node configuration reference](https://mmesh.io/docs/platform/reference/mmesh-node.yml) to find all the configuration options.

4. Install and start the mmesh-node agent as a system service.

    ```shell
    sudo /usr/local/libexec/mmesh-node service-install
    ```

5. Check the service status.

    ```shell
    launchctl print system/io.mmesh.mmesh-node
    ```

    You should get an output like this:

    ```console
    system/io.mmesh.mmesh-node = {
        active count = 1
        path = /Library/LaunchDaemons/io.mmesh.mmesh-node.plist
        state = running

        program = /usr/local/libexec/mmesh-node
        arguments = {
            /usr/local/libexec/mmesh-node
            service-start
        }

        working directory = /var/tmp

        stdout path = /usr/local/var/log/io.mmesh.mmesh-node.out.log
        stderr path = /usr/local/var/log/io.mmesh.mmesh-node.err.log
        default environment = {
            PATH => /usr/bin:/bin:/usr/sbin:/sbin
        }

        environment = {
            XPC_SERVICE_NAME => io.mmesh.mmesh-node
        }

        domain = system
        minimum runtime = 10
        exit timeout = 5
        runs = 1
        pid = 3925
        immediate reason = speculative
        forks = 28
        execs = 1
        initialized = 1
        trampolined = 1
        started suspended = 0
        proxy started suspended = 0
        last exit code = (never exited)

        spawn type = daemon (3)
        jetsam priority = 4
        jetsam memory limit (active) = (unlimited)
        jetsam memory limit (inactive) = (unlimited)
        jetsamproperties category = daemon
        submitted job. ignore execute allowed
        jetsam thread limit = 32
        cpumon = default

        properties = keepalive | runatload | inferred program
    }
    ```

##### Uninstall macOS mmesh-node

To remove `mmesh-node` from the system, use the following commands:

```shell
sudo /usr/local/libexec/mmesh-node service-uninstall
sudo rm /usr/local/libexec/mmesh-node
sudo rm /etc/mmesh/mmesh-node.yml
sudo rmdir /etc/mmesh
```

### Windows Installation

#### Windows binary installation with curl

1. Open the Command Prompt as Administrator and create a folder for mmesh.

    ```shell
    mkdir 'C:\Program Files\mmesh'
    ```

2. Download the latest release into the mmesh folder.

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/windows/amd64/mmesh-node.exe"
    ```

3. Validate the binary (optional).

    Download the mmesh-node.exe checksum file:

    ```shell
    curl -LO "https://mmesh.s3.nl-ams.scw.cloud/binaries/stable/latest/windows/amd64/mmesh-node.exe_checksum.sha256"
    ```

    Validate the mmesh-node.exe binary against the checksum file:

    - Using Command Prompt to manually compare CertUtil's output to the checksum file downloaded:

         ```shell
         CertUtil -hashfile mmesh-node.exe SHA256
         type mmesh-node.exe_checksum.sha256
         ```

    - Using PowerShell to automate the verification using the -eq operator to get a `True` or `False` result:

         ```powershell
         $($(CertUtil -hashfile .\mmesh-node.exe SHA256)[1] -replace " ", "") -eq $(type .\mmesh-node.exe_checksum.sha256).split(" ")[0]
         ```

4. Download the `wintun` driver from <https://wintun.net>.

5. Unzip the wintun archive and copy the AMD64 binary `wintun.dll` to `C:\Program Files\mmesh`.

6. Use an editor to create the mmesh-node configuration file `C:\Program Files\mmesh\mmesh-node.yml`.

    See the [mmesh-node configuration reference](https://mmesh.io/docs/platform/reference/mmesh-node.yml) to find all the configuration options.

7. Install the mmesh-node agent as a Windows service.

    > The instructions below assume that the `wintun.dll`, `mmesh-node.exe` and `mmesh-node.yml` files are stored in `C:\Program Files\mmesh`.

    ```shell
    .\mmesh-node.exe service-install --config "C:\Program Files\mmesh\mmesh-node.yml"
    ```

    Make sure to provide the absolute path of the mmesh-node.yml configuration file, otherwise the Windows service may fail to start.

8. Start the service.

    ```shell
    net start "mmesh-node"
    ```

##### Uninstall Windows mmesh-node

To remove `mmesh-node` from the system, open the Command Prompt as Administrator and use the following commands:

```shell
net stop "mmesh-node"
cd 'C:\Program Files\mmesh'
.\mmesh-node.exe service-uninstall
del *.*
cd ..
rmdir 'C:\Program Files\mmesh'
```

## Artifacts Verification

### Binaries

All artifacts are checksummed and the checksum file is signed with [cosign](https://github.com/sigstore/cosign).

1. Download the files you want and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [Releases](https://github.com/mmesh/m-node/releases) page:

2. Verify the signature:

    ```shell
    cosign verify-blob \
        --cert checksums.txt.pem \
        --signature checksums.txt.sig \
        checksums.txt
    ```

3. If the signature is valid, you can then verify the SHA256 sums match with the downloaded binary:

    ```shell
    sha256sum --ignore-missing -c checksums.txt
    ```

### Docker Images

Our Docker images are signed with [cosign](https://github.com/sigstore/cosign).

Verify the signatures:

```console
COSIGN_EXPERIMENTAL=1 cosign verify mmeshdev/mmesh-node
```

## Configuration

See the [mmesh-node configuration reference](https://mmesh.io/docs/platform/reference/mmesh-node.yml) to find all the configuration options.

## Running with Docker

You can also run the `mmesh-node` agent as a Docker container. See examples below.

Registries:

- `mmeshdev/mmesh-node`
- `ghcr.io/mmesh/mmesh-node`

Example usage:

```shell
docker run -d --restart=always \
  --net=host \
  --cap-add=net_admin \
  --device=/dev/net/tun \
  --name mmesh-node \
  -v /etc/mmesh:/etc/mmesh:ro \
  mmeshdev/mmesh-node:latest start
```

## Community

Have questions, need support and or just want to talk about mmesh?

Get in touch with the mmesh community!

[![Discord](https://img.shields.io/badge/Join_us_on_Discord-5865F2?style=flat&logo=discord&logoColor=white)](https://mmesh.io/discord)
[![GitHub Discussions](https://img.shields.io/badge/GitHub_Discussions-181717?style=flat&logo=github&logoColor=white)](https://github.com/orgs/mmesh/discussions)
[![Twitter](https://img.shields.io/badge/Follow_on_Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white)](https://twitter.com/mmesh_io)

## Code of Conduct

Participation in the mmesh community is governed by the Contributor Covenant [Code of Conduct](https://github.com/mmesh/.github/blob/HEAD/CODE_OF_CONDUCT.md). Please make sure to read and observe this document.

Please make sure to read and observe this document. By participating, you are ex
pected to uphold this code.

## License

The mmesh open source projects are licensed under the [Apache 2.0 License](/LICENSE).

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmmesh%2Fm-node.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmmesh%2Fm-node?ref=badge_large)
