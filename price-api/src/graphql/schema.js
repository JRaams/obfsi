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
  type tasks {
    getAll: [Task]
    get(id: Int!): Task
  }
  type assetGroups {
    getAll: [AssetGroup]
    getById(id: Int!): AssetGroup
    getByName(name: String!): AssetGroup
  }
  type Query {
    tasks: tasks
    assetGroups: assetGroups
  }
`);

module.exports = schema;
