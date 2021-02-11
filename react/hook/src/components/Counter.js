import { useCounter } from '../reducers';

export default function Counter() {
  console.log('====rerender Counter');

  const [state, dispatch] = useCounter();

  return (
    <>
      Count: {state.count}
      <button
        onClick={() => dispatch({type: 'reset'})}>
          Reset
      </button>
      <button onClick={() => dispatch({type: 'decrement'})}>-</button>
      <button onClick={() => dispatch({type: 'increment'})}>+</button>
    </>
  );
}