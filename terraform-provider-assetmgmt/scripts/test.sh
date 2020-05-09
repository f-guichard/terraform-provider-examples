#!/bin/bash

python ../mock-hpam-server/fake_hpam.py #Bloquant
##SI backend http
terraform init ../tf_examples
##
terraform apply --auto-approve ../tf_examples
