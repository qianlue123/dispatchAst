package qian_channel

// ----- ----- ----- ----- -----
//  data from Ast CLI + Bash
// ----- ----- ----- ----- -----

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"

	Ast "dispatchAst/internal/consts"
)

func CMD(cmd string) string {
	out, _ := exec.Command("bash", "-c", "whoami").Output()
	if out = out[:len(out)-1]; string(out) != "root" {
		return fmt.Sprintf("sudo asterisk -rx '%s' ", cmd)
	}
	return fmt.Sprintf("asterisk -rx '%s' ", cmd)
}

var mapBash = map[int]string{
	/*
			  ...
			  Caller ID: 2028
			  Caller ID Name: 2028
			  Connected Line ID: 2025
			  Connected Line ID Name: (N/A)
				Eff. Connected Line ID: 2025
		    Eff. Connected Line ID Name: (N/A)
			  ...
			  BRIDGEPEER=PJSIP/2025-00000044
			  DIALEDPEERNUMBER=2025
			  DIALEDPEERNAME=PJSIP/2025-00000044
			  ...
			  SIPDOMAIN=192.168.100.211
			  ...
			  level 1: src=2028
			  level 1: dst=2025
			  level 1: dcontext=qian-ctx
			  level 1: channel=PJSIP/2028-00000043
			  level 1: dstchannel=PJSIP/2025-00000044
			  ...
	*/
	571: CMD(Ast.RX[57]) + "| grep --ignore-case 'Caller ID' " +
		"| grep --ignore-case --invert-match 'name' | awk -F ': ' '{print $2}'",
	// grep -i "xxx" | grep -ivE "aaa|bbb"
	572: CMD(Ast.RX[57]) + "| grep --ignore-case 'Connected Line ID' " +
		"| grep --invert-match --extended-regexp 'Name|Eff' | awk -F ': ' '{print $2}'",
	573: CMD(Ast.RX[57]) + "| grep --ignore-case 'bridgePeer' | awk -F '=' '{print $2}' ", // right leg
	574: CMD(Ast.RX[57]) + "| grep --ignore-case 'dialedPeerNumber' | awk -F '=' '{print $2}' ",
	575: CMD(Ast.RX[57]) + "| grep --ignore-case 'dialedPeerName' | awk -F '=' '{print $2}' ",
	576: CMD(Ast.RX[57]) + "| grep --ignore-case 'sipDomain' | awk -F '=' '{print $2}' ",
	577: CMD(Ast.RX[57]) + "| grep --ignore-case 'src=' | awk -F '=' '{print $2}' ",       // = 571
	578: CMD(Ast.RX[57]) + "| grep --ignore-case 'dst=' | awk -F '=' '{print $2}' ",       // = 572
	579: CMD(Ast.RX[57]) + "| grep --ignore-case 'dstchannel' | awk -F '=' '{print $2}' ", // = 573

	/*
		Bridge-ID                            Chans Type     Technology    Duration
		4c612a9b-1e30-44c2-8cf6-ee9a2950b3ba     2 basic    native_rtp    20:24:56
	*/
	2521: CMD(Ast.RX[252]) + "| tail --lines +2 ",
	2522: CMD(Ast.RX[252]) + "| tail --lines +2 | awk '{print $1}' ",
	2523: CMD(Ast.RX[252]) + "| tail --lines +2 | awk '{print $2}' ",
	2524: CMD(Ast.RX[252]) + "| tail --lines +2 | awk '{print $3}' ",
	2525: CMD(Ast.RX[252]) + "| tail --lines +2 | awk '{print $4}' ",

	/*
		Id: 4c612a9b-1e30-44c2-8cf6-ee9a2950b3ba
		Type: basic
		Technology: native_rtp
		Subclass: basic
		Creator:
		Name:
		Video-Mode: none
		Video-Source-Id:
		Num-Channels: 2
		Num-Active: 2
		Duration: 50:00:06
		Channel: PJSIP/6661-0000001f
		Channel: PJSIP/901-00000020
	*/
	2531: CMD(Ast.RX[253]) + "| grep 'Type' | awk '{print $2}' ",
	2532: CMD(Ast.RX[253]) + "| grep 'Technology' | awk '{print $2}' ",
	2533: CMD(Ast.RX[253]) + "| grep 'Subclass' | awk '{print $2}' ",
	2534: CMD(Ast.RX[253]) + "| grep --ignore-case 'video-mode' | awk '{print $2}' ",
	2535: CMD(Ast.RX[253]) + "| grep --ignore-case 'video-source-id' | awk '{print $2}' ",
	2536: CMD(Ast.RX[253]) + "| grep --ignore-case 'num-channels' | awk '{print $2}' ",
	2537: CMD(Ast.RX[253]) + "| grep --ignore-case 'num-active' | awk '{print $2}' ",
	2538: CMD(Ast.RX[253]) + "| grep --ignore-case 'Duration' | awk '{print $2}' ",

	// e.g asterisk -rx 'bridge show <tab birdgeid>' | tail --lines 2| awk '{print $2}'
	2539: CMD(Ast.RX[253]) + "| tail --lines %d | awk '{print $2}' ",
}

