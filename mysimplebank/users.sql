
/*

MY SCHEMA FOR BANKING APP

*/

/* CREATE TABLE IN DATABASE */
CREATE TABLE IF NOT EXISTS USERS(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    FULLNAME TEXT NOTNULL,
    EMAIL TEXT NOTNULL
);


/* GET ALL USERS DATA */
SELECT * FROM USERS;

/* INSERT DATA INTO USERS TABLE */
INSERT INTO USERS(FULLNAME , EMAIL)VALES(?,?);

/* GET ALL USERS DATA */
(VIKASH PARASHAR , VIKASHPARASHAR111@GMAIL.COM)
(KHUSHBOO PARASHAR  ,  KHUSHBOOPARASHAR1@GMAIL.COM)
(NIYATI SHARMA ,NOEMAIL@GMAIL.COM)
(RITIKA SHARMA , RITIKASH@GMAIL.COM)
(RAMPATI DEVI  , RAMPATIDEVI03@GMAIL.COM)





/* GET SINGLE USER DATA FROM TABLE */
COMSELECT * FROM USERS WHERE ID = ?

