module "ec2" {
  source = "terraform-aws-modules/ec2-instance/aws"

  name          = "manager-ec2"
  instance_type = "t2.micro"
  key_name      = "manager-key"
  ami           = "ami-0c55b159cbfafe1f0"
  subnet_id     = module.vpc.public_subnets[0]

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}
