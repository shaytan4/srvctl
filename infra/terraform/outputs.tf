#####################################
##            Output               ##
#####################################

output "ec2_ip" {
  description = "EC2 public ip addr"
  value       = aws_instance.web[*].public_ip
}

output "ec2_fqdn" {
  description = "EC2 public fqdn"
  value       = aws_instance.web[*].public_dns
}

