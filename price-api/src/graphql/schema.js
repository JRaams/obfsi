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
    assetGroupByID(id: Int!): AssetGroup
    assetGroupByName(name: String!): AssetGroup
  }
`);

module.exports = schema;
