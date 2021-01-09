const { buildSchema } = require('graphql');

const schema = buildSchema(`
  type Task {
    id: Int!
    name: String!
    complete: Boolean!
  }
  type Query {
    tasks: [Task]
    task(id: Int!): Task
  }
`);

module.exports = schema;
