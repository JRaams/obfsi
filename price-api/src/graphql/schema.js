const { buildSchema } = require('graphql');

const schema = buildSchema(`
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

  type assetGroups {
    getAll: [AssetGroup]
    getById(id: Int!): AssetGroup
    getByName(name: String!): AssetGroup
  }
  type assets {
    getAll: [Asset]
    getById(id: Int!): Asset
    getByName(name: String!): Asset
  }

  type Query {
    assetGroups: assetGroups
    assets: assets
  }
`);

module.exports = schema;