// 功能: 由channel名获取本身对应的话机名字
func GetSelfExtNameWithName(channelName string) string {
	cmd := fmt.Sprintf(mapBash[571], channelName)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	out = out[:len(out)-1]

	return string(out)
}

// 功能: 由channel名获取对端的话机名字
func GetPeerExtNameWithName(channelName string) string {
	cmd := fmt.Sprintf(mapBash[572], channelName)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	out = out[:len(out)-1]

	return string(out)
}

// 功能: 获取系统所有桥接
func GetBridgeName() []string {
	// 系统没有设备在通讯中, 返回空
	if v, _ := exec.Command("bash", "-c", mapBash[2521]).Output(); len(v) == 0 {
		return nil
	}

	nameArr := make([]string, 0)

	out, _ := exec.Command("bash", "-c", mapBash[2522]).Output()
	// 先去掉末尾换行再按换行切分
	out = out[:len(out)-1]

	outs := bytes.Split(out, []byte("\n"))
	for i, content := range outs {
		fmt.Printf("%d %v \n", i, string(content))
		nameArr = append(nameArr, string(content))
	}

	return nameArr
}

// 功能: 由桥接名查询桥接的类型
func GetBridgeTypeWithName(bridgeName string) string {
	cmd := fmt.Sprintf(mapBash[2531], bridgeName)

	out, _ := exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]

	return string(out)
}

// 功能: 由桥接名查询含有的 channel 数量
func CountChannelsFromBridge(bridgeName string) (t int) {
	cmd := fmt.Sprintf(mapBash[2536], bridgeName)

	out, _ := exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]

	if len(out) == 0 {
		t = 0
	} else {
		t, _ = strconv.Atoi(string(out))
	}

	return t
}

// 功能: 由桥接名查询所有的 channels 名字集
func GetChannelNamesFromBridge(bridgeName string, c int) ([]string, error) {
	if c%2 == 0 {
		c = CountChannelsFromBridge(bridgeName) // check again
	}

	arr, cmd := make([]string, 0), fmt.Sprintf(mapBash[2539], bridgeName, c)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}
	out = out[:len(out)-1]

	outs := bytes.Split(out, []byte("\n"))
	for _, content := range outs {
		arr = append(arr, string(content))
	}

	return arr, nil
}

// 功能: 获取具体channel内所有话机id集合
func getExtNamesFromChannel(channelName string) ([]string, error) {
	arr := make([]string, 0)

	arr1, arr2 := GetSelfExtNameWithName(channelName),
		GetPeerExtNameWithName(channelName)

	arr = append(arr, arr1, arr2)

	fmt.Println("arr: ", arr)

	return arr, nil
}
