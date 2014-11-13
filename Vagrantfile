# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = '2'

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.ssh.private_key_path = [ '~/.vagrant.d/insecure_private_key', '~/.ssh/id_rsa' ]
  config.ssh.forward_agent = true

  config.vm.synced_folder '.', '/home/vagrant/src/github.com/mperham/inspeqtor'

  install_go_script = <<-SCRIPT
    curl https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz -O
    tar -C /usr/local -xvf go1.3.3.linux-amd64.tar.gz
    rm go1.3.3.linux-amd64.tar.gz
  SCRIPT

  environment_script = <<-SCRIPT
    echo export PATH=$PATH:/usr/local/go/bin >> ~/.profile
    echo export GOPATH=/home/vagrant >> ~/.profile
  SCRIPT

  ownership_script = <<-SCRIPT
    chown -R vagrant:vagrant ~/src
  SCRIPT

  config.vm.provision :shell, inline: install_go_script
  config.vm.provision :shell, inline: environment_script
  config.vm.provision :shell, inline: ownership_script

  config.vm.define 'centos-6.5' do |centos|
    centos.vm.box = 'chef/centos-6.5'
  end

  config.vm.define 'ubuntu-12.04' do |ubuntu|
    ubuntu.vm.box = 'hashicorp/precise64'
  end

  config.vm.define 'ubuntu-14.04' do |ubuntu|
    ubuntu.vm.box = 'ubuntu/trusty64'
  end
end
