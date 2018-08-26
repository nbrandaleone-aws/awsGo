# awsGo
This repository demonstrates how to deploy a GoLang lambda function
on AWS.  This particular function interacts with EKS, and is meant for testing only.

It uses the latest Go SDK (version 2).
It does not use a Go context, since it is not triggered by an event.
However, it does accept environmental variables.

Finally, it has a Makefile, which compiles the code for
a Linux target, even though development is on a Mac. This is
necessary for execution on Lambda.
