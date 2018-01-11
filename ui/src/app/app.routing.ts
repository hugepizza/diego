import { RouterModule, Routes } from '@angular/router';

// import { AuthComponent } from './layout/auth/auth.component';
import { SharedModule } from './shared/shared.module';
import { Error404Component } from './pages/error404/error404.component';

@NgModule({
  imports: [
    SharedModule
  ],
  declarations: [Error404Component]
})

export const AppRoutes: Routes = [
  { path: '', component: Error404Component },
  { path: '**', component: Error404Component }
];
