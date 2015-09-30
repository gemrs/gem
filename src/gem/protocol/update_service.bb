/* UpdateRequest is a request for a file within the cache to be sent
   Index 0 corresponds to a file in idx1 (rather than idx0)
   Priority identifies how urgent the request is, and the server prioritizes requests accordingly.
*/
type UpdateRequest struct {
    Index int8
    File  int16
    /* Priority is an integer between 0-2
         0: Urgent request (ie. client needs resources for a new area)
         1: Preload request (pre-login client update)
         2: Background request (low priority) */
    Priority  int8
}

/* The server's response to a request is chunked. One chunk is sent for every 500
   bytes of the file.
   The last chunk of a file is not padded out to the full 500 bytes
*/
type UpdateResponse struct {
    Index int8
    File  int16
    Size  int16
    Chunk int8 /* chunk's sequence id, starting from 0 */
    Data  byte[500]
}
