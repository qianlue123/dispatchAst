package consts

/* Asterisk CLI sets */

// 约定 双数全, 单数传参
var RX = map[int]string{

	/* 1-50 pjsip series
	 */
	1: "pjsip show version",

	2:  "pjsip show channels",    // 当前信道状况
	4:  "pjsip show aors",        // 查看设备的注册状态
	6:  "pjsip show auths",       // 身份认证
	8:  "pjsip show endpoints",   // 显示终端
	10: "pjsip show transpoints", // 显示连接协议
	12: "pjsip show version",
	14: "pjsip list endpoints",
	16: "pjsip show contacts", // 分机的局域网IP

	// 单数传参, 和从数据库中获取值等效
	3:  "pjsip show channel '%s' ",
	5:  "pjsip show aor '%s' ",
	7:  "pjsip show auth '%s' ",
	9:  "pjsip show endpoint '%s' ",
	11: "pjsip show transpoint '%s' ",
	17: "pjsip show contact '%s' ",

	/** 51-100 core series
	 */
	51: "core show version",

	52: "core show calls",
	54: "core show channels",
	// hints 无法判断呼叫(直接当作是 inuse ) , 捕捉 ring 需要 channel 或者 pjsip
	56: "core show hints",

	53: "core show calls uptime",
	55: "core show hint '%s' ",
	57: "core show channel '%s' ",

	/** 101-150 database series
	 */
	101: "database show pbx",

	// family
	102: "database show ampuser",
	104: "database show device",
	106: "database show CW",
	108: "database show CustomDevstate",
	110: "database show CustomPresence",
	112: "database show registrar",
	115: "database show registrar/contact", // k: {json} 格式
	116: "database show RECconf",

	// keytree
	103: "database show ampuser '%s' ",
	105: "database show device '%s' ",
	107: "database show CW '%s' ",
	109: "database show CustomDevstate '%s' ",
	111: "database show CustomPresence '%s' ",
	113: "database show registrar '%s' ",
	117: "database show RECconf '%s' ",

	/** 151-200 confbridge series
	 */
	151: "confbridge list",

	152: "confbridge kick %s all",   // 关房
	154: "confbridge lock %s",       // 锁门, 只有 admin 用户能进
	155: "confbridge unlock %s",     // 不锁
	156: "confbridge mute %s all",   // 全禁
	158: "confbridge unmute %s all", // 全解

	153: "confbridge kick %s %s",   // 踢某个
	157: "confbridge mute %s %s",   // 禁言某人
	159: "confbridge unmute %s %s", // 解禁某人

	/** 201-250
	 */
	201: "channel originate local/%s@from-internal extension %s@from-internal", // freepbx 通例

	// 老版本的 soft hangup
	202: "channel request hangup all",

	203: "channel request hangup %s ",                           // e.g. PJSIP/2024-00000048
	205: "channel originate local/%s@%s extension %s@%s",        //
	207: "channel originate PJSIP/%s application Dial PJSIP/%s", // 忽略具体信道

}

// manipulate module
var RX_module = map[int]string{
	1: "module reload res_pjsip.so",
}
