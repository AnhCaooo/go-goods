# 1.2.3 
- Add new method `DeleteAll` to cache to allow delete all caches that contains given key 
- Add `NotFound` as new translation key

# 1.2.2
- Improve removeDuplicate to skip empty item

# 1.2.1 
- Add func to remove duplicated items from list
# 1.2.0 

# 1.1.2
- Improve how Middleware Auth can read and store UserID and SessionID

# 1.1.1
- Create Translation Keys enumerations for common http error response 
 
# 1.1.0 
- Authenticate middleware 
- Extended version of http.Error to include translation field.

# 1.0.0 
- cache in-memory
- encryption & decryption
- handle access token (JWT)
- logger (customize from [Uber Zap logger](https://github.com/uber-go/zap))
- standard encode and decode HTTP request and HTTP response
- standard map `interface{}` to specific struct 
- prometheus configuration
- middleware 