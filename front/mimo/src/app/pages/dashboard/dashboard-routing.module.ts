import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './dashboard.component';
import { ArchiveComponent } from './archive/archive.component';
import { KnowledgeComponent } from './knowledge/knowledge.component';
import { SquareComponent } from './square/square.component';
import { WorkstationComponent } from './workstation/workstation.component';

const routes: Routes = [
  { path: '', component: DashboardComponent, children: [
    { path: 'square', component: SquareComponent },
    { path: 'workstation', component: WorkstationComponent,},
    { path: 'archive', component: ArchiveComponent },
    { path: 'knowledge', component: KnowledgeComponent },
    { path: '', pathMatch: 'full', redirectTo: 'knowledge' },
    // { path: '**', pathMatch: 'full', redirectTo: 'knowledge' },
   ]
  }

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
