package conntrack

// Filter TCP SYN packet
func (conn *Connection) invalidTCPConn() bool {
	if conn.protoInfo != nil && conn.protoInfo.tcp != nil {
		if !conn.protoInfo.tcp.SYN {
			return true // only tcp syn packets are permitted, drop the pkt
		}
	} else {
		return true // missing protoInfo, drop the pkt
	}

	return false // tcp connection request is valid, accept the pkt
}
