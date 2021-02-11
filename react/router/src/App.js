import React from 'react';
import {
  Link,
  useParams
} from 'react-router-dom';

export default function App() {
  return (
    <div>
      <ul>
        <li>
          <Link to="/home">home</Link>
        </li>
        <li>
          <Link to="/topics">Topics</Link>
        </li>
      </ul>
    </div>
  )
}

export function Home() {
  return (
    <div>
      <h2>Home</h2>
    </div>
  )
}

export function Topics() {
  return (
    <div>
      <h2>Topics</h2>

      <ul>
        <li>
          <Link to="/topics/test1">test1</Link>
        </li>
        <li>
          <Link to="/topics/test2">test2</Link>
        </li>
        <li>
          <Link to="/topics/test3">test3</Link>
        </li>
      </ul>
    </div>
  )
}

export function Topic() {
  const { topicId } = useParams();

  return (
    <div>
      <h2>Topic: {topicId}</h2>
    </div>
  )
}