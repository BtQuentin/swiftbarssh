package ssh

import (
	"fmt"
	"github.com/mikkeloscar/sshconfig"
	"strings"
)

func PrintHosts(configFile string) {
	hosts, err := sshconfig.Parse(configFile)
	if err != nil {
		fmt.Println(err)
	}

	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")

	for _, host := range hosts {
		displayHost := replacer.Replace(strings.Join(host.Host, " "))
		line := fmt.Sprintf("%s | bash='ssh' param0=%s@%s param1=-i param2=%s terminal=true", displayHost, host.User, host.HostName, host.IdentityFile)
		fmt.Println(line)
	}
}
