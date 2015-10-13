package gcm

// Message is used by the application server to send a message to
// the GCM server. See the documentation for GCM Architectural
// Overview for more information:
// http://developer.android.com/google/gcm/gcm.html#send-msg
type Message struct {
	To					  string				 `json:"to,omitempty"`
	RegistrationIDs       []string               `json:"registration_ids,omitempty"`
	CollapseKey           string                 `json:"collapse_key,omitempty"`
	Data                  map[string]interface{} `json:"data,omitempty"`
	Notification		  map[string]interface{} `json:"notification,omitempty"`
	DelayWhileIdle        bool                   `json:"delay_while_idle,omitempty"`
	TimeToLive            int                    `json:"time_to_live,omitempty"`
	RestrictedPackageName string                 `json:"restricted_package_name,omitempty"`
	DryRun                bool                   `json:"dry_run,omitempty"`
}

// NewMessage returns a new Message with the specified payload
// and registration IDs.
func NewMessage(data map[string]interface{},notification map[string]interface{}, regIDs ...string) *Message {
	if len(regIDs) == 1{
		return &Message{To: regIDs[0], Data: data, Notification: notification}
	}else if len(regIDs) > 1{
		return &Message{RegistrationIDs: regIDs, Data: data, Notification: notification}
	}else{
		return nil
	}
	
}
