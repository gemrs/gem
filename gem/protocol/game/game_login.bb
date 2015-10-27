/* The unencrypted portion of the login block */
type InboundLoginBlock struct {
    LoginType   uint8
    LoginLen    uint8 /* Should be SecureBlockSize + (9 * 4) + 2 + (1 * 4) */
    Magic       uint8 /* always 255 */
    Revision    uint16 /* always 317 */
    MemType     uint8
    ArchiveCRCs uint32[9]
    SecureBlockSize uint8 /* Technically part of the header to ClientSecureLoginBlock */
}

/* RSA Encrypted portion of the login block */
type InboundSecureLoginBlock struct {
    Magic     uint8 /* always 10 */
    ISAACSeed uint32[4] /* the complete isaac seed */
    ClientUID uint32
    Username  string
    Password  string
}

/* Server responds to the above two blocks with one of the following two responses */
type OutboundLoginResponse struct {
    Response uint8
    Rights   uint8
    Flagged  uint8 /* bot detection. enabled extra tracking packets to be sent from the client (i think) */
}

type OutboundLoginResponseUnsuccessful struct {
    Response uint8
}
