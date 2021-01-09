const service = require('./taskService');

const tasks = () => {
  return service.getTasks();
};

const task = (args) => {
  return service.getTask(args.id);
};

module.exports = {
  tasks,
  task,
};
