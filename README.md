# fitworks-backend
Backend for FitWorks

## Things I learned

### psql user
When I installed psql via homebrew, homebrew set root user as my linux user.  
because of this, when I created a new user in psql and created a db through the new user,
gorm cannot connect.  
when i switched to the unix root user, setup password properly using psql, and updated my credentials,
i was able to connect to my proper db.  

## roadmap
[] design db
[] develop data service for each db
    [] user db
    [] homework db
    * domain driven approach since it does not make sense for multiple relational db to have its own mini microservice.  
[] host db and data service
[] develop business service
    [] fitworks business service - perhaps just stick to this repo
[] develop minimal frontend
[] implement authentication via OAuth2.0
[] handle jwt
[] upgrade frontned

### optional roadmap
[] properly setup API gateway