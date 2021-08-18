import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TurnPageService } from 'src/app/services/turn-page.service';

@Component({
  selector: 'app-knowledge',
  templateUrl: './knowledge.component.html',
  styleUrls: ['./knowledge.component.sass']
})
export class KnowledgeComponent implements OnInit {

  constructor(
    private turnPageService: TurnPageService,
    private route: ActivatedRoute,
    ) { }

  ngOnInit(): void {
  }

    actionToArchive = (aid?: number) => {

        this.turnPageService.goNavigate(['../archive', ], { queryParams: { aid: aid }, relativeTo: this.route})
    }
}
