const service = require('./taskService');

const tasks = () => {
  return service.getTasks();
};

const task = ({ id }) => {
  const t = service.getTask(id);
  if (t === undefined) {
    throw new Error(`No task found with id: ${id}`);
  }
  return t;
};

module.exports = {
  tasks,
  task,
};
