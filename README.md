# golam

A sample starting point for an AWS lambda implemented in golang.

Yes, `aws sam` `cloudformation` `serverless` are other options. This was written for smaller scale hobby use.

### usage

* Setup ~/.aws/config and ~/.aws/credentials according
  to [AWS Documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)
* Function name will be derived from the last part of the `go.mod` module name.
* Install [mage](https://magefile.org/) and run `mage` to deploy or `mage delete` to remove the function and associated
  roles.
* Customize additional roles in [mage.go#L180](magefiles/mage.go#L180).
* Run mage with the `-v` switch to see additional debugging.

[![Go Report Card](https://goreportcard.com/badge/github.com/mlctrez/golam)](https://goreportcard.com/report/github.com/mlctrez/golam)

created by [tigwen](https://github.com/mlctrez/tigwen)
