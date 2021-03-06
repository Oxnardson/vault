Vault [![Build Status](https://travis-ci.org/hashicorp/vault.svg)](https://travis-ci.org/hashicorp/vault)
=========

-	Website: https://www.vaultproject.io
-	IRC: `#vault-tool` on Freenode
-	Mailing list: [Google Groups](https://groups.google.com/group/vault-tool)

![Vault](https://raw.githubusercontent.com/hashicorp/vault/master/website/source/assets/images/logo-big.png?token=AAAFE8XmW6YF5TNuk3cosDGBK-sUGPEjks5VSAa2wA%3D%3D)

Vault is a tool for securely accessing secrets. A secret is anything that you want to tightly control access to, such as API keys, passwords, certificates, and more. Vault provides a unified interface to any secret, while providing tight access control and recording a detailed audit log.

A modern system requires access to a multitude of secrets: database credentials, API keys for external services, credentials for service-oriented architecture communication, etc. Understanding who is accessing what secrets is already very difficult and platform-specific. Adding on key rolling, secure storage, and detailed audit logs is almost impossible without a custom solution. This is where Vault steps in.

The key features of Vault are:

* **Secure Secret Storage**: Arbitrary key/value secrets can be stored
  in Vault. Vault encrypts these secrets prior to writing them to persistent
  storage, so gaining access to the raw storage isn't enough to access
  your secrets. Vault can write to disk, [Consul](https://www.consul.io),
  and more.

* **Dynamic Secrets**: Vault can generate secrets on-demand for some
  systems, such as AWS or SQL databases. For example, when an application
  needs to access an S3 bucket, it asks Vault for credentials, and Vault
  will generate an AWS keypair with valid permissions on demand. After
  creating these dynamic secrets, Vault will also automatically revoke them
  after the lease is up.

* **Data Encryption**: Vault can encrypt and decrypt data without storing
  it. This allows security teams to define encryption parameters and
  developers to store encrypted data in a location such as SQL without
  having to design their own encryption methods.

* **Leasing and Renewal**: All secrets in Vault have a _lease_ associated
  with it. At the end of the lease, Vault will automatically revoke that
  secret. Clients are able to renew leases via built-in renew APIs.

* **Revocation**: Vault has built-in support for secret revocation. Vault
  can revoke not only single secrets, but a tree of secrets, for example
  all secrets read by a specific user, or all secrets of a particular type.
  Revocation assists in key rolling as well as locking down systems in the
  case of an intrusion.

For more information, see the [introduction section](https://www.vaultproject.io/intro)
of the Vault website.

Getting Started & Documentation
-------------------------------

All documentation is available on the [Vault website](https://www.vaultproject.io).

Developing Vault
--------------------

If you wish to work on Vault itself or any of its built-in systems,
you'll first need [Go](https://www.golang.org) installed on your
machine (version 1.4+ is *required*). Alternatively, you can use the
Vagrantfile in the root of this repo to stand up a virtual machine with
the appropriate dev tooling already set up for you.

For local dev first make sure Go is properly installed, including setting up a
[GOPATH](https://golang.org/doc/code.html#GOPATH). After setting up Go,
install Godeps, a tool we use for vendoring dependencies:

```sh
$ go get github.com/tools/godep
...
```

Next, clone this repository into `$GOPATH/src/github.com/hashicorp/vault`.
Then type `make`. This will run the tests. If this exits with exit status 0,
then everything is working!

```sh
$ make
...
```

To compile a development version of Vault, run `make dev`. This will put the
Vault binary in the `bin` and `$GOPATH/bin` folders:

```sh
$ make dev
...
$ bin/vault
...
```

If you're developing a specific package, you can run tests for just that
package by specifying the `TEST` variable. For example below, only
`vault` package tests will be run.

```sh
$ make test TEST=./vault
...
```

### Acceptance Tests

Vault has comprehensive [acceptance tests](https://en.wikipedia.org/wiki/Acceptance_testing)
covering most of the features of the secret and auth backends.

If you're working on a feature of a secret or auth backend and want to
verify it is functioning (and also hasn't broken anything else), we recommend
running the acceptance tests.

**Warning:** The acceptance tests create/destroy/modify *real resources*, which
may incur real costs in some cases. In the presence of a bug, it is technically
possible that broken backends could leave dangling data behind. Therefore,
please run the acceptance tests at your own risk. At the very least,
we recommend running them in their own private account for whatever backend
you're testing.

To run the acceptance tests, invoke `make testacc`:

```sh
$ make testacc TEST=./builtin/logical/consul
...
```

The `TEST` variable is required, and you should specify the folder where the
backend is. The `TESTARGS` variable is recommended to filter down to a specific
resource to test, since testing all of them at once can sometimes take a very
long time.

Acceptance tests typically require other environment variables to be set for
things such as access keys. The test itself should error early and tell
you what to set, so it is not documented here.
