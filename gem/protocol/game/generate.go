package game

/* generate login structures */
//go:generate bbc game_login.bb game_login.bb.go

/* generate server -> client packets */
//go:generate bbc outbound_packets.bb outbound_packets.bb.go

/* generate client -> server packets */
//go:generate bbc inbound_packets.bb inbound_packets.bb.go

//go:generate bbc player_appearance_update.bb player_appearance_update.bb.go
