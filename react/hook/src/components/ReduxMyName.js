import { connect } from 'react-redux';
import { getName } from '../redux/actions';

function MyName({ name, count, getName }) {
  console.log('====rerender ReduxMyName');

  function fetch() {
    async function fetchData() {
      function getRandomInt(max) {
        return Math.floor(Math.random() * Math.floor(max)) + 1;
      }

      const resp = await window.fetch(`https://jsonplaceholder.typicode.com/users/${getRandomInt(10)}`);
      const data = await resp.json();

      getName(data.name);
    }

    fetchData();
  }

  return (
    <>
      <h2>My Name: {name}</h2>
      <h2>Count: {count}</h2>
      <button onClick={fetch}>Click to get name</button>
    </>
  )
}

export default connect(
  ({ testing }) => ({
    name: testing.name,
    count: testing.count
  }),
  { getName }
)(MyName);