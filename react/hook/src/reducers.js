import { useReducer } from 'react';

const initialState = {
  count: 0,
  name: 'test'
};

const initialName = {
  name: 'test'
};

function reducer(state, action) {
  switch (action.type) {
    case 'increment':
      return {...state, count: state.count + 1};
    case 'decrement':
      return {...state, count: state.count - 1};
    case 'reset':
      return reset();
    case 'getname':
      return {...state, name: action.payload};
    default:
      throw new Error();
  }
}

function nameReducer(state, action) {
  switch (action.type) {
    case 'getname':
      return {name: action.payload};
    default:
      throw new Error();
  }
}

function reset() {
  return initialState;
}

export function useCounter() {
  const [state, dispatch] = useReducer(reducer, initialState, reset);
  return [state, dispatch];
}

export function useName() {
  const [state, dispatch] = useReducer(nameReducer, initialName);
  return [state, dispatch];
}