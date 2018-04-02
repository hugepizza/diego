import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld.vue'
import Layout from '@/components/Layout.vue'

// File
import File from '@/components/file/File.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/vue',
      name: 'HelloWorld',
      component: HelloWorld
    }, {
      path: '/',
      component: Layout,
      children: [
        {
          path: '',
          name: 'File',
          component: File
        }
      ]
    }
  ]
})
