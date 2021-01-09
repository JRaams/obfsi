const { assetGroupResolver } = require('../entities/assetgroups');
const { assetResolver } = require('../entities/assets');

module.exports = {
  ...assetGroupResolver,
  ...assetResolver,
};
