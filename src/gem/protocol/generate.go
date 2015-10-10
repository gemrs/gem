package protocol

/* generate handshake/login/update structures */
//go:generate bbc handshake.bb handshake.bb.go
//go:generate bbc update_service.bb update_service.bb.go
//go:generate bbc game_login.bb game_login.bb.go

/* generate server -> client packets */
//go:generate bbc outbound_packets.bb outbound_packets.bb.go

/* generate client -> server packets */
//go:generate bbc inbound_packets.bb inbound_packets.bb.go
