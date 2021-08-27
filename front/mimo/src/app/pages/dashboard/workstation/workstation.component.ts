import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { NgxAstilectronService } from 'src/app/share/ngx-astilectron/ngx-astilectron.service';

@Component({
  selector: 'app-workstation',
  templateUrl: './workstation.component.html',
  styleUrls: ['./workstation.component.sass']
})
export class WorkstationComponent implements OnInit {
  MenuStatusMappings: {[index: string]: (string|boolean)[]}
  knowledgeList: {cover: string, creation_time: string, name: string, store: string, brief: string}[] = [];

  constructor(
    private astilectronSrv: NgxAstilectronService,
    private cd: ChangeDetectorRef
  ) {
    this.MenuStatusMappings = {
      knowledge: ['我的知识库', false],
      recently: ['最近使用', false],
      space: ['我的知识空间', false],
      attention: ['我的关注', false],
    };
   }

  ngOnInit(): void {
    this.ActionSelectMenu('knowledge');
  }

  StatusSelectMenu = (menu: string) => {
    if (!Object.keys(this.MenuStatusMappings).includes(menu)) { this.MenuStatusMappings[menu][1] = false; }
    return this.MenuStatusMappings[menu][1];
  }

  TextSelectMenu = (menu: string) => {
    return this.MenuStatusMappings[menu][0];
  }

  ActionSelectMenu = (menu: string) => {
    // tslint:disable-next-line: variable-name
    Object.keys(this.MenuStatusMappings).forEach(_menu => this.MenuStatusMappings[_menu][1] = false);
    this.MenuStatusMappings[menu][1] = true;
    if (menu === 'knowledge') {
      this.astilectronSrv.sendMessage({name: 'knowledge', payload: 'list', cb: msg => {

        console.log('knowledgeList:', msg.payload.data)
        this.knowledgeList = [...msg.payload.data];
        console.log('knowledgeList: ', this.knowledgeList)
        this.cd.markForCheck();    // 进行标注
        this.cd.detectChanges();   // 要多加一行这个 执行一次变化检测
      }})
    }
  }
}
