const service = require('./taskService');

const getAll = () => {
  return service.getTasks();
};

const get = ({ id }) => {
  const t = service.getTask(id);
  if (t === undefined) {
    throw new Error(`No task found with id: ${id}`);
  }
  return t;
};

module.exports = {
  tasks: {
    getAll,
    get,
  },
};
