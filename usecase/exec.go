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
	Name string       `json:"channel"`
	Org  []ChannelOrg `json:"orgs"`
}

type ChannelOrg struct {
	Name  string   `json:"name"`
	Peers []string `json:"peers"`
}

type ChainCode struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Lang        string `json:"lang"`
	Channel     string `json:"channel"`
	Init        string `json:"init"`
	Endorsement string `json:"endorsement"`
	Directory   string `json:"directory"`
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

func DataOrg() (dataOrg []byte) {
	//Assume organation retrieve from another source because this have random field. It can be document database or call API
	dataOrg = []byte(`[
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
	return
}

func AddMultipleOrgsInChannel(dataOrg []byte) []Channel {
	orgs := []Orgs{}
	if err := json.Unmarshal([]byte(dataOrg), &orgs); err != nil {
		panic(err)
	}

	items := []string{}

	for _, v := range orgs {
		items = append(items, v.Organization.Name)
	}

	var err error

	channelOrgs := []ChannelOrg{}

	for {
		prompt := promptui.Select{
			Label: "Please choose organization",
			Items: items,
		}
		idx, res, err := prompt.Run()
		if err != nil {
			break
		}
		if res == "exit" {
			break
		}
		chOrg := ChannelOrg{
			Name: orgs[idx].Organization.Name,
		}
		channelOrgs = append(channelOrgs, chOrg)

		pContinue := promptui.Select{
			Label: "Do you want add more organization?",
			Items: []string{"yes", "no"},
		}
		idx, _, err = pContinue.Run()
		if err != nil {
			break
		}

		if idx == 1 {
			break
		}

	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return []Channel{
		{Name: "channel-1",
			Org: channelOrgs,
		},
	}
}

func AddMultipleChainCode(ch []Channel) []ChainCode {
	chCodes := []ChainCode{}
	chItems := []string{}
	var err error

	for _, v := range ch {
		chItems = append(chItems, v.Name)
	}
	for {
		pName := promptui.Prompt{
			Label: "Input chaincode name",
		}
		pVersion := promptui.Prompt{
			Label: "Input chaincode version",
		}
		pLang := promptui.Prompt{
			Label: "Input chaincode languange",
		}
		var name string
		var version string
		var lang string

		for name == "" {
			name, err = pName.Run()
			if err != nil {
				break
			}
		}

		for version == "" {
			version, err = pVersion.Run()
			if err != nil {
				break
			}
		}

		for lang == "" {
			lang, err = pLang.Run()
			if err != nil {
				break
			}
		}

		pChannel := promptui.Select{
			Label: "Choose your channel?",
			Items: chItems,
		}
		_, channel, err := pChannel.Run()
		if err != nil {
			break
		}

		code := ChainCode{
			Name:        name,
			Version:     version,
			Lang:        lang,
			Channel:     channel,
			Init:        "{\"Args\":[]}",
			Endorsement: "AND ('Org1MSP.member')",
			Directory:   "",
		}
		chCodes = append(chCodes, code)

		pContinue := promptui.Select{
			Label: "Do you want add more chaincode?",
			Items: []string{"yes", "no"},
		}
		idx, _, err := pContinue.Run()
		if err != nil {
			break
		}

		if idx == 1 {
			break
		}

	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	return chCodes
}
