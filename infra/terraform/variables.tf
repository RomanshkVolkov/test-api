variable "access_key" {
  type        = string
  description = "AWS access key"
}

variable "secret_key" {
  type        = string
  description = "AWS secret key"
}

variable "region" {
  type        = string
  default     = "us-east-1"
  description = "AWS region"
}
