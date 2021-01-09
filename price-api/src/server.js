const polka = require('polka');
const { json } = require('body-parser');
const { graphqlHTTP } = require('express-graphql');
const { schema, rootValue } = require('./graphql');

const PORT = Number(process.env.PORT || 3000);
const PRODUCTION = process.env.PRODUCTION === 'true';

polka()
  .use(json())
  .use(
    '/graphql',
    graphqlHTTP({
      schema,
      rootValue,
      graphiql: !PRODUCTION,
      customFormatErrorFn: (error) => {
        const format = {
          message: error.message,
          path: error.path,
        };
        if (!PRODUCTION) {
          format.locations = error.locations;
          format.stack = error.stack ? error.stack.split('\n') : [];
        }
        return format;
      },
    })
  )
  .listen(PORT, (err) => {
    if (err) throw err;
    console.log('API Server listening on:', PORT);
  });
