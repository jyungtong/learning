export function getName({ dispatch }) {
  async function fetchData() {
    const resp = await window.fetch('https://jsonplaceholder.typicode.com/users/1');
    const data = await resp.json();

    dispatch({
      type: 'getname',
      payload: data.name
    });
  }

  fetchData();
}
