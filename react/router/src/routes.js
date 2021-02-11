import { Home, Topics, Topic } from './App';

export const routes = [
  {
    path: '/home',
    component: Home
  },
  {
    path: '/topics',
    component: Topics
  }
];

export const topicRoutes = [
  {
    path: '/topics/:topicId',
    component: Topic
  }
]