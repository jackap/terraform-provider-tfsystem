all: 
	go build -o terraform-provider-example
	terraform init
	TF_LOG=TRACE	terraform apply
