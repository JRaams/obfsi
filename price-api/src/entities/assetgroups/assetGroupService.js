const assetGroups = [{ id: 1, name: 'test', stars: 5 }];

module.exports.getAssetGroups = () => assetGroups;
module.exports.getAssetGroup = (id) => assetGroups.find((ag) => ag.id === id);
