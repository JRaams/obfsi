const tasks = [
  { id: 1, name: 'Go to Market', complete: false },
  { id: 2, name: 'Walk the dog', complete: true },
  { id: 3, name: 'Take a nap', complete: false },
];

module.exports.getTasks = () => tasks;
module.exports.getTask = (id) => tasks.find((o) => o.id === id);
