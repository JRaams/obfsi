const polka = require('polka');
const { json } = require('body-parser');
const { graphqlHTTP } = require('express-graphql');
const { schema, rootValue } = require('./graphql');

const { PORT = 3000, PRODUCTION = false } = process.env;

polka()
  .use(json())
  .use(
    '/graphql',
    graphqlHTTP({
      schema,
      rootValue,
      graphiql: !PRODUCTION,
    })
  )
  .listen(PORT, (err) => {
    if (err) throw err;
    console.log('API Server listening on :', PORT);
  });
