const service = require('./assetGroupService');

const assetGroups = () => {
  return service.getAssetGroups();
};

const assetGroup = ({ id }) => {
  const ag = service.getAssetGroup(id);
  if (ag === undefined) {
    throw new Error(`No assetgroup found with id: ${id}`);
  }
  return ag;
};

module.exports = {
  assetGroups,
  assetGroup,
};
