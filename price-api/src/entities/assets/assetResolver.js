const service = require('./assetService');

const getAll = () => {
  return service.getAll();
};

const getById = ({ id }) => {
  const ag = service.getById(id);
  if (ag === undefined) {
    throw new Error(`No asset found with id: ${id}`);
  }
  return ag;
};

const getByName = ({ name }) => {
  const ag = service.getByName(name);
  if (ag === undefined) {
    throw new Error(`No asset found with name: ${name}`);
  }
  return ag;
};

module.exports = {
  assets: {
    getAll,
    getById,
    getByName,
  },
};
