# AlertChain

AlertChain is programmable SOAR (Security Orchestration, Automation and Response) platform and universal alert manager.

## Installation

Build binary with `go >= 1.16` and `npm >= 7.18.1`.

```sh
% git clone https://github.com/m-mizutani/alertchain.git
% cd alertchain
% make
% cp alertchain /path/to/bin
```

## Usage

Write your workflow in your github project.

```go
package main

import (
	"github.com/m-mizutani/alertchain"
	"github.com/m-mizutani/alertchain/types"
)

type myEvaluator struct{}

func (x *myEvaluator) Name() string                              { return "myEvaluator" }
func (x *myEvaluator) Description() string                       { return "Eval alert" }
func (x *myEvaluator) IsExecutable(alert *alertchain.Alert) bool { return false }

func (x *myEvaluator) Execute(alert *alertchain.Alert) error {
	if alert.Title == "Something wrong" {
		alert.Severity = types.SevAffected
	}
	if err := alert.Commit(); err != nil {
		return err
	}
	return nil
}

func Chain() *alertchain.Chain {
	return &alertchain.Chain{
		Stages: []alertchain.Tasks{
			{&myEvaluator{}},
		},
	}
}
```

Then, compile it as plugin.

```sh
$ go build -buildmode=plugin -o mychain .
```

Finally, run `alertchain` with your plugin.

```sh
$ alertchain -p mychain
```
