import { INCREMENT, DECREMENT, GET_NAME } from './store';

export const increment = () => ({
  type: INCREMENT
});

export const decrement = () => ({
  type: DECREMENT
});

export const getName = (name) => ({
  type: GET_NAME,
  payload: name
});