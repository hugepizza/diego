import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DirectoryComponent } from './directory.component';
import { RouterModule, Routes } from '@angular/router';
import { SharedModule } from '../../shared/shared.module';

export const DirectoryRoutes: Routes = [
  {
    path: '',
    component: DirectoryComponent,
    data: {
      heading: '文件夹'
    }
  }
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(DirectoryRoutes),
    SharedModule
  ],
  declarations: [DirectoryComponent]
})
export class DirectoryModule { }
