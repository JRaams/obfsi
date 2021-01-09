const { taskResolver } = require('../entities/tasks');
const { assetGroupResolver } = require('../entities/assetgroups');
const { assetResolver } = require('../entities/assets');

module.exports = {
  ...taskResolver,
  ...assetGroupResolver,
  ...assetResolver,
};
