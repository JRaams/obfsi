const schema = require('./schema');
const resolvers = require('./resolvers');

module.exports = {
  schema,
  rootValue: resolvers
}