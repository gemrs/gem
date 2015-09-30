/* The unencrypted portion of the login block */
type ClientLoginBlock struct {
    LoginType   int8
    LoginLen    int8 /* Should be SecureBlockSize + (9 * 4) + 2 + (1 * 4) */
    Magic       int8 /* always 255 */
    Revision    int16 /* always 317 */
    MemType     int8
    ArchiveCRCs int32[9]
    SecureBlockSize int8 /* Technically part of the header to ClientSecureLoginBlock */
}

/* RSA Encrypted portion of the login block */
type ClientSecureLoginBlock struct {
    Magic     int8 /* always 10 */
    ISAACSeed int32[4] /* the complete isaac seed */
    ClientUID int32
    Username  string
    Password  string
}

/* Server responds to the above two blocks with one of the following two responses */
type ServerLoginResponse struct {
    Response int8
    Rights   int8
    Flagged  int8 /* bot detection. enabled extra tracking packets to be sent from the client (i think) */
}

type ServerLoginResponseUnsuccessful struct {
    Response int8
}
