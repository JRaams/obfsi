const polka = require('polka')
const { json } = require('body-parser')
const { buildSchema } = require('graphql')
const { graphqlHTTP } = require('express-graphql')

const PORT = process.env.PORT || 3000;
const PRODUCTION = !process.env.PRODUCTION || false;

const tasks = [
	{ id:1, name:'Go to Market', complete:false },
	{ id:2, name:'Walk the dog', complete:true },
	{ id:3, name:'Take a nap', complete:false }
];

const schema = buildSchema(`
	type Task {
		id: Int!
		name: String!
		complete: Boolean!
	}
	type Query {
		tasks: [Task]
		task(id: Int!): Task
	}
`);

let rootValue = {
  tasks: () => tasks,
  task: (args) => tasks.find(o => o.id === args.id)
}

polka()
  .use(json())
  .use('/graphql', graphqlHTTP({
    schema,
    rootValue,
    graphiql: PRODUCTION
  }))
  .listen(PORT, err => {
    if (err) throw err;
    console.log('API Server listening on :', PORT)
  })