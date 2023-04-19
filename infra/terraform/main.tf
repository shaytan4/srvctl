# Specify the provider and region
provider "aws" {
  region = var.aws_region
}

# Create a security group that allows SSH and HTTP access
resource "aws_security_group" "web" {
  name        = var.aws_sec_gr
  description = "Allow in out traffic to our app"

  ingress {
    description = "Allow ssh acces , 22 port"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    #security_groups = [aws_security_group.web.id]
  }

  ingress {
    description = "Allow http traffic"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow https traffic"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow extra web traffic 8080 tcp port"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow traffic from me to any"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Create an EC2 instance using the AMI and security group
resource "aws_instance" "web" {
  ami             = lookup(var.aminame, var.aws_region)
  instance_type   = var.inst_type
  security_groups = [aws_security_group.web.name]
  key_name        = var.key_name
  count           = var.master_count
  #ebs_optimized        = true
  #iam_instance_profile = "test"
  #encrypted     = true
  monitoring = true

  tags = {
    Name = "${var.inst_name}-${count.index + 1}"
  }

  # https://developer.hashicorp.com/terraform/language/resources/provisioners/connection
  connection {
    # Host name
    host = self.public_ip
    # The default username for our AMI
    user = "ec2-user"
    # Private key for connection    
    private_key = file(var.private_key)
    # Type of connection
    type = "ssh"
  }

  # Installing docker
  provisioner "remote-exec" {
    inline = [
      "sudo yum update -y",
      "sudo yum install -y yum-utils",
      "sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo",
      "sudo yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y",
      "sudo systemctl start docker",
      "sudo systemctl enable docker",
      "sudo yum install -y git"
    ]
    on_failure = continue
  }
}