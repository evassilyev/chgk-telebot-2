# chgk-telebot-2
Телеграмм-бот для игры в "Что? Где? Когда?". Версия 2.

---
# Installation  

## Prerequiements
The goal of this section is to configure the server (VPS) for working with commands from the Makefile.

### Server
1. Create new user on the server with home directory. Set password for him. Add him to sudoers.
2. Generate ssh keys pair on your local (development) machine and copy the public key into the server (VPS).  
3. In the file `/etc/ssh/sshd_config` on the server (VPS)  change the settings listed below:
- Change default `Port` value to YOUR_PORT_VALUE
- `PermitRootLogin` `no`
- `PubkeyAuthentication` `yes`
- Set `AuthorizedKeysFile` to the path where you stored the public ssh key from your local (development) machine
### Local machine
Configure SSH client on your local (development) machine (`.ssh/config`) by setting values from the server (VPS)
Example:
```shell
Host my_remote_vps
    Port YOUR_PORT_VALUE
    User vps_user
    HostName 8.8.8.8
    StrictHostKeyChecking no
    IdentityFile ~/.ssh/my_private_ssh_key
```
### Domain name
Create domain or subdomain and link it to your VPS IP address
