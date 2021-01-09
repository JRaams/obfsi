const { buildSchema } = require('graphql');

const schema = buildSchema(`
  type Task {
    id: Int!
    name: String!
    complete: Boolean!
  }
  type AssetGroup {
    id: Int!
    name: String!
    stars: Int! 
  }
  type Query {
    tasks: [Task]
    task(id: Int!): Task
    assetGroups: [AssetGroup]
    assetGroup(id: Int!): AssetGroup
  }
`);

module.exports = schema;
