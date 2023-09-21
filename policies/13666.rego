package suricata

default alert := {}

# Trigger an alert if the Signature field contains the specific string "Suspicious connection to port 13666".
alert := {"src_ip": inp.SrcIP} {
	inp := input[_]
	inp.Type == "alert"
	contains(inp.Alert.Signature, "Suspicious connection to port 13666")
}