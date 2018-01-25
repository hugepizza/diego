import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';
import { LocationStrategy, PathLocationStrategy } from '@angular/common';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { SharedModule } from './shared/shared.module';
import { AppRoutes } from './app.routing';
import { AppComponent } from './app.component';
import { Error404Component } from './pages/error404/error404.component';
import { AuthComponent } from './layout/auth/auth.component'
import { AdminLayoutComponent } from './layout/admin/admin-layout.component';

@NgModule({
  declarations: [
    AppComponent,
    Error404Component,
    AuthComponent,
    AdminLayoutComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    RouterModule.forRoot(AppRoutes),
    SharedModule,
    HttpModule,
    FormsModule,
  ],
  providers: [
    { provide: LocationStrategy, useClass: PathLocationStrategy }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
