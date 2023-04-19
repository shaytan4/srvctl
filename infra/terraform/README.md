<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=1.4.4 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 4.63.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 4.63.0 |

## Resources

| Name | Type |
|------|------|
| [aws_instance.web](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance) | resource |
| [aws_security_group.web](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/security_group) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aminame"></a> [aminame](#input\_aminame) | OS for installation | `map(string)` | ```{ "eu-north-1": "ami-0a6351192ce04d50c" }``` | no |
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | Set aws region | `string` | `"eu-north-1"` | no |
| <a name="input_aws_sec_gr"></a> [aws\_sec\_gr](#input\_aws\_sec\_gr) | Name for security group | `string` | `"appsecgroup3"` | no |
| <a name="input_inst_name"></a> [inst\_name](#input\_inst\_name) | Name for instance | `string` | `"mycoolapp3"` | no |
| <a name="input_inst_type"></a> [inst\_type](#input\_inst\_type) | Set type of instance | `string` | `"t3.micro"` | no |
| <a name="input_key_name"></a> [key\_name](#input\_key\_name) | Desired name of AWS key pair, ssh key loaded to AWS | `string` | `"KonstHome"` | no |
| <a name="input_master_count"></a> [master\_count](#input\_master\_count) | Number of instances to deploy | `number` | `1` | no |
| <a name="input_private_key"></a> [private\_key](#input\_private\_key) | Path to privat ssh key | `string` | `"~/.ssh/id_rsa"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_ec2_fqdn"></a> [ec2\_fqdn](#output\_ec2\_fqdn) | EC2 public fqdn |
| <a name="output_ec2_ip"></a> [ec2\_ip](#output\_ec2\_ip) | EC2 public ip addr |
<!-- END_TF_DOCS -->