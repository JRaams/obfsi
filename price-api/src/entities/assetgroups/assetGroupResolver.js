const service = require('./assetGroupService');

const getAll = () => {
  return service.getAssetGroups();
};

const getById = ({ id }) => {
  const ag = service.getAssetGroupByID(id);
  if (ag === undefined) {
    throw new Error(`No assetgroup found with id: ${id}`);
  }
  return ag;
};

const getByName = ({ name }) => {
  const ag = service.getAssetGroupByName(name);
  if (ag === undefined) {
    throw new Error(`No assetgroup found with name: ${name}`);
  }
  return ag;
};

module.exports = {
  assetGroups: {
    getAll,
    getById,
    getByName,
  },
};
