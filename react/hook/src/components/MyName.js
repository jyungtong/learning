import { useName } from '../reducers';
import { getName } from '../actions';

export default function MyName() {
  console.log('====rerender MyName');

  const [state, dispatch] = useName();

  return (
    <>
      <h2>My Name: {state.name}</h2>
      <button onClick={() => getName({ dispatch })}>Click to get name</button>
    </>
  )
}