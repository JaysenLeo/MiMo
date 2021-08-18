import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { DashboardComponent } from './dashboard.component';
import { KnowledgeComponent } from './knowledge/knowledge.component';
import { SquareComponent } from './square/square.component';

import { MarkdownModule } from 'ngx-markdown';

import { NzDividerModule } from 'ng-zorro-antd/divider';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzAvatarModule } from 'ng-zorro-antd/avatar';
import {
  UserOutline, MoreOutline, FieldTimeOutline, HeartOutline, CodeSandboxOutline,
  PlusOutline
} from '@ant-design/icons-angular/icons';
import { KListComponent } from './knowledge/k-list/k-list.component';
import { WorkstationComponent } from './workstation/workstation.component';
import { AppHeaderModule } from 'src/app/share/component/app-header/app-header.module';
@NgModule({
  declarations: [
    DashboardComponent,
    KnowledgeComponent,
    SquareComponent,
    KListComponent,
    WorkstationComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MarkdownModule.forChild(),
    NzGridModule,
    NzButtonModule,
    NzAvatarModule,
    NzDividerModule,
    NzIconModule.forChild([UserOutline, MoreOutline, FieldTimeOutline, HeartOutline, CodeSandboxOutline,
    PlusOutline]),
    AppHeaderModule
  ]
})
export class DashboardModule { }
