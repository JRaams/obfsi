const service = require('./assetService');

const getAll = () => {
  return service.getAll();
};

module.exports = {
  assets: {
    getAll,
  },
};
