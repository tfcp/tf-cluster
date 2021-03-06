import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: '集群管理',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '主页', icon: 'dashboard' }
    }]
  },

  {
    path: '/system',
    component: Layout,
    redirect: '/system/user-list',
    name: 'System',
    meta: { title: '系统', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'user-list',
        name: '用户列表',
        component: () => import('@/views/system/user-list'),
        meta: { title: '用户列表', icon: 'el-icon-user-solid\n' }
      },
      {
        path: 'pwd-edit',
        name: '密码修改',
        component: () => import('@/views/system/pwd-edit'),
        meta: { title: '密码修改', icon: 'el-icon-paperclip' }
      },
      {
        path: 'user-add',
        name: '用户创建',
        hidden: true,
        component: () => import('@/views/system/user-add'),
        meta: { title: '用户创建', icon: 'tree' }
      }
    ]
  },

  {
    path: '/cluster',
    component: Layout,
    redirect: '/system/user-list',
    name: 'Cluster',
    meta: { title: '集群管理', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'config-add',
        name: '配置添加',
        hidden: true,
        component: () => import('@/views/cluster/config-add'),
      },
      {
        path: 'config-list',
        name: '配置列表',
        component: () => import('@/views/cluster/config-list'),
        meta: { title: '配置列表', icon: 'el-icon-s-management' }
      },
      {
        path: 'node-list',
        name: '节点列表',
        component: () => import('@/views/cluster/node-list'),
        meta: { title: '节点列表', icon: 'el-icon-monitor' }
      },
      {
        path: 'node-add',
        name: '节点详情',
        hidden: true,
        component: () => import('@/views/cluster/node-add')
      },
      {
        path: 'app-list',
        name: '应用列表',
        component: () => import('@/views/cluster/app-list'),
        meta: { title: '应用列表', icon: 'el-icon-takeaway-box' }
      }
    ]
  },

  {
    path: '/form',
    component: Layout,
    // hidden: true,
    children: [
      {
        path: 'index',
        name: 'Form',
        component: () => import('@/views/form/index'),
        meta: { title: 'Form', icon: 'form' }
      }
    ]
  },
  //
  // {
  //   path: 'external-link',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'https://panjiachen.github.io/vue-element-admin-site/#/',
  //       meta: { title: 'External Link', icon: 'link' }
  //     }
  //   ]
  // },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
