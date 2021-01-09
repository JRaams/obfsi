# price-api

GraphQL Api using [polka](https://github.com/lukeed/polka) to present steam asset price data.

### Running

1. Install npm dependencies
   `$ yarn install`

2. Export DBSOURCE env
   ` $ export DBSOURCE='postgres://obfsi@0.0.0.0:5432/obfsi?sslmode=disable'`

3. Start
   `$ yarn start`

4. Visit the [graphiql explorer](http://0.0.0.0:3000/graphql)

### Production

1. EXPORT PRODUCTION=true
   This will disable the graphiql explorer

### Linting and formatting

1. Linting (read-only)
   `$ yarn lint`

2. Linting + fix
   `$ yarn lint-fix`

3. Formatting
   `$ yarn format`
