package protocol

/* generate handshake/login/update structures */
//go:generate bbc handshake.bb handshake.bb.go
//go:generate bbc update_service.bb update_service.bb.go
//go:generate bbc game_login.bb game_login.bb.go

/* generate server -> client packets */
//go:generate bbc chat_message.bb chat_message.bb.go
