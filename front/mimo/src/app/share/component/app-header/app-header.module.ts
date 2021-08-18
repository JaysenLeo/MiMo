import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzButtonModule } from 'ng-zorro-antd/button';

import { AppHeaderComponent } from './app-header.component';



@NgModule({
  declarations: [
    AppHeaderComponent
  ],
  imports: [
    CommonModule,
    NzIconModule,
    NzButtonModule,
  ],
  exports:[
    AppHeaderComponent
  ]
})
export class AppHeaderModule { }
