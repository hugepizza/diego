import { Routes } from '@angular/router';

import { AdminLayoutComponent } from './layout/admin/admin-layout.component';
import { AuthLayoutComponent } from './layout/auth/auth-layout.component';
import { Error404Component } from './pages/error404/error404.component';

export const AppRoutes: Routes = [{
   { path: '**', component: Error404Component }
}];
