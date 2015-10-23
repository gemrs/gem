# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "jimmidyson/centos-7.1"

  config.vm.network "forwarded_port", guest: 80, host: 80
  config.vm.network "forwarded_port", guest: 43594, host: 43594
  config.vm.network "forwarded_port", guest: 43595, host: 43595

  config.vm.synced_folder ".", "/vagrant/src/github.com/sinusoids/gem"
  config.ssh.forward_x11 = true

  config.vm.provider "virtualbox" do |v|
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.provider "libvirt" do |lv, override|
    override.vm.synced_folder ".", "/vagrant/src/github.com/sinusoids/gem", create: true, :nfs => true, :mount_options => ['nolock,vers=3,tcp,noatime'], id: "vagrant-root"
  end

  config.vm.provision "shell", inline: <<-SHELL
    sudo yum -y install epel-release
    sudo yum -y install git gdb python-devel pytest python-pip libffi-devel
    sudo yum -y groupinstall "Development Tools"
    sudo pip install yapsy

    mkdir -p /vagrant/src/
    mkdir -p /home/vagrant/go/
    chown -R vagrant:vagrant /home/vagrant/go/
    chown -R vagrant:vagrant /vagrant
    wget -nv https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz
    echo 'export GOROOT=/usr/local/go' >> /home/vagrant/.profile
    echo 'export GOPATH=/vagrant' >> /home/vagrant/.profile
    echo 'export PATH=$GOROOT/bin:$GOPATH/bin:$PATH' >> /home/vagrant/.profile
    echo 'export GO15VENDOREXPERIMENT=1' >> /home/vagrant/.profile
    echo 'source $HOME/.profile' >> /home/vagrant/.bashrc
    # Would be nice if gb could handle this, but it doesn't seem to build binaries from vendored packages
    su - vagrant -c 'go get github.com/tgascoigne/gopygen'
    su - vagrant -c 'go get github.com/blynn/nex'
  SHELL
end
