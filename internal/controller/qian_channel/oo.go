package qian_channel

import "fmt"

// 话机
type extension struct {
	CID     string
	Desc    string // 用户名
	ExtName string // 电话号
	ExtAge  int    // 使用年限
	ExtPwd  string // 密码
	State   int    // 状态
}

type Channel struct {
	name    string
	extSelf extension
	extPeer extension
}

type Bridge struct {
	name     string
	bType    string
	channels []Channel
}

type TwoWayBridge struct {
	*Bridge
}

type MultiBridge struct {
	*Bridge
}

type Ivoip interface {
	Count() int
	ShowChannel()
	ShowChannelName() []string
	ShowExtension() []extension
}

func (b *Bridge) SetName(name string) {
	b.name = name
}

func (b *Bridge) SetType(name string) {
	b.bType = GetBridgeTypeWithName(name)
}

func (b *Bridge) GetName() string {
	return b.name
}

func (b *Bridge) GetType() string {
	return b.bType
}

func (c *Channel) SetName(name string) {
	c.name = name
}

func (c *Channel) GetName() string {
	return c.name
}

func (c *Channel) SetExtSelf(ext extension) {
	c.extSelf = ext
}

func (c *Channel) GetExtSelf() extension {
	return c.extSelf
}

func (c *Channel) SetExtPeer(ext extension) {
	c.extPeer = ext
}

func (c *Channel) GetExtPeer() extension {
	return c.extPeer
}

// 功能: Channel 的构造函数
func NewChannel(name string) *Channel {
	extNameArr, err := getExtNamesFromChannel(name)
	if err != nil {
		extNameArr = nil
	}

	fmt.Println("extensions: ", extNameArr)

	extSelf := extension{ExtName: GetSelfExtNameWithName(name)}
	extPeer := extension{ExtName: GetPeerExtNameWithName(name)}

	return &Channel{name, extSelf, extPeer}
}

// 功能: Bridge 的构造函数
func NewBridge(name string, options ...string) *Bridge {
	bType := ""

	if len(options) == 1 {
		bType = options[0]
	} else {
		bType = GetBridgeTypeWithName(name)
	}

	c := CountChannelsFromBridge(name)
	channelNameArr, _ := GetChannelNamesFromBridge(name, c)

	channels := make([]Channel, 0)
	for _, cName := range channelNameArr {
		cOne := NewChannel(cName)
		channels = append(channels, *cOne)
	}

	return &Bridge{name, bType, channels}
}

// 功能: 一个桥接中含有的信道数量
func (b *Bridge) Count() int {
	return CountChannelsFromBridge(b.name)
}

// 功能: 一个桥接中含有的信道名
func (b *Bridge) ShowChannelName() []string {
	nameArr, err := GetChannelNamesFromBridge(b.name, b.Count())
	if err != nil {
		return nil
	}
	return nameArr
}

// 功能: 一个桥接中含有的所有话机
// note 必须先初始化好桥接中的channels
func (b *Bridge) ShowExtension() []extension {
	extArr := make([]extension, 0)

	// var extMap map[string]extension
	extMap := make(map[string]extension)

	for _, c := range b.channels {

		ext, ext2 := c.GetExtSelf(), c.GetExtPeer()

		// 如果分机还没放到 map , 说明它还没被统计, 可收入集合中
		if _, exist := extMap[ext.ExtName]; !exist {
			extArr = append(extArr, ext)
			extMap[ext.ExtName] = ext
		}

		if _, exist := extMap[ext2.ExtName]; !exist {
			extArr = append(extArr, ext2)
			extMap[ext2.ExtName] = ext2
		}
	}

	fmt.Println("extArr:", extArr)
	return extArr
}
