package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/manifoldco/promptui"
)

type Orgs struct {
	Organization Organization           `json:"organization"`
	AnotherField map[string]interface{} `json:"-"`
}

type Organization struct {
	Name    string `json:"name"`
	Domain  string `json:"domain"`
	MspName string `json:"mspName,omitempty"`
}

type Channel struct {
	Name string     `json:"channel"`
	Org  ChannelOrg `json:"orgs"`
}

type ChannelOrg struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}

func ChooseFabricVersion() string {
	prompt := promptui.Select{
		Label: "Please choose fabric version",
		Items: []string{"1.4.6", "2.2.4"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}

func ChooseMonitorLog() string {
	prompt := promptui.Select{
		Label: "Please choose monitor log level",
		Items: []string{"off", "debug", "info", "trace", "all"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}

func AddMultipleOrgsInChannel() []Channel {
	//Assume organation retrieve from another source because this have random field. It can be document database or call API
	dataOrg := []byte(`[
		{
		  "organization": {
			"name": "Orderer",
			"domain": "orderer.example.com"
		  },
		  "orderers": [
			{
			  "groupName": "group1",
			  "prefix": "orderer",
			  "type": "raft",
			  "instances": 1
			}
		  ]
		},
		{
		  "organization": {
			"name": "Org1",
			"mspName": "Org1MSP",
			"domain": "org1.example.com"
		  },
		  "ca": {
			"prefix": "ca"
		  },
		  "peer": {
			"prefix": "peer",
			"instances": 2,
			"db": "LevelDb"
		  }
		}
	  ]`)

	orgs := []Orgs{}
	if err := json.Unmarshal([]byte(dataOrg), &orgs); err != nil {
		panic(err)
	}

	items := []string{}

	for _, v := range orgs {
		items = append(items, v.Organization.Name)
	}

	idx := 0
	var result string
	var err error
	items = append(items, "exit")

	channels := []Channel{}

	for {
		prompt := promptui.Select{
			Label: "Please choose organization",
			Items: items,
		}
		idx, result, err = prompt.Run()
		if idx == len(items)-1 {
			break
		}
		ch := Channel{
			Name: "my-channel-1",
			Org: ChannelOrg{
				Name: orgs[idx].Organization.Name,
			},
		}
		channels = append(channels, ch)
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	fmt.Printf("You choose %q\n", result)

	return channels
}
