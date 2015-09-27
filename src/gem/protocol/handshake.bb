type ServiceSelect struct {
    Service uint8
}

type GameHandshake struct {
    NameHash int8
}

type UpdateHandshakeResponse struct {
    ignored int8[8]
}

type GameHandshakeResponse struct {
    ignored         uint8[8]
    loginRequest    uint8 /* always 0 */
    ServerISAACSeed int32[2]
}