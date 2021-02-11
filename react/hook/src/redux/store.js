import { createStore, combineReducers } from 'redux';

export const INCREMENT = 'increment';
export const DECREMENT = 'decrement';
export const GET_NAME = 'getname';

const initialState = {
  count: 0,
  name: 'test-redux'
};

function testing(state = initialState, action) {
  switch (action.type) {
    case INCREMENT:
      return {
        ...state,
        count: state.count + 1
      }
    case DECREMENT:
      return {
        ...state,
        count: state.count - 1
      }
    case GET_NAME:
      return {
        ...state,
        name: action.payload
      }
    default:
      return state;
  }
}

export default createStore(combineReducers({ testing }));