Package inmemorydb creates a tcp session with a map to accept and store key-value pairs required for immediate execution in runtime
syntax to use is as follows



SET "key" "value"

GET "key"

DEL "key"


Example:
SET cat persian
GET cat
DEL cat
common applications like cache usage, session specific environmental variables etc.,
