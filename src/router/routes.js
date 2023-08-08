const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        name: 'dashboard',
        component: () => import('pages/DashboardPage.vue'),
      },
      {
        path: '/charts',
        name: 'charts',
        component: () => import('pages/ChartsPage.vue'),
      },
      {
        path: '/controls',
        name: 'controls',
        component: () => import('pages/ControlsPage.vue'),
      },
      {
        path: '/account',
        name: 'account',
        component: () => import('pages/AccountPage.vue'),
      },
    ],
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
]

export default routes
