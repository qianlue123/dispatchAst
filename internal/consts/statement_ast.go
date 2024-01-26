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

	// 单数传参, 和从数据库中获取值等效
	3:  "pjsip show channel '%s' ",
	5:  "pjsip show aor '%s' ",
	7:  "pjsip show auth '%s' ",
	9:  "pjsip show endpoint '%s' ",
	11: "pjsip show transpoint '%s' ",

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
}

// manipulate module
var RX_module = map[int]string{
	1: "module reload res_pjsip.so",
}
