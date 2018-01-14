import { RouterModule, Routes } from '@angular/router';

import { AuthComponent } from './layout/auth/auth.component';
import { AdminLayoutComponent } from './layout/admin/admin-layout.component';
import { SharedModule } from './shared/shared.module';
import { Error404Component } from './pages/error404/error404.component';

export const AppRoutes: Routes = [{
  path: '',
  component: AdminLayoutComponent,
  children: [
    { path: '', redirectTo: 'directory', pathMatch: 'full' },
    { path: 'directory', loadChildren: './pages/directory/directory.module#DirectoryModule' },
  ]
}, {
  path: '',
  component: AuthComponent,
  children: [
    { path: 'error404', component: Error404Component },
    { path: '**', redirectTo: 'error404' },
  ]
}];
