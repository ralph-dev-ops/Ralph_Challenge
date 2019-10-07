# Infrastructure

I chose to use AWS Elastic Beanstalk for the task at hand as it provides everything needed to host a static web application with minimal effort.

I initially was going to use cloud formation to for Elastic Beanstalk, but was able to leverage the AWS CLI commands for Elastic Beanstalk.
https://github.com/aws/aws-elastic-beanstalk-cli-setup

After initial configuration of the Elastic Beanstalk tools, the `01_create.sh` script is used to create the initial web stack as well as an initial deploy.

For server validation, the `02_validate.sh` script is used to check for an HTTP redirect and also the served content of the website.

The validation script will only work correctly after the target group has completed registration in AWS. This script can be combined with a wait loop to automatically check for target registration, but for purposes of this demo it is run independently.

https://ralph-secnet.us-east-1.elasticbeanstalk.com/

# Code

I chose to use Go as it has primitive types for complex numbers, which makes complex number operations trivial. There is an included Makefile to build the code.
