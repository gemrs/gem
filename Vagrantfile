# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "jimmidyson/centos-7.1"

  config.vm.network "forwarded_port", guest: 80, host: 80 
  config.vm.network "forwarded_port", guest: 43594, host: 43594
  config.vm.network "forwarded_port", guest: 43595, host: 43595

  config.vm.synced_folder ".", "/vagrant"

  config.vm.provider "libvirt" do |lv, override|
    override.vm.synced_folder ".", "/vagrant", create: true, :nfs => true, :mount_options => ['nolock,vers=3,tcp,noatime'], id: "vagrant-root"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo yum -y install epel-release
    sudo yum -y install git gdb golang python-devel pytest python-pip libffi-devel
    sudo pip install yapsy

    mkdir -p /home/vagrant/go/
    chown -R vagrant:vagrant /home/vagrant/go/
    echo 'export GOPATH=$HOME/go' >> /home/vagrant/.profile
    echo 'export PATH=$PATH:$GOPATH/bin' >> /home/vagrant/.profile
    echo 'source $HOME/.profile' >> /home/vagrant/.bashrc
    su - vagrant -c 'go get github.com/constabulary/gb/...'
  SHELL
end
