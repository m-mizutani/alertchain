package cli

import (
	"github.com/m-mizutani/alertchain/pkg/chain"
	"github.com/m-mizutani/alertchain/pkg/chain/core"
	"github.com/m-mizutani/alertchain/pkg/controller/cli/config"
	"github.com/m-mizutani/alertchain/pkg/utils"
)

func buildChain(policy *config.Policy, options ...core.Option) (*chain.Chain, error) {
	if policy.Print() {
		utils.Logger().Info("enable print mode")
		options = append(options, core.WithEnablePrint())
	}

	alertPolicy, err := policy.Load("alert")
	if err != nil {
		return nil, err
	}
	options = append(options, core.WithPolicyAlert(alertPolicy))

	actionPolicy, err := policy.Load("action")
	if err != nil {
		return nil, err
	}
	options = append(options, core.WithPolicyAction(actionPolicy))

	return chain.New(options...)
}
