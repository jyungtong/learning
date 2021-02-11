import React, { useState } from 'react';
import BoxComponent from './components/BoxComponent';
import Counter from './components/Counter';
import MyName from './components/MyName';
import { useThemeUpdate } from './contexts/ThemeContext';
import './App.css';

function App() {
  const [count, setCount] = useState(0);
  const [name, setName] = useState('');
  const toggleTheme = useThemeUpdate();

  console.log('=====App.js rerender');

  function onClick() {
    return setCount(count + 1);
  }

  async function onClickGetName() {
    const resp = await window.fetch('https://jsonplaceholder.typicode.com/users/1');
    const data = await resp.json();

    setName(data.name);
  }

  // On first load data
  // useEffect(() => {
  //   async function fetchData() {
  //     const resp = await window.fetch('https://jsonplaceholder.typicode.com/users/1');
  //     const data = await resp.json();

  //     setName(data.name);
  //   }

  //   fetchData();
  // }, [])

  console.log('====rerender App.js');

  return (
    <div className="App">
      <h1>Clicked {count} times</h1>
      <button onClick={onClick}>Click me</button>

      <h1>User: {name}</h1>
      <button onClick={onClickGetName}>Click to get name</button>

      <BoxComponent/>
      <button onClick={toggleTheme}>Toggle theme</button>

      <hr/>
      <Counter />

      <hr/>
      <MyName />
    </div>
  );
}

export default App;
