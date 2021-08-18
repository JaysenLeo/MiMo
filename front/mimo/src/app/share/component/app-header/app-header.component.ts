import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TurnPageService } from 'src/app/services/turn-page.service';
interface MenuStatusMapping {
  Label:    string,
  Text:     string,
  Url:      string,
  Icon:     string,
  IsActive: boolean
}
@Component({
  selector: 'app-app-header',
  templateUrl: './app-header.component.html',
  styleUrls: ['./app-header.component.sass']
})
export class AppHeaderComponent implements OnInit {

    menuStatusMapping: MenuStatusMapping[] = []
    constructor(
      private turnPageService: TurnPageService,
      private route: ActivatedRoute,
    ) {
      this.menuStatusMapping = [
        { Label: "workstation", Text: "我的工作台", Url: "", Icon: "",              IsActive: false },
        { Label: "knowledge",   Text: "知识库", Url: "", Icon: "",              IsActive: false },
        { Label: "square",      Text: "知识广场",   Url: "", Icon: "ng-zorro:antd", IsActive: false },
      ]
    }

    ngOnInit(): void {
    }

    ActionSelectMenu = (menu: string) => {
      this.menuStatusMapping.map(_menu => {
        if(_menu.Label == menu) {
          _menu.IsActive = true
        } else {
          _menu.IsActive = false
        }
      })
        this.turnPageService.goNavigate([`./${menu}`, ], {relativeTo: this.route})
    }
}
