# V2 MODE

(A)  we create a user with fullname , email and password.
(A) we need to hash&salt the password

(A) User, role based access, we need to retrieve the create users and I the super admin should be able to give admin access to this USER. 
(A) which means we need a way to create super admin
(A) role( super admin, admin, user) // CRUD to give or remove this role to a user. 
(A) a User should always have a  role and only one role.
(A) user has many todos

(A) a user can update is password (forgot password scenario)
(A) we also need to introduce JWT, that token needs to expire in 1 week. and since we dont have a frontend, we need a token table. 
(A) the super admin token expiration is 30 days.
(A) admin and user token expiration is 7 days (1 week)

(A) A user should have only one role and should have only one token when he has login at least ounce. After the expiration date, the user will need to log again.
(A) which means a user should be able to login and logout. the logout won't trigger a deletion or expiration of the token.
(A) when the token expire for a user, the token will be set to null.
