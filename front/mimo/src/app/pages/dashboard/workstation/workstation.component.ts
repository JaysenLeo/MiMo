import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-workstation',
  templateUrl: './workstation.component.html',
  styleUrls: ['./workstation.component.sass']
})
export class WorkstationComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }
  MenuStatusMappings: {[index:string]: boolean} = {}
  StatusSelectMenu = (menu: string) => {
    if(!Object.keys(this.MenuStatusMappings).includes(menu)) this.MenuStatusMappings[menu] = false
    return this.MenuStatusMappings[menu]
  }
  ActionSelectMenu = (menu: string) => {
    Object.keys(this.MenuStatusMappings).forEach(_menu => this.MenuStatusMappings[_menu] = false)
    this.MenuStatusMappings[menu] = true
  }
}
