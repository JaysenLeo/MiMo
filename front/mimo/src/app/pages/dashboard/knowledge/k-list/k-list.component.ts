import { Component, OnInit } from '@angular/core';
import { OneMyKDataBriefResp } from './k-list.interface';

@Component({
  selector: 'app-k-list',
  templateUrl: './k-list.component.html',
  styleUrls: ['./k-list.component.sass']
})
export class KListComponent implements OnInit {

  MyKListDataR: OneMyKDataBriefResp[] = [
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
    { UAvatarUrl: "//zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png", NikeName: "李成杰1", KTitle: "标题1", KCoverImgUrl: "http://placehold.it/70x100", KBrief: "概述"},
  ]
  constructor() { }

  ngOnInit(): void {
  }

}
