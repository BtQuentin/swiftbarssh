package ssh

import (
	"fmt"
	"strings"

	"github.com/mikkeloscar/sshconfig"
)

func PrintHosts(configFile string) {
	hosts, err := sshconfig.Parse(configFile)
	if err != nil {
		fmt.Println(err)
	}

	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")

	for _, host := range hosts {
		displayHost := replacer.Replace(strings.Join(host.Host, " "))
		userName := replacer.Replace(strings.TrimSpace(host.User))
		hostName := replacer.Replace(strings.TrimSpace(host.HostName))
		identityFile := replacer.Replace(strings.TrimSpace(host.IdentityFile))
		line := fmt.Sprintf("%s | bash='ssh' param0=%s@%s param1=-i param2=%s terminal=true", displayHost, userName, hostName, identityFile)
		fmt.Println(line)
	}
}
