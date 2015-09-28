type ClientLoginBlock struct {
    LoginType   int8
    LoginLen    int8
    Magic       int8 /* always 255 */
    Revision    int16
    MemType     int8
    ArchiveCRCs int32[9]
    SecureBlockSize int8
}

type ClientSecureLoginBlock struct {
    Magic     int8 /* always 10 */
    ISAACSeed int32[4]
    ClientUID int32
    Username  string
    Password  string
}

type ServerLoginResponse struct {
    Response int8
    Rights   int8
    Flagged  int8
}

type ServerLoginResponseUnsuccessful struct {
    Response int8
}
