# REPORT
  Первым делом, я инициализировал Vagrant
проект командой: `vagrant init`.

Далее, я приступил к конфигурации виртуалки.
В качестве образа я выбрал CentOS, выдал имя хоста потому, что мне так захотелось 0.0
```ruby
  config.vm.box = "centos/7"
  config.vm.hostname = "centos-slave"
```

Отконфигурировал количество используемых ядер и объём используемого ОЗУ, а также
выдал имя для виртуалки для удобства дальнейшего использования.
```ruby
  config.vm.provider "virtualbox" do |v|
    v.name = "centos-slave-1"
    v.memory = 1024
    v.cpus = 2
  end
```

Далее перешёл к конфигурации SSH.
```ruby
  config.ssh.host = "127.0.0.1"
  config.ssh.port = 4444
  config.vm.network "forwarded_port", guest: 22,
  host: 4444
```
Важно! Нужно пробрасывать SSH порт вручную,
иначе вы не будете иметь возможности
подключится по SSH.
А во время запуска командой `vagrant up` вы увидете сообщение:
```bash
  default: Warning: Connection refused. Retrying...
```

Покончили с конфигурацией, далее я запустил виртуалку командой: vagrant up.
```
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Importing base box 'centos/7'...
==> default: Matching MAC address for NAT networking...
==> default: Checking if box 'centos/7' version '2004.01' is up to date...
==> default: Setting the name of the VM: centos-slave-1
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
==> default: Forwarding ports...
    default: 22 (guest) => 4444 (host) (adapter 1)
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:4444
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: 
    default: Vagrant insecure key detected. Vagrant will automatically replace
    default: this with a newly generated keypair for better security.
    default: 
    default: Inserting generated public key within guest...
    default: Removing insecure key from the guest if it's present...
    default: Key inserted! Disconnecting and reconnecting using new SSH key...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
    default: No guest additions were detected on the base box for this VM! Guest
    default: additions are required for forwarded ports, shared folders, host only
    default: networking, and more. If SSH fails on this machine, please install
    default: the guest additions and repackage the box to continue.
    default: 
    default: This is not an error message; everything may continue to work properly,
    default: in which case you may ignore this message.
==> default: Setting hostname...
==> default: Rsyncing folder: /home/paprika/paprika/itransition/DevOps_courses/2.Virtualization_and_Containerization/task1/ 
=> /vagrant
```
Уже после запуска я SSH-нулся в виртуалку командой: vagrant ssh
```bash
task1 · master± ⟩ vagrant ssh
Last login: Thu Oct 12 13:14:15 2023 from 10.0.2.2
[vagrant@centos-slave ~]$ 
```

Внутри я так же решил поиграться с консолью, лишь бы проверить "а всё ли собственно нормально?"

```bash
[vagrant@centos-slave ~]$ ls -a
.  ..  .bash_history  .bash_logout  .bash_profile  .bashrc  .ssh
[vagrant@centos-slave ~]$ hostname
centos-slave
[vagrant@centos-slave ~]$ 
```

Всё близится к концу, а значит самое время
отключить(vagrant halt) и удалить(vagrant
destroy) виртуалку.

```bash
task1 · master± ⟩ vagrant halt
==> default: Attempting graceful shutdown of VM...

task1 · master± ⟩ vagrant destroy
    default: Are you sure you want to destroy the 'default' VM? [y/N] y
==> default: Destroying VM and associated drives...

task1 · master± ⟩ 
```

Если вы удалили директорию с вашей виртуалкой,
можно выполнить эти же команды,
но получив ID виртуалки из таблицы полученной
из команды `vagrant global-status`.
```bash
task1 · master± ⟩ vagrant global-status
id       name   provider state  directory                           
--------------------------------------------------------------------
There are no active Vagrant environments on this computer! Or,
you haven't destroyed and recreated Vagrant environments that were
started with an older version of Vagrant.
```