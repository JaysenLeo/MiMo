import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { DashboardComponent } from './dashboard.component';
import { KnowledgeComponent } from './knowledge/knowledge.component';
import { SquareComponent } from './square/square.component';


@NgModule({
  declarations: [
    DashboardComponent,
    KnowledgeComponent,
    SquareComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule
  ]
})
export class DashboardModule { }
