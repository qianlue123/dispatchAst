package consts

/* Info in Freepbx 16 */

/* freepbx 的 bulk handler 模块聚集话机实例
 *
 * note: 在gui创建分机后下载的数据, 比源码提供的测试数据字段要多一点
 * reference:
 *    github.com/FreePBX/bulkhandler/blob/release/16.0/utests/extensions.csv
 */
type ExtensionUnit struct {
	Extension                   string `json:"extension"` // "2024"
	Password                    string `json:"password"`
	Name                        string `json:"name"`      // "qianlue"
	Voicemail                   string `json:"voicemail"` // "novm"
	RingTime                    string `json:"ringtimer"` // "0"
	NoAnswer                    string `json:"noanswer"`
	Recording                   string `json:"recording"`
	OutboundCid                 string `json:"outboundcid"` // "20241"
	Sipname                     string `json:"sipname"`
	NoanswerCid                 string `json:"noanswer_cid"`
	BusyCid                     string `json:"busy_cid"`
	ChanunavailCid              string `json:"chanunavail_cid"`
	NoanswerDest                string `json:"noanswer_dest"`
	BusyDest                    string `json:"busy_dest"`
	ChanunavailDest             string `json:"chanunavail_dest"`
	Mohclass                    string `json:"mohclass"`               // "default"
	ID                          string `json:"id"`                     // "2024"
	Tech                        string `json:"tech"`                   // "PJSIP"
	Dial                        string `json:"dial"`                   // "PJSIP\/2024",
	Devicetype                  string `json:"devicetype"`             // "fixed"
	User                        string `json:"user"`                   // "2024"
	Description                 string `json:"description"`            // "qianlue"
	Emergency_cid               string `json:"emergency_cid"`          // "20242"
	HintOverride                string `json:"hint_override"`          // new
	Cwtone                      string `json:"cwtone"`                 // new
	RecordingInExternal         string `json:"recording_in_external"`  // dontcare
	RecordingOutExternal        string `json:"recording_out_external"` // dontcare
	RecordingInInternal         string `json:"recording_in_internal"`  // dontcare
	RecordingOutInternal        string `json:"recording_out_internal"` // dontcare
	RecordingOudemand           string `json:"recording_ondemand"`     // "disable"
	RecordingPriority           string `json:"recording_priority"`     // "10"
	Answermode                  string `json:"answermode"`
	Intercom                    string `json:"intercom"`
	Accountcode                 string `json:"accountcode"`
	Allow                       string `json:"allow"`
	Avpf                        string `json:"avpf"`
	Callerid                    string `json:"callerid"`
	CanReinvite                 string `json:"canreinvite"` // "no"
	Context                     string `json:"context"`     // "from-internal"
	Deny                        string `json:"deny"`        // "0.0.0.0\/0.0.0.0",
	DisAllow                    string `json:"disallow"`
	Dtmfmode                    string `json:"dtmfmode"`
	Encryption                  string `json:"encryption"`
	ForceAvp                    string `json:"force_avp"`
	Host                        string `json:"dynamic"`
	Icesupport                  string `json:"icesupport"`
	Mailbox                     string `json:"mailbox"`
	Namedcallgroup              string `json:"namedcallgroup"`
	Namedpickupgroup            string `json:"namedpickupgroup"`
	Nat                         string `json:"nat"`
	Permit                      string `json:"permit"` // "0.0.0.0\/0.0.0.0",
	Port                        string `json:"port"`
	Qualify                     string `json:"qualify"`
	Qualifyfreq                 string `json:"qualifyfreq"`
	Secret                      string `json:"secret"`
	Sendrpid                    string `json:"sendrpid"`
	Sessiontimers               string `json:"sessiontimers"`
	Sipdriver                   string `json:"sipdriver"` // chan_pjsip
	Transport                   string `json:"transport"` // ""
	Trustrpid                   string `json:"trustrpid"` // "yes"
	Type                        string `json:"type"`
	Videosupport                string `json:"videosupport"`
	Sendpid                     string `json:"sendpid"`
	Trustpid                    string `json:"trustpid"`
	CallwaitingEnable           string `json:"callwaiting_enable"`
	Findmefollow_strategy       string `json:"findmefollow_strategy"`
	Findmefollow_grptime        string `json:"findmefollow_grptime"`
	Findmefollow_grppre         string `json:"findmefollow_grppre"`
	Findmefollow_grplist        string `json:"findmefollow_grplist"`
	Findmefollow_annmsg_id      string `json:"findmefollow_annmsg_id"`
	Findmefollow_postdest       string `json:"findmefollow_postdest"`
	Findmefollow_dring          string `json:"findmefollow_dring"`
	Findmefollow_needsconf      string `json:"findmefollow_needsconf"`
	Findmefollow_remotealert_id string `json:"findmefollow_remotealert_id"`
	Findmefollow_toolate_id     string `json:"findmefollow_toolate_id"`
	Findmefollow_ringing        string `json:"findmefollow_ringing"`
	Findmefollow_pre_ring       string `json:"findmefollow_pre_ring"`
	Findmefollow_voicemail      string `json:"findmefollow_voicemail"`
	Findmefollow_changecid      string `json:"findmefollow_changecid"`
	Findmefollow_fixedcid       string `json:"findmefollow_fixedcid"`
	Findmefollow_enabled        string `json:"findmefollow_enabled"`
	Languages_language          string `json:"languages_language"`
	Voicemail_enable            string `json:"voicemail_enable"`
	Voicemail_vmpwd             string `json:"voicemail_vmpwd"`
	Voicemail_email             string `json:"voicemail_email"`
	Voicemail_pager             string `json:"voicemail_pager"`
	Voicemail_options           string `json:"voicemail_options"`
}
