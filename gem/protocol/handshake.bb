/* ServiceSelect is the very first byte sent by the client, and is used to identify
   which service (update or game) the this session is for.
*/
type InboundServiceSelect struct {
    Service uint8
}

/* If Service is 14, it is followed up by a byte containing a part of the username
   This is probably used to select the login server or something, but we have no real
   use for it
*/
type InboundGameHandshake struct {
    NameHash uint8
}

/* Next, the server sends it's response.
   ServerISAACSeed is the server's contribution to the isaac key, and should be
   64 bits of randomness.
   From here, the client sends the ClientLoginBlock (see game_login.bb)
*/
type OutboundGameHandshake struct {
    ignored         uint8[8]
    loginRequest    uint8 /* always 0 */
    ServerISAACSeed uint32[2]
}

/* If ServiceSelect.Service is 15, the server sends 8 bytes, which is ignored by the client
   From here, the client begins making it's UpdateRequest's (see update_service.bb)
*/
type OutboundUpdateHandshake struct {
    ignored uint8[8]
}
