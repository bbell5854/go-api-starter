# go-api-starter

This is starter boilerplate for building a go api inside a lightweight docker container with Fiber.

## Boilerplate Features
* Fiber
* Auth0 middleware
* MySql database integration

## Environment Variables
* `DB_USERNAME`: (required) Database username.
* `DB_PASSWORD`: (required) Database password.
* `DB_HOST`: (required) Database host.
* `DB_PORT`: (required) Database port.

* `AUTH_AUDIENCE`: (required) Auth0 audience url. This can be found in your auth 0 account configuration.
* `AUTH_DOMAIN`: (required) Auth0 domain url. This can be found in your auth0 account configuration.

* `ENVIRONMENT=production` (optional) Flag for production environment. If not set, allows you to specify environment variables in a .env file.
