import { createPlugin, createRouteRef } from '@backstage/core';
import ExampleComponent from './components/welcome';
import Login from './components/Login';
import Create from './components/BookBorrow';

export const rootRouteRef = createRouteRef({
  path: '/welcome',
  title: 'welcome',
});

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.addRoute(rootRouteRef, ExampleComponent);
    router.registerRoute('/BookBorrow', Create);
    router.registerRoute('/Login', Login);
  },
});
