Readme!

# Install 

docker-compose up --force-recreate

Note: after quit, beankstalk container (queue server) may hang up. It is good idea to execute "docker-compose down" manually.

# Design:
	- Will approach via clean architecture strategy, based on microservices.
	- Two real time services (apis):
		- User: will handle user-related things (register, login, follow/unfollow).
		- Content: will handle the tweets (create, read).
	- One background worker:
		- Will handle cache updates based on user and content messages.
	- Cache:
		- Will store user personal feed (for login) and user's top tweets (most recent ones).
	- Messaging:
		- Will use queue to send updates from apis to worker.
	- Database:
		- Will use relational DB engine, as a lot of interactions (relations) will take place between User and Content.
	- Code:
		- Each component will have its own directory into repository.
		- Isolation must be preserved across apis. Worker can import models from apis.
		- Woker will expose a client to make it easier to integrate.
	- Testing:
		- Will apply TDD, with guidence on formal specifications

# Trade-offs:
	- It is better to use one DB per service (user and content), but for simplify reasons will share DB and separate domains by tables.
	- Will check (and create) databases, tables and indexes at api startup, but would be better doing this outside code.
	- As we don't have login method, both user and password will be required via header. Using password instead of some nice and secure token. Also, for checking users' repository db is imported into contents (not the best way, but will do the trick for ex. purpose)
	- Cache for feed should be clean up after some time. The best trigger would be "after X time from last login", but as we are not implementing login now, this will be an important todo.
	- Error handling in src/enums/errors.go should have its own package. Not doing this because of time and scope.
	- Handling users endpoints via username as params instead of id, because we don't allow same user name and it is more readable when testing via postman or similar

# Tools:
	- Specifications written in: https://z-editor.github.io/
	- Queue will use beanstalkd (with bodsch/docker-beanstalkd container)
	- DB will implement PostgreSQL (official docker image)
	- Cache server will be Redis (official docker image)
	- Project language will be Go, with Fiber framework and "gorm" as ORM.
	- "deploy" process will be via docker-compose

# Steps for develop:
	1. Create design principles (this doc).
	2. Create specifications based on requirements pdf (stored at `docs/` directory).
	3. Create docker-compose yaml to build up the environment.
	4. 

# TODOS:
	- Setup contents api
	- Setup cache worker
	- Remove redis connection from users
	- Implement smoke test for main flows: create user, follow, post content, read feed, unfollow
	- Implement pagination in service and repo content read
