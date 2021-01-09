const { taskResolver } = require('../entities/tasks');
const { assetGroupResolver } = require('../entities/assetgroups');

module.exports = {
  ...taskResolver,
  ...assetGroupResolver,
};
