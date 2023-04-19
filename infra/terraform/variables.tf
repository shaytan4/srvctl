# Defining Region
variable "aws_region" {
  type        = string
  default     = "eu-north-1"
  description = "Set aws region"
}

variable "inst_name" {
  type        = string
  default     = "mycoolapp3"
  description = "Name for instance"
}

variable "aws_sec_gr" {
  type        = string
  default     = "appsecgroup3"
  description = "Name for security group"
}

variable "aminame" {
  type = map(string)
  default = {
    eu-north-1 = "ami-0a6351192ce04d50c" # RedHat
  }
  description = "OS for installation"
}

variable "master_count" {
  type        = number
  default     = 1
  description = "Number of instances to deploy"
}

variable "inst_type" {
  type        = string
  default     = "t3.micro"
  description = "Set type of instance"
}

variable "key_name" {
  type        = string
  default     = "KonstHome"
  description = "Desired name of AWS key pair, ssh key loaded to AWS"
}

# Defining Private Key
variable "private_key" {
  type        = string
  default     = "~/.ssh/id_rsa"
  description = "Path to privat ssh key"
}
