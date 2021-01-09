const service = require('./assetGroupService');

const assetGroups = () => {
  return service.getAssetGroups();
};

const assetGroupByID = ({ id }) => {
  const ag = service.getAssetGroupByID(id);
  if (ag === undefined) {
    throw new Error(`No assetgroup found with id: ${id}`);
  }
  return ag;
};

const assetGroupByName = ({ name }) => {
  const ag = service.getAssetGroupByName(name);
  if (ag === undefined) {
    throw new Error(`No assetgroup found with name: ${name}`);
  }
  return ag;
};

module.exports = {
  assetGroups,
  assetGroupByID,
  assetGroupByName,
};
