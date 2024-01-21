package qian_force

// 展示信道状态信息
type Channel struct {
	ChannelID string
	CallerID  string // 主动打的分机号
	Location  string
	States    string // 呼叫、被叫、起来
	Duration  string
}

var asteriskrx = map[int]string{
	1: "asterisk -rx 'core show help' ",

	// [Channel  Location  State  Application(Data)]
	// 列2列4数据没用, 太长时自动截断
	2:  "asterisk -rx 'core show channels' ",
	21: "asterisk -rx 'core show channels count' ",
	22: "asterisk -rx 'core show channels verbose' ",
	23: "asterisk -rx 'core show channels conciss' ",

	// 比 core show channels 多获得信道存在的时间信息
	3:  "asterisk -rx 'pjsip show channels' ",
	31: "asterisk -rx 'pjsip show channel %d' ", // channelID
	// 显示 channelID 后面带 /Dial | /AppDial
	4: "asterisk -rx 'pjsip list channels' ",

	// 原先的 soft hangup
	5: "asterisk -rx 'channel request hangup all' ",
	// 拼装型命令, 配合 go sprintf
	51: "asterisk -rx 'channel request hangup %s ' ", // e.g. PJSIP/2024-00000048

	61: "asterisk -rx 'channel originate local/%s@from-internal extension %s@from-internal' ",
}

var mapBash = map[int]string{
	// core show channels 先去除标题行和尾部统计信息
	21: asteriskrx[2] + "| tail --lines +2 | head --lines -3 ",
	22: asteriskrx[2] + "| tail --lines +2 | head --lines -3 | awk '{print $1}' ",
	23: asteriskrx[2] + "| tail --lines +2 | head --lines -3 | awk '{print $1, $3}' ",

	/** example of 'core show channels count'
	 *   4 active channels
	 *   2 actice call
	 *   ? calls processed
	 */
	211: asteriskrx[21] + "| grep 'active call' | awk '{print $1}' ",
	212: asteriskrx[21] + "| grep channels | awk '{print $1}' ",
}
