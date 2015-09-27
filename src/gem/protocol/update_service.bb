type UpdateRequest struct {
    Index int8
    File  int16
    /* Priority is an integer between 0-2
         0: Urgent request (ie. client needs resources for a new area)
         1: Preload request (pre-login client update)
         2: Background request (low priority) */
    Priority  int8
}

type UpdateResponse struct {
    Index int8
    File  int16
    Size  int16
    Chunk int8
    Data  byte[500]
}