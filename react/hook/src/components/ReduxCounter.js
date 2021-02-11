import { connect } from 'react-redux';
import { increment, decrement } from '../redux/actions';

function ReduxCounter({ count, increment, decrement }) {
  console.log('====rerender ReduxCounter');

  return (
    <>
      Count: {count}
      {/* <button
        onClick={() => dispatch({type: 'reset'})}>
          Reset
      </button> */}
      <button onClick={decrement}>-</button>
      <button onClick={increment}>+</button>
    </>
  );
}

export default connect(
  ({ testing })=> ({
    count: testing.count
  }),
  { increment, decrement }
)(ReduxCounter);