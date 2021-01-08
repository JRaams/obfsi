const polka = require('polka')
const { json } = require('body-parser')
const { graphqlHTTP } = require('express-graphql')
const { schema, rootValue } = require('./graphql');

const PORT = process.env.PORT || 3000;
const PRODUCTION = !process.env.PRODUCTION || false;

polka()
  .use(json())
  .use('/graphql', graphqlHTTP({
    schema,
    rootValue,
    graphiql: PRODUCTION
  }))
  .listen(PORT, err => {
    if (err) throw err;
    console.log('API Server listening on :', PORT)
  })