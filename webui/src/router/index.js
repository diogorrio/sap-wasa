/*
MAIN STRUCTURE:

| WELCOME SCREEN |
|    (LOGIN)     |

| LOGIN SCREEN  |

| STREAM SCREEN |
|    (NEXT)     |

| USER PROFILE SCREEN |

| ERROR SCREEN |

 */

import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '.../views/LoginView.vue'
import StreamView from '.../views/StreamView.vue'
import UserProfileView from '.../views/UserProfileView.vue'
import ErrorView from '.../views/ErrorView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/session', component: LoginView},
		{path: '/users/:user_id/stream', component: StreamView},
		{path: '/users/:user_id/user_profile', component: UserProfileView},
		{path: '/:catchAll(.*)', component: ErrorView},
	]
})

export default router
