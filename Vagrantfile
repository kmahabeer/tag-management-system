# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.
  
  project_name            = "tagging-service"
  vm_variant              = "A"
  vm_environment          = "dev"
  vm_number               = 1
  vm_memory               = 4096
  box_name                = "ubuntu/jammy64"
  box_version             = "20241002.0.0"
  synced_folder_host      = "."
  synced_folder_guest     = "/home/vagrant/#{project_name}" # Note: Do NOT set the guest to `/home/vagrant/` otherwise it will remove your ssh keys
  vm_name                 = "#{project_name}-#{vm_variant}-#{vm_environment}-#{format('%02d', vm_number)}"
  

  config.vm.define vm_name do |machine|
    machine.vm.hostname = vm_name
    machine.vm.box = box_name
    machine.vm.box_version = box_version

    # Port forwarding
    ports = [
      8000
    ]
    ports.each do |port|
      machine.vm.network "forwarded_port", guest: port, host: port
    end

    # Shared folder
    machine.vm.synced_folder synced_folder_host, synced_folder_guest

    # VM provider settings
    machine.vm.provider "virtualbox" do |vb|
      vb.name = vm_name
      vb.memory = vm_memory
      vb.gui = true
    end

    # The following will install 
    machine.vm.provision "shell", inline: <<-SHELL
      apt-get update -y
      apt install -y python3-pip
      apt install -y python3-venv
      apt install -y openjdk-11-jdk
      pip install --no-cache-dir -r requirements.txt
    SHELL

    machine.vm.provision "shell", inline: <<-SHELL
      export PATH=$PATH:/home/vagrant/.local/bin
      export JAVA_HOME=/usr/lib/jvm/java-11-openjdk-amd64
      echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
      source ~/.bashrc
    SHELL
  end
end