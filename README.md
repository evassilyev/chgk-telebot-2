# chgk-telebot-2
Телеграмм-бот для игры в "Что? Где? Когда?". Версия 2.

---
# Installation  

## Prerequiements
The goal of this section is to configure the server (VPS) for working with commands from the Makefile.

### Server
1. Create new user with name[^1] and password[^2] on the server. Set password for him. Add him to sudoers.
2. Generate ssh keys pair on your local (development) machine and copy the public key into the server (VPS).  
3. In the file `/etc/ssh/sshd_config` on the server (VPS)  change the settings listed below:
- Change default `Port` value to YOUR_PORT_VALUE
- `PermitRootLogin` `no`
- `PubkeyAuthentication` `yes`
- Set `AuthorizedKeysFile` to the path where you stored the public ssh key from your local (development) machine  
[^1]: Used in configuration file `vars.mk` as `USER` value  
[^2]: Used in configuration file `vars.mk` as `PASSWORD` value
### Local machine
Configure host[^3] for your SSH client on your local (development) machine (`.ssh/config`)
Example:
```shell
Host my_remote_vps
    Port YOUR_PORT_VALUE
    User vps_user
    HostName 8.8.8.8
    StrictHostKeyChecking no
    IdentityFile ~/.ssh/my_private_ssh_key
```
[^3]: Used in configuration file `vars.mk` as `SERVER` value
### Domain name
Create subdomain[^4] and link it to your VPS IP address  
[^4]: Used in configuration file `vars.mk` as `VPS` value

## Configuration files
### vars.mk
Configuration file used in the Makefile. Must be placed in the root of repository.  
Values and description:  
```shell
DAYS?=30 # duration of the certificate validity in days
VPS?=SUBDOMAIN # subdomain used in certificate
SERVER?=my_remote_vps # name for ssh connection
COMPANY?=Vector # Company name used in certificate generation
USER=MYUSER # VPS user name
SUDOPASS=password # VPS user password
BASE_VERSION?=0.0a # bot version
DOCKER_NETWORK_NAME=dockernet # random name for docker internal communications
ENV_PATH=cmd/bot/.env # path to environment file that bot use
REMOTE_DB_IP=8.8.8.8 # remote database IP
REMOTE_DB_PORT=222 # remote database port
SERVICE_NAME=mysrv # service name used in Docker image & container building
```
