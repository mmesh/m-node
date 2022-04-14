package ops

import "github.com/spf13/viper"

var disabledOps bool = viper.GetBool("agent.management.disableOps")
