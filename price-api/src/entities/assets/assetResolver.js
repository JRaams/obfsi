const service = require('./assetService');

const getAll = () => {
  return service.getAssets();
};

module.exports = {
  assets: {
    getAll,
  },
};
