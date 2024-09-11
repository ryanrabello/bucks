* Add types to the profile object (if it makes sense to do in go)
* Create a user model in the database (maybe)
* Create other models
    * Transaction
      * Sender FK (User)
      * Receiver FK (User)
      * Date (time.Time)
      * Amount (int)
      * Note optional (string)
      * Code FK (Code)
    * Code
      * Code (string)
      * Sender (User)
      * Amount (int)
      * claimedDate (time.Time)
      * createdDate (time.Time)?
    * User
      * Id (string)
      * Name unique? (string)
      * ProfilePicture (string)
      