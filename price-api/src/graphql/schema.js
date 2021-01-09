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
  type Asset {
    id: Int!
    assetGroup: AssetGroup
    name: String!
    namecolor: String!
    iconurl: String!
    type: String!
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
  type assets {
    getAll: [Asset]
  }
  type Query {
    tasks: tasks
    assetGroups: assetGroups
    assets: assets
  }
`);

module.exports = schema;
